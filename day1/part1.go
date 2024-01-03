package main

type FirstSolution struct{}

func (f1 *FirstSolution) Extract(chunk string) (int, bool) {
	value := chunk[0]
	if value > '0' && value <= '9' {
		return int(value) - '0', true
	}

	return 0, false
}
