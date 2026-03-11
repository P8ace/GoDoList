package service

import (
	"github.com/P8ace/GoDoList/internal/adapters/database/repo"
)

type Service struct {
	repo *repo.Queries
}
