package main

import (
	"fmt"
	"sync"
	"time"
)

// DatabaseConnection represents a mock database connection
type DatabaseConnection struct {
	ID int
}

// ConnectionPoolManager manages database connections
type ConnectionPoolManager struct {
	pool       chan *DatabaseConnection
	maxConn    int
	activeConn int
	mutex      sync.Mutex
}

// Singleton instance and sync object
var (
	instance *ConnectionPoolManager
	once     sync.Once
)

// GetInstance - Provides the singleton instance of the ConnectionPoolManager
func GetInstance(maxConnections int) *ConnectionPoolManager {
	once.Do(func() {
		instance = &ConnectionPoolManager{
			pool:    make(chan *DatabaseConnection, maxConnections),
			maxConn: maxConnections,
		}
	})
	return instance
}

// GetConnection - Provides a database connection
func (cpm *ConnectionPoolManager) GetConnection() (*DatabaseConnection, error) {
	cpm.mutex.Lock()
	defer cpm.mutex.Unlock()

	// Check if max limit is reached
	if cpm.activeConn >= cpm.maxConn {
		return nil, fmt.Errorf("maximum connections reached, try again later")
	}

	// Create a new connection if pool is empty
	select {
	case conn := <-cpm.pool:
		return conn, nil
	default:
		cpm.activeConn++
		newConn := &DatabaseConnection{ID: cpm.activeConn}
		fmt.Printf("[INFO] New Connection Created: %d\n", newConn.ID)
		return newConn, nil
	}
}

// ReleaseConnection - Releases a database connection back to the pool
func (cpm *ConnectionPoolManager) ReleaseConnection(conn *DatabaseConnection) {
	cpm.mutex.Lock()
	defer cpm.mutex.Unlock()

	select {
	case cpm.pool <- conn:
		fmt.Printf("[INFO] Connection %d released back to pool\n", conn.ID)
	default:
		// Close the connection if the pool is full
		fmt.Printf("[INFO] Connection %d closed (Pool Full)\n", conn.ID)
		cpm.activeConn--
	}
}

// DisplayStatus - Shows the current pool status
func (cpm *ConnectionPoolManager) DisplayStatus() {
	fmt.Printf("[INFO] Active Connections: %d | Pool Size: %d\n", cpm.activeConn, len(cpm.pool))
}

func main() {
	const maxConnections = 3
	poolManager := GetInstance(maxConnections)

	// Simulate concurrent requests
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := poolManager.GetConnection()
			if err != nil {
				fmt.Printf("[ERROR] Goroutine %d: %v\n", id, err)
				return
			}

			// Simulate database work
			fmt.Printf("[INFO] Goroutine %d acquired connection %d\n", id, conn.ID)
			time.Sleep(2 * time.Second)

			// Release the connection
			poolManager.ReleaseConnection(conn)
			poolManager.DisplayStatus()
		}(i + 1)
	}

	wg.Wait()
	fmt.Println("[INFO] All Goroutines Finished")
}
