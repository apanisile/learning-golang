package iteration

import "testing"

func Testing (t *testing.T){
	repeated := Repeat("a")
	expected  := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, &repeated)
	}
}

func Repeat(character string) string {
	return ""
}