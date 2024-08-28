package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSlug(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
