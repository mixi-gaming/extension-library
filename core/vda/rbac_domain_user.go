package vda

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mixi-gaming/extension-library/transport"
)

func AddRoleForUserInDomain(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.add_role_for_user"
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
		return map[string]interface{}{"code": "10", "message": "AddRoleForUserInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "AddRoleForUserInDomain: " + err.Error()}
	}

	return resp
}

func DeleteRoleOfUserInDomain(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.delete_role_of_user"
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
		return map[string]interface{}{"code": "10", "message": "DeleteRoleForUserInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRoleForUserInDomain: " + err.Error()}
	}

	return resp
}

func GetUsersByRoleInDomain(apiKey, role string, page, limit int) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.get_users_by_role"
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
		return map[string]interface{}{"code": "10", "message": "GetUsersByRoleInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByRoleInDomain: " + err.Error()}
	}

	return resp
}

func GetRolesOfUserInDomain(apiKey, userID string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.get_roles_of_user"
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
		return map[string]interface{}{"code": "10", "message": "GetRolesOfUserInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetRolesOfUserInDomain: " + err.Error()}
	}

	return resp
}

func GetAllPermissionsOfUserInDomain(apiKey, userID string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.get_permissions_of_user"
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
		return map[string]interface{}{"code": "10", "message": "GetAllPermissionsOfUserInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetAllPermissionsOfUserInDomain: " + err.Error()}
	}

	return resp
}

func HasPermissionForUser(apiKey, userID string, permission []string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.has_permission_for_user"
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

func HasRoleForUserInDomain(apiKey, role, userID string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.has_role_for_user"
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
		return map[string]interface{}{"code": "10", "message": "HasRoleForUserInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "HasRoleForUserInDomain: " + err.Error()}
	}

	return resp
}

func GetUsersByPermissionInDomain(apiKey string, permission []string) map[string]interface{} {
	subject := "vda_request.rbac_domain.user.get_users_by_permission"
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
		return map[string]interface{}{"code": "10", "message": "GetUsersByPermissionInDomain: " + err.Error()}
	}

	resp := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &resp); err != nil {
		return map[string]interface{}{"code": "10", "message": "GetUsersByPermissionInDomain: " + err.Error()}
	}

	return resp
}
