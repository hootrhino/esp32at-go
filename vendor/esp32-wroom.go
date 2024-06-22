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

package vendor

import (
	"esp32at-go/device"
	"os"
	"time"
)

func NewEsp32Wroom(tty string) device.Device {
	return &Esp32Wroom{tty: tty}
}

type Esp32Wroom struct {
	tty string
}

func (Esp32 *Esp32Wroom) AT(command string, timeout time.Duration) (string, error) {
	file, err := os.OpenFile(Esp32.tty, os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.WriteString(command)
	if err != nil {
		return "", err
	}
	buffer := [512]byte{}
	deadline := time.Now().Add(timeout)
	file.SetReadDeadline(deadline)
	n, err := file.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
