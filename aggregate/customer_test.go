package aggregate_test

import (
	"errors"
	"testing"

	"github.com/landersonalmeida/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "A valid name",
			name:        "",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error [%s], got [%s]", tc.expectedErr, err)
			}
		})
	}
}
