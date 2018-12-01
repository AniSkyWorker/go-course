package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		fizzBuzzRange int
		want          string
	}{
		{2, "1, 2"},
		{4, "1, 2, Fizz, 4"},
		{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz"},
	}
	for _, testCase := range cases {
		got := FizzBuzz(testCase.fizzBuzzRange)
		if got != testCase.want {
			t.Errorf("FizzBuzz(%v) = %q, want %q", testCase.fizzBuzzRange, got, testCase.want)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	FizzBuzz(b.N)
}

func BenchmarkFizzBuzzBuffer(b *testing.B) {
	FizzBuzzBuffer(b.N)
}
