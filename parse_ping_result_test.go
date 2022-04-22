package exec_ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strReceivedDebian = `PING google.com (127.0.0.1) 56(84) bytes of data.
64 bytes from 127.0.0.1 (127.0.0.1): icmp_seq=1 ttl=57 time=43.8 ms
64 bytes from 127.0.0.1 (127.0.0.1): icmp_seq=2 ttl=57 time=43.6 ms
64 bytes from 127.0.0.1 (127.0.0.1): icmp_seq=3 ttl=57 time=43.6 ms

--- google.com ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 43.594/43.687/43.830/0.102 ms`

var strLostMacOs = `PING test.domain.go (127.0.0.1): 56 data bytes
	64 bytes from 127.0.0.1: icmp_seq=0 ttl=48 time=39.846 ms
	64 bytes from 127.0.0.1: icmp_seq=1 ttl=48 time=58.014 ms
	^C
	--- test.domain.go ping statistics ---
	2 packets transmitted, 0 packets received, 100.0% packet loss
	round-trip min/avg/max/stddev = 39.846/48.930/58.014/9.084 ms`

func TestParseReceivedPackets(t *testing.T) {
	result, _, err := ParsePingResult(strReceivedDebian)
	assert.Nil(t, err)
	var pingResult *PingResult = &PingResult{
		Received: 2,
		Loss:     0,
	}
	assert.Equal(t, result, pingResult)
}

func TestParseLossPingPackets(t *testing.T) {
	result, _, err := ParsePingResult(strLostMacOs)
	assert.Nil(t, err)
	var pingResult *PingResult = &PingResult{
		Received: 0,
		Loss:     2,
	}
	assert.Equal(t, result, pingResult)
}
