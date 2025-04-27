package models

// VehicleType enum
type VehicleType int

const (
	CarType VehicleType = iota
	MotorcycleType
	TruckType
)

func (vt VehicleType) String() string {
	switch vt {
	case CarType:
		return "Car"
	case MotorcycleType:
		return "Motorcycle"
	case TruckType:
		return "Truck"
	default:
		return "Unknown"
	}
}

// Vehicle abstracts any vehicle
type Vehicle interface {
	Type() VehicleType
	LicensePlate() string
}

// Car, Motorcycle, Truck concrete types
type Car struct{ License string }

func (c *Car) Type() VehicleType    { return CarType }
func (c *Car) LicensePlate() string { return c.License }

type Motorcycle struct{ License string }

func (m *Motorcycle) Type() VehicleType    { return MotorcycleType }
func (m *Motorcycle) LicensePlate() string { return m.License }

type Truck struct{ License string }

func (t *Truck) Type() VehicleType    { return TruckType }
func (t *Truck) LicensePlate() string { return t.License }
