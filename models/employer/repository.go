package employer

import (
	"errors"

	"github.com/google/uuid"
)

var(
	ErrEmployerNotFound = errors.New("employer not found")
	ErrEmployerCreateFailed = errors.New("employer create failed")
	ErrEmployerUpdateFailed = errors.New("employer update failed")
)

type EmployerRepository interface {
	Get(uuid.UUID) (Employer, error)
	Create(Employer) (Employer, error)
	Update(Employer) (Employer, error)
}