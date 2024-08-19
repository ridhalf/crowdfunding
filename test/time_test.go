package test

import (
	"crowdfunding/helper"
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	loc, errTime := time.LoadLocation("Asia/Jakarta")
	helper.PanicIfError(errTime)
	now := (time.Now().In(loc))
	nowFormat := now.Format("2006-01-02 15:04:05")
	fmt.Println(nowFormat)

}
