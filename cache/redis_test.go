package cache

import (
	"Robot2019/myUtil"
	"Robot2019/webServer/inspectionRecord/realtimeRecord"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host: "192.168.10.27:6379",
	}
	theRedis := NewRedis(opts)
	var err error
	timeoutDuration := 10 * time.Second

	if err = theRedis.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !theRedis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := theRedis.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = theRedis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}

	////////////////////////////////////////////////////////

	for i := 0; i < 5; i++ {
		rti := realtimeRecord.RealTimeInfo{
			InspectionID: i,
			RecordID:     i * 10,
			Level:        "Info",
			DateTime:     myUtil.FormatTime(time.Now()),
			TextContent:  "test",
			ImageUrl:     "image url",
		}
		if err = theRedis.ListMaxLenRPush("testList", rti, 3); err != nil {
			t.Error("ListMaxLenRPush Error", err)
		}

		if length, err := theRedis.ListLen("testList"); err != nil {
			t.Error("ListLen Error", err)
		} else {
			t.Logf("list length = %d", length)
		}

		if strArray, err := theRedis.ListLRange("testList", 0, -1); err != nil {
			t.Error("ListLRange Error", err)
		} else {
			t.Logf("string array = %s", strArray)
		}
	}

}
