package main_test

import (
	"fmt"
	"testing"

	. "github.com/nikolausreza131192/tdd/basic"
	"github.com/stretchr/testify/assert"
)

func TestPrintProfile(t *testing.T) {
	tcs := []struct {
		name           string
		profileName    string
		profileAge     int
		profileAddress string
		expectedResult string
	}{
		{
			name:           "Success; Only profile name is fulfilled",
			profileName:    "Wewe",
			expectedResult: fmt.Sprintf("Profile:\nName: Wewe"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			result := PrintProfile(tc.profileName, tc.profileAge, tc.profileAddress)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
