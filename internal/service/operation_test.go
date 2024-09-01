package service

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 1)
	if result != 2 {
		t.Error("Fail")
	}
}
