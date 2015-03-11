package request

import (
	"fmt"
)

func NewRequest(url string, resType string, method string, postdata string) *Req {

	fmt.Print(url, resType, method, postdata)

	return &Req{url, resType, method, postdata}
}
