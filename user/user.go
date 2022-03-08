package user

import "github.com/edlorenzo/users-api/model"

type Store interface {
	GetUserByIDs(ids *model.NewUser) (*model.NewUser, error)
	ListLimitOffset(offset, limit int) (model.Profile, int64, error)
	List() (model.ProfileList, error)
}
