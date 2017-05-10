package main

const (
	StateSwimmin = "Quack-quack"
	StateChillin = "Go swim yourself!"
	StateUnknown = ""
)

type CautiousSwimmer struct {
	thermometer   WaterThermometer
	okTemperature int
	place         string
	state         string
}

func NewCautiousSwimmer(place string, okTemperature int, thermometer WaterThermometer) *CautiousSwimmer {
	return &CautiousSwimmer{
		place:         place,
		okTemperature: okTemperature,
		thermometer:   thermometer,
		state:         StateUnknown,
	}
}

func (s *CautiousSwimmer) Swim() error {
	t, err := s.tryOutWater()

	if err != nil {
		return err
	}

	if s.swimmable(t) {
		s.swim()
	} else {
		s.chill()
	}

	return nil
}

func (s *CautiousSwimmer) State() string {
	return s.state
}

func (s *CautiousSwimmer) tryOutWater() (int, error) {
	return s.thermometer.MeasureAt(s.place)
}

func (s *CautiousSwimmer) swimmable(t int) bool {
	return t >= s.okTemperature
}

func (s *CautiousSwimmer) swim() {
	s.state = StateSwimmin
}

func (s *CautiousSwimmer) chill() {
	s.state = StateChillin
}
