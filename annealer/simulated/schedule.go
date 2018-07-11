package simulated

type TemperatureSchedule struct {
	initialTemperature float64
	lastTemperature    float64
	scale              float64
	currentTemperature float64
}

func NewTemperatureSchedule(scale, initialTemperature, lastTemperature float64) *TemperatureSchedule {
	if scale < 0 {
		scale = 0.95
	}
	if initialTemperature < 0 {
		initialTemperature = 5.0
	}
	if lastTemperature < 0 {
		lastTemperature = 0.2
	}

	return &TemperatureSchedule{
		initialTemperature: initialTemperature,
		lastTemperature:    lastTemperature,
		scale:              scale,
		currentTemperature: initialTemperature,
	}
}

func (t *TemperatureSchedule) Reset() {
	t.currentTemperature = t.initialTemperature
}

func (t *TemperatureSchedule) Next() float64 {
	if t.currentTemperature <= t.lastTemperature {
		return -1
	}
	preCurrentTemperature := t.currentTemperature
	t.currentTemperature *= t.scale
	return preCurrentTemperature
}
