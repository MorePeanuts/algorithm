package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"pos case", "anagram", "nagaram", true},
		{"neg case", "rat", "car", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			actual := isAnagram(tst.s, tst.t)
			if actual != tst.expected {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tst.s, tst.t, actual, tst.expected)
			}
		})
	}
}
