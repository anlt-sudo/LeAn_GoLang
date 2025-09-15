package service

import (
	"errors"
	"go-shop-api/internal/model"
	"reflect"
	"testing"
)

type mockCategoryRepo struct {
	categories []model.Category
	findByID   *model.Category
	findByIDErr error
	createErr  error
	updateErr  error
	deleteErr  error
	searchRes  []model.Category
	searchErr  error
}

func (m *mockCategoryRepo) FindAll() ([]model.Category, error) {
	return m.categories, nil
}
func (m *mockCategoryRepo) FindByID(id uint) (*model.Category, error) {
	if m.findByIDErr != nil {
		return nil, m.findByIDErr
	}
	return m.findByID, nil
}
func (m *mockCategoryRepo) Create(c *model.Category) error {
	return m.createErr
}
func (m *mockCategoryRepo) Update(c *model.Category) error {
	return m.updateErr
}
func (m *mockCategoryRepo) Delete(id uint) error {
	return m.deleteErr
}
func (m *mockCategoryRepo) SearchByName(keyword string) ([]model.Category, error) {
	return m.searchRes, m.searchErr
}

func TestGetAll(t *testing.T) {
	repo := &mockCategoryRepo{categories: []model.Category{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}}
	svc := NewCategoryService(repo)

	cats, err := svc.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(cats, repo.categories) {
		t.Errorf("expected %v, got %v", repo.categories, cats)
	}
}

func TestGetByID(t *testing.T) {
	repo := &mockCategoryRepo{findByID: &model.Category{ID: 1, Name: "A"}}
	svc := NewCategoryService(repo)
	cat, err := svc.GetByID(1)
	if err != nil || cat.ID != 1 {
		t.Errorf("expected category, got %v, err %v", cat, err)
	}

	repo.findByIDErr = errors.New("not found")
	_, err = svc.GetByID(2)
	if err == nil {
		t.Error("expected error for not found")
	}
}

func TestCreate(t *testing.T) {
	repo := &mockCategoryRepo{}
	svc := NewCategoryService(repo)
	err := svc.Create(&model.Category{ID: 1, Name: "A"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	repo.createErr = errors.New("fail")
	err = svc.Create(&model.Category{ID: 2, Name: "B"})
	if err == nil {
		t.Error("expected error for create fail")
	}
}

func TestUpdate(t *testing.T) {
	repo := &mockCategoryRepo{}
	svc := NewCategoryService(repo)
	err := svc.Update(&model.Category{ID: 1, Name: "A"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	repo.updateErr = errors.New("fail")
	err = svc.Update(&model.Category{ID: 2, Name: "B"})
	if err == nil {
		t.Error("expected error for update fail")
	}
}

func TestDelete(t *testing.T) {
	repo := &mockCategoryRepo{}
	svc := NewCategoryService(repo)
	err := svc.Delete(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	repo.deleteErr = errors.New("fail")
	err = svc.Delete(2)
	if err == nil {
		t.Error("expected error for delete fail")
	}
}

func TestSearchByName(t *testing.T) {
	repo := &mockCategoryRepo{searchRes: []model.Category{{ID: 1, Name: "A"}}}
	svc := NewCategoryService(repo)
	cats, err := svc.SearchByName("A")
	if err != nil || len(cats) != 1 {
		t.Errorf("expected 1 result, got %v, err %v", cats, err)
	}

	repo.searchErr = errors.New("fail")
	_, err = svc.SearchByName("B")
	if err == nil {
		t.Error("expected error for search fail")
	}
}
