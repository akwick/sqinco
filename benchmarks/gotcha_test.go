package benchmarks

import (
	"os/exec"
	"testing"
)

func BenchmarkSafeSQL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = exec.Command("safesql", "github.com/akwick/gotcha")
	}
}

func BenchmarkGas(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = exec.Command("gas", "-include=G201,G202", "../../gotcha/...")
	}
}

func BenchmarkGotcha(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = exec.Command("gotcha", "-src=../../gotcha/start.go", "-ssf=../sqlQuery.txt", "-pkgs=github.com/akwick/gotcha/lattice/taint,github.com/akwick/gotcha/worklist")
	}
}
