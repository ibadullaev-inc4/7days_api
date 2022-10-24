package store

import "restAPI/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (user *model.User, err error) {
	if err := r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1,$2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FundByEmail(email string) (*model.User, error) {
	return nil, nil
}
