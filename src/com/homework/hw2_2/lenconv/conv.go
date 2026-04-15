package lenconv

func FToM(l Foot) Meter {
	return Meter(l * 0.3048)
}

func MToF(l Meter) Foot {
	return Foot(l * 3.28084)
}
