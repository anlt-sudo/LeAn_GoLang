package database

import "fmt"

// DB is a tiny in-memory demo "connection"
type DB struct {
    connected bool
}

func Connect() *DB {
    fmt.Println("Connecting to demo DB...")
    return &DB{connected: true}
}


func (d *DB) Close() {
    d.connected = false
    fmt.Println("Demo DB closed.")
}

// IsConnected returns true if the DB is connected
func (d *DB) IsConnected() bool {
    return d.connected
}
