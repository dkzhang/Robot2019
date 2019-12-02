package cache

import (
	"Robot2019/myUtil"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host: "192.168.10.27:6379",
	}
	redis := NewRedis(opts)
	var err error
	timeoutDuration := 10 * time.Second

	if err = redis.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}

	rti := RealTimeInfo{
		InspectionID: 1,
		RecordID:     "",
		Level:        "Info",
		DateTime:     myUtil.FormatTime(time.Now()),
		TextContent:  "test",
		ImageUrl:     "image url",
	}
	if err = redis.SetStream("test", &rti); err != nil {
		t.Errorf("SetStream Error , err=%v", err)
	}

	getRti, err := redis.GetAllStream("test")
	if err != nil {
		t.Errorf("GetAllStream error, err=%v", err)
	} else {
		t.Logf("GetAllStream replay = %v", getRti)
	}
}
