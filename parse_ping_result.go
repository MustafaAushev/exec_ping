package exec_ping

import (
	"regexp"
	"strconv"
)

func ParsePingResult(str string) (*PingResult, string, error) {
	transmitted, received, errorStr, err := ParsePackets(str)
	if err != nil {
		return nil, errorStr, err
	}
	pingResult := PingResult{
		Received: received,
		Loss:     transmitted - received,
	}
	return &pingResult, errorStr, nil
}

func ParsePackets(str string) (int16, int16, string, error) {
	regPackets := regexp.MustCompile(`\d+[\s]*\w*[\s]*transmitted,[\s]*\d+[\s]*\w*[\s]*received`)
	packets := regPackets.FindString(str)
	if packets == "" {
		return 0, 0, "Not parsing ping output", nil
	}
	regPacketsNumber := regexp.MustCompile(`\d+`)
	packetsNumber := regPacketsNumber.FindAllString(packets, 2)
	transmitted, err := strconv.Atoi(packetsNumber[0])
	if err != nil {
		return 0, 0, "", err
	}
	received, err := strconv.Atoi(packetsNumber[1])
	if err != nil {
		return 0, 0, "", err
	}

	return int16(transmitted), int16(received), "", nil
}
