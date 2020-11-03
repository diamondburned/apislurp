package testutil

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
)

func ReadFile(t *testing.T, file string) []byte {
	path := filepath.Join("..", "discord-api-docs", "docs", file)

	f, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("failed to read md file:", err)
	}

	return f
}

func Equal(t *testing.T, v1, v2 interface{}, desc string) {
	if eq := deep.Equal(v1, v2); eq != nil {
		t.Errorf("Returned %s document is different: %v", desc, eq)
	}
}
