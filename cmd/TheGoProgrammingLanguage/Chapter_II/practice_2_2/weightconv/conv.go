package weightconv

func (k Kilogram) ToLB() Pound {
	return Pound(k / 0.45359237)
}

func (p Pound) ToKG() Kilogram {
	return Kilogram(p * 0.45359237)
}
