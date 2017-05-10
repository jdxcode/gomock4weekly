package main

type WaterThermometer interface {
	MeasureAt(place string) (int, error)
}