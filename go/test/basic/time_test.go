package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	layout := "2006-01-02"
	// 将当前时间转换为Unix时间戳
	now, _ := time.Parse(layout, time.Now().Format(layout))
	fmt.Println(now.Unix())
	// 将Unix时间戳转换为当前时间
	unix := int64(1575763200)
	fmt.Println(time.Unix(unix, 0).Format(layout))
}
