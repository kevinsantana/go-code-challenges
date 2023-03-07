package share

import (
	"net/http"
)

type AccessToken struct {
	AccessToken string `json:"access_token" example:"01GF442ATTVP4M6M0XGHQYT544"`
	Expiration  int64  `json:"expiration" example:"1666116857"`
} //@name AuthorizeResponse

type User struct {
	Id string
	Secret string
	Session *http.Cookie
	AccessToken string
	ExpirationTime int64
}