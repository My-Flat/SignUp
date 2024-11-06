package models

type User struct {
	ID       string `firestore:"id,omitempty" json:"id"`
	Email    string `firestore:"email" json:"email"`
	Password string `firestore:"password" json:"password"`
	// ... other fields
}
