package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/xariez/registry-service/models/employer"
)

type MemoryRepository struct {
	employers map[uuid.UUID]employer.Employer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		employers: make(map[uuid.UUID]employer.Employer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (employer.Employer, error) {
	if employer, ok := mr.employers[id]; ok {
		return employer, nil
	}
	return employer.Employer{}, employer.ErrEmployerNotFound
}

func (mr *MemoryRepository) Create(e employer.Employer) (employer.Employer, error) {
	if mr.employers == nil {
		mr.Lock()
		mr.employers = make(map[uuid.UUID]employer.Employer)
		mr.Unlock()
	}
	if _, ok := mr.employers[e.ID()]; ok {
		return employer.Employer{}, fmt.Errorf("employer already exists: %w", employer.ErrEmployerCreateFailed)
	}
	mr.Lock()
	mr.employers[e.ID()] = e
	mr.Unlock()
	return e, nil
}

func (mr *MemoryRepository) Update(e employer.Employer) (employer.Employer, error) {
	if _, ok := mr.employers[e.ID()]; ok {
		return employer.Employer{}, fmt.Errorf("employer already exists: %w", employer.ErrEmployerUpdateFailed)
	}
	mr.Lock()
	mr.employers[e.ID()] = e
	mr.Unlock()
	return e, nil
}
