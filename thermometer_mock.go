package main

import "github.com/stretchr/testify/mock"

type WaterThermometerMock struct {
	mock.Mock
}

func (t *WaterThermometerMock) MeasureAt(place string) (int, error) {
	args := t.Called(place)
	return args.Int(0), args.Error(1)
}
