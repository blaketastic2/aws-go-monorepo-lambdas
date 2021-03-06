package main

import (
	"testing"

	atl "github.com/yogeshlonkar/aws-lambda-go-test/local"
)

func TestMain(t *testing.T) {
	response, err := atl.Run(atl.Input{})
	if err != nil {
		panic(err)
	}
	if string(response) != "\"yogesh\"" {
		t.Errorf("response=%s, want \"yogesh\"", response)
	}
}
