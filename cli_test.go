package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	setRunAppEnv("KEY1=val1 KEY2=val2")

	assert.Equal(t, os.Getenv("KEY1"), "val1")
	assert.Equal(t, os.Getenv("KEY2"), "val2")
}
