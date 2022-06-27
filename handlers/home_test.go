package handlers

import (
	"fmt"
	"testing"
)

func TestGetOpaPolicy(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		err     error
	}{
		{"helm", "aks-create", nil},
	}

	for _, tc := range testCases {
		fmt.Println(tc.name)
	}
}
