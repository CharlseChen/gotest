package main

import "fmt"

type Cooker struct{}

func (c *Cooker) MakeChicken() {
	fmt.Println("做烤鸡")
}

func (c *Cooker) MakeBBQ() {
	fmt.Println("做烧烤")
}

type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

type CommandCookBBQ struct {
	cooker *Cooker
}

func (cmd CommandCookBBQ) Make() {
	cmd.cooker.MakeBBQ()
}

type Waiter struct {
	CmdList []Command
}

func (w *Waiter) Notify() {
	if len(w.CmdList) == 0 {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker: cooker}
	cmdBBQ := CommandCookBBQ{cooker: cooker}
	w := new(Waiter)
	w.CmdList = append(w.CmdList, cmdBBQ, cmdChicken)
	w.Notify()
}
