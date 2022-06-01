package main

import (
	"fmt"
	"math"
	"time"
)

func Do(attempts int) time.Duration {
	if attempts > 13 {
		return 2 * time.Minute
	}
	return time.Duration(math.Pow(float64(attempts), math.E)) * time.Millisecond * 100
}

type Human interface {
	Say(s string) error
	Eat(s string) error
	Walk(s string) error
}

type TestA string

func (t TestA) Say(s string) error {
	fmt.Printf("Human Say: %s\n", s)
	return nil
}

func (t TestA) Eat(s string) error {
	fmt.Printf("Human Eat: %s\n", s)
	return nil
}

func (t TestA) Walk(s string) error {
	fmt.Printf("Human Walk: %s\n", s)
	return nil
}

func main() {
	var h Human
	var t TestA
	h = t
	_ = h.Eat("烤羊排")
	_ = h.Say("炸鸡翅")
	_ = h.Walk("去炸鸡翅")
}
