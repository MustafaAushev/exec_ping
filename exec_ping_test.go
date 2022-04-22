package exec_ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecPingSuccess(t *testing.T) {
	pingResult, err := Run("google.com", "-c", "3")
	assert.Nil(t, err)
	var pr *PingResult = &PingResult{
		Received: 3,
		Loss:     0,
	}
	assert.Equal(t, pingResult, pr)
}

func TestExecPingFailed(t *testing.T) {
	_, err := Run("use-case.ru", "-c", "3")
	assert.NotNil(t, err)
}
