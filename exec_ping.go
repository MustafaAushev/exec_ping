package exec_ping

import (
	"bytes"
	"fmt"
	"os/exec"
)

type PingResult struct {
	Received int16
	Loss     int16
}

func Run(host string, attempt int16, timeoutSeconds int16) (*PingResult, error) {
	cmd := exec.Command("ping", host, "-c", fmt.Sprint(attempt), "-t", fmt.Sprint(timeoutSeconds))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	pingResult, err := ParsePingResult(out.String())
	if err != nil {
		return pingResult, err
	}
	return pingResult, nil
}