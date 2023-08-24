package repositories

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"newglo/internals/domain"
	"strconv"
)

type Database struct {
	*sql.DB
	*redis.Client
	context.Context
}

func New(db *sql.DB, rdb *redis.Client, ctx context.Context) *Database {
	return &Database{DB: db, Client: rdb, Context: ctx}
}

func (d *Database) DoLogin(user domain.LoginCredentials) (int, error) {
	fmt.Print(user.Username)
	//TODO implement me
	var hashedPassword string
	var userlogged domain.User
	err := d.DB.QueryRow("select pass,pID,uID from users where userName = ?", user.Username).Scan(&hashedPassword, &userlogged.PID, &userlogged.UID)
	if err != nil {
		return 0, err
	}
	password := []byte(user.Password)
	hasher := md5.New()
	hasher.Write([]byte(user.Password))

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		if hashedPassword != hex.EncodeToString(hasher.Sum(nil)) {
			return 0, err
		}

	}
	err = d.DB.QueryRow("select cID,firstName,lastName,aID,daID from profile where pID = ?", userlogged.PID).Scan(&userlogged.CID, &userlogged.FullName, &userlogged.Email, &userlogged.AID, &userlogged.DaID)
	if err != nil {
		return 0, err
	}
	//store in in hset redis with user + pid as key
	err = d.Client.HSet(d.Context, user.Username+strconv.Itoa(int(userlogged.PID)), "CID", userlogged.CID, "FullName", userlogged.FullName, "Email", userlogged.Email, "AID", userlogged.AID, "DaID", userlogged.DaID).Err()

	return int(userlogged.PID), err
}
