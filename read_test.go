package cuttergo

import (
	"testing"
)

// TestReadRunes tests the ReadRunes function
func TestReadRunes(testing *testing.T) {
	dict, arr, err := ReadRunes("./dict.txt")
	if err != nil {
		testing.Errorf("Error: %v", err)
	}

	if len(dict) != 181263 {
		testing.Errorf("Error: dict length = %v, not 181265", len(dict))
	}

	if (dict["再见"] + 6.971266636694509) > 0.000001 {
		testing.Errorf("Error: dict[再见] = %v, not -6.971266636694509", dict["再见"])
	}

	if len(arr) != 181263 {
		testing.Errorf("Error: arr length = %v, not 181265", len(arr))
	}

}
