package service

import (
	"fmt"
	"strings"

	"github.com/anlt-sudo/student-management/internal/domain"
)

type SchoolService struct {
	classes   map[string]*domain.Class
	students  []*domain.Student
}

func NewSchoolService() *SchoolService {
	return &SchoolService{
		classes:  make(map[string]*domain.Class),
		students: []*domain.Student{},
	}
}

func (s *SchoolService) AddClass(name string) error {
	key := strings.ToLower(name)
	if _, exists := s.classes[key]; exists {
		return fmt.Errorf("❌ Lớp %s đã tồn tại", name)
	}
	s.classes[key] = &domain.Class{Name: name, NumStudents: 0}
	return nil
}

func (s *SchoolService) AddStudent(name string, className string) error {
	key := strings.ToLower(className)
	class, exists := s.classes[key]
	if !exists {
		return fmt.Errorf("❌ Lớp %s không tồn tại", className)
	}

	student := &domain.Student{Name: name, ClassName: class.Name}
	s.students = append(s.students, student)
	class.NumStudents++
	return nil
}

func (s *SchoolService) GetAllData() []*domain.Class {
	result := []*domain.Class{}
	for _, c := range s.classes {
		result = append(result, c)
	}
	return result
}

func (s *SchoolService) GetStudentsByClass(className string) ([]*domain.Student, error) {
	key := strings.ToLower(className)
	_, exists := s.classes[key]
	if !exists {
		return nil, fmt.Errorf("❌ Lớp %s không tồn tại", className)
	}

	var result []*domain.Student
	for _, st := range s.students {
		if strings.EqualFold(st.ClassName, className) {
			result = append(result, st)
		}
	}
	return result, nil
}
