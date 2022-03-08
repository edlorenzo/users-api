package redis

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type Conf struct {
	Address  string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"300"`
}

var client *redis.Client

func InitRedis(cfg *Conf) (*redis.Client, error) {
	addr := cfg.Address
	host := strings.Split(addr, ",")
	client = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_URL", host[0]),
		Password: getEnv("REDIS_PASSWORD", cfg.Password),
		DB:       0,
		// ReadTimeout: 15 * time.Second,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return client, err
	}

	log.Println("Redis Host: ", cfg.Address)
	log.Println("Redis Success!: ", pong, err)

	return client, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetSession(cookie string) (bool, error) {
	key := fmt.Sprintf("%s%s", cookie, "session_timeout_key")
	cmd := client.Get(key)
	if cmd.Err() != nil {
		log.Printf("[error] get session failed, cookie: %s, error: %s", cookie, cmd.Err().Error())
		return false, cmd.Err()
	}

	return true, nil
}

func ResetSession(cookie string) error {
	key := fmt.Sprintf("%s%s", cookie, "session_timeout_key")
	cmd := client.Set(key, key, time.Minute*2)
	if cmd.Err() != nil {
		log.Printf("[error] get session failed, cookie: %s, error: %s", cookie, cmd.Err().Error())
		time.Sleep(time.Millisecond * 10)
		cmd = client.Set(key, key, time.Minute*2)
		if cmd.Err() != nil {
			log.Printf("[error] get session failed again, cookie: %s, error: %s", cookie, cmd.Err().Error())
			return cmd.Err()
		}
	}

	return nil
}
