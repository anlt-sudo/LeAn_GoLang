package user

import "github.com/anlt-sudo/my-project/pkg/database"

// Save simulates saving user into DB, returns true on success
func Save(db *database.DB, u *User) bool {
    if db == nil || !db.IsConnected() {
        return false
    }
    // demo: always succeed
    return true
}
