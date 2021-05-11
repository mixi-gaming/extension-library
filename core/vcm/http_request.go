package vcm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mixi-gaming/extension-library/transport"
)

func HttpScheduleNoti(ctx context.Context, apiKey, id, timestamp, timezone string, dataMsg interface{}) map[string]interface{} {
	url := transport.Domain + "/api/notification/v2/schedule/activeSchedule/" + id
	apiKey = "Bearer " + apiKey

	tmp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(tmp, 0)

	bodyRequest := map[string]interface{}{
		"year":     tm.Year(),
		"month":    int(tm.Month()),
		"day":      tm.Day(),
		"hour":     tm.Hour(),
		"minutes":  tm.Minute(),
		"second":   tm.Second(),
		"timezone": timezone,
		"data":     dataMsg,
	}

	bodyBytes, _ := json.Marshal(bodyRequest)

	respBytes, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Send HTTP Request Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Read HTTP Response Failed: " + err.Error()}
	}
	return resp
}

func HttpDeleteScheduleNoti(ctx context.Context, apiKey, id string) map[string]interface{} {
	fmt.Println("Begin HttpDeleteScheduleNoti")
	url := transport.Domain + "/api/notification/v2/schedule/deactiveSchedule/" + id

	respBytes, err := transport.MakeHTTPDeleteRequest(ctx, url, []string{"Authorization"}, []string{apiKey})
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Send HTTP Request Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Read HTTP Response Failed: " + err.Error()}
	}
	return resp
}
