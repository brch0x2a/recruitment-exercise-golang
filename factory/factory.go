package factory

import (
	"fmt"

	".main.go/assemblyspot"
	".main.go/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	vehicleList := f.generateVehicleLots(amountOfVehicles)

	for _, vehicle := range vehicleList {
		fmt.Println("Assembling vehicle...")



		select {

			case idleSpot := <-f.AssemblingSpots:
				idleSpot.SetVehicle(&vehicle)

				vehicle, err := idleSpot.AssembleVehicle()
		
				if err != nil {
					continue
				}
		
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot.GetAssembledLogs()
		
				idleSpot.SetVehicle(nil)
		
				f.AssemblingSpots <- idleSpot

			case idleSpot2 := <-f.AssemblingSpots:
				idleSpot2.SetVehicle(&vehicle)

				vehicle, err := idleSpot2.AssembleVehicle()
		
				if err != nil {
					continue
				}
		
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot2.GetAssembledLogs()
		
				idleSpot2.SetVehicle(nil)
		
				f.AssemblingSpots <- idleSpot2

			case idleSpot3 := <-f.AssemblingSpots:
				idleSpot3.SetVehicle(&vehicle)

				vehicle, err := idleSpot3.AssembleVehicle()
		
				if err != nil {
					continue
				}
		
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot3.GetAssembledLogs()
		
				idleSpot3.SetVehicle(nil)
		
				f.AssemblingSpots <- idleSpot3

			case idleSpot4 := <-f.AssemblingSpots:
				idleSpot4.SetVehicle(&vehicle)

				vehicle, err := idleSpot4.AssembleVehicle()
		
				if err != nil {
					continue
				}
		
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot4.GetAssembledLogs()
		
				idleSpot4.SetVehicle(nil)
		
				f.AssemblingSpots <- idleSpot4

			case idleSpot5 := <-f.AssemblingSpots:
				idleSpot5.SetVehicle(&vehicle)

				vehicle, err := idleSpot5.AssembleVehicle()
		
				if err != nil {
					continue
				}
		
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot5.GetAssembledLogs()
		
				idleSpot5.SetVehicle(nil)
		
				f.AssemblingSpots <- idleSpot5
		default:
			continue
	}




	}
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	var vehicles = []vehicle.Car{}
	var index = 0

	for {
		vehicles = append(vehicles, vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
