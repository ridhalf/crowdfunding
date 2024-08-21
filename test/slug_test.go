package test

import (
	"fmt"
	"github.com/gosimple/slug"
	"strconv"
	"testing"
)

func TestSlug(t *testing.T) {
	s := slug.Make("Foo Faa" + "-" + strconv.Itoa(1))
	fmt.Println(s)
}
