package main

import (
	"errors"
	"fmt"
)

// -----------------------------
// Domain Model
// -----------------------------
type User struct {
	ID   int
	Name string
	Age  int
}

// -----------------------------
// Abstraction (Repository Interface)
// -----------------------------
type UserRepository interface {
	Create(u User) (int, error)
	GetByID(id int) (User, error)
	Update(u User) error
	Delete(id int) error
}

// -----------------------------
// High-Level Module (Service)
// -----------------------------
type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) RegisterUser(name string, age int) (int, error) {
	user := User{ID: 0, Name: name, Age: age}
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(id int, name string, age int) error {
	user := User{ID: id, Name: name, Age: age}
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

// -----------------------------
// Low-Level Module: In-Memory Repository
// -----------------------------
type InMemoryUserRepo struct {
	data   map[int]User
	nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data:   make(map[int]User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepo) Create(u User) (int, error) {
	u.ID = r.nextID
	r.data[u.ID] = u
	r.nextID++
	return u.ID, nil
}

func (r *InMemoryUserRepo) GetByID(id int) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepo) Update(u User) error {
	if _, ok := r.data[u.ID]; !ok {
		return errors.New("user not found")
	}
	r.data[u.ID] = u
	return nil
}

func (r *InMemoryUserRepo) Delete(id int) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.data, id)
	return nil
}

// -----------------------------
// (Future) Low-Level Module: SQL Repository
// -----------------------------
// type SQLUserRepo struct { /* db connection */ }
// func (r *SQLUserRepo) Create(u User) (int, error)    { /* ... */ }
// func (r *SQLUserRepo) GetByID(id int) (User, error) { /* ... */ }
// func (r *SQLUserRepo) Update(u User) error          { /* ... */ }
// func (r *SQLUserRepo) Delete(id int) error          { /* ... */ }

// -----------------------------
// Main: Demonstration
// -----------------------------
func main() {
	// Wire up In-Memory repository
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)

	// Create
	id, _ := service.RegisterUser("Alice", 30)
	fmt.Println("Created User with ID:", id)

	// Read
	user, _ := service.GetUser(id)
	fmt.Printf("Fetched User: %+v\n", user)

	// Update
	_ = service.UpdateUser(id, "Alice Smith", 31)
	updated, _ := service.GetUser(id)
	fmt.Printf("Updated User: %+v\n", updated)

	// Delete
	_ = service.DeleteUser(id)
	_, err := service.GetUser(id)
	fmt.Println("After delete, fetch error:", err)
}
