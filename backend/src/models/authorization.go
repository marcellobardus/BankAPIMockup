package models

import (
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
	RequestMethod HttpMethod `bson:"requestmethod" json:"requestmethod"`
	Endpoint      string     `bson:"endpoint" json:"endpoint"`
	Token         string     `bson:"token" json:"token"`
	Expiration    time.Time  `bson:"expiration" json:"expiration"`
}

func NewAuthorization(requestmethod HttpMethod, endpoint string, expiration time.Time) *Authorization {
	a := new(Authorization)
	a.RequestMethod = requestmethod
	a.Endpoint = endpoint
	a.Expiration = expiration
	a.Token = setToken(requestmethod, endpoint, expiration)
	return a
}

func setToken(requestmethod HttpMethod, endpoint string, expiration time.Time) string {
	return ""
}
