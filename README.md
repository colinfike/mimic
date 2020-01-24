# Mimic
[![Documentation](https://godoc.org/github.com/colinfike/mimic?status.svg)](https://godoc.org/github.com/colinfike/mimic)

This package exposes an easy to use interface to train Markov chains based on sentences you provide. You can generate sentences just as easily.

## Getting Started

Install Mimic

```bash
go get github.com/colinfike/mimic
```

Generate a new MarkovChain object.

```Go
markov := mimic.NewMarkovChain()
```

Train it with sentences.

```Go
markov.Train([]string{"Hello there Obi Wan", "Then he said hello there jon", "Obi Wan he said you wouldn't make it"})
```

Generate a sentence.

```Go
for i := 0; i < 10; i++ {
    fmt.Println(markov.Generate())
}

// then he said hello there obi wan he said hello there obi wan
// then he said you wouldn't make it
// obi wan he said hello there jon
// obi wan he said you wouldn't make it
// hello there jon
// hello there obi wan he said hello there jon
// obi wan he said you wouldn't make it
// obi wan he said hello there jon
// hello there jon
// hello there jon
 ```

## Limitations

Currently this is in a limited v1.0 state. The only processing done on the input is downcasing everything. There could be undefined behavior with certain inputs. There is no weighting on how frequently a suffix occurs after a specific prefix, they are all weighted equal at the moment when generating a sentence. Currently delimited by spaces, not customizable at this time.
