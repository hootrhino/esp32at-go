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

package atcmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hootrhino/rhilex-goat/device"
)

/*
*
* 测试: AT- OK
*
 */
func AT(Esp32 device.Device) bool {
	ATResponse, err := Esp32.AT("AT\r\n", 100)
	if err != nil {
		return false
	}
	if len(ATResponse.Data) < 1 {
		return false
	}
	if ATResponse.Data[0] == "OK" {
		return true
	}
	return false
}

/*
*
* 重启
*
 */
func RST(Esp32 device.Device) bool {
	ATResponse, err := Esp32.AT("AT+RST\r\n", 100)
	if err != nil {
		return false
	}
	if len(ATResponse.Data) != 1 {
		return false
	}
	if ATResponse.Data[0] == "OK" {
		return true
	}
	return false
}

/*
*
<AT version info>
<SDK version info>
<compile time>
<Bin version>
OK
*
*/
type GMRResponse struct {
	AtVersion   string `json:"atVVersion"`
	SDKVersion  string `json:"sdkVersion"`
	CompileTime string `json:"compileTime"`
	BinVersion  string `json:"binVersion"`
}

func (O GMRResponse) String() string {
	if bytes, err := json.Marshal(O); err != nil {
		return ""
	} else {
		return string(bytes)
	}
}
func GMR(Esp32 device.Device) (GMRResponse, error) {
	GMRResponse := GMRResponse{}
	ATResponse, err := Esp32.AT("AT+GMR\r\n", 200*time.Millisecond)
	if err != nil {
		return GMRResponse, err
	}
	if len(ATResponse.Data) != 5 {
		return GMRResponse, fmt.Errorf("request GMR error:%v", ATResponse.Data)
	}
	GMRResponse.AtVersion = ATResponse.Data[0]
	GMRResponse.SDKVersion = ATResponse.Data[1]
	GMRResponse.CompileTime = ATResponse.Data[2]
	GMRResponse.BinVersion = ATResponse.Data[3]
	return GMRResponse, nil
}
