package dbfunction

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type TableData struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
	Count int `json:"count"`
	Data  any `json:"data"`
}
type CustomerTypeDD struct {
	CustomerId   int    `json:"customer_id"`
	CustomerName string `json:"customer_name"`
}

type CustomerTypeDetail struct {
	CustomerId   int       `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	GroupName    string    `json:"group_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetCustomerType(GroupId *int) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	customerDetails := make([]CustomerTypeDD, 0)

	rows, err := db.QueryContext(ctx, "select customer_id,customer_type from customer_type where group_id = coalesce(?,group_id)", GroupId)
	if err != nil {
		return customerDetails, err
	}
	defer rows.Close()

	for rows.Next() {
		var CustomerTypeDDData CustomerTypeDD
		err := rows.Scan(&CustomerTypeDDData.CustomerId, &CustomerTypeDDData.CustomerName)
		if err != nil {
			return customerDetails, err
		}

		customerDetails = append(customerDetails, CustomerTypeDDData)
	}

	return customerDetails, nil
}

func GetCustomerTypeTable(GroupId *int, searchTerm string, start int, limit int, orderBy int) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	customerDetails := make([]CustomerTypeDetail, 0)
	var tableData TableData

	searchStr := "%" + strings.ToLower(searchTerm) + "%"

	if orderBy == 0 {
		orderBy = 1
	}

	orderColumn := map[int]string{
		1: "c.customer_id",
		2: "c.customer_type",
		3: "g.group_name",
		4: "c.created_at",
	}
	// Code to count the customer types
	countQuery := `
	SELECT COUNT('X') 
	FROM customer_type c
	LEFT JOIN group_master g ON c.group_id = g.group_id
	WHERE c.group_id = COALESCE(?, c.group_id)
	AND (LOWER(c.customer_type) LIKE (?) OR LOWER(g.group_name) LIKE (?))
	`

	err := db.QueryRowContext(ctx, countQuery, GroupId, searchStr, searchStr).Scan(&tableData.Count)
	if err != nil {
		return tableData, err
	}

	// Code to fetch the customer types
	selectQuery := fmt.Sprintf(`
	SELECT c.customer_id, c.customer_type, g.group_name, c.created_at 
	FROM customer_type c
	LEFT JOIN group_master g ON c.group_id = g.group_id
	WHERE c.group_id = COALESCE(?, c.group_id)
	AND (LOWER(c.customer_type) LIKE (?) OR LOWER(g.group_name) LIKE (?))
	ORDER BY %s 
	LIMIT ?,?
	`, orderColumn[orderBy])

	// fmt.Println("select query : ", selectQuery)
	rows, err := db.QueryContext(ctx, selectQuery, GroupId, searchStr, searchStr, start, limit)
	if err != nil {
		return tableData, err
	}
	defer rows.Close()

	for rows.Next() {
		var customerType CustomerTypeDetail
		var CreatedTimeString []byte
		err := rows.Scan(&customerType.CustomerId, &customerType.CustomerName, &customerType.GroupName, &CreatedTimeString)
		if err != nil {
			return tableData, err
		}

		createdStr := string(CreatedTimeString)
		customerType.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdStr)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			// You can also skip or return error here
		}
		// customerType.CreatedAt = createdAt
		customerDetails = append(customerDetails, customerType)
	}

	tableData.Start = start
	tableData.Limit = limit
	tableData.Data = customerDetails

	return tableData, nil

}
