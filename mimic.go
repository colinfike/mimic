package mimic

import (
	"math/rand"
	"strings"
	"time"
)

const delim string = " "

// MarkovChain stores the markov chain and is the main interface into the package.
type MarkovChain struct {
	chain map[string]map[string]struct{}
}

// NewMarkovChain returns a new MarkovChain
func NewMarkovChain() *MarkovChain {
	return &MarkovChain{chain: make(map[string]map[string]struct{})}
}

// Train consumes the trainingText and updates the Markov chain accordingly
func (m *MarkovChain) Train(trainingText []string) {
	var prefix, suffix string

	for _, sentence := range trainingText {
		prefix = strings.Repeat(delim, 2)
		sentence = strings.ToLower(sentence) + delim
		words := strings.Split(sentence, delim)

		for _, suffix = range words {
			if suffixMap, ok := m.chain[prefix]; ok {
				if _, ok := suffixMap[suffix]; !ok {
					suffixMap[suffix] = struct{}{}
				}
			} else {
				m.chain[prefix] = map[string]struct{}{suffix: struct{}{}}
			}
			prefix = nextPrefix(prefix, suffix)
		}

	}
}

// Generate returns a generated sentence from the Markov chain
func (m *MarkovChain) Generate() string {
	var sentence, prefix string
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix = strings.Repeat(delim, 2)
	for {
		suffixMap := m.chain[prefix]
		var keys []string
		for key := range suffixMap {
			keys = append(keys, key)
		}

		suffix := keys[rand.Intn(len(keys))]
		if len(suffix) == 0 {
			break
		}

		sentence = sentence + delim + suffix
		prefix = nextPrefix(prefix, suffix)
	}
	return sentence
}

func nextPrefix(prefix, suffix string) string {
	var newPrefix string
	splitPre := strings.Split(prefix, delim)
	if len(splitPre[1]) == 0 {
		newPrefix = delim + suffix
	} else {
		newPrefix = splitPre[1] + delim + suffix
	}
	return newPrefix
}
