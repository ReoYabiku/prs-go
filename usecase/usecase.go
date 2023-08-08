package usecase

import (
	"prs-go/repository"

	"github.com/morikuni/failure"
)

type Usecase struct {
	gr repository.GitHub
}

func NewUsecase(gr repository.GitHub) *Usecase {
	return &Usecase{
		gr: gr,
	}
}

func (u *Usecase) OpenURLs() error {
	urls, err := u.gr.ListURL()
	if err != nil {
		return failure.Wrap(err)
	}

	for _, url := range urls {
		err := url.Call()
		if err != nil {
			return failure.Wrap(err)
		}
	}
	return nil
}
