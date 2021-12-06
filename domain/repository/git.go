package repository

import "github.com/masibw/gifc/domain/entity"

type Git interface {
	Get() (*entity.Git, error)
}
