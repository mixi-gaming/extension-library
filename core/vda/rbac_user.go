package vda

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mixi-gaming/extension-library/transport"
)

func AddRoleForUser(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac.user.add_role_for_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role":    role,
				"user_id": userID,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "AddRoleForUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "AddRoleForUser: " + err.Error()}
	}

	return resp
}

func DeleteRoleOfUser(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac.user.delete_role_of_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role":    role,
				"user_id": userID,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRoleForUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRoleForUser: " + err.Error()}
	}

	return resp
}

func GetUsersByRole(apiKey, role string, page, limit int) map[string]interface{} {
	subject := "vda_request.rbac.user.get_users_by_role"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role":  role,
				"page":  page,
				"limit": limit,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByRole: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByRole: " + err.Error()}
	}

	return resp
}

func GetRolesOfUser(apiKey, userID string) map[string]interface{} {
	subject := "vda_request.rbac.user.get_roles_of_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"user_id": userID,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetRolesOfUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetRolesOfUser: " + err.Error()}
	}

	return resp
}

func GetAllPermissionsOfUser(apiKey, userID string) map[string]interface{} {
	subject := "vda_request.rbac.user.get_permissions_of_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"user_id": userID,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllPermissionsOfUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllPermissionsOfUser: " + err.Error()}
	}

	return resp
}

func HasPermissionForUser(apiKey, userID string, permission []string) map[string]interface{} {
	subject := "vda_request.rbac.user.has_permission_for_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"user_id": userID,
			},
			"body": map[string]interface{}{
				"permission": permission,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "HasPermissionForUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "HasPermissionForUser: " + err.Error()}
	}

	return resp
}

func HasRoleForUser(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac.user.has_role_for_user"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"param": map[string]interface{}{
				"role":    role,
				"user_id": userID,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "HasRoleForUser: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "HasRoleForUser: " + err.Error()}
	}

	return resp
}

func GetUsersByPermission(apiKey string, permission []string) map[string]interface{} {
	subject := "vda_request.rbac.user.get_users_by_permission"
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"permission": permission,
			},
		},
	)

	msg, err := transport.Nc.Request(subject, requestBody, transport.Timeout)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByPermission: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByPermission: " + err.Error()}
	}

	return resp
}
