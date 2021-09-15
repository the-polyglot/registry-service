package member

import (
	"errors"

	"github.com/google/uuid"
)

var(
	ErrMemberNotFound = errors.New("member not found")
	ErrMemberRegistrationFailed = errors.New("member registration failed")
	ErrMemberDetailsUpdateFailed = errors.New("member details update failed")
)

type MemberRepository interface {
	Get(uuid.UUID) (*Member, error)
	Register(Member) (*Member, error)
	Update(Member) (*Member, error)
}