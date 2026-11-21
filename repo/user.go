package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)


type UserRepo interface {
    user.UserRepo
}


type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}


func (r *userRepo) Create(user domain.User) (*domain.User, error) {
    query := `
        INSERT INTO users (
            first_name,
            last_name,
            email,
            password,
            is_shop_owner
        )
        VALUES (
            :first_name,
            :last_name,
            :email,
            :password,
            :is_shop_owner
        )
        RETURNING id
    `

    rows, err := r.db.NamedQuery(query, user)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    if !rows.Next() {
        if err := rows.Err(); err != nil {
            return nil, err
        }
        return nil, fmt.Errorf("failed to create user: no ID returned")
    }

    if err := rows.Scan(&user.ID); err != nil {
        return nil, err
    }

    return &user, nil
}



func (r *userRepo) Find(email, pass string) (*domain.User, error) {
    var user domain.User

    query := `
        SELECT
            id,
            first_name,
            last_name,
            email,
            password,
            is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
        LIMIT 1
    `

    err := r.db.Get(&user, query, email, pass)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &user, nil
}

