package exec_ping

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type PingResult struct {
	Received int16
	Loss     int16
}

func Run(host string, attempt int16, timeoutSeconds int16) (*PingResult, error, error) {
	cmd := exec.Command("ping", host, "-c", fmt.Sprint(attempt), "-t", fmt.Sprint(timeoutSeconds))
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	errRuning := cmd.Run()
	outString := out.String()
	outErrorString := outErr.String()
	fmt.Print(outString, outErrorString)
	if outString == "" {
		return nil, errRuning, errors.New("Empty output ping")
	}
	pingResult, err := ParsePingResult(outString)
	fmt.Print(pingResult.Received, pingResult.Loss, "\n")
	if err != nil {
		return pingResult, errRuning, err
	}
	return pingResult, errRuning, nil
}
