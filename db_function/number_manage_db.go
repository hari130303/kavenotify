package dbfunction

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type NumberDetail struct {
	Number       string    `json:"number"`
	Username     string    `json:"user_name"`
	CustomerName string    `json:"Customer_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func NumberManageTable(GroupId *int, searchTerm string, start int, limit int, orderBy int) (any, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	numberDetails := make([]NumberDetail, 0)
	var tableData TableData

	searchStr := "%" + strings.ToLower(searchTerm) + "%"

	if orderBy == 0 {
		orderBy = 1
	}

	orderColumn := map[int]string{
		1: "n.number",
		2: "n.user_name",
		3: "c.customer_type",
		4: "n.created_at",
	}
	// Code to count the customer types
	countQuery := `
	SELECT COUNT('X') 
	from number_master n
	left join customer_type c on n.customer_id = c.customer_id
	where c.group_id = coalesce(?,c.group_id)
	AND (LOWER(n.user_name) LIKE (?) OR LOWER(c.customer_type) LIKE (?) OR LOWER(n.number) LIKE (?))
	`

	err := db.QueryRowContext(ctx, countQuery, GroupId, searchStr, searchStr, searchStr).Scan(&tableData.Count)
	if err != nil {
		return tableData, err
	}

	// Code to fetch the customer types
	selectQuery := fmt.Sprintf(`
	SELECT n.number,n.user_name,c.customer_type,n.created_at  
	from number_master n
	left join customer_type c on n.customer_id = c.customer_id
	where c.group_id = coalesce(?,c.group_id)
	AND (LOWER(n.user_name) LIKE (?) OR LOWER(c.customer_type) LIKE (?) OR LOWER(n.number) LIKE (?))
	ORDER BY %s 
	LIMIT ?,?
	`, orderColumn[orderBy])

	// fmt.Println("select query : ", selectQuery)
	rows, err := db.QueryContext(ctx, selectQuery, GroupId, searchStr, searchStr, searchStr, start, limit)
	if err != nil {
		return tableData, err
	}
	defer rows.Close()

	for rows.Next() {
		var numberDetail NumberDetail
		var CreatedTimeString []byte
		err := rows.Scan(&numberDetail.Number, &numberDetail.Username, &numberDetail.CustomerName, &CreatedTimeString)
		if err != nil {
			return tableData, err
		}

		createdStr := string(CreatedTimeString)
		numberDetail.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdStr)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			// You can also skip or return error here
		}
		// customerType.CreatedAt = createdAt
		numberDetails = append(numberDetails, numberDetail)
	}

	tableData.Start = start
	tableData.Limit = limit
	tableData.Data = numberDetails

	return tableData, nil

}
