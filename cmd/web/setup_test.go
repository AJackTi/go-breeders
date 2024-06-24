package main

import (
	"os"
	"testing"

	"go-breeders/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
