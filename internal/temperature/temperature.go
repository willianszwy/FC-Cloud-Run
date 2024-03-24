package temperature

type Temperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

func New(celsius, fahrenheit float64) *Temperature {
	return &Temperature{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     celsius + 273,
	}
}
