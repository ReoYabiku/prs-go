package repository

import "prs-go/entity"

type GitHub interface {
	ListURL() ([]*entity.URL, error)
}