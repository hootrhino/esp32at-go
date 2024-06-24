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

package test

import (
	"testing"

	"fmt"
	esp32wroom "rhilex-goat/bsp/esp32wroom"
	esp32wroomAt "rhilex-goat/bsp/esp32wroom/atcmd"
	"time"

	serial "github.com/hootrhino/goserial"
)

// go test -timeout 30s -run ^Test_Esp32Wroom_AT_Mac$ rhilex-goat/test -v -count=1
func Test_Esp32Wroom_AT_Mac(t *testing.T) {
	SerialPeerRwTimeout := 50 * time.Millisecond
	config := serial.Config{
		Address:  "COM3",
		BaudRate: 115200,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
		Timeout:  SerialPeerRwTimeout,
	}
	serialPort, err := serial.Open(&config)
	if err != nil {
		panic(err)
	}
	Esp32 := esp32wroom.NewEsp32Wroom("ESP32-WROOM", serialPort)
	Esp32.Flush()
	GMRResponse, err := esp32wroomAt.GMR(Esp32)
	if err != nil {
		panic(err)
	}
	fmt.Println("AT GMRResponse=", GMRResponse)
	RSTResponse := esp32wroomAt.RST(Esp32)
	fmt.Println("AT RSTResponse=", RSTResponse)
	serialPort.Close()
}
