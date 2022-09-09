package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/morozovcookie/change-management/nanoid"
)

func main() {
	var (
		alphabet string
		size     int
	)

	flag.StringVar(&alphabet, "alphabet", nanoid.DefaultAlphabet, "alphabet used for identifier characters")
	flag.IntVar(&size, "size", nanoid.DefaultSize, "length of identifier")

	id := nanoid.NewIdentifierGenerator(nanoid.WithAlphabet(alphabet), nanoid.WithSize(size)).
		GenerateIdentifier(context.Background())

	_, _ = fmt.Fprintf(os.Stdout, "%s\n", id)
}
