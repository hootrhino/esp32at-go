// Copyright (C) 2024 wwhai
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package bsp

import (
	"context"
	"espressif-goat/device"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func NewEsp32Wroom(name string, io io.ReadWriteCloser) device.Device {
	return &Esp32Wroom{name: name, io: io}
}

type Esp32Wroom struct {
	name string
	io   io.ReadWriteCloser
}

func (Esp32 *Esp32Wroom) Init(config map[string]any) error {
	return nil
}
func (Esp32 *Esp32Wroom) Close() error {
	return nil
}
func (Esp32 *Esp32Wroom) Flush() {
	var responseData [1]byte
	for {
		N, _ := Esp32.io.Read(responseData[:])
		if N == 0 {
			return
		}
	}
}
func (Esp32 *Esp32Wroom) AT(AtCmd string, HwCardResponseTimeout time.Duration) (device.ATResponse, error) {
	ATResponse := device.ATResponse{Command: AtCmd}
	_, errWrite := Esp32.io.Write([]byte(AtCmd))
	if errWrite != nil {
		return ATResponse, errWrite
	}
	var responseData [256]byte
	acc := 0
	Ctx, Cancel := context.WithTimeout(context.Background(), HwCardResponseTimeout)
	wg := sync.WaitGroup{}
	wg.Add(1)
	var errRw error
	go func(io io.ReadWriteCloser) {
		defer wg.Done()
		defer Cancel()
		for {
			select {
			case <-Ctx.Done():
				return
			default:
				N, errRead := Esp32.io.Read(responseData[acc:])
				if errRead != nil {
					if strings.Contains(errRead.Error(), "timeout") {
						if N > 0 {
							acc += N
						}
						continue
					}
					errRw = errRead
					return
				} else {
					acc += N
				}
			}
		}
	}(Esp32.io)
	wg.Wait()
	if (len(AtCmd) <= len(responseData)) && (len(AtCmd) <= acc) {
		ResponseId := string(responseData[:len(AtCmd)])
		atReturn := []string{}
		if ResponseId != AtCmd {
			return ATResponse, fmt.Errorf("AT command execute error")
		}
		for _, s := range strings.Split(string(responseData[len(AtCmd):acc]), "\r\n") {
			if s != "" {
				atReturn = append(atReturn, s)
			}
		}
		ATResponse.Data = atReturn
	}
	return ATResponse, errRw
}
