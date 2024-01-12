package main

import "fmt"

func mainpanic() {

	fmt.Println("Theory Panic 1")
	fmt.Println("Starting...")

	or := &Orchestrator{
		handlers: map[string]func(){
			"handlerOne":   handlerOne,
			"handlerTwo":   handlerTwo,
			"handlerThree": handlerThree,
		},
	}

	or.RunHandler("handlerOne")
	or.RunHandler("handlerTwo")
	or.RunHandler("handlerThree")
	or.RunHandler("handlerTwo")

	fmt.Println("Ending...")

}

type Orchestrator struct {
	handlers map[string]func()
}

func (o *Orchestrator) RunHandler(name string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	hd, ok := o.handlers[name]
	if !ok {
		return
	}
	hd()

}

func handlerOne() {
	fmt.Println("Handler one")
}

func handlerTwo() {
	fmt.Println("Handler two")
}

func handlerThree() {
	panic("Handler three failure")
}
