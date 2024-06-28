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
	"errors"
	"fmt"
	"regexp"
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
* AT+MAC=<MAC>
*设置MAC地址，需要重启之后生效
 */
// 判断MAC是否在指定范围内
func IsMAC(MAC string) error {
	// 检查输入的长度是否为12位
	if len(MAC) != 12 {
		return errors.New("mac length error")
	}
	//检测字符串是否符合16进制的格式
	var hexPattern = `^[0-9a-fA-F]+$`
	match, _ := regexp.MatchString(hexPattern, MAC)
	if !match {
		return errors.New("uuid is not hex")
	}
	return nil
}

func SetMAC(Mx01 device.Device, MAC string) (bool, error) {
	err := IsMAC(MAC)
	if err != nil {
		return false, err
	}
	cmd := fmt.Sprintf("AT+MAC=%s\r\n", MAC)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetMAC error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetMAC error:%v", ATResponse.Data)
	}
	return true, nil
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
* AT+NAME=<NAME>设置NAME地址
*
 */

func SetNAME(Mx01 device.Device, NAME string) (bool, error) {
	if len(NAME) > 20 {
		return false, errors.New("name length error")
	}
	cmd := fmt.Sprintf("AT+NAME=%s\r\n", NAME)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetNAME error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetNAME error:%v", ATResponse.Data)
	}
	return true, nil
}

/*
*
* AT+ADV?查询设备蓝牙广播状态
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
*立即生效，复位重启后恢复广播。
*AT+ADV=<NUM>设置广播状态
*
 */
func SetADV(Mx01 device.Device, NUM int) (bool, error) {
	if NUM != 1 && NUM != 0 {
		return false, errors.New("num value error")
	}
	cmd := fmt.Sprintf("AT+ADV=%d\r\n", NUM)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetADV error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetADV error:%v", ATResponse.Data)
	}
	return true, nil
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
* AT+UART=<NUM>设置设备波特率
*<NUM>:0:9600/ 1:14400/ 2:19200/ 3:38400/ 4:57600/ 5:115200
 */
func SetUART(Mx01 device.Device, NUM int) (bool, error) {
	if NUM < 0 || NUM > 5 {
		return false, errors.New("num value error")
	}
	cmd := fmt.Sprintf("AT+UART=%d\r\n", NUM)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetUART error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetUART error:%v", ATResponse.Data)
	}
	return true, nil
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
*断开蓝牙连接
* AT+DISCONN=<NUM>
*0-断开所有连接的从设备 1-主动断开与主机端设备的连接,
 */
