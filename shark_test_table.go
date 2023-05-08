package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test all possible errors from the Hunt method of the Shark struct.
func TestShark_Hunt_Errors(t *testing.T) {
	// Define a test table struct to test all possible errors from the Hunt method

	// arrange
	type errorTests struct {
		name        string
		shark       *Shark
		prey        *Prey
		expectedErr error
	}

	// Define the test table
	tests := []errorTests{
		{
			name: "shark is tired",
			shark: &Shark{
				hungry: false,
				tired:  true,
				speed:  20,
			},
			prey: &Prey{
				name:  "fish",
				speed: 10,
			},
			expectedErr: ErrTired,
		},
		{
			name: "shark is not hungry",
			shark: &Shark{
				hungry: false,
				tired:  false,
				speed:  20,
			},
			prey: &Prey{
				name:  "",
				speed: 0,
			},
			expectedErr: ErrNotHungry,
		},
		{
			name: "shark cannot reach the prey",
			shark: &Shark{
				hungry: true,
				tired:  false,
				speed:  10,
			},
			prey: &Prey{
				name:  "fish",
				speed: 20,
			},
			expectedErr: ErrPreyEscaped,
		},
		{
			name: "prey is nil",
			shark: &Shark{
				hungry: true,
				tired:  false,
				speed:  20,
			},
			prey:        nil,
			expectedErr: ErrNotPrey,
		},
	}

	// Loop over the test table
	for _, test := range tests {
		// act
		err := test.shark.Hunt(test.prey)

		// assert
		
		// Check that there is an error from the Hunt method
		if assert.Error(t, err, "the Hunt method should return an error") {
			// Check that the error message is correct
			assert.ErrorIs(t, err, test.expectedErr, "the error message is not correct")
		}
	}
}