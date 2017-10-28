package event

import (
	"sync"
)

// ListType event list type
type ListType map[string][]*Event

// Manager ..
type Manager struct {
	List ListType
	sync.Mutex
}

// Add event
func (manager *Manager) Add(event *Event) {
	manager.Lock()
	defer manager.Unlock()
	manager.List[event.Type] = append(manager.List[event.Type], event)
}

// RunALL events
func (manager *Manager) RunALL(concurrency bool) {
	for _, events := range manager.List {
		for _, event := range events {
			if concurrency {
				go event.FN(event.Args...)
			} else {
				event.FN(event.Args...)
			}

		}
	}
}

// RunType events
func (manager *Manager) RunType(typ string, concurrency bool) {
	if events, oke := manager.List[typ]; oke {
		for _, event := range events {
			if concurrency {
				go event.FN(event.Args...)

			} else {
				event.FN(event.Args...)
			}
		}
	}
}

// Run event
func (manager *Manager) Run(name string) {
	for _, events := range manager.List {
		for _, event := range events {
			if event.Name == name {
				event.FN(event.Args...)
			}
		}
	}
}

// Event ..
type Event struct {
	FN   func(...interface{})
	Args []interface{}
	Name string
	Type string
}

// New Manager
func New() *Manager {
	return &Manager{List: make(ListType)}
}
