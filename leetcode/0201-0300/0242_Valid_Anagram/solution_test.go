package solution

import "testing"

func TestIsAnagramExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			"example_1",
			"anagram",
			"nagaram",
			true,
		},
		{
			"example_2",
			"rat",
			"car",
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isAnagram(tst.s, tst.t)
			if got != tst.want {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tst.s, tst.t, got, tst.want)
			}
		})
	}
}
