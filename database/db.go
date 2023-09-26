package database

import (
	"api/api/entities"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize the database connection
func InitDatabase() {
	connStr := "host=localhost port=5432 user=postgres password=2198 dbname=user_management sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
	} else {
		fmt.Println("Database connection established.")
	}

}

func CreateUser(user *entities.User) error {
	_, err := db.Exec("INSERT INTO users (id, name, surname, endereco, celular) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Name, user.Surname, user.Endereco, user.NumeroCelular)
	if err != nil {
		fmt.Printf("Error inserting user into the database: %v\n", err)
		return err
	}

	fmt.Println("Should've worked", user.ID)
	return nil
}

func GetUser(userID string) (*entities.User, error) {
	var user entities.User
	err := db.QueryRow("SELECT id, name, surname, endereco, celular FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Name, &user.Surname, &user.Endereco, &user.NumeroCelular)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Here", userID)
			return nil, nil // User not found
		}
		fmt.Println("Here aaaa")
		return nil, err // Other error
	}

	fmt.Println(&user)

	return &user, nil
}

func UpdateUser(userID string, user *entities.User) (*entities.User, error) {
	_, err := db.Exec("UPDATE users SET name = $1, surname = $2, endereco = $3, celular = $4 WHERE id = $5",
		user.Name, user.Surname, user.Endereco, user.NumeroCelular, userID)
	if err != nil {
		return nil, err
	}

	// Return the updated user after a successful update
	return user, nil
}

func DeleteUser(userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}

// GetAllUsers retrieves all user records from the database and returns them as a slice of User entities.
func GetAllUsers() ([]entities.User, error) {
	var users []entities.User

	rows, err := db.Query("SELECT id, name, surname, endereco, celular FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Endereco, &user.NumeroCelular)
		if err != nil {
			return nil, err
		}

		// Append the retrieved user to the users slice
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func CreateAccount(account *entities.Account) error {
	_, err := db.Exec("INSERT INTO accounts (numeroconta, saldo, tipoconta, donoid) VALUES ($1, $2, $3, $4)", account.NumeroConta, account.Saldo, account.TipoConta, account.Dono.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAccount(numeroConta string) (*entities.Account, error) {
	var account entities.Account

	row := db.QueryRow("SELECT numeroconta, saldo, tipoconta FROM accounts WHERE numeroconta = $1", numeroConta)
	err := row.Scan(&account.NumeroConta, &account.Saldo, &account.TipoConta)

	if err == sql.ErrNoRows {
		return nil, errors.New("Account not found")
	} else if err != nil {
		return nil, err
	}

	// Initialize the Dono field with an empty user
	account.Dono = &entities.User{}

	return &account, nil
}

func UpdateAccount(numeroConta string, saldo, tipoConta string) error {
	_, err := db.Exec("UPDATE accounts SET saldo = $1, tipoconta = $2 WHERE numeroconta = $3", saldo, tipoConta, numeroConta)
	return err
}

func DeleteAccount(numeroConta string) error {
	_, err := db.Exec("DELETE FROM accounts WHERE numeroconta = $1", numeroConta)
	return err
}

func GetAllAccounts() ([]*entities.Account, error) {
	var accounts []*entities.Account

	rows, err := db.Query("SELECT numeroconta, saldo, tipoconta FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account entities.Account
		err := rows.Scan(&account.NumeroConta, &account.Saldo, &account.TipoConta)
		if err != nil {
			return nil, err
		}

		// Append the retrieved account to the accounts slice
		accounts = append(accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
