# Stable matching

Gale-Shapley algoritm implementation in Go. Inspired by [Numberphile video](https://www.youtube.com/watch?v=Qcv1IqHWAzg). See the [explanation on Wikipedia](https://en.wikipedia.org/wiki/Gale–Shapley_algorithm)

## Inputs

- N×N table of proposors preferences
- N×N table of acceptors preferences

The mather relies on the same number of proposors and acceptors.

```go
proposors := Table{
        "p1": {"a4", "a3", "a2", "a1"},
        "p2": {"a4", "a3", "a2", "a1"},
        "p3": {"a4", "a3", "a2", "a1"},
        "p4": {"a4", "a3", "a2", "a1"},
}

acceptors := Table{
        "a1": {"p1", "p2", "p3", "p4"},
        "a2": {"p1", "p2", "p3", "p4"},
        "a3": {"p1", "p2", "p3", "p4"},
        "a4": {"p1", "p2", "p3", "p4"},
}

matcher := NewMatcher(proposors, acceptors)
matches := matcher.Match() // map[p1:a4 p2:a3 p3:a2 p4:a1]
```

## License

MIT
