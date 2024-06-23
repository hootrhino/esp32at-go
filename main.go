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
