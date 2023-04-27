package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:generate go run asm.go -out bytes.s -stubs bytes.go

func TestSimpleIndex(t *testing.T) {
	for _, tt := range []struct {
		needle   []byte
		haystack []byte 
		index    int64 
	}{
		{
			[]byte{4, 1, 3},
			[]byte{
				0, 0, 0, 4, 1, 3, 0, 0,
				0, 4, 1, 3, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			int64(3),
		},
		{
			[]byte(`amet`),
			[]byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit integer.`),
			int64(22),
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("`%s` in `%s`", tt.needle, tt.haystack), func(t *testing.T) {
			i := Index(tt.haystack, tt.needle)
			require.Equal(t, tt.index, i)
		})
	}
}

func TestMask(t *testing.T) {
	for _, tt := range []struct {
		name     string
		needle   []byte
		haystack []byte 
		index    int64 
	}{
		{
			"chunk second match",
			[]byte{4, 1, 3},
			[]byte{
				0, 0, 0, 4, 2, 3, 0, 0,
				0, 4, 1, 3, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			int64(9),
		},
		{
			"chunk first match",
			[]byte{4, 1, 3},
			[]byte{
				0, 0, 0, 4, 1, 3, 0, 0,
				0, 4, 1, 3, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			int64(3),
		},
		{
			"longer chunk",
			[]byte{4, 1, 3, 3},
			[]byte{
				0, 0, 0, 4, 1, 3, 0, 0,
				0, 4, 1, 3, 3, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			int64(9),
		},
		{
			"text chunk",
			[]byte(`amet`),
			[]byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit integer.`),
			int64(22),
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf(tt.name), func(t *testing.T) {
			index := findInChunk(tt.needle, tt.haystack)
			require.Equal(t, tt.index, index)
			if index != -1 {
				end := index + int64(len(tt.needle))
				require.ElementsMatch(t, tt.needle, tt.haystack[index:end])
			}
		})
	}
}
