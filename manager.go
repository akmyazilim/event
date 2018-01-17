package event

import (
	"sync"
)

type logger interface {
	Printf(format string, v ...interface{})
}

// ListType event list type
type ListType map[string][]*Event

// Manager ..
type Manager struct {
	List ListType
	sync.Mutex
	Logger logger
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
				go manager.runEvent(event)
				// go event.FN(event.Args...)
			} else {
				manager.runEvent(event)
				// event.FN(event.Args...)
			}

		}
	}
}

// RunType events
func (manager *Manager) RunType(typ string, concurrency bool) {
	if events, oke := manager.List[typ]; oke {
		for _, event := range events {
			if concurrency {
				// go event.FN(event.Args...)
				go manager.runEvent(event)
			} else {
				manager.runEvent(event)
				// event.FN(event.Args...)
			}
		}
	}
}

// Run event
func (manager *Manager) Run(name string) {
	for _, events := range manager.List {
		for _, event := range events {
			if event.Name == name {
				manager.runEvent(event)
			}
		}
	}
}
func (manager *Manager) runEvent(event *Event) {
	if manager.Logger != nil {
		manager.Logger.Printf("Event => Name: %s Type: %s called", event.Name, event.Type)
	}
	event.Run()
}

// SetLogger for debug event manager
func (manager *Manager) SetLogger(logger logger) *Manager {
	manager.Logger = logger
	return manager
}

// Event ..
type Event struct {
	FN   func(...interface{})
	Args []interface{}
	Name string
	Type string
}

// Run your self
func (event *Event) Run() {
	event.FN(event.Args...)
}

// New Manager
func New() *Manager {
	return &Manager{List: make(ListType)}
}
