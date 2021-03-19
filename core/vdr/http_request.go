package vdr

import (
	"context"
	"encoding/json"

	"github.com/mixi-gaming/extension-library/transport"
)

// HttpCreateRecord - HttpCreateRecord
func HttpCreateRecord(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := transport.Domain + "/api/report/v1/secure/record/create/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := transport.MakeHTTPPutRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Create VDR Record Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Create VDR Record Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}

// HttpDataIncrement - HttpDataIncrement
func HttpDataIncrement(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := transport.Domain + "/api/report/v1/secure/record/data/increment/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Increment Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Increment Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}

// HttpDataDecrement - HttpDataDecrement
func HttpDataDecrement(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := transport.Domain + "/api/report/v1/secure/record/data/decrement/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Decrement Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Decrement Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}

// HttpRetrieveStatistics - HttpRetrieveStatistics
func HttpRetrieveStatistics(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := transport.Domain + "/api/report/v1/secure/statistics/retrieve/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Retrieve Statistics Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Retrieve Statistics Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}
