package snippet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringToMap(t *testing.T) {
	str := string("transfer=ComTransfer;com=/dev/ttys1;baudrate=9600;timeout=1000;addr=01;parity=None;databits=8;stopbits=1;flow_control=None")
	m := StringToMap(str)
	assert.Equal(t, m["transfer"], "ComTransfer")
	assert.Equal(t, m["com"], "/dev/ttys1")
	assert.NotEqual(t, m["stopbits"], "2")
}