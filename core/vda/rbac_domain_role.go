package vda

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mixi-gaming/extension-library/transport"
)

func CreateRoleInDomain(apiKey string, role string, permissions [][]string, otherData interface{}) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.create_role"
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
		return map[string]interface{}{"code": "10", "message": "CreateRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "CreateRoleInDomain: " + err.Error()}
	}

	return resp
}

func GetAllRolesInDomain(apiKey string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.get_all_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesInDomain: " + err.Error()}
	}

	return resp
}

func GetAllRolesWithPermissionsInDomain(apiKey string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.get_all_roles_with_permissions"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesWithPermissionsInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllRolesWithPermissionsInDomain: " + err.Error()}
	}

	return resp
}

func GetDetailARoleInDomain(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.get_detail_role"
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
		return map[string]interface{}{"code": "10", "message": "GetDetailARoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetDetailARoleInDomain: " + err.Error()}
	}

	return resp
}

func UpdateRoleInDomain(apiKey string, role string, permissions [][]string, otherData interface{}) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.update_role"
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
		return map[string]interface{}{"code": "10", "message": "UpdateRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRoleInDomain: " + err.Error()}
	}

	return resp
}

func AddPermissionsForRole(apiKey string, role string, permissions [][]string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.add_permissions_for_role"
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

func DeleteRoleInDomain(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.delete_role"
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
		return map[string]interface{}{"code": "10", "message": "DeleteRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRoleInDomain: " + err.Error()}
	}

	return resp
}

func DeletePermissionsForRoleInDomain(apiKey string, role string, permissions [][]string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.delete_permissions_for_role"
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
		return map[string]interface{}{"code": "10", "message": "DeletePermissionsForRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeletePermissionsForRoleInDomain: " + err.Error()}
	}

	return resp
}

func DeleteAllPermissionsForRoleInDomain(apiKey string, role string) map[string]interface{} {
	subject := "vda_request.rbac_domain.role.delete_all_permissions_for_role"
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
		return map[string]interface{}{"code": "10", "message": "DeleteAllPermissionsForRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteAllPermissionsForRoleInDomain: " + err.Error()}
	}

	return resp
}
