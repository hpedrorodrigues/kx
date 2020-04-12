package fuzzyfinder

import "github.com/ktr0731/go-fuzzyfinder"

func Find(slice interface{}, itemFunc func(i int) string) (int, error) {
	return fuzzyfinder.Find(slice, itemFunc)
}
