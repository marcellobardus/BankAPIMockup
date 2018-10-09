package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dchest/uniuri"
	"time"

	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/utils"
)

type HttpMethod int

const (
	GET    HttpMethod = 0
	POST   HttpMethod = 1
	PUT    HttpMethod = 2
	DELETE HttpMethod = 3
	PATCH  HttpMethod = 4
)

type Authorization struct {
	Token             string    `bson:"token" json:"token"`
	Expiration        time.Time `bson:"expiration" json:"expiration"`
	AuthorizedAccount *Account  `bson:"authorizedaccount" json:"authorizedaccount"`
}

func NewAuthorization(expiration time.Time, account *Account) *Authorization {
	a := new(Authorization)
	a.Expiration = expiration
	a.AuthorizedAccount = account
	a.Token = setToken(expiration)
	return a
}

func setToken(expiration time.Time) string {
	randomString := uniuri.New()

	hashData := utils.StringConcatenation(randomString, expiration.String())
	md5Hash := md5.Sum([]byte(hashData))
	token := hex.EncodeToString(md5Hash[:])
	return token
}
