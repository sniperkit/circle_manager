package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemEnv(t *testing.T) {
	setRunAppEnv("KEY1=val1 KEY2=val2")

	assert.Equal(t, os.Getenv("KEY1"), "val1")
	assert.Equal(t, os.Getenv("KEY2"), "val2")
}

func TestSetEnv(t *testing.T) {
	envs = &Envs{
		Mode: "add",
		Name: "",
	}
	assert.NotNil(t, envsValid())

	envs = &Envs{
		Mode: "delete",
		Name: "",
	}
	assert.NotNil(t, envsValid())

	envs = &Envs{
		Mode: "add",
		Name: "test",
	}
	assert.Nil(t, envsValid())

	envs = &Envs{
		Mode: "delete",
		Name: "test",
	}
	assert.Nil(t, envsValid())
}
