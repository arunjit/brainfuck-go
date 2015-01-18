// brainfuck_test
package brainfuck_test

import (
	"testing"

	"github.com/arunjit/brainfuck-go"
)

func TestEvalToString(t *testing.T) {
	input := `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---
.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`
	output := brainfuck.EvalToString(input)
	if output != "Hello World!" {
		t.Errorf("Unexpected output: <<%s>>\n", output)
	}
}
