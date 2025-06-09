package dbfunction

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

func SessionManager() {
	for {
		_, err := db.Exec("DELETE FROM session_master WHERE session_time <= NOW() - INTERVAL 1 HOUR;")
		if err != nil {
			fmt.Println("error during delete expired sessions :", err)
		}

		// fmt.Println("deleted")
		time.Sleep(1 * time.Hour)
	}
}

func SessionCheck(sessionId string) (UserDetail, error) {
	var UserDetail UserDetail
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	SessionDataSelect := `select encrypted_data from session_master where session_id = ?`
	var SessionData string
	err := db.QueryRowContext(ctx, SessionDataSelect, sessionId).Scan(&SessionData)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserDetail, fmt.Errorf("session expired or invalid")
		} else {
			return UserDetail, err
		}
	}

	DecryptData, err := GetAESDecrypted(SessionData)
	if err != nil {
		return UserDetail, err
	}

	// fmt.Println("decrypted Data : ", DecryptData)

	err = json.Unmarshal(DecryptData, &UserDetail)
	if err != nil {
		return UserDetail, err
	}

	fmt.Println("user data from session check : ", UserDetail)
	return UserDetail, nil
}
