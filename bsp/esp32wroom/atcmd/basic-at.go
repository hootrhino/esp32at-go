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
	"errors"
	"fmt"
	"net"
	"rhilex-goat/device"
	"time"
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

/*
*
进入 Deep-sleep 模式
OK
*
*/
func Deep_sleep(Esp32 device.Device, sleepTime int) bool {
	var cmd string = fmt.Sprintf("AT+GSLP=%d\r\n", sleepTime)
	ATResponse, err := Esp32.AT(cmd, 200*time.Millisecond)
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
关闭 AT 回显功能
OK
*
*/
func ATE0(Esp32 device.Device) bool {
	ATResponse, err := Esp32.AT("ATE0\r\n", 200*time.Millisecond)
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
开启 AT 回显功能
OK
*
*/
func ATE1(Esp32 device.Device) bool {
	ATResponse, err := Esp32.AT("ATE1\r\n", 200*time.Millisecond)
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
AT+SAVETRANSLINK 设置透传模式
*
*/

/*
设置开机进入 TCP/SSL 透传模式 信息
AT+SAVETRANSLINK=
<mode>,
<"remote host">,
<remote port>,
<"type">,
<keep_alive>
*/

type TcpSslSTLRequest struct {
	Mode        int    `json:"Mode"`
	Remote_host string `json:"Remote_host"`
	Remote_port int    `json:"Remote_port"`
	STL_type    string `json:"STL_type"`
	Keep_alive  int    `json:"Keep_alive"`
}

func NewTcpSslSTLRequest(request TcpSslSTLRequest) error {
	if request.Mode == 0 {
		if request.Remote_host != "" || request.Remote_port != 0 ||
			request.STL_type != "" || request.Keep_alive != 0 {
			return errors.New("when mode is 0, other fields must be empty")
		}
		return nil
	}
	if request.Mode != 1 {
		return errors.New("mode must be 0 or 1")
	}

	if len([]byte(request.Remote_host)) > 64 {
		return errors.New("remote_host must be at most 64 bytes")
	}
	if ip := net.ParseIP(request.Remote_host); ip == nil {
		// 如果remoteHost不是有效的IP地址，尝试解析为域名
		if _, err := net.LookupHost(request.Remote_host); err != nil {
			return errors.New("remote_host must be a valid IPv4, IPv6 address, or domain name")
		}
	}
	if request.Remote_port < 0 || request.Remote_port > 65535 {
		return errors.New("remote port must be between 0 and 65535")
	}

	if request.STL_type != "TCP" && request.STL_type != "TCPv6" &&
		request.STL_type != "SSL" && request.STL_type != "SSLv6" {
		return errors.New("stl_type must be one of TCP, TCPv6, SSL, SSLv6")
	}

	if request.Keep_alive < 0 || request.Keep_alive > 7200 {
		return errors.New("keep_alive must be between 0 and 7200")
	}
	return nil
}

func TcpSslSTL(Esp32 device.Device, request TcpSslSTLRequest) (bool, error) {
	err := NewTcpSslSTLRequest(request)
	if err != nil {
		return false, err
	}
	var cmd string
	if request.Mode == 0 {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d\r\n", request.Mode)
	} else {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d,\"%s\",%d,\"%s\",%d\r\n",
			request.Mode,
			request.Remote_host,
			request.Remote_port,
			request.STL_type,
			request.Keep_alive)
	}

	ATResponse, err := Esp32.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request TcpSslSTL error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] == "OK" {
		return true, nil
	}
	return false, fmt.Errorf("request TcpSslSTL error:%v", ATResponse.Data)
}

/*
设置开机进入 UDP 透传模式 信息
AT+SAVETRANSLINK=
<mode>,
<"remote host">,
<remote port>,
<"type">,
<local port>
*/
type UDPRequest struct {
	Mode        int    `json:"Mode"`
	Remote_host string `json:"Remote_host"`
	Remote_port int    `json:"Remote"`
	STL_type    string `json:"STL_type"`
	Local_port  int    `json:"Local_port"`
}

func NewUDPRequest(request UDPRequest) error {
	if request.Mode == 0 {
		if request.Remote_host != "" || request.Remote_port != 0 ||
			request.STL_type != "" || request.Local_port != 0 {
			return errors.New("when mode is 0, other fields must be empty")
		}
		return nil
	}
	if request.Mode != 1 {
		return errors.New("mode must be 0 or 1")
	}

	if len([]byte(request.Remote_host)) > 64 {
		return errors.New("remote_host must be at most 64 bytes")
	}
	if ip := net.ParseIP(request.Remote_host); ip == nil {
		// 如果remoteHost不是有效的IP地址，尝试解析为域名
		if _, err := net.LookupHost(request.Remote_host); err != nil {
			return errors.New("remote_host must be a valid IPv4, IPv6 address, or domain name")
		}
	}
	if request.Remote_port < 0 || request.Remote_port > 65535 {
		return errors.New("remote port must be between 0 and 65535")
	}

	if request.STL_type != "UDP" && request.STL_type != "UDPv6" {
		return errors.New("stl_type must be one of UDP, UDPv6")
	}

	if request.Local_port < 0 || request.Local_port > 65535 {
		return errors.New("remote port must be between 0 and 65535")
	}

	return nil
}

func UdpSTL(Esp32 device.Device, request UDPRequest) (bool, error) {
	err := NewUDPRequest(request)
	if err != nil {
		return false, err
	}
	var cmd string
	if request.Mode == 0 {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d\r\n", request.Mode)
	} else {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d,\"%s\",%d,\"%s\",%d\r\n",
			request.Mode,
			request.Remote_host,
			request.Remote_port,
			request.STL_type,
			request.Local_port)
	}

	ATResponse, err := Esp32.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request UdpSTL error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request UdpSTL error:%v", ATResponse.Data)
	}
	return true, nil
}

/*
设置开机进入 BLE 透传模式
AT+SAVETRANSLINK=
<mode>,
<role>,
<tx_srv>,
<tx_char>,
<rx_srv>,
<rx_char>,
<peer_addr>
*/
type BLERequest struct {
	Mode      int    `json:"Mode"`
	Role      int    `json:"Role"`
	Tx_srv    int    `json:"Tx_srv"`
	Tx_char   int    `json:"Tx_char"`
	Rx_srv    int    `json:"STL_type"`
	Rx_char   int    `json:"Local_port"`
	Peer_addr string `json:"Peer_addr"`
}

func BleSTL(Esp32 device.Device, request BLERequest) (bool, error) {
	var cmd string
	if request.Mode == 0 {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d\r\n", request.Mode)
	} else {
		cmd = fmt.Sprintf("AT+SAVETRANSLINK=%d,%d,%d,%d,%d,%d,\"%s\"\r\n",
			request.Mode,
			request.Role,
			request.Tx_srv,
			request.Tx_char,
			request.Rx_srv,
			request.Rx_char,
			request.Peer_addr)
	}

	ATResponse, err := Esp32.AT(cmd, 200*time.Millisecond)
	if err != nil {
		return false, err
	}
	if len(ATResponse.Data) != 1 {
		return false, fmt.Errorf("request UdpSTL error:%v", ATResponse.Data)
	}
	if ATResponse.Data[0] != "OK" {
		return false, fmt.Errorf("request UdpSTL error:%v", ATResponse.Data)
	}
	return true, nil
}
