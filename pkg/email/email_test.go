package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) { // nolint
	err := SendEmail(`
	<p>your email code is: </p>
	` + "12345")
	assert.Nil(t, err)
}
