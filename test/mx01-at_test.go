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

	mx01 "github.com/hootrhino/rhilex-goat/bsp/mx01"

	"fmt"
	mx01At "github.com/hootrhino/rhilex-goat/bsp/mx01/atcmd"
	"time"

	serial "github.com/hootrhino/goserial"
)

// go test -timeout 30s -run ^Test_MX01_AT_Mac$ rhilex-goat/test -v -count=1
func Test_MX01_AT_Mac(t *testing.T) {
	config := serial.Config{
		Address:  "COM3",
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
		Timeout:  time.Duration(50) * time.Millisecond,
	}
	serialPort, err := serial.Open(&config)
	if err != nil {
		panic(err)
	}
	mx01 := mx01.NewMX01("mx01", serialPort)
	mx01.Flush()
	MACResponse, err := mx01At.MAC(mx01)
	if err != nil {
		panic(err)
	}
	fmt.Println("AT MACResponse=", MACResponse)
	serialPort.Close()
}

// go test -timeout 30s -run ^Test_MX01_AT_Name$ rhilex-goat/test -v -count=1
func Test_MX01_AT_Name(t *testing.T) {
	config := serial.Config{
		Address:  "COM3",
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "N",
		StopBits: 1,
		Timeout:  time.Duration(50) * time.Millisecond,
	}
	serialPort, err := serial.Open(&config)
	if err != nil {
		panic(err)
	}
	mx01 := mx01.NewMX01("mx01", serialPort)
	mx01.Flush()
	MACResponse, err := mx01At.NAME(mx01)
	if err != nil {
		panic(err)
	}
	fmt.Println("AT MACResponse=", MACResponse)
	serialPort.Close()
}
