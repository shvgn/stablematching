package stablematching

// Table is the preference table for both proposors and acceptors.
//
// For now both tables have to have the same set members:
//
//   proposors := Table{
//           "p1": {"a4", "a3", "a2", "a1"},
//           "p2": {"a4", "a3", "a2", "a1"},
//           "p3": {"a4", "a3", "a2", "a1"},
//           "p4": {"a4", "a3", "a2", "a1"},
//   }
//
//   acceptors := Table{
//           "a1": {"p1", "p2", "p3", "p4"},
//           "a2": {"p1", "p2", "p3", "p4"},
//           "a3": {"p1", "p2", "p3", "p4"},
//           "a4": {"p1", "p2", "p3", "p4"},
//   }
//
//   matcher := NewMatcher(proposors, acceptors)
//   matches := matcher.Match() // map[p1:a4 p2:a3 p3:a2 p4:a1]
type Table map[string][]string

// NewMatcher creates a Matcher struct
func NewMatcher(proposorRanks, acceptorRanks Table) *Matcher {
	return &Matcher{
		pairs:     make(map[string]string),
		proposors: newProposors(proposorRanks),
		acceptors: newAcceptors(acceptorRanks),
		free:      make(chan string),
	}
}

// Matcher matches string values according to preferences of proposors ad acceptors (both are inner
// types).
type Matcher struct {
	// pairs map name-by-name from both sides (p -> a and a -> p)
	pairs map[string]string

	// maps by name
	proposors map[string]*proposor
	acceptors map[string]*acceptor

	// the free proposers queue
	free chan string
}

// Match calculates the mathes by Gale-Shapley algorithm
func (m *Matcher) Match() map[string]string {
	if len(m.proposors) == 0 {
		return make(map[string]string)
	}

	go func() {
		for pname := range m.proposors {
			m.free <- pname
		}
	}()

	for pname := range m.free {
		p := m.proposors[pname]
		a := m.acceptors[p.next()]

		m.propose(pname, a.name)

		if len(m.pairs)/2 == len(m.proposors) {
			close(m.free)
			break
		}
	}

	matches := make(map[string]string)
	for pname := range m.proposors {
		aname := m.pairs[pname]
		matches[pname] = aname
	}

	return matches
}

func (m *Matcher) propose(pname, aname string) {
	curName, ok := m.pairs[aname]
	if !ok {
		// No pair
		m.pair(pname, aname)
		return
	}

	a := m.acceptors[aname]
	if a.firstIsPreferred(pname, curName) {
		m.pair(pname, aname)
		m.enqueue(curName)
	} else {
		m.enqueue(pname)
	}
}

func (m *Matcher) enqueue(pname string) {
	go func() {
		m.free <- pname
	}()
}

// pair creates a pair for proposer and acceptor
func (m *Matcher) pair(pname, aname string) {
	prevname, ok := m.pairs[aname]
	if ok {
		delete(m.pairs, aname)
		delete(m.pairs, prevname)
	}
	m.pairs[pname] = aname
	m.pairs[aname] = pname
}

func newProposors(table Table) map[string]*proposor {
	ps := make(map[string]*proposor)
	for name, ranks := range table {
		ps[name] = newProposor(name, ranks)
	}
	return ps
}

func newProposor(name string, toSee []string) *proposor {
	return &proposor{
		name:  name,
		seen:  -1,
		toSee: toSee,
	}
}

type proposor struct {
	name  string
	seen  int
	toSee []string
}

func (r *proposor) next() string {
	r.seen++
	return r.toSee[r.seen]
}

func newAcceptors(table Table) map[string]*acceptor {
	as := make(map[string]*acceptor)
	for name, ranks := range table {
		as[name] = newAcceptor(name, ranks)
	}
	return as
}

func newAcceptor(name string, ranks []string) *acceptor {
	mapified := make(map[string]int)
	for i, pname := range ranks {
		mapified[pname] = i
	}
	return &acceptor{
		name:  name,
		ranks: mapified,
	}
}

type acceptor struct {
	name  string
	ranks map[string]int
}

func (a *acceptor) firstIsPreferred(first, second string) bool {
	rank1, ok := a.ranks[first]
	if !ok {
		return false
	}

	rank2, ok := a.ranks[second]
	if !ok {
		return true
	}

	return rank1 < rank2
}
