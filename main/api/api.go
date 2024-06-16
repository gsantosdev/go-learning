package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{"Ferrari", 100})
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
}

func main() {
	createCars()
	e := echo.New()
	e.GET("/cards", getCars)
	e.Logger.Fatal(e.Start(":8080"))
}


func getCars (c echo.Context) error{
	return c.JSON(200, cars)
}
