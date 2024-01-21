package dbutils

import "database/sql"

func GetQueryRows[T any](rows *sql.Rows, handler func(rowData *T) error) ([]T, error) {
	defer rows.Close()

	var rowsData []T
	for rows.Next() {
		var rowData T
		err := handler(&rowData)
		if err != nil {
			return nil, err
		}
		rowsData = append(rowsData, rowData)
	}
	return rowsData, nil
}
