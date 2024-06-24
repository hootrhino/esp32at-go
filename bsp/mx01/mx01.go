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
	"fmt"
	"io"
	"rhilex-goat/device"
	"strings"
	"sync"
	"time"
)

func NewMX01(name string, io io.ReadWriteCloser) device.Device {
	return &MX01{name: name, io: io}
}

type MX01 struct {
	name string
	io   io.ReadWriteCloser
}

func (MX01 *MX01) Init(config map[string]any) error {
	return nil
}
func (MX01 *MX01) Close() error {
	return nil
}
func (MX01 *MX01) Flush() {
	var responseData [1]byte
	for {
		N, _ := MX01.io.Read(responseData[:])
		if N == 0 {
			return
		}
	}
}
func (MX01 *MX01) AT(AtCmd string, HwCardResponseTimeout time.Duration) (device.ATResponse, error) {
	ATResponse := device.ATResponse{Command: AtCmd}
	_, errWrite := MX01.io.Write([]byte(AtCmd))
	if errWrite != nil {
		return ATResponse, errWrite
	}
	var responseData [256]byte
	Ctx, Cancel := context.WithTimeout(context.Background(), HwCardResponseTimeout)
	wg := sync.WaitGroup{}
	wg.Add(1)
	var errRw error
	acc := 0
	go func(io io.ReadWriteCloser) {
		defer wg.Done()
		defer Cancel()
		for {
			select {
			case <-Ctx.Done():
				return
			default:
				N, errRead := MX01.io.Read(responseData[acc:])
				if errRead != nil {
					if strings.Contains(errRead.Error(), "timeout") {
						if N > 0 {
							acc += N
						}
						continue
					}
					errRw = errRead
					return
				}
				if N > 0 {
					acc += N
				}
			}
		}
	}(MX01.io)
	wg.Wait()
	atReturn := []string{}
	if len(AtCmd)-4 > 0 {
		// AT+NAME?\r\n
		// +NAME:XXXXX
		returnId := fmt.Sprintf("AT%s", string(responseData[:len(AtCmd)-4]))
		ValidId := (returnId[:len((returnId))-1] == AtCmd[:len((AtCmd))-3])
		if ValidId {
			finalByte := responseData[len(AtCmd)-4 : acc]
			for _, s := range strings.Split(string(finalByte), "\r\n") {
				if s != "" {
					atReturn = append(atReturn, s)
				}
			}
		}
		ATResponse.Data = atReturn
	}

	return ATResponse, errRw
}
