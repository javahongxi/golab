package pipeline

import (
	"fmt"
	"testing"
)

func TestNodes(t *testing.T) {
	p := Merge(
		InMemSort(ArraySource(3, 2, 6, 7, 4)),
		InMemSort(ArraySource(7, 4, 0, 3, 2, 13, 8)),
	)
	for v := range p {
		fmt.Println(v)
	}
}
