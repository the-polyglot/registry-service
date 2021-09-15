package member

type nextOfKin struct {
	name         string
	idCardNumber string
	address      string
	relationship uint8
}

func NewNextOfKin(name, idCardNumber, address string, relationship uint8) nextOfKin {
	return nextOfKin{name, idCardNumber, address, relationship}
}
