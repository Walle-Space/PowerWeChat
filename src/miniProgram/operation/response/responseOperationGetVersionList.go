package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	*response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
