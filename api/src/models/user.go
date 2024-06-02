package models

type User struct {
	ID        uint64 `db:"id,omitempty" json:"id"`
	Name      string `db:"name" json:"name"`
	Nick      string `db:"nick" json:"nick"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"createdAt" json:"createdAt"`
}
