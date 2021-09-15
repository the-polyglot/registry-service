package employer

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
	"github.com/xariez/registry-service/models"
)

type Employer struct {
	id   uuid.UUID
	name string
	models.Address
	phoneInfo models.PhoneInfo
}

func NewEmployer(name string) (Employer, error) {
	if name == "" {
		return Employer{}, errors.New("name is required")
	}

	ne := Employer{uuid.New(), name, models.Address{}, models.PhoneInfo{}}
	return ne, nil
}

func (e *Employer) ID() uuid.UUID {
	return e.id
}

func (e *Employer) SetID(id uuid.UUID) {
	e.id = id
}

func (e *Employer) Name() string {
	return e.name
}

func (e *Employer) SetName(name string) {
	e.name = name
}

func (e *Employer) GetAddress() models.Address {
	return e.Address
}

func (e *Employer) SetAddress(addressLine, city, street, postalCode string, email string) (bool, error) {
	e.Address = models.NewAddress(addressLine, city, street, postalCode, email)
	if reflect.DeepEqual(e.Address, models.Address{}) {
		return false, errors.New("error setting address")
	}
	return true, nil
}

func (e *Employer) RemoveAddress() (*Employer, error) {
	e.Address = models.Address{}
	return e, nil
}

func (e *Employer) AddPhoneInfo(primaryLine, otherLine string) (*Employer, error) {
	e.phoneInfo = models.NewPhoneInfo(primaryLine, otherLine)
	return e, nil
}

func (e *Employer) RemovePhoneInfo() (*Employer, error) {
	e.phoneInfo = models.PhoneInfo{}
	return e, nil
}
