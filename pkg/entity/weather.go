package entity

type Weather struct {
	TempC float64
}

func (w *Weather) TempF() float64 {
	return w.TempC*1.8 + 32
}

func (w *Weather) TempK() float64 {
	return w.TempC + 273
}
