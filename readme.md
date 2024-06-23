<!--
 Copyright (C) 2024 wwhai

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as
 published by the Free Software Foundation, either version 3 of the
 License, or (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <https://www.gnu.org/licenses/>.
-->

# ESP32 AT 指令封装
官方手册：
- https://robu.in/wp-content/uploads/2019/12/esp32_at_instruction_set_and_examples_en_1.0.pdf
- https://docs.espressif.com/projects/esp-at/en/latest/esp32/AT_Command_Set/BLE_AT_Commands.html

## 使用
下面是个简单的例子：

```go
package main

import (
	esp32wroom "espressif-goat/bsp/esp32wroom"
	esp32wroomAt "espressif-goat/bsp/esp32wroom/atcmd"
	"fmt"
	"time"

	serial "github.com/hootrhino/goserial"
)

func main() {
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
	fmt.Println("AT=", GMRResponse)
	serialPort.Close()
}

```
输出：
```json
{
    "atVVersion": "AT version:3.2.0.0(s-ec2dec2 - ESP32 - Jul 28 2023 07:05:28)",
    "sdkVersion": "SDK version:v5.0.2-376-g24b9d38a24-dirty",
    "compileTime": "compile time(6118fc22):Jul 28 2023 09:47:28",
    "binVersion": "Bin version:v3.2.0.0(WROOM-32)"
}
```

注意：
- `SerialPeerRwTimeout`: 指的是系统句柄读取周期，通常和MCU的反应时间有关，50-100ms左右最佳。
- `HwCardResponseTimeout`：指的是**本次指令期望响应时间**，指令返回数据越多， 等待时间越久。取决于AT指令手册里面写的具体时间。
上面这两个参数一定要设置合理的范围。