package evmSigVerifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyEvmSingedMessage(t *testing.T) {
	signature := "0x728e6fc7d34fc8b60c96e47332735a542edd4bcc59e381c4055fad123f3eb348094a69412d99da05156a804d604e05e653c866576544881296f6b5789efc57b21c"
	address := "0xf17b474e76e6f11fbd51f5b7040ec3fa82e61ec2"
	success, err := VerifyEvmSingedMessage("RyeHarvest", signature, address)
	if err != nil {
		t.Error("this is tx 0: ", err)
	}
	assert.Equal(t, success, true)
}
