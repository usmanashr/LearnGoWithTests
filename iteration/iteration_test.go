package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("u", 5)
	expected := "uuuuu"
	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}
