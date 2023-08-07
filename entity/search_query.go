package entity

type prStatus int

const (
	PRStatusClosed = iota
	PRStatusMerged
	PRStatusOpen
)

func (s prStatus) String() string {
	switch s {
	case PRStatusClosed:
		return "CLOSED"
	case PRStatusMerged:
		return "MERGED"
	case PRStatusOpen:
		return "OPEN"
	default: // TODO: 不正値が混入した場合の振る舞いを考える
		return "OPEN"
	}
}

type SearchQuery struct {
	Owner           string
	Status          prStatus
	RepositoryOwner string
}
