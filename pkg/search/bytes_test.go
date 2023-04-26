package search

import (
	"fmt"
	"testing"
	"strconv"

	"github.com/stretchr/testify/require"
)

//go:generate go run asm.go -out bytes.s -stubs bytes.go

func TestSimpleSearch(t *testing.T) {
	for _, tt := range []struct {
		haystack []byte 
		needle   []byte
		match    bool
	}{
		//{[]byte(`foobar`), []byte(`foobaz`), false},
		//{[]byte(`foobar`), []byte(`foobar`), true},
		//{[]byte(`foo`), []byte(`foobar`), false},
		//{[]byte(`foobar`), []byte(`foo`), false},
		//{[]byte(`a cat tries`), []byte(`cat`), true},
		{
			[]byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit integer.`),
			[]byte(`amet`),
			true,
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("`%s` in `%s`", tt.needle, tt.haystack), func(t *testing.T) {
			r := Search(tt.haystack, tt.needle)
			require.Equal(t, tt.match, r)
		})
	}
}

func TestMask(t *testing.T) {
	first := byte(4)
	array := [32]byte{
		0, 4, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	mask := Mask(first, array[:])
	require.Equal(t, "10", strconv.FormatInt(int64(mask), 2))
}
