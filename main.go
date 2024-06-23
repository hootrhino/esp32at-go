package main

import (
	esp32wroom "espressif-goat/bsp/esp32wroom"
	"fmt"
	"time"

	serial "github.com/hootrhino/goserial"
)

func main() {
	SerialPeerRwTimeout := 50 * time.Millisecond
	HwCardResponseTimeout := 300 * time.Millisecond
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
	{
		Response, errAt := Esp32.AT("AT\r\n", HwCardResponseTimeout)
		if errAt != nil {
			panic(errAt)
		}
		fmt.Println("AT=", Response)

	}
	{
		Response, errAt := Esp32.AT("AT+GMR\r\n", HwCardResponseTimeout)
		if errAt != nil {
			panic(errAt)
		}
		fmt.Println("AT+GMR=", Response)
	}
	serialPort.Close()
}
