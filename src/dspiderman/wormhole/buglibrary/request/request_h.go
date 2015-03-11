package request

import (
	"net/http"
)

type Req struct {
	Request_URL    string
	Remote_Address string
	Request_Method string
	Status_Code    int

	Request_Headers http.Header
	Cookie          []*http.Cookie

	Postdata		string

	RespType		string

}

