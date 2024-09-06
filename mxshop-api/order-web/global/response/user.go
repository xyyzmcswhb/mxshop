package response

import (
	"fmt"
	"time"
)

type UserResponse struct {
	Id       int32  `json:"id"`
	Nickname string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
}
type JsonTime time.Time

// 日期格式转换
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}
