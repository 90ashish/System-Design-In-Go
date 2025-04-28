# Parking Lot System

## Overview

A thread-safe, multi-level parking lot management system written in Go. Supports Cars, Motorcycles, and Trucks; assigns and releases spots; tracks real-time availability; and handles concurrent entry/exit across multiple gates.

---

## Table of Contents

1. [Problem Statement](#problem-statement)  
2. [Requirements](#requirements)  
3. [Core Components](#core-components)  
4. [Design Considerations](#design-considerations)   
5. [How to Run](#how-to-run)  
6. [Questions](#questions)  

---

## Problem Statement

Design a parking lot system that:

- Has multiple levels, each with a fixed number of spots per vehicle type  
- Supports different vehicle types (Car, Motorcycle, Truck)  
- Assigns a spot on entry and releases it on exit  
- Tracks spot availability in real time  
- Handles concurrent access from multiple entry/exit points  

---

## Requirements

1. **Levels & Spots**  
   - Multiple levels, each with N spots for each vehicle type  
2. **Vehicle Types**  
   - Car, Motorcycle, Truck  
3. **Spot Allocation**  
   - Each spot only accepts matching vehicle type  
4. **Ticketing**  
   - Generate a unique ticket on park; use it to unpark  
5. **Availability Tracking**  
   - Real-time counts of free spots per level/type  
6. **Concurrency**  
   - Safe under concurrent park/unpark calls  

---

## Core Components

- **ParkingLot** (singleton)  
  - Manages all levels, dispatches park/unpark requests  
- **Level**  
  - Contains a collection of ParkingSpots and availability counters  
- **ParkingSpot**  
  - Knows its type, occupancy status, and parked vehicle  
- **Vehicle** (interface)  
  - Abstracts common behavior of Car, Motorcycle, Truck  
- **Ticket**  
  - Records spot ID, level, vehicle info, entry time, and a unique ID  

---

## Design Considerations

- **Thread Safety**  
  - `sync.Once` for singleton initialization  
  - `sync.Mutex` in each spot, level, and the lot for mutual exclusion  
  - `atomic` counter for lock-free ticket ID generation  
- **Patterns Used**  
  - **Singleton** for `ParkingLot`  
  - **Factory** (optional) for vehicle instantiation  
  - **Observer** (optional) for notifying availability changes  
- **Scalability**  
  - Modular package layout under `internal/`  
  - Easily add new vehicle types or spot allocation strategies  
- **Extensibility**  
  - Plug in different spot-selection algorithms (e.g., nearest, random)  
  - Integrate external monitoring or alerting  

---


## How to Run

1. Clone the repository  
2. `cd project-root`  
3. `go build ./cmd/app`  
4. `./app`  

---

## Questions

1. **What are the functional vs. non-functional requirements of this system?**  
2. **How would you model levels, spots, and vehicles as Go types?**  
3. **How can you ensure thread safety when multiple goroutines park or unpark simultaneously?**  
4. **How does the Singleton pattern guarantee only one `ParkingLot` instance?**  
5. **What strategy generates unique, collision-free ticket IDs?**  
6. **How is real-time availability tracked and exposed to clients?**  
7. **How would you handle multiple entry/exit gates in the design?**  
8. **Which design patterns (beyond Singleton) could improve extensibility?**  
9. **How would you structure the codebase to support future features (e.g., reservations, dynamic pricing)?**  
10. **What trade-offs exist between different spot-selection algorithms (first-available, nearest, random)?**  

---
