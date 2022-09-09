package nanoid

// Option used to configure IdentifierGenerator object.
type Option interface {
	apply(ig *IdentifierGenerator)
}

var _ Option = (*optionFunc)(nil)

type optionFunc func(ig *IdentifierGenerator)

func (fn optionFunc) apply(ig *IdentifierGenerator) {
	fn(ig)
}

// WithAlphabet sets up the alphabet used for identifier characters.
func WithAlphabet(alphabet string) Option {
	return optionFunc(func(ig *IdentifierGenerator) {
		alphabetSize := len(alphabet)
		if alphabetSize < 2 || alphabetSize > 255 {
			return
		}

		ig.alphabet, ig.alphabetSize = []rune(alphabet), alphabetSize
	})
}

// WithSize sets up the length of identifier.
func WithSize(size int) Option {
	return optionFunc(func(ig *IdentifierGenerator) {
		if size <= 0 {
			return
		}

		ig.size = size
	})
}
