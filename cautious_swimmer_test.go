package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const okTemperature = 26

func TestCautiousSwimmer(t *testing.T) {
	test := func(temperature int, state string) {
		name := fmt.Sprintf("%d˚: %s", temperature, state)
		t.Run(name, func(t *testing.T) {
			place := "bath"
			thermometer := &WaterThermometerMock{}
			thermometer.On("MeasureAt", place).Return(temperature, nil)

			subject := NewCautiousSwimmer(place, okTemperature, thermometer)
			err := subject.Swim()

			if err != nil {
				panic(err)
			}
			assert.Equal(t, subject.State(), state)
		})
	}

	test(30, StateSwimmin)
	test(26, StateSwimmin)
	test(20, StateChillin)
}

func TestBrokenThermometer(t *testing.T) {
	place := "bath"
	mockErr := errors.New("Too lazy to measure")
	thermometer := &WaterThermometerMock{}
	thermometer.On("MeasureAt", place).Return(0, mockErr)

	subject := NewCautiousSwimmer(place, okTemperature, thermometer)
	err := subject.Swim()

	assert.Equal(t, err, mockErr)
	assert.Equal(t, subject.State(), StateUnknown)
}
