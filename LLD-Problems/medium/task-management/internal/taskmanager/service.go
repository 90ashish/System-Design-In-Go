package taskmanager

import (
	"errors"
	"sort"
	"time"
)

// TaskFilter defines a filtering strategy
// (Strategy Pattern)
type TaskFilter func(*Task) bool

// FilterByStatus returns a filter for Status
func FilterByStatus(status Status) TaskFilter {
	return func(t *Task) bool {
		return t.Status == status
	}
}

// FilterByPriority returns a filter for Priority
func FilterByPriority(p Priority) TaskFilter {
	return func(t *Task) bool {
		return t.Priority == p
	}
}

// FilterByDueBefore returns a filter for due-date
func FilterByDueBefore(deadline time.Time) TaskFilter {
	return func(t *Task) bool {
		return t.DueDate.Before(deadline)
	}
}

// TaskService provides a high-level API
// (Facade & Dependency Inversion Patterns)
type TaskService struct {
	repo TaskRepository
}

// NewTaskService injects the repository
// (Dependency Injection)
func NewTaskService(r TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

// CreateTask wraps NewTask + persistence
func (s *TaskService) CreateTask(title, desc string, pr Priority, due time.Time) *Task {
	t := NewTask("", title, desc, pr, due)
	s.repo.Create(t)
	return t
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(t *Task) error {
	return s.repo.Update(t)
}

// DeleteTask removes a task
func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// GetTasks applies any number of filters
func (s *TaskService) GetTasks(filters ...TaskFilter) ([]*Task, error) {
	all, _ := s.repo.GetAll()
	var out []*Task
	for _, t := range all {
		ok := true
		for _, f := range filters {
			if !f(t) {
				ok = false
				break
			}
		}
		if ok {
			out = append(out, t)
		}
	}
	return out, nil
}

// GetHighestPriorityTask finds the task with highest Priority
// and nearest DueDate
func (s *TaskService) GetHighestPriorityTask() (*Task, error) {
	all, _ := s.repo.GetAll()
	if len(all) == 0 {
		return nil, errors.New("no tasks")
	}
	// pick by priority then by due date
	best := all[0]
	for _, t := range all[1:] {
		if t.Priority > best.Priority ||
			(t.Priority == best.Priority && t.DueDate.Before(best.DueDate)) {
			best = t
		}
	}
	return best, nil
}

// SortTasksByDueDate returns all tasks sorted ascending
func (s *TaskService) SortTasksByDueDate() ([]*Task, error) {
	all, _ := s.repo.GetAll()
	sort.Slice(all, func(i, j int) bool {
		return all[i].DueDate.Before(all[j].DueDate)
	})
	return all, nil
}
