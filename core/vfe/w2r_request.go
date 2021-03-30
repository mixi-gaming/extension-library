package vfe

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mixi-gaming/extension-library/model"
	"github.com/mixi-gaming/extension-library/transport"
)

// NatsUploadFiles - NatsUploadFiles
func NatsUploadFiles(apiKey, bucketID string, files []*model.FileStorage) map[string]interface{} {
	subj := "fs_request.file_storage.upload_files"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.ApiKey = apiKey
	nReq.BucketID = bucketID
	nReq.Files = files

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		fmt.Println("Failed at transport.Nc.Request:", err)
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		fmt.Println("Failed at Unmarshal response data:", err)
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}

// NatsDeleteFiles - NatsDeleteFiles
func NatsDeleteFiles(apiKey, bucketID string, recordIDs []interface{}) map[string]interface{} {
	subj := "fs_request.file_storage.delete_files"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.ApiKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordIDs = recordIDs

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}

// NatsRetrieveObjsFolder - Nats Retrieve all object in folder
func NatsRetrieveObjsFolder(apiKey, bucketID, folder string) map[string]interface{} {
	subj := "fs_request.file_storage.retrieve_objs_folder"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.ApiKey = apiKey
	nReq.BucketID = bucketID
	nReq.Folder = folder

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}
