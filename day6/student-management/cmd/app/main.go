package main

import (
	service "github.com/anlt-sudo/student-management/internal/services"
	"github.com/anlt-sudo/student-management/pkg/ui"
)

func main() {
	school := service.NewSchoolService()
	ui.ShowMenu(school)
}
