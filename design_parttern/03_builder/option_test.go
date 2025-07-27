package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_option_builder(t *testing.T) {
	client, err := NewRedisClient(
		"134.5.4.3",
		8978,
		WithDb(10),
		WithPassword("abc"),
		WithRegion("hd"),
	)

	assert.Nil(t, err)
	assert.Equal(t, client.addr, "134.5.4.3")
	assert.Equal(t, client.port, 8978)
	assert.Equal(t, client.db, 10)
	assert.Equal(t, client.password, "abc")
	assert.Equal(t, client.region, "hd")
}
