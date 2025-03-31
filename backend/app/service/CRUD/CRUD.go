package service

import (
	"database/sql"
	"errors"
	"reflect"
)

// DBService defines the methods for common database operations.
type DBService interface {
	Create(query string, data ...interface{}) (int64, error)
	SelectSingle(query string, condition []interface{}, dest ...interface{}) error
	SelectAllForPrimitive(query string, condition []interface{}, dest interface{}) error
	SelectAllForStruct(query string, condition []interface{}, dest interface{}, fields []string) error
	Update(query string, data ...interface{}) (int64, error)
	Delete(query string, data ...interface{}) error
}

// DB implements the DBService interface.
type DB struct {
	instanceOfDB *sql.DB
}

func SqlService(db *sql.DB) DBService {
	return &DB{
		instanceOfDB: db,
	}
}

func (db *DB) Create(query string, data ...interface{}) (int64, error) {
	result, err := db.instanceOfDB.Exec(query, data...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) SelectSingle(query string, condition []interface{}, dest ...interface{}) error {
	// Prepare and execute the query
	err := db.instanceOfDB.QueryRow(query, condition...).Scan(dest...)
	return err
}

func (db *DB) SelectAllForPrimitive(query string, condition []interface{}, dest interface{}) error {
	rows, err := db.instanceOfDB.Query(query, condition...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Ensure dest is a slice
	sliceValue := reflect.ValueOf(dest)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}

	sliceElemType := sliceValue.Elem().Type().Elem()

	for rows.Next() {
		// Create a new instance of the slice element type
		elem := reflect.New(sliceElemType).Interface()
		err = rows.Scan(elem)
		if err != nil {
			return err
		}

		// Append the scanned element to the destination slice
		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), reflect.ValueOf(elem).Elem()))
	}

	return nil
}

func (db *DB) SelectAllForStruct(query string, condition []interface{}, dest interface{}, fields []string) error {
	rows, err := db.instanceOfDB.Query(query, condition...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Get the type of the destination struct
	typeDest := reflect.TypeOf(dest).Elem().Elem()

	for rows.Next() {
		// Create a new instance of the destination struct
		entrie := reflect.New(typeDest).Elem()

		// Create a slice to hold field values
		values := make([]interface{}, len(fields))

		// Map field names to their indexes
		fieldIndex := make(map[string]int)
		for i, fieldName := range fields {
			fieldIndex[fieldName] = i
		}

		// Iterate over struct fields and set their pointers
		for i := 0; i < typeDest.NumField(); i++ {
			field := typeDest.Field(i)
			fieldName := field.Name
			fieldPtr := entrie.Field(i).Addr().Interface()
			if index, ok := fieldIndex[fieldName]; ok {
				values[index] = fieldPtr
			}
		}

		// Scan the row into the pointers of the struct fields
		if err := rows.Scan(values...); err != nil {
			return err
		}

		// Append the populated struct to the destination slice
		reflect.ValueOf(dest).Elem().Set(reflect.Append(reflect.ValueOf(dest).Elem(), entrie))
	}

	return nil
}

func (db *DB) Update(query string, data ...interface{}) (int64, error) {
	result, err := db.instanceOfDB.Exec(query, data...)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DB) Delete(query string, data ...interface{}) error {
	_, err := db.instanceOfDB.Exec(query, data...)
	return err
}
