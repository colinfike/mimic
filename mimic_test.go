package mimic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrainSingleValid(t *testing.T) {
	testSet := []string{"hello there"}

	markov := NewMarkovChain(2)
	markov.Train(testSet)

	expectedMap := map[string]map[string]struct{}{
		"  ": map[string]struct{}{
			"hello": struct{}{},
		},
		" hello": map[string]struct{}{
			"there": struct{}{},
		},
		"hello there": map[string]struct{}{
			"": struct{}{},
		},
	}
	assert.Equal(t, markov.chain, expectedMap)
}

func TestTrainTwoValid(t *testing.T) {
	testSet := []string{"hello there", "hello there jon"}

	markov := NewMarkovChain(2)
	markov.Train(testSet)

	expectedMap := map[string]map[string]struct{}{
		"  ": map[string]struct{}{
			"hello": struct{}{},
		},
		" hello": map[string]struct{}{
			"there": struct{}{},
		},
		"hello there": map[string]struct{}{
			"":    struct{}{},
			"jon": struct{}{},
		},
		"there jon": map[string]struct{}{
			"": struct{}{},
		},
	}
	assert.Equal(t, markov.chain, expectedMap)
}

func TestTrainThreeValid(t *testing.T) {
	testSet := []string{"hello there", "hello there jon", "Run there jon quickly!", "ignore"}

	markov := NewMarkovChain(2)
	markov.Train(testSet)

	expectedMap := map[string]map[string]struct{}{
		"  ": map[string]struct{}{
			"hello": struct{}{},
			"run":   struct{}{},
		},
		" run": map[string]struct{}{
			"there": struct{}{},
		},

		" hello": map[string]struct{}{
			"there": struct{}{},
		},
		"run there": map[string]struct{}{
			"jon": struct{}{},
		},
		"hello there": map[string]struct{}{
			"":    struct{}{},
			"jon": struct{}{},
		},
		"there jon": map[string]struct{}{
			"quickly!": struct{}{},
			"":         struct{}{},
		},
		"jon quickly!": map[string]struct{}{
			"": struct{}{},
		},
	}
	assert.Equal(t, markov.chain, expectedMap)
}

func TestGenerate(t *testing.T) {
	markov := NewMarkovChain(2)
	testSet := []string{"Hello there jon"}
	markov.Train(testSet)

	assert.Equal(t, "hello there jon", markov.Generate())
}

var nextSuffixTests = []struct {
	prefix         string
	suffix         string
	expectedPrefix string
}{
	{"  ", "nonempty", " nonempty"},
	{" prefix", "suffix", "prefix suffix"},
	{"ignored prefix", "suffix", "prefix suffix"},
}

func TestNextSuffix(t *testing.T) {
	for _, testCase := range nextSuffixTests {
		newPrefix := nextPrefix(testCase.prefix, testCase.suffix)
		assert.Equal(t, testCase.expectedPrefix, newPrefix)
	}
}
