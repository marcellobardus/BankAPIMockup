package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/utils"
	"math/rand"
	"time"
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
	Endpoints         []string  `bson:"endpoints" json:"endpoints"`
	Token             string    `bson:"token" json:"token"`
	Expiration        time.Time `bson:"expiration" json:"expiration"`
	AuthorizedAccount *Account  `bson:"authorizedaccount" json:"authorizedaccount"`
}

func NewAuthorization(endpoints []string, expiration time.Time, account *Account) *Authorization {
	a := new(Authorization)
	a.Endpoints = endpoints
	a.Expiration = expiration
	a.AuthorizedAccount = account
	a.Token = setToken(endpoints, expiration)
	return a
}

func setToken(endpoints []string, expiration time.Time) string {
	rand.Seed(time.Now().Unix())
	var bytesArray = make([]byte, 0)
	bytesArraySize := rand.Intn(255-10) + 10
	for i := 0; i < bytesArraySize; i++ {
		random := rand.Intn(127-0) + 0
		bytesArray = append(bytesArray, byte(random))
	}

	var endpointsConc string

	for i := range endpoints {
		endpointsConc = utils.StringConcatenation(endpoints[i], "")
	}

	hashData := utils.StringConcatenation(string(bytesArray), endpointsConc, expiration.String())
	md5Hash := md5.Sum([]byte(hashData))
	token := hex.EncodeToString(md5Hash[:])
	return token
}
