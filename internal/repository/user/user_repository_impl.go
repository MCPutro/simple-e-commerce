package user

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type userRepository struct {
}

func NewUserRepository() Repository {
	return &userRepository{}
}

func (r *userRepository) Write(cxt context.Context, tx *sql.Tx, user *domain.User) error {
	query := "INSERT INTO e_commerce.users (id,name,email,password,`role`) VALUES (?, ?, ?, ?, ?);"

	result, err := tx.ExecContext(cxt, query, user.Id, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}

	nRow, err := result.RowsAffected()
	if nRow == 0 {
		return err
	}

	// check address user
	lenAddress := len(user.Address)
	if len(user.Address) > 0 {
		for i := 0; i < lenAddress; i++ {
			query1 := "INSERT INTO e_commerce.user_addresses (user_id,seq,address,city,postal_code) VALUES (?, ?, ?, ?, ?);"
			_, err = tx.ExecContext(cxt, query1, user.Id, (i + 1), user.Address[i].Address, user.Address[i].City, user.Address[i].PostalCode)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *userRepository) ReadById(cxt context.Context, tx *sql.Tx, id string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT u.id, u.name, u.email, u.password, u.role, 
		       ua.address, ua.city, ua.postal_code 
		FROM e_commerce.users u 
		LEFT JOIN e_commerce.user_addresses ua ON u.id = ua.user_id 
		WHERE u.id = ?`

	rows, err := tx.QueryContext(cxt, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := []domain.UserAddress{}
	for rows.Next() {
		var address domain.UserAddress
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &address.Address, &address.City, &address.PostalCode); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	user.Address = addresses
	return &user, nil
}

func (r *userRepository) ReadByEmail(cxt context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT u.id, u.name, u.email, u.password, u.role, 
		       ua.address, ua.city, ua.postal_code 
		FROM e_commerce.users u 
		LEFT JOIN e_commerce.user_addresses ua ON u.id = ua.user_id 
		WHERE u.email = ?`

	rows, err := tx.QueryContext(cxt, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := []domain.UserAddress{}
	for rows.Next() {
		var address domain.UserAddress
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &address.Address, &address.City, &address.PostalCode); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	user.Address = addresses
	return &user, nil
}

func (r *userRepository) ReadAll(cxt context.Context, tx *sql.Tx) ([]domain.User, error) {
	query := "SELECT id, name, email, password, role FROM e_commerce.users"
	rows, err := tx.QueryContext(cxt, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Update(cxt context.Context, tx *sql.Tx, user *domain.User) error {
	query := "UPDATE e_commerce.users SET name = ?, email = ?, password = ?, role = ? WHERE id = ?"
	_, err := tx.ExecContext(cxt, query, user.Name, user.Email, user.Password, user.Role, user.Id)
	return err
}

func (r *userRepository) Delete(cxt context.Context, tx *sql.Tx, id string) error {
	query := "DELETE FROM e_commerce.users WHERE id = ?"
	_, err := tx.ExecContext(cxt, query, id)
	return err
}
