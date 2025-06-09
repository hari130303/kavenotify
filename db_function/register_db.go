package dbfunction

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func RegisterApi(GroupName string, GroupDesc *string, contactNumber string, mailId string, UserName string, Password string, customerType string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("error:", err)

		return nil, fmt.Errorf("internal server error")
	}
	defer tx.Rollback()
	var DBgroupID int64
	result, err := tx.ExecContext(ctx,
		"INSERT INTO group_master (group_name, group_desc, contact_number, mail_id) VALUES (?, ?, ?, ?);",
		GroupName, GroupDesc, contactNumber, mailId)
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error inserting into group_master: %v", err)
		return nil, fmt.Errorf("internal server error")
	}

	DBgroupID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error fetching last insert ID: %v", err)
		return nil, fmt.Errorf("internal server error")
	}

	// add the new customer_type for new group
	_, err = tx.ExecContext(ctx, "INSERT INTO customer_type (customer_type,customer_preference,group_id) VALUES (?, ?, ?);",
		customerType, true, DBgroupID)
	if err != nil {
		fmt.Println("error:", err)

		tx.Rollback()
		return nil, fmt.Errorf("internal server error")
	}

	var userExistCheck string
	if err = tx.QueryRowContext(ctx, "SELECT user_name from user_master where user_name = ?;",
		UserName).Scan(&userExistCheck); err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			fmt.Println("error:", err)

			return nil, fmt.Errorf("internal server error")
		}
	}
	if userExistCheck != "" {
		return nil, fmt.Errorf("user already exist")
	}

	HashPass, err := HashPassword(Password)
	if err != nil {
		tx.Rollback()
		fmt.Println("error:", err)

		return nil, fmt.Errorf("internal server error")
	}
	// add the user with group_admin role
	_, err = tx.ExecContext(ctx, "INSERT INTO user_master (user_name,password,role,group_id,contact_number) VALUES (?, ?, ?, ?, ?);",
		UserName, HashPass, "GROUP_ADMIN", DBgroupID, contactNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("error:", err)

		return nil, fmt.Errorf("internal server error")
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("error:", err)

		return nil, fmt.Errorf("internal server error")
	}
	return nil, nil
}
