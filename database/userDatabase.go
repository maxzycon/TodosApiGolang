package database

type User struct {
	Username string `gorm:"primaryKey;unique;not null;type:varchar;size:255"`
	Hash_password string
}