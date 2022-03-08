package handler

import (
	"sort"

	_ "github.com/gofiber/fiber/v2"

	"github.com/edlorenzo/users-api/model"
	"github.com/edlorenzo/users-api/user"
)

type userDataListResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []*DataList `json:"data"`
}

type userListResponse struct {
	User userDataListResponse `json:"user"`
}

type DataList struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}

func newUserListResponse(us user.Store, profile model.ProfileList, status int, message string) *userListResponse {
	r := new(userListResponse)
	r.User.Status = status
	r.User.Message = message
	r.User.Data = make([]*DataList, 0)
	for _, a := range profile {
		ar := new(DataList)
		ar.Name = a.Name
		ar.Login = a.Login
		ar.Company = a.Company
		ar.Followers = a.Followers
		ar.PublicRepos = a.PublicRepos
		r.User.Data = append(r.User.Data, ar)
	}

	sort.Slice(r.User.Data[:], func(i, j int) bool {
		return r.User.Data[i].Name < r.User.Data[j].Name
	})

	return r
}