//还没有正确实现
func DISCONN(Mx01 device.Device, NUM int) (string, error) {
	if NUM != 1 && NUM != 0 {
		return "", errors.New("num value error")
	}
	//cmd := fmt.Sprintf("AT+DISCONN=%d\r\n", NUM)
	ATResponse, err := Mx01.AT("AT+DISCONN=1\r\n", 200*time.Millisecond)
	if err != nil {
		return "", err
	}
	if len(ATResponse.Data) != 1 {
		return "", fmt.Errorf("request DISCONN error:%v", ATResponse.Data)
	}

	return ATResponse.Data[0], nil
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
*重启后生效。
* AT+AINTVL= <NUM>设置广播间隔
*<NUM>：20-10000 单位毫秒
 */
func SetAINTVL(Mx01 device.Device, NUM int) (bool, error) {
	if NUM < 20 || NUM > 10000 {
		return false, errors.New("num value error")
	}
	cmd := fmt.Sprintf("AT+AINTVL=%d\r\n", NUM)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetAINTVL error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetAINTVL error:%v", ATResponse.Data)
	}
	return true, nil
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
*重启生效，MAC 地址修改后不可恢复。
* AT+RESET=1 恢复出厂设置
*
 */
func RESET(Mx01 device.Device) (bool, error) {
	ATResponse, err := Mx01.AT("AT+RESET=1\r\n", 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request RESET error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request RESET error:%v", ATResponse.Data)
	}
	return true, nil
}

/*
*
* AT+REBOOT=1 设置模组重启。
*
 */
func REBOOT(Mx01 device.Device) (bool, error) {
	ATResponse, err := Mx01.AT("AT+REBOOT=1\r\n", 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request REBOOT error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request REBOOT error:%v", ATResponse.Data)
	}
	return true, nil
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
*重启后生效。
* AT+TXPOWER=<NUM> 设置模组的发射功率
*NUM:0:5dbm/ 1:4dbm/ 2:3dbm/ 3:0dbm/ 4:-2dbm/ 5:-5dbm/ 6:-6dbm/ 7:-10dbm/ 8:-15dbm/
 */
func SetTXPOWER(Mx01 device.Device, NUM int) (bool, error) {
	if NUM < 0 || NUM > 8 {
		return false, errors.New("num value error")
	}
	cmd := fmt.Sprintf("AT+TXPOWER=%d\r\n", NUM)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetTXPOWER error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetTXPOWER error:%v", ATResponse.Data)
	}
	return true, nil
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
*重启后生效
* AT+UUIDS=<UUID> 设置 BLE 主服务通道
*支持参数：16bit 格式或 128bit 格式的 UUID
 */
// 检查输入的长度
func IsUUID(UUID string) error {
	if len(UUID) != 4 && len(UUID) != 32 {
		return errors.New("uuid length error")
	}

	var hexPattern = `^[0-9a-fA-F]+$`
	match, _ := regexp.MatchString(hexPattern, UUID)
	if !match {
		return errors.New("uuid is not hex")
	}
	return nil
}

func SetUUIDS(Mx01 device.Device, UUID string) (bool, error) {
	err := IsUUID(UUID)
	if err != nil {
		return false, err
	}
	cmd := fmt.Sprintf("AT+UUIDS=%s\r\n", UUID)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetUUIDS error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetUUIDS error:%v", ATResponse.Data)
	}
	return true, nil
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
*重启后生效
* AT+UUIDN=<UUID> 设置 BLE 读服务通道
*支持参数：16bit 格式或 128bit 格式的 UUID
 */
func SetUUIDN(Mx01 device.Device, UUID string) (bool, error) {
	err := IsUUID(UUID)
	if err != nil {
		return false, err
	}
	cmd := fmt.Sprintf("AT+UUIDN=%s\r\n", UUID)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetUUIDN error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetUUIDN error:%v", ATResponse.Data)
	}
	return true, nil
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
*重启后生效
* AT+UUIDW=<UUID> 设置 BLE 写服务通道
*支持参数：16bit 格式或 128bit 格式的 UUID
 */
func SetUUIDW(Mx01 device.Device, UUID string) (bool, error) {
	err := IsUUID(UUID)
	if err != nil {
		return false, err
	}
	cmd := fmt.Sprintf("AT+UUIDW=%s\r\n", UUID)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetUUIDW error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetUUIDW error:%v", ATResponse.Data)
	}
	return true, nil
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

/*
*
* AT+AMDATA=<HEX> 设置自定义广播数据
*
 */
func IsAMDATA(UUID string) error {
	if len(UUID)%2 != 0 {
		return errors.New("uuid length error")
	}

	var hexPattern = `^[0-9a-fA-F]+$`
	match, _ := regexp.MatchString(hexPattern, UUID)
	if !match {
		return errors.New("uuid is not hex")
	}
	return nil
}

func SetAMDATA(Mx01 device.Device, AMDATA string) (bool, error) {
	err := IsAMDATA(AMDATA)
	if err != nil {
		return false, err
	}
	cmd := fmt.Sprintf("AT+AMDATA=%s\r\n", AMDATA)
	ATResponse, err := Mx01.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request SetAMDATA error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request SetAMDATA error:%v", ATResponse.Data)
	}
	return true, nil
}
