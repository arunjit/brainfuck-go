// Package brainfuck is a Brainfuck interpreter.
// Currently does not support input.
package brainfuck

// EvalToString evaluates brainfuck code to outputs to string
func EvalToString(input string) string {
	array := make([]rune, 256)
	ptr := 0

	skip := false
	loopIndex := -1
	loops := []int{-1, -1, -1, -1, -1, -1, -1, -1}

	var out []rune

	for i := 0; i < len(input); i++ {
		if skip {
			continue
		}
		ch := input[i]
		switch ch {
		case '>':
			ptr++ // TODO: overflow
		case '<':
			ptr-- // TODO: underflow
		case '+':
			array[ptr]++
		case '-':
			array[ptr]--
		case '.':
			out = append(out, array[ptr])
		case ',':
			// Do nothing
		case '[':
			if array[ptr] != 0 {
				if loopIndex == -1 || loops[loopIndex] != i {
					loopIndex++
					loops[loopIndex] = i
				}
			} else {
				skip = true
			}
		case ']':
			skip = false
			if array[ptr] != 0 {
				i = loops[loopIndex] - 1
			} else {
				loops[loopIndex] = -1
				loopIndex--
			}
		}
	}

	return string(out)
}
