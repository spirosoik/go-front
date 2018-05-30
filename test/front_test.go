package test

import (
	"testing"

	"github.com/spirosoik/go-front"
	"github.com/stretchr/testify/assert"
)

func TestNewNilCfg(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(nil)

	assert.Nil(c)
	assert.NotNil(err)
}

func TestNewNilBaseUrl(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(&front.Config{
		APIToken: "test_token",
	})

	assert.Nil(c)
	assert.NotNil(err)
}

func TestNewNilApiToken(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(&front.Config{
		BaseURL: "http://test.com",
	})

	assert.Nil(c)
	assert.NotNil(err)
}
func TestNewEmptyConfig(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(&front.Config{})

	assert.Nil(c)
	assert.NotNil(err)
}
func TestNewBadUrl(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(&front.Config{
		BaseURL: "file://test.txt",
	})

	assert.Nil(c)
	assert.NotNil(err)
}

func TestNewSuccess(t *testing.T) {
	assert := assert.New(t)

	c, err := front.New(&front.Config{
		BaseURL:  "http://test.com",
		APIToken: "test_token",
	})

	assert.Nil(err)
	assert.NotNil(c)
}
