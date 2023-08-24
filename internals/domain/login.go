package domain

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Username string `json:"username"`
	PID      int8   `json:"pID"`
	AID      int8   `json:"aID"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	DaID     int8   `json:"daID"`
	UID      int8   `json:"uID"`
	CID      int8   `json:"cID"`
}
