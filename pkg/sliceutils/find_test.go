package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type myType int

func TestFind(t *testing.T) {
	t.Parallel()

	slice := []myType{1, 3, 7}
	assert.Equal(t, -1, Find(slice, 3))
	assert.Equal(t, 1, Find(slice, myType(3)))
}

type myWeirdStruct struct {
	a string
	b int
}

func TestFindMatching(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	slice := []myWeirdStruct{
		{"twinkle twinkle", 1},
		{"little star", 2},
	}
	a.Equal(-1, FindMatching(slice, func(weirdStruct myWeirdStruct) bool {
		return weirdStruct.b > 2
	}))
	a.Equal(0, FindMatching(slice, func(weirdStruct myWeirdStruct) bool {
		return weirdStruct.b > 0
	}))
	a.Equal(1, FindMatching(slice, func(weirdStruct myWeirdStruct) bool {
		return weirdStruct.b > 1
	}))
	a.Equal(1, FindMatching(slice, func(weirdStruct *myWeirdStruct) bool {
		return weirdStruct.b > 1
	}))
	a.Panics(func() {
		FindMatching([]string{"1", "2"}, func(int) bool {
			return false
		})
	})
	a.Panics(func() {
		FindMatching("1", func(string) bool {
			return false
		})
	})
}
