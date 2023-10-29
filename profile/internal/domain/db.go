package domain

import (
	"sync"
	"time"
)

type DB struct {
	mx       sync.RWMutex
	users    map[int64]*User
	logins   map[string]int64
	uniqueID int64
}

func NewDB() *DB {
	return &DB{
		users:    make(map[int64]*User),
		logins:   make(map[string]int64),
		uniqueID: 0,
	}
}

func (db *DB) StoreUser(login string) *User {
	db.mx.Lock()
	defer db.mx.Unlock()
	db.uniqueID++
	user := &User{
		Id:        db.uniqueID,
		Login:     login,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.users[user.Id] = user
	db.logins[login] = user.Id
	return user
}

func (db *DB) GetUser(id int64) (*User, bool) {
	db.mx.RLock()
	defer db.mx.RUnlock()
	user, exists := db.users[id]
	return user, exists
}

func (db *DB) FindLogin(login string) (int64, bool) {
	db.mx.RLock()
	defer db.mx.RUnlock()
	userID, exists := db.logins[login]
	return userID, exists
}
