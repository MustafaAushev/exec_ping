package exec_ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecPing(t *testing.T) {
	pingResult, errRuning, err := Run("google.com", 3, 5)
	assert.Nil(t, err)
	assert.Nil(t, errRuning)
	var pr *PingResult = &PingResult{
		Received: 3,
		Loss:     0,
	}
	assert.Equal(t, pingResult, pr)
}
