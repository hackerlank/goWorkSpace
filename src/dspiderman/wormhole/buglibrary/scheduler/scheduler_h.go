/**
	schedule crawler tasks (req cache)
 */
package scheduler

import (
	"dspiderman/wormhole/buglibrary/request"
)

type scheduler interface {
	Push(req *request.Req)
	Pull() *request.Req
}
