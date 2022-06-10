package phils

import (
	"fmt"
	"sync"
)

type LunchTable struct {
	Table []interface{}
}

func NewLunchTable(philsNumber int) LunchTable {
	philsNumber *= 2
	table := make([]interface{}, philsNumber)
	for i := 0; i < philsNumber; i += 2 {
		table[i] = new(sync.Mutex)
		// нужно для поиска "места за столом"
		table[i+1] = true
	}
	return LunchTable{table}
}

// AddPhilosopher возвращает индекс добавленного элемента
func (l *LunchTable) AddPhilosopher(philosopher Philosopher) (int, error) {
	for i, element := range l.Table {
		if element == true {
			l.Table[i] = philosopher
			return i, nil
		}
	}
	return 0, fmt.Errorf("место для философа не найдено, возможно, не был вызван конструктор")
}
