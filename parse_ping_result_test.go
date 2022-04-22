package exec_ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strReceived = `PING test.domain.go (127.0.0.1): 56 data bytes
	64 bytes from 127.0.0.1: icmp_seq=0 ttl=48 time=39.846 ms
	64 bytes from 127.0.0.1: icmp_seq=1 ttl=48 time=58.014 ms
	^C
	--- test.domain.go ping statistics ---
	2 packets transmitted, 2 packets received, 0.0% packet loss
	round-trip min/avg/max/stddev = 39.846/48.930/58.014/9.084 ms`

var strLost = `PING test.domain.go (127.0.0.1): 56 data bytes
	64 bytes from 127.0.0.1: icmp_seq=0 ttl=48 time=39.846 ms
	64 bytes from 127.0.0.1: icmp_seq=1 ttl=48 time=58.014 ms
	^C
	--- test.domain.go ping statistics ---
	2 packets transmitted, 0 packets received, 100.0% packet loss
	round-trip min/avg/max/stddev = 39.846/48.930/58.014/9.084 ms`

func TestParseReceivedPackets(t *testing.T) {
	result, err := ParsePingResult(strReceived)
	assert.Nil(t, err)
	var pingResult *PingResult = &PingResult{
		Received: 2,
		Loss:     0,
	}
	assert.Equal(t, result, pingResult)
}

func TestParseLossPingPackets(t *testing.T) {
	result, err := ParsePingResult(strLost)
	assert.Nil(t, err)
	var pingResult *PingResult = &PingResult{
		Received: 0,
		Loss:     2,
	}
	assert.Equal(t, result, pingResult)
}
