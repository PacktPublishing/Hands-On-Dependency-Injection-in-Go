package fixedv1

func Go(vehicle actions) {
	switch concrete := vehicle.(type) {
	case poweredActions:
		concrete.startEngine()

	case unpoweredActions:
		concrete.pushStart()
	}

	vehicle.drive()
}

type actions interface {
	drive()
}

type poweredActions interface {
	actions
	startEngine()
	stopEngine()
}

type unpoweredActions interface {
	actions
	pushStart()
}

type Vehicle struct {
}

func (v Vehicle) drive() {
	// TODO: implement
}

type PoweredVehicle struct {
	Vehicle
}

func (v PoweredVehicle) startEngine() {
	// common engine start code
}

type Car struct {
	PoweredVehicle
}

type Sled struct {
	Vehicle
}

func (s Sled) pushStart() {
	// do nothing
}
