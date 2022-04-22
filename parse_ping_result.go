package exec_ping

import (
	"regexp"
	"strconv"
)

func ParsePingResult(str string) (*PingResult, error) {
	transmitted, received, err := ParsePackets(str)
	if err != nil {
		return nil, err
	}
	pingResult := PingResult{
		Received: received,
		Loss:     transmitted - received,
	}
	return &pingResult, nil
}

func ParsePackets(str string) (int16, int16, error) {
	regPackets := regexp.MustCompile(`\d+ packets transmitted, \d+ packets received`)
	packets := regPackets.FindString(str)
	regPacketsNumber := regexp.MustCompile(`\d+`)
	packetsNumber := regPacketsNumber.FindAllString(packets, 2)
	transmitted, err := strconv.Atoi(packetsNumber[0])
	if err != nil {
		return 0, 0, err
	}
	received, err := strconv.Atoi(packetsNumber[1])
	if err != nil {
		return 0, 0, err
	}

	return int16(transmitted), int16(received), nil
}
