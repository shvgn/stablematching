package main

func main() {

}

type entry struct {
	id      int
	weights []int
}

func (e entry) Next() int {
	return e.weights[0]
}

type match struct {
	proposed  int
	responded int
}

func findMatch(proposing, responding []entry) []match {
	respondingById := make(map[int]entry)
	for _, r := range responding {
		respondingById[r.id] = r
	}

	matches := make([]match, 0)
	// map[responding]proposing
	accepted := make(map[int]int)

	for {
		for _, p := range proposing {
			want := p.Next()
			prevAccepted, ok := accepted[want]
			if !ok {
				accepted[want] = p.id
				continue
			}
			if prevAccepted == p.id {
				continue
			}

			responder := respondingById[want]
			if responder.NewIsBetter(prevAccepted, p.id) {
				accepted[want] = p.id
			}
			p.Pop()
		}
	}

	return matches
}
