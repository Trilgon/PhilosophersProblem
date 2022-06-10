package phils

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name string
}

func (p Philosopher) Live(leftFork *sync.Mutex, rightFork *sync.Mutex) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Философ совсем изголодался - ", rec)
			return
		}
	}()
	var thinking *time.Timer = nil
	isOutTime := false
	for {
		if isOutTime {
			panic("Паника, философ думает больше одной секунды")
		}
		if thinking == nil {
			thinking = time.AfterFunc(60000*time.Millisecond, func() {
				isOutTime = true
			})
		}
		if leftFork.TryLock() {
			if rightFork.TryLock() {
				thinking.Stop()
				thinking = nil
				fmt.Printf("Философ: %s ест\n", p.Name)
				time.Sleep(200 * time.Millisecond)
				leftFork.Unlock()
				rightFork.Unlock()
				fmt.Printf("Философ: %s спит\n", p.Name)
				time.Sleep(200 * time.Millisecond)
				continue
			}
			leftFork.Unlock()
		}
		fmt.Printf("Философ: %s думает\n", p.Name)
	}
}
