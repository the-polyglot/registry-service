package member

// import "github.com/xariez/registry-service/models"

type Nominee struct {
	//models.ModelBase
	name         string
	memberId     string
	idCardNumber string
	address      string
	phoneNumber  string
	percentage   float32
	relationship uint
}

func NewNominee(name, memberId, idCardNumber, address, phoneNumber string, percentage float32, relationship uint) (*Nominee, error) {
	n := &Nominee{name, memberId, idCardNumber, address, phoneNumber, percentage, relationship}
	return n, nil
}
