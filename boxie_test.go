package boxie

import (
	"strconv"
	"testing"

	"fmt"
)

func TestBasic(t *testing.T) {
	b := New()
	for i := 0; i < 54; i++ {
		b.Append(strconv.Itoa(i))
	}

	b.Insert(51, "w00t")
	b.Insert(2, "scoot")
	b.Prepend("beginning")

	b.ForEach(func(i int, val interface{}) (end bool) {
		fmt.Println(i, val)
		return
	})
}
