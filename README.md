# Stable matching

Gale-Shapley algoritm implementation in Go. Inspired by [Numberphile video](https://www.youtube.com/watch?v=Qcv1IqHWAzg). See the [explanation on Wikipedia](https://en.wikipedia.org/wiki/Gale–Shapley_algorithm)

## Inputs

- N Proposers
- N Acceptors
- N×N table of proposers preferences
- N×N table of acceptors preferences

### Interfaces

You define the table yourself, and pass proposers and acceptors as Rankers

```go
type Ranker inteface {
        Rank(Ranker) int
}
```

## Outputs

The list of matched pairs sorted by the match rank descendetly

```go
[]Match{
        {Proposer, Acceptor}, // higest rank
        {Proposer, Acceptor},
        {Proposer, Acceptor},
        ...
}
```
