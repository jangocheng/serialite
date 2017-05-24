package serial

import (
	"golang.org/x/sys/unix"
	"os"
	"sys"
)

type termios2 struct {
	Test bool
}

func confTermios2(sc *SerialConfig) (*termios2, error) {
	t2 := new(termios2)
	return t2, nil
}

func open(sc *SerialConfig) (Port, error) {
	f, err := os.OpenFile(
		sc.Port,
		os.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK,
		0600)
	if err != nil {
		return nil, err
	}

	err = unix.SetNonblock(int(f.Fd()), false)
	if err != nil {
		f.Close()
		return nil, err
	}

	t2, err := confTermios2(sc)
	if err != nil {
		f.Close()
		return nil, err
	}

	r, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		file.Fd(),
		unix.TCSETS2,
		t2)

	return f, nil
}
