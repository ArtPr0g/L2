package pattern

/*
	Данный патерн используется для создания объектов без необходимости указывать конкретный класс объекта, который нужно создать.
*/

import "fmt"


type Transport interface {
    Drive()
}

// Структура автомобиля
type Car struct{}

func (c Car) Drive() {
    fmt.Println("Driving a car 🚗")
}

// Структура самолета
type Plane struct{}

func (p Plane) Drive() {
    fmt.Println("Flying a plane ✈️")
}

// Фабрика транспорта
type TransportFactory interface {
    CreateTransport() Transport
}

type CarFactory struct{}

func (cf CarFactory) CreateTransport() Transport {
    return Car{}
}

// Фабрика для создания самолетов
type PlaneFactory struct{}

func (pf PlaneFactory) CreateTransport() Transport {
    return Plane{}
}

func main() {
    carFactory := CarFactory{}
    car := carFactory.CreateTransport()
    car.Drive()

    planeFactory := PlaneFactory{}
    plane := planeFactory.CreateTransport()
    plane.Drive()
}
