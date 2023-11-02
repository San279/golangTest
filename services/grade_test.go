package services

import (
	"fmt"
	"testing"
)

// go test -v
func TestCheckGrade(t *testing.T) {

	type testList struct {
		testCasetag string
		input       int
		expected    string
	}

	cases := []testList{
		{testCasetag: "a", input: 80, expected: "A"},
		{testCasetag: "b", input: 70, expected: "B"},

		{testCasetag: "c", input: 60, expected: "C"},
		{testCasetag: "d", input: 50, expected: "D"},
		{testCasetag: "f", input: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.testCasetag, func(t *testing.T) {
			output := CheckGrade(c.input)

			if output != c.expected {
				t.Errorf("got %v expected %v", output, c.expected)
			}
		})
	}
}

//go test -bench=.

func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CheckGrade((i + 5) * 10)
	}
}

// documents
// godoc -http=:6060
func ExampleCheckGrade() {
	grade := CheckGrade(60)
	fmt.Println("document grade checking")
	fmt.Println(grade)
}
