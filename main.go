package main

import (
	"philosophers/phils"
	"sync"
)

func main() {
	lunchTable := phils.NewLunchTable(6)
	// создаём срез индексов всех философов, чтобы потом не искать их в основном
	philsIndexes := make([]int, 6)
	var err error
	if philsIndexes[0], err = lunchTable.AddPhilosopher(phils.Philosopher{"Аристотель"}); err != nil {
		panic(err)
	}
	if philsIndexes[1], err = lunchTable.AddPhilosopher(phils.Philosopher{"Платон"}); err != nil {
		panic(err)
	}
	if philsIndexes[2], err = lunchTable.AddPhilosopher(phils.Philosopher{"Сократ"}); err != nil {
		panic(err)
	}
	if philsIndexes[3], err = lunchTable.AddPhilosopher(phils.Philosopher{"Демокрит"}); err != nil {
		panic(err)
	}
	if philsIndexes[4], err = lunchTable.AddPhilosopher(phils.Philosopher{"Маркс"}); err != nil {
		panic(err)
	}
	if philsIndexes[5], err = lunchTable.AddPhilosopher(phils.Philosopher{"Кант"}); err != nil {
		panic(err)
	}

	for _, index := range philsIndexes {
		go run(index, &lunchTable)
	}
	for {
	}
}

func run(index int, lunchTable *phils.LunchTable) {
	var leftFork *sync.Mutex
	var rightFork *sync.Mutex
	switch val := lunchTable.Table[index-1].(type) {
	case *sync.Mutex:
		leftFork = val
	}
	if index+1 < len(lunchTable.Table) {
		switch val := lunchTable.Table[index+1].(type) {
		case *sync.Mutex:
			rightFork = val
		}
	} else {
		switch val := lunchTable.Table[0].(type) {
		case *sync.Mutex:
			rightFork = val
		}
	}
	lunchTable.Table[index].(phils.Philosopher).Live(leftFork, rightFork)
}
