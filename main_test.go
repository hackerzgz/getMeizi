package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if e := TestOnly(); e != nil {
		t.Error(e)
	}
	t.Log("Passed")
}
