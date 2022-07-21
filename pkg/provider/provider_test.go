package provider

import (
	"fmt"
	"testing"
)

func TestGetProvider(t *testing.T) {
	testCases := []struct {
		name      string
		namespace string
		err       error
	}{
		{"podtato-head-podtato-kustomize", "podtato-kustomize", nil},
	}

	for _, tc := range testCases {
		//payload := GetCloudEvent(tc.name, tc.namespace)
		//fmt.Println(payload)
		fmt.Printf("%s%s", tc.name, tc.namespace)
	}
}
