package serial

import (
	"testing"
)

func TestSerial(t *testing.T) {
	var sc SerialConfig()
	f, err := Open(&sc)
	if err != nil {
		t.Error("Failed to open")
		return
	}
	defer f.Close()
}
