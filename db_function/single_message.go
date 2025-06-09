package dbfunction

import (
	"context"
	"time"
)

func MessageSendAndSave(phoneNumber string, MessageContent string, CustomerId int, userName string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Get a Tx for making transaction requests.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "insert into "+SchemaName+"number_master (number,user_name,customer_id) values (?,?,?)", phoneNumber, userName, CustomerId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return nil, nil
}
