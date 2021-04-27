package vda

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mixi-gaming/extension-library/transport"
)

func CreateRole(apiKey string, role string, permissions [][]string, otherData interface{}) map[string]interface{} {
	subject := "vda_request.rbac.role.create_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"role":        role,
				"permissions": permissions,
				"other_data":  otherData,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "CreateRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "CreateRole: " + err.Error()}
	}

	return resp
}

func GetAllRoles(apiKey string) map[string]interface{} {
	subject := "vda_request.rbac.role.get_all_roles"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRoles: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRoles: " + err.Error()}
	}

	return resp
}

func GetAllRolesWithPermissions(apiKey string) map[string]interface{} {
	subject := "vda_request.rbac.role.get_all_roles_with_permissions"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesWithPermissions: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesWithPermissions: " + err.Error()}
	}

	return resp
}

func GetDetailARole(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac.role.get_detail_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role": role,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetDetailARole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetDetailARole: " + err.Error()}
	}

	return resp
}

func UpdateRole(apiKey string, role string, permissions [][]string, otherData interface{}) map[string]interface{} {
	subject := "vda_request.rbac.role.update_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"role":        role,
				"permissions": permissions,
				"other_data":  otherData,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRole: " + err.Error()}
	}

	return resp
}

func AddPermissionsForRole(apiKey string, role string, permissions [][]string) map[string]interface{} {
	subject := "vda_request.rbac.role.add_permissions_for_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"role":        role,
				"permissions": permissions,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "AddPermissionsForRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "AddPermissionsForRole: " + err.Error()}
	}

	return resp
}

func DeleteRole(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac.role.delete_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role": role,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRole: " + err.Error()}
	}

	return resp
}

func DeletePermissionsForRole(apiKey string, role string, permissions [][]string) map[string]interface{} {
	subject := "vda_request.rbac.role.delete_permissions_for_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"role":        role,
				"permissions": permissions,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeletePermissionsForRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeletePermissionsForRole: " + err.Error()}
	}

	return resp
}

func DeleteAllPermissionsForRole(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac.role.delete_all_permissions_for_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role": role,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteAllPermissionsForRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteAllPermissionsForRole: " + err.Error()}
	}

	return resp
}
