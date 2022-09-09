package nanoid

import (
	"bytes"
	"context"
	"crypto/rand"
	"math"
	"math/bits"

	cm "github.com/morozovcookie/change-management"
)

var _ cm.IdentifierGenerator = (*IdentifierGenerator)(nil)

const (
	DefaultAlphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	DefaultSize     = 32
)

// IdentifierGenerator represents a service for generating unique identifier
// value.
type IdentifierGenerator struct {
	alphabet []rune
	size     int

	alphabetSize int
	mask         int
	step         int
}

func NewIdentifierGenerator(opts ...Option) *IdentifierGenerator {
	ig := &IdentifierGenerator{
		alphabet: []rune(DefaultAlphabet),
		size:     DefaultSize,
	}

	ig.alphabetSize = len(ig.alphabet)

	for _, opt := range opts {
		opt.apply(ig)
	}

	ig.mask = (2 << (31 - bits.LeadingZeros32((uint32(ig.alphabetSize)-1)|1))) - 1
	ig.step = int(math.Ceil((1.6 * float64(ig.mask*ig.size)) / float64(ig.alphabetSize)))

	return ig
}

// GenerateIdentifier returns a unique identifier value.
func (ig *IdentifierGenerator) GenerateIdentifier(_ context.Context) cm.ID {
	var buf bytes.Buffer

	buf.Grow(ig.size)

	bb := make([]byte, ig.step)

	for buf.Len() < ig.size {
		if _, err := rand.Read(bb); err != nil {
			panic(err)
		}

		for i := 0; i < len(bb) && buf.Len() < ig.size; i++ {
			idx := int(bb[i]) & ig.mask
			if idx >= ig.alphabetSize {
				continue
			}

			if _, err := buf.WriteRune(ig.alphabet[idx]); err != nil {
				panic(err)
			}
		}
	}

	return cm.ID(buf.String())
}
