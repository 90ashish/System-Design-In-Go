package main

import (
	"fmt"
	"time"

	"task-system/internal/taskmanager"
)

func main() {
	// Setup: Repository (Repository Pattern) + Service (Facade Pattern)
	repo := taskmanager.NewInMemoryTaskRepository()
	svc := taskmanager.NewTaskService(repo)

	// Create some tasks
	t1 := svc.CreateTask("Write report", "Annual financial report", taskmanager.High, time.Now().Add(48*time.Hour))
	_ = svc.CreateTask("Team meeting", "Project status update", taskmanager.Medium, time.Now().Add(24*time.Hour))
	t3 := svc.CreateTask("Pay bills", "Electricity & Internet", taskmanager.Low, time.Now().Add(72*time.Hour))

	// Update a task
	t1.Status = taskmanager.InProgress
	svc.UpdateTask(t1)

	// Fetch all OPEN tasks (Strategy Pattern for filtering)
	openTasks, _ := svc.GetTasks(taskmanager.FilterByStatus(taskmanager.Open))
	fmt.Println("Open tasks:")
	for _, t := range openTasks {
		fmt.Printf(" • %s [%s]\n", t.Title, t.Status)
	}

	// Get highest priority task
	top, _ := svc.GetHighestPriorityTask()
	fmt.Printf("\nTop priority task: %s (due %s)\n", top.Title, top.DueDate.Format(time.RFC822))

	// List tasks sorted by due date
	sorted, _ := svc.SortTasksByDueDate()
	fmt.Println("\nTasks sorted by due date:")
	for _, t := range sorted {
		fmt.Printf(" • %s – due %s\n", t.Title, t.DueDate.Format(time.RFC822))
	}

	// Delete a task
	svc.DeleteTask(t3.ID)
}
