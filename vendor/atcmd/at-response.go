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
	"io"
	"strings"
)

type ATCommandResponse struct {
	Command string
	Data    []byte
}

// parseATCommandResponse
func parseATCommandResponse(reader io.Reader) (*ATCommandResponse, error) {
	response := &ATCommandResponse{}
	line, err := readLine(reader)
	if err != nil {
		return nil, err
	}
	response.Command = line
	response.Data, err = readRemainingData(reader)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func readLine(reader io.Reader) (string, error) {
	var buffer [256]byte
	_, err := reader.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buffer[:])), nil
}

func readRemainingData(reader io.Reader) ([]byte, error) {
	var buffer []byte
	for {
		chunk, err := readLine(reader)
		if err == io.EOF {
			break
		}
		buffer = append(buffer, []byte(chunk+"\r\n")...)
	}
	return buffer, nil
}

func TestAt() {
	response := `
AT+CIPSTATUS\r\n
STATE: IP STA\r\n
IPADDR: 192.168.1.100\r\n
NETMASK: 255.255.255.0\r\n
GW: 192.168.1.1\r\n
RSSI: -69\r\n\r\nOK\r\n
`
	resp, err := parseATCommandResponse(strings.NewReader(response))
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// 输出解析结果
	fmt.Printf("Command: %s\n", resp.Command)
	fmt.Printf("Data: %s\n", string(resp.Data))
}
