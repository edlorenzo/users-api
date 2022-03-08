package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"

	"github.com/edlorenzo/users-api/logs"
	"github.com/edlorenzo/users-api/model"
)

type UserStore struct {
	db  *redis.Client
	api string
}

func NewUserStore(db *redis.Client, uri string) *UserStore {
	return &UserStore{
		db:  db,
		api: uri,
	}
}

func (as *UserStore) GetUserByIDs(user *model.NewUser) (*model.NewUser, error) {
	return user, nil
}

func (as *UserStore) ListLimitOffset(offset, limit int) (model.Profile, int64, error) {
	var (
		profile model.Profile
		count   int64
	)
	return profile, count, nil
}

func (as *UserStore) List() (model.ProfileList, error) {
	var (
		users, newUsers model.NewUser
	)

	res, err := http.Get(as.api)
	if err != nil {
		log.Fatal(err)
	}

	logs.Logs.Infof("[Info] User List() Status: %d", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Logs.Errorf("[Error] User List() ReadAll err: %v", err)
	}
	err = res.Body.Close()
	if err != nil {
		logs.Logs.Errorf("[Error] User List() Body.Close() err: %v", err)
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		logs.Logs.Errorf("[Error] User List() Unmarshal err: %v", err)
	}

	for i := 0; i < 10; i++ {
		newUsers = append(newUsers, users[i])
	}

	newProfile := as.userProfile(newUsers)

	return newProfile, nil
}

func (as *UserStore) userProfile(newUsers model.NewUser) model.ProfileList {
	var (
		profile, deserialized model.Profile
		newProfile            model.ProfileList
	)

	for _, user := range newUsers {
		key := fmt.Sprintf("%s%s", user.Login, "_user_cache")

		val, err := as.db.Get(key).Bytes()
		if err != nil {
			prof, err := http.Get(as.api + "/" + user.Login)
			if err != nil {
				logs.Logs.Infof("[Info] Unable to retrieve user: %s", user.Login)
				continue
			}

			profileBody, err := ioutil.ReadAll(prof.Body)
			if err != nil {
				logs.Logs.Errorf("[Error] User prifile ReadAll: %v", err)
			}
			err = prof.Body.Close()
			if err != nil {
				logs.Logs.Errorf("[Error] User prifile Body.Close(): %v", err)
			}

			err = json.Unmarshal(profileBody, &profile)
			if err != nil {
				logs.Logs.Errorf("[Error] User prifile Unmarshal: %v", err)
			}

			p := &model.Profile{
				Name:        profile.Name,
				Login:       profile.Login,
				Company:     profile.Company,
				Followers:   profile.Followers,
				PublicRepos: profile.PublicRepos,
			}

			serialized, err := json.Marshal(p)
			if err != nil {
				logs.Logs.Errorf("[Error] invalid serialization: %v", err)
			} else {
				// fmt.Println("serialized data: ", string(serialized))
				err := as.db.Set(key, serialized, 2*time.Minute).Err()
				if err != nil {
					log.Panic(err)
				}
			}

			err = json.Unmarshal(serialized, &deserialized)
			if err != nil {
				logs.Logs.Errorf("[Error] invalid deserialization: %v", err)
			}
			newProfile = append(newProfile, deserialized)
			continue
		}

		userProfile := model.Profile{}
		err = json.Unmarshal(val, &userProfile)
		if err != nil {
			logs.Logs.Errorf("[Error] User prifile Redis Unmarshal: %v", err)
		}

		newProfile = append(newProfile, userProfile)
	}

	return newProfile
}
