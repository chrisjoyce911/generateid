package generateid

import (
	"testing"
)

func Test_randomchar(t *testing.T) {

	actualResult := randomchar(4, "AAAAA")
	expectedResult := "AAAA"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}

}

func Test_randomchar_len(t *testing.T) {

	actualResult := randomchar(4, "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789")
	expectedResult := "AAAA"

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected %d but got %d", len(expectedResult), len(actualResult))
	}

}

func TestGenerateID(t *testing.T) {
	actualResult := GenerateID()
	expectedResult := "9AGZ-HQ4H-16TR-9HDF"

	if len(actualResult) != len(expectedResult) {
		t.Fatalf("Expected %d but got %d", len(expectedResult), len(actualResult))
	}
}
