package vcm

import (
	"encoding/json"
	"fmt"

	"github.com/mixi-gaming/extension-library/transport"

	"github.com/google/uuid"
)

// NatsSendEmailDefault - NatsSendEmailDefault
func NatsSendEmailDefault(to []string, subject, body string) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"is_default_email": true,
			"body": map[string]interface{}{
				"to":      to,
				"subject": subject,
				"body":    body,
			},
		})

	emailSubject := "vcm_email_request.send_email_default"
	msg, err := transport.Nc.Request(emailSubject, requestBody, transport.Timeout)
	if err != nil {
		fmt.Println("Request VCM Send Email Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}

// NatsPushIOSNotification - NatsPushIOSNotification
func NatsPushIOSNotification(apiKey, to, title, body string, data interface{}) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"notification": map[string]interface{}{
					"title": title,
					"body":  body,
				},
				"data": data,
				"to":   to,
			},
		},
	)

	subject := "vcm_ios_request.push_notification"
	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		fmt.Println("Request VCM Push IOS Notification Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}

// NatsPushAndroidNotification - NatsPushAndroidNotification
func NatsPushAndroidNotification(apiKey, to, title, body string, data interface{}) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"notification": map[string]interface{}{
					"title":   title,
					"message": body,
				},
				"data": data,
				"to":   to,
			},
		},
	)

	subject := "vcm_android_request.push_notification"
	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		fmt.Println("Request VCM Push Android Notification Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}

func NatsGetAllIOSHistory(apiKey string, start int64, stop int64) map[string]interface{} {
	subject := "vcm_ios_request.get_all_history"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"start":      start,
			"stop":       stop,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All IOS History FAILED: " + err.Error()}
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All IOS History FAILED: " + err.Error()}
	}

	return result
}

func NatsGetAllAndroidHistory(apiKey string, start int64, stop int64) map[string]interface{} {
	subject := "vcm_android_request.get_all_history"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"start":      start,
			"stop":       stop,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Android History FAILED: " + err.Error()}
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Android History FAILED: " + err.Error()}
	}

	return result
}

func NatsGetAndroidNotisWithFilter(apiKey string, page, limit int, filterBody map[string]interface{}) map[string]interface{} {
	subject := "vcm_android_request.get_notis_with_filter"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id":  uuid.New().String(),
			"api_key":     apiKey,
			"page":        page,
			"limit":       limit,
			"filter_body": filterBody,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Android History FAILED: " + err.Error()}
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Android History FAILED: " + err.Error()}
	}

	return result
}
