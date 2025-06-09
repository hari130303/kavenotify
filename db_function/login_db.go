package dbfunction

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserDetail struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	RoleId   int    `json:"role_id"`
	GroupId  *int   `json:"group_id"`
}

func UserCheck(userName string, Password string) (any, error) {
	var userDetail UserDetail

	var DBpass string

	// fmt.Println("schema_name :", SchemaName)
	err := db.QueryRow("SELECT u.user_id,u.user_name,u.password,u.role,r.role_id,u.group_id FROM "+SchemaName+"user_master u left outer join  "+SchemaName+"role_master r on u.role = r.role_name where u.user_name = ?", userName).Scan(
		&userDetail.UserId, &userDetail.UserName, &DBpass, &userDetail.Role, &userDetail.RoleId, &userDetail.GroupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	PassCheck := CheckPasswordHash(Password, DBpass)

	if !PassCheck {
		return nil, fmt.Errorf("incorrect password")
	}

	uuid := uuid.New().String()

	userDetailJson, err := json.Marshal(userDetail)
	if err != nil {
		return nil, err
	}

	EncryptDetail, err := GetAESEncrypted(string(userDetailJson))
	if err != nil {
		return nil, err
	}

	ExpireTime := time.Now().Add(time.Hour * 1)

	// fmt.Println("uuid : ", uuid)
	// fmt.Println("encrypt data : ", EncryptDetail)
	// fmt.Println("expire time : ", ExpireTime)

	query := "INSERT INTO " + SchemaName + "session_master (session_id,encrypted_data,session_time) VALUES (?, ?, ?)"
	_, err = db.ExecContext(context.Background(), query, uuid, EncryptDetail, ExpireTime)
	if err != nil {
		return nil, err
	}
	var ReturnData struct {
		RedirectPage string `json:"redirect_page"`
		SessionId    string `json:"session_id"`
	}

	if userDetail.Role == "SUPER_ADMIN" {
		ReturnData.RedirectPage = "dashboard/single_message"
	} else if userDetail.Role == "GROUP_ADMIN" {
		ReturnData.RedirectPage = "dashboard/single_message"
	} else if userDetail.Role == "SYSTEM_ADMIN" {
		ReturnData.RedirectPage = "dashboard/single_message"
	} else {
		return nil, fmt.Errorf("error during assign dashboard page : invalid role")
	}
	ReturnData.SessionId = uuid
	return ReturnData, nil
}
