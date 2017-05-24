package serial

import (
	"io"
	"time"
)

const DefaultLength = 8 // Default length of the data frame.

type StopBits uint8

const (
	Stop1Bit  StopBits = 1
	Stop15Bit StopBits = 15
	Stop2Bit  StopBits = 2
)

type Parity byte

const (
	ParityNone Parity = 'N'
	ParityOdd  Parity = 'O'
	ParityEven Parity = 'E'
)

type SerialConfig struct {
	Port       string        // Name of the /dev/tty port to use.
	BaudRate   uint32        // BaudRate of the device.
	DataLength int8          // Length of the data to receive (default 8).
	Parity     Parity        // Parity of the connection.
	StopBits   StopBits      // Number of stop bits at the end of the frame.
	Timeout    time.Duration // Timeout to cancel the connection. Disabled if 0.
}

type Port io.ReadWriteCloser

func Open(sc *SerialConfig) (Port, error) {
	if sc.Parity == 0 {
		sc.Parity = ParityNone
	}
	if sc.DataLength == 0 {
		sc.DataLength = DefaultLength
	}
	if sc.StopBits == 0 {
		// sc.StopBits == Stop1Bit
	}
	return open(sc)
}
