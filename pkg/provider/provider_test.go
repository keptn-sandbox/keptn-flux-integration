package provider

import (
	"fmt"
	"testing"
)

func TestGetProvider(t *testing.T) {
	testCases := []struct {
		name string
		err  error
	}{
		{"podtato-head-podtato-kustomize", nil},
	}

	for _, tc := range testCases {
		// payload := GetCloudEvent(tc.name)
		// fmt.Println(payload)
		fmt.Println(tc.name)
	}
}
