package main

import (
	"fmt"
	"my-first-go-project/domain"
	"time"
)

func main() {
	a := domain.NewAccount("123", "zxie", 10000, time.Now())
	a.SetName("mwang")
	fmt.Println(a.GetAccountName())
}
