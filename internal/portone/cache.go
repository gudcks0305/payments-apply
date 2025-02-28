package portone

import (
	"sync"
	"time"
)

type TokenCache struct {
	token      string
	expireTime time.Time
	mutex      sync.RWMutex
}

func NewTokenCache() *TokenCache {
	return &TokenCache{}
}

func (c *TokenCache) Get() (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.token == "" || time.Now().After(c.expireTime) {
		return "", false
	}
	return c.token, true
}

func (c *TokenCache) Set(token string, expireTime time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.token = token
	c.expireTime = expireTime
}

func (c *TokenCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.token = ""
	c.expireTime = time.Time{}
}
