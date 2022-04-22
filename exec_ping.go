package exec_ping

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var debug = os.Getenv("DEBUG")

type PingResult struct {
	Received int16
	Loss     int16
}

var EmptyPingResult = &PingResult{Received: 0, Loss: 0}

func getErrorStr(stdout string, strerr string, anyStr string) string {
	return fmt.Sprintf("stdout: %s; stderr: %s; more: %s", stdout, strerr, anyStr)
}

func Run(host string, args ...string) (*PingResult, error) {
	args = append(args, host)
	cmd := exec.Command("ping", args...)
	var stdout bytes.Buffer
	var strerr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &strerr
	errRuning := cmd.Run()
	outString := stdout.String()
	outErrorString := strerr.String()
	if outString == "" {
		if debug == "true" {
			fmt.Println(getErrorStr(outString, outErrorString, ""))
		}
		return EmptyPingResult, errRuning
	}
	pingResult, errorStr, err := ParsePingResult(outString)
	if debug == "true" {
		fmt.Println(getErrorStr(outString, outErrorString, errorStr))
	}
	if err != nil {
		return pingResult, errRuning
	}
	return pingResult, errRuning
}
