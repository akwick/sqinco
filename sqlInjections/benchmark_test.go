package sqlInjections

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSafeSQL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// don't assert that err is nil because the program returns something even if the result is correct
		out, _ := exec.Command("safesql", "github.com/akwick/sqinco").CombinedOutput()
		assert.Contains(b, string(out), "Found 8 potentially unsafe SQL statements", "")
	}
}

func BenchmarkGas(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// don't assert that err is nil because the program returns something even if the result is correct
		out, _ := exec.Command("gas", "-include=G201,G202", "../...").CombinedOutput()
		assert.Contains(b, string(out), "Issues: 9", "")
	}
}

func BenchmarkGotcha(b *testing.B) {
	for n := 0; n < b.N; n++ {
		out, err := exec.Command("gotcha", "-src=../main.go", "-ssf=../sqlQuery.txt", "-pkgs=github.com/akwick/sqinco/sqlInjections").CombinedOutput()
		assert.Nil(b, err)
		assert.Contains(b, string(out), "err.NumberOfFlows: 1 ", "")
	}
}
