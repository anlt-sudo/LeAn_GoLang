package main

import "fmt"

type ContactInfo struct {
    Email string
    Phone string
}

type Employee struct {
    Name     string
    Position string
    Contact  ContactInfo
}

func main() {
    emp := Employee{
        Name:     "Alice",
        Position: "Software Engineer",
        Contact: ContactInfo{
            Email: "alice@example.com",
            Phone: "123-456-7890",
        },
    }

    fmt.Println("--- Thông tin ban đầu của nhân viên ---")
    fmt.Printf("Tên: %s\n", emp.Name)
    fmt.Printf("Chức vụ: %s\n", emp.Position)
    fmt.Printf("Email: %s\n", emp.Contact.Email)
    fmt.Printf("SĐT: %s\n", emp.Contact.Phone)

    fmt.Println("\n--- Cập nhật thông tin ---")

    emp.Position = "Senior Software Engineer"

    emp.Contact.Email = "alice.s@newcorp.com"

    fmt.Printf("Chức vụ mới: %s\n", emp.Position)
    fmt.Printf("Email mới: %s\n", emp.Contact.Email)
}