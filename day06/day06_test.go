package day06

import (
	"github.com/its-felix/AdventOfCode2024/inputs"
	"testing"
)

const input = "dayxx.txt"

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines(input)))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines(input)))
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart1(inputs.GetInputLines(input)) != 0 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart2(inputs.GetInputLines(input)) != 0 {
			b.FailNow()
		}
	}
}
