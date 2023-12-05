package day3

import "testing"

func TestPartOne(t *testing.T) {
	const testFile = "test_input.txt"

	got, err := PartOne(testFile)
	if err != nil {
		t.Fatal(err)
	}

	want := 4361
	if got != want {
		t.Fatalf("PartOne(%q) = %q; want = %q", testFile, got, want)
	}
}
