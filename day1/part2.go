package main

type SecondSolution struct{}

func (f1 *SecondSolution) Extract(chunk string) (int, bool) {
	first := FirstSolution{}

	if v, ok := first.Extract(chunk); ok {
		return v, ok
	}

	if StartWith(chunk, "one") {
		return 1, true
	}

	if StartWith(chunk, "two") {
		return 2, true
	}

	if StartWith(chunk, "three") {
		return 3, true
	}

	if StartWith(chunk, "four") {
		return 4, true
	}

	if StartWith(chunk, "five") {
		return 5, true
	}

	if StartWith(chunk, "six") {
		return 6, true
	}

	if StartWith(chunk, "seven") {
		return 7, true
	}

	if StartWith(chunk, "eight") {
		return 8, true
	}

	if StartWith(chunk, "nine") {
		return 9, true
	}

	return 0, false
}

func StartWith(target, chunk string) bool {
	if len(target) < len(chunk) {
		return false
	}

	for i, c := range chunk {
		if c != rune(target[i]) {
			return false
		}
	}

	return true
}
