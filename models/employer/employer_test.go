package employer

import (
	"fmt"
	"testing"
)

func TestEmployer_NewEmployer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: fmt.Errorf("name is required"),
		},
		{
			test:        "Valid description",
			name:        "Teknokraft Limited",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewEmployer(tc.name)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

// func TestEmployer_AddAddress(t *testing.T) {
// 	type testCase struct {
// 		test 	   string
// 		expectedErr error
// 	}

// 	ne, _  := NewEmployer("Teknokraft Limited")
// 	ne.SetAddress("103577", "Nairobi", "Woodvale Groove", "00101", "info@tkl.co.ke")

// 	testCases := []testCase {
// 		{
// 			test: "Address provided",
// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tc := range testCases{
// 		t.Run(tc.test, func(t *testing.T) {


// 			_, err := 

// 			if err != tc.expectedErr {
// 				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
// 			}
// 		})
// 	}
// }