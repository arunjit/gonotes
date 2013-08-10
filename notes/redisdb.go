// +build !appengine

package notes

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisDatabase struct {
	Pool *redis.Pool
}

const (
	maxIdleConnections    = 3
	idleConnectionTimeout = 45 * time.Second
)

func NewRedisDatabase(redisAddr string) *RedisDatabase {
	pool := &redis.Pool{
		MaxIdle:     maxIdleConnections,
		IdleTimeout: idleConnectionTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
	return &RedisDatabase{
		Pool: pool,
	}
}

func (s *RedisDatabase) Search(r *http.Request, opts *SearchOptions) ([]*Note, error) {
	return []*Note{}, nil
}

func (s *RedisDatabase) Update(r *http.Request, note *Note) (*Note, error) {
	if note.ID == 0 {
		note.ID = createID()
	}
	// set(id, encode(note))
	return note, nil
}

func (s *RedisDatabase) Delete(r *http.Request, id int64) error {
	return nil
}

func createID() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}
