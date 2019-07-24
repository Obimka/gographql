package mutations

const (
	ADD_USER = `INSERT INTO users (firstname, lastname) VALUES ($1, $2)`
)
