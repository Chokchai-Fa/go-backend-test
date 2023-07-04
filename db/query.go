package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "Phukao98765"
	dbname   = "svc_backend_team"
)

type InsertNewUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
}

func openDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateNewUser(arg InsertNewUserParams) (*User, error) {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	query1 := "INSERT INTO users(name,email,tel) VALUES($1,$2,$3)"
	rows1, err := db.Query(query1, arg.Name, arg.Email, arg.Tel)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows1.Close()

	user := &User{}
	query2 := "SELECT * FROM users WHERE name = $1"
	rows2, err := db.Query(query2, arg.Name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows2.Next() {
		err = rows2.Scan(&user.ID, &user.Name, &user.Email, &user.Tel)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, sql.ErrNoRows
	}

	return user, nil
}

type GetUserByIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetUserById(ctx context.Context, arg GetUserByIdParams) (*User, error) {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	user := &User{}

	query := "SELECT * FROM users  WHERE id = $1"
	rows, err := db.QueryContext(ctx, query, arg.ID)
	fmt.Print(err)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Tel)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, sql.ErrNoRows
	}

	defer rows.Close()

	return user, nil

}

func GetAllUser(ctx context.Context) ([]User, error) {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	userList := []User{}

	query := "SELECT * FROM users ORDER BY id ASC"
	rows, err := db.QueryContext(ctx, query)
	fmt.Print(err)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Tel)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	return userList, nil

}

type UpdateAllDetailParams struct {
	ID    int64
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
}

func UpdateAllDetail(ctx context.Context, arg UpdateAllDetailParams) (*User, error) {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	query1 := "UPDATE users SET name = $2, email = $3, tel = $4 WHERE id = $1"
	rows1, err := db.QueryContext(ctx, query1, arg.ID, arg.Name, arg.Email, arg.Tel)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows1.Close()

	user := &User{}

	query2 := "SELECT * FROM users WHERE id = $1"
	rows2, err := db.QueryContext(ctx, query2, arg.ID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows2.Next() {
		err = rows2.Scan(&user.ID, &user.Name, &user.Email, &user.Tel)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, sql.ErrNoRows
	}

	defer rows2.Close()

	return user, nil
}

type UpdateSomeDetailParams struct {
	ID    int64
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
}

func UpdateSomeDetail(ctx context.Context, arg UpdateSomeDetailParams) (*User, error) {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	query1 := "UPDATE users SET"
	queryParams := []interface{}{}
	paramCount := 2 // Starting index for the next parameter
	queryParams = append(queryParams, arg.ID)

	if arg.Name != "" {
		query1 += " name = $" + strconv.Itoa(paramCount) + ","
		queryParams = append(queryParams, arg.Name)
		paramCount++
	}

	if arg.Email != "" {
		query1 += " email = $" + strconv.Itoa(paramCount) + ","
		queryParams = append(queryParams, arg.Email)
		paramCount++
	}

	if arg.Tel != "" {
		query1 += " tel = $" + strconv.Itoa(paramCount) + ","
		queryParams = append(queryParams, arg.Tel)
		paramCount++
	}

	query1 = strings.TrimSuffix(query1, ",") + " WHERE id = $1"

	fmt.Print(query1)

	rows1, err := db.QueryContext(ctx, query1, queryParams...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows1.Close()

	user := &User{}

	query2 := "SELECT * FROM users WHERE id = $1"
	rows2, err := db.QueryContext(ctx, query2, arg.ID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows2.Next() {
		err = rows2.Scan(&user.ID, &user.Name, &user.Email, &user.Tel)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, sql.ErrNoRows
	}

	defer rows2.Close()

	return user, nil

}

type DeleteUserbyIdParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func DeleteUserbyId(ctx context.Context, arg DeleteUserbyIdParams) error {
	db, err := openDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	// check id exist
	query1 := "SELECT * FROM users WHERE id = $1"
	rows1, err := db.QueryContext(ctx, query1, arg.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rows1.Next() {
		if err != nil {
			return err
		}

	} else {
		return sql.ErrNoRows
	}

	defer rows1.Close()

	// if exist -> delete it
	query2 := "DELETE FROM users WHERE id = $1"
	rows2, err := db.QueryContext(ctx, query2, arg.ID)
	fmt.Print(err)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows2.Close()

	return nil
}
