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
	ATResponse, err := Mx01.AT("AT+MAC?\r\n", 200*time.Millisecond)
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
	ATResponse, err := Mx01.AT("AT+NAME?\r\n", 200*time.Millisecond)
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
* AT+ADV?
*
 */

func ADV(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+ADV?\r\n", 200*time.Millisecond)
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
* AT+UART?
*
 */

func UART(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+UART?\r\n", 200*time.Millisecond)
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
* AT+DEV? 查询当前已连接的设备
*
 */

func DEV(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+DEV?\r\n", 200*time.Millisecond)
	if err != nil {
		return "", err
	}
	if len(ATResponse.Data) == 0 {
		return "No device connected", nil
	}
	if len(ATResponse.Data) == 1 {
		return ATResponse.Data[0], nil
	}
	return "", fmt.Errorf("invalid result:%v", ATResponse.String())
}

/*
*
* AT+AINTVL? 查询广播间隔
*
 */

func AINTVL(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+AINTVL?\r\n", 200*time.Millisecond)
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
* AT+VER? 读取软件版本
*
 */

func VER(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+VER?\r\n", 200*time.Millisecond)
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
* AT+TXPOWER? 查询模组的发射功率
*
 */

func TXPOWER(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+TXPOWER?\r\n", 200*time.Millisecond)
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
* AT+UUIDS? 查询 BLE 主服务通道
*
 */

func UUIDS(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+UUIDS?\r\n", 200*time.Millisecond)
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
* AT+UUIDN? 查询 BLE 读服务通道
*
 */

func UUIDN(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+UUIDN?\r\n", 200*time.Millisecond)
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
* AT+UUIDW? 查询 BLE 写服务通道
*
 */

func UUIDW(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+UUIDW?\r\n", 200*time.Millisecond)
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
* AT+AMDATA? 查询自定义广播数据
*
 */

func AMDATA(Mx01 device.Device) (string, error) {
	ATResponse, err := Mx01.AT("AT+AMDATA?\r\n", 200*time.Millisecond)
	if err != nil {
		return "", err
	}
	if len(ATResponse.Data) == 1 {
		return ATResponse.Data[0], nil
	}
	return "", fmt.Errorf("invalid result:%v", ATResponse.String())
}
