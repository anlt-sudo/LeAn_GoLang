package auth

// Simple auth demo (do NOT use in production)
func Authenticate(email, password string) bool {
    // demo: accept any password length >= 6 for emails containing '@'
    if len(password) >= 6 && len(email) > 3 {
        return true
    }
    return false
}
