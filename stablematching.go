package stablematching

import (
	"math"
)

type Match struct {
	Proposer, Acceptor string
}

type ranker struct {
	name  string
	seen  int
	toSee []string
}

func (r *ranker) next() string {
	r.seen++
	return r.toSee[r.seen]
}

func newRanker(name string, toSee []string) *ranker {
	return &ranker{
		name:  name,
		seen:  -1,
		toSee: toSee,
	}
}

func newRankers(table map[string][]string) []*ranker {
	rs := make([]*ranker, len(table))
	for name, ranks := range table {
		rs = append(rs, newRanker(name, ranks))
	}
	return rs
}

// Rank returs the rank, the lesser the better
func (r *ranker) Rank(other ranker) int {
	for i, name := range r.toSee {
		if name == other.name {
			return i
		}
	}
	return math.MaxInt
}

func (r *ranker) FirstIsPreferred(first, second string) bool {
	rank1, rank2 := -1, -1
	for i, name := range r.toSee {
		if rank1 > -1 && rank2 > -1 {
			break
		}
		if name == first {
			rank1 = i
			continue
		}
		if name == second {
			rank2 = i
			continue
		}
	}
	return rank1 < rank2
}

type Matcher struct {
	// pairs map name-by-name from both sides (p -> a and a -> p)
	pairs map[string]string

	// maps by name
	proposers map[string]*ranker
	acceptors map[string]*ranker

	// the free proposers queue
	free chan string
}

func (m *Matcher) Match() []Match {
	go func() {
		for pname := range m.proposers {
			m.free <- pname
		}
	}()

	for pname := range m.free {
		p := m.proposers[pname]
		a := m.acceptors[p.next()]

		m.Propose(pname, a.name)

		if len(m.pairs)/2 == len(m.proposers) {
			close(m.free)
			break
		}
	}

	matches := make([]Match, 0, len(m.proposers))
	for pname := range m.proposers {
		aname := m.pairs[pname]
		matches = append(matches, Match{Proposer: pname, Acceptor: aname})
	}

	return matches
}

func (m *Matcher) Propose(pname, aname string) {
	curName, ok := m.pairs[aname]
	if !ok {
		// No pair
		m.Pair(pname, aname)
		return
	}

	a := m.acceptors[aname]
	if a.FirstIsPreferred(pname, curName) {
		m.Pair(pname, aname)
		m.Enqueue(curName)
	} else {
		m.Enqueue(pname)
	}
}

func (m *Matcher) Enqueue(pname string) {
	go func() {
		m.free <- pname
	}()
}

// func (m *Matcher) HasPair(pname string) bool {
// 	_, ok := m.pairs[pname]
// 	return ok
// }

// Pair creates a pair for proposer and acceptor
func (m *Matcher) Pair(pname, aname string) {
	prevname, ok := m.pairs[aname]
	if ok {
		delete(m.pairs, aname)
		delete(m.pairs, prevname)
	}
	m.pairs[pname] = aname
	m.pairs[aname] = pname
}
