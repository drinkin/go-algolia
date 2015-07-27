package algolia_test

import (
	"os"
	"testing"
)

const (
	TestIndexName = "go_test"
)

func TestMain(m *testing.M) {
	// Setup

	os.Exit(m.Run())
}
