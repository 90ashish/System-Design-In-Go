package taskmanager

import (
	"errors"
	"fmt"
	"sync"
)

// TaskRepository defines CRUD operations
// (Repository Pattern)
type TaskRepository interface {
	Create(task *Task) (*Task, error)
	Update(task *Task) error
	Delete(id string) error
	GetAll() ([]*Task, error)
	GetByID(id string) (*Task, error)
}

// InMemoryTaskRepository is a thread-safe in-memory store
type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  map[string]*Task
	nextID int
}

// NewInMemoryTaskRepository constructs the repository
// (Factory Pattern)
func NewInMemoryTaskRepository() TaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[string]*Task),
		nextID: 1,
	}
}

// Create assigns a new ID and stores the task
func (r *InMemoryTaskRepository) Create(t *Task) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("%d", r.nextID)
	r.nextID++
	t.ID = id
	r.tasks[id] = t
	return t, nil
}

// Update replaces an existing task
func (r *InMemoryTaskRepository) Update(t *Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[t.ID]; !ok {
		return errors.New("task not found")
	}
	r.tasks[t.ID] = t
	return nil
}

// Delete removes a task by ID
func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

// GetAll returns all tasks
func (r *InMemoryTaskRepository) GetAll() ([]*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		list = append(list, t)
	}
	return list, nil
}

// GetByID retrieves a task by ID
func (r *InMemoryTaskRepository) GetByID(id string) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return t, nil
}
