package main

import (
	"fmt"
	"log"

	"github.com/anlt-sudo/my-project/internal/auth"
	"github.com/anlt-sudo/my-project/internal/user"
	"github.com/anlt-sudo/my-project/pkg/database"
)

func main() {
    // load config (simple static demo)
    cfg := map[string]string{"port": "8080"}

    // connect to "database"
    db := database.Connect()
    defer db.Close()

    // create user sample
    u := user.NewUser("alice@example.com", "Alice")
    created := user.Save(db, u)
    fmt.Printf("Created user? %v\n", created)

    // auth demo
    ok := auth.Authenticate("alice@example.com", "password123")
    log.Printf("Authenticate alice -> %v (config port=%s)\n", ok, cfg["port"])
}
