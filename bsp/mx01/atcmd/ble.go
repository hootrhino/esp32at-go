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
	"fmt"
	"rhilex-goat/device"
	"time"
)

/*
*
* AT+MAC?
*
 */
func MAC(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+MAC?\r\n", 1000*time.Millisecond)
	if err != nil {
		return "", err
	}
	if len(ATResponse.Data) == 1 {
		return ATResponse.Data[0], nil
	}
	return "", fmt.Errorf("invalid result:%v", ATResponse.String())
}

/*
*
* AT+NAME?
*
 */
func NAME(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+NAME?\r\n", 1000*time.Millisecond)
	if err != nil {
		return "", err
	}
	if len(ATResponse.Data) == 1 {
		return ATResponse.Data[0], nil
	}
	return "", fmt.Errorf("invalid result:%v", ATResponse.String())
}
