package member

import (
	"time"

	"github.com/xariez/registry-service/models"
	"github.com/xariez/registry-service/models/member/enums"
)

type Member struct {
	//models.ModelBase
	surname       string
	otherName     string
	payrollNumber string
	idCardNumber  string
	pin           string
	salutation    enums.Salutation
	gender        enums.Gender
	maritalStatus enums.MaritalStatus
	dateOfBirth   time.Time
	registeredAt  time.Time
	address       models.Address
	phoneInfo     models.PhoneInfo
	EmploymentInfo
	nextOfKin
	nominees []*Nominee
}

func NewMember(surname, otherName, payrollNumber, idCardNumber, pin string,
	salutation enums.Salutation, gender enums.Gender, maritalStatus enums.MaritalStatus,
	dateOfBirth time.Time, registeredAt time.Time) *Member {
	return &Member{
		surname:        surname,
		otherName:      otherName,
		payrollNumber:  payrollNumber,
		idCardNumber:   idCardNumber,
		pin:            pin,
		salutation:     salutation,
		gender:         gender,
		maritalStatus:  maritalStatus,
		dateOfBirth:    dateOfBirth,
		registeredAt:   registeredAt,
		address:        models.Address{},
		phoneInfo:      models.PhoneInfo{},
		EmploymentInfo: EmploymentInfo{},
		nextOfKin:      nextOfKin{},
		nominees:       make([]*Nominee, 0),
	}
}

// func (m *Member) register() (*Member, error) {
// 	newMember := &Member{}
// }

// func (m *Member) update() (*Member, error) {

// }

// func (m *Member) addNextOfKin(name, idCardNumber, address string, relationship uint32) (*Member, error) {
// 	m.nextOfKin = NewNextOfKin(name, idCardNumber, address, relationship)
// 	return m, nil
// }

// func (m *Member) removeNextOfKin() bool {
// 	m.nextOfKin = NewNextOfKin("", "", "", 0)
// 	return true
// }

// func (m *Member) addEmploymentInfo(employerId, designation string, termsOfService uint) (*Member, error) {
// 	m.employmentInfo = NewEmploymentInfo(employerId, designation, termsOfService)
// 	return m, nil
// }

// func (m *Member) removeEmploymentInfo() bool {
// 	c.employmentInfo = NewEmploymentInfo("", "", 0)
// 	return true
// }

// func (m *Member) addAddress(addressLine, city, street, postalCode string, email string) (*Member, error) {
// 	m.Address = models.NewAddress(addressLine, city, street, postalCode, email)
// 	return m, nil
// }

// func (m *Member) removeAddress() (*Member, error) {
// 	m.Address = models.NewAddress("", "", "", "", "")
// 	return m, nil
// }

// func (m *Member) addNominee(name, memberId, idCardNumber, address, phoneNumber string, percentage float32, relationship uint) (*Member, error) {
// 	allocated := float32(0)
// 	newNominee, err := NewNominee(name, memberId, idCardNumber, address, phoneNumber, percentage, relationship)

// 	allocated += newNominee.Percentage

// 	if allocated < float32(100) {
// 		m.nominees = append(c.nominees, newNominee)
// 	}else {
// 		// exception thrown
// 		// log error
// 	}

// 	if err != nil {
// 		log.Print("sorry, failed to add nominee")
// 	}
// 	return m, nil
// }
