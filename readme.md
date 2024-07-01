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
# RHILEX AT 操作库
为了方便RHILEX用户使用串口指令集，RHILEX团队封装好了一些常见模块的AT指令库，方便集成使用。
## 支持模块
### ESP32-Wroom

ESP32 AT固件是一种专为ESP32开发板设计的固件，它允许开发者通过AT指令集来控制ESP32模块。这种固件支持Wi-Fi和蓝牙低功耗（BLE）操作，基本应用可以满足大多数需求。ESP32 AT固件包含了一系列特定功能的二进制文件，如启动加载器、AT应用固件、出厂配置参数等，这些文件共同工作以提供一个完整的解决方案，使ESP32能够作为一个独立的通信节点或与其他设备通信。

官方手册：
- https://docs.espressif.com/projects/esp-at/en/latest/esp32/AT_Command_Set/BLE_AT_Commands.html

### MX-01
MX-01 蓝牙模组是一款支持低功耗蓝牙协议的串口透传模组；模组具有小体积、高性能、高性价
比、低功耗、平台兼容性强等优点；可以帮助用户快速掌握蓝牙技术，加速产品开发；模组已兼容的
软件平台包括：IOS 应用程序、 Android 应用程序、微信小程序等。MCU 通过串口连接模组，可与手
机、平板等设备进行数据通讯，轻松实现智能无线控制和数据采集；模组广泛应用在智能家居、共享
售货机等领域。

这款模块售价只有2元，很适合平时测试使用，因此在本仓库专门支持了一个库。

官方手册：
- `doc` 路径下mx-01.pdf。

## 使用
下面是个简单的例子：

```go
package main

import (
	esp32wroom "github.com/hootrhino/rhilex-goat/bsp/esp32wroom"
	esp32wroomAt "github.com/hootrhino/rhilex-goat/bsp/esp32wroom/atcmd"
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