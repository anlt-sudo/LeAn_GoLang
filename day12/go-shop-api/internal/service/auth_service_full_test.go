package service_test

import (
	"errors"
	"testing"
	"time"

	"go-shop-api/internal/model"
	"go-shop-api/internal/service"
)

type mockUserRepo struct {
	users map[string]*model.User
	byID  map[uint]*model.User
	createErr error
}

func (m *mockUserRepo) FindByEmail(email string) (*model.User, error) {
	u, ok := m.users[email]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}
func (m *mockUserRepo) FindByID(id uint) (*model.User, error) {
	u, ok := m.byID[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}
func (m *mockUserRepo) Create(u *model.User) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.users[u.Email] = u
	m.byID[u.ID] = u
	return nil
}

type mockRefreshRepo struct {
	tokens map[string]*model.RefreshToken
	revokeErr error
	saveErr   error
}

func (m *mockRefreshRepo) Save(rt *model.RefreshToken) error {
	if m.saveErr != nil {
		return m.saveErr
	}
	m.tokens[rt.TokenHash] = rt
	return nil
}
func (m *mockRefreshRepo) FindValid(tokenHash string) (*model.RefreshToken, error) {
	rt, ok := m.tokens[tokenHash]
	if !ok || rt.Revoked || rt.ExpiredAt.Before(time.Now()) {
		return nil, errors.New("not found")
	}
	return rt, nil
}
func (m *mockRefreshRepo) RevokeByUserID(userID uint) error {
	if m.revokeErr != nil {
		return m.revokeErr
	}
	for _, rt := range m.tokens {
		if rt.UserID == userID {
			rt.Revoked = true
		}
	}
	return nil
}

func TestRegisterAndAuthenticate(t *testing.T) {
	userRepo := &mockUserRepo{users: map[string]*model.User{}, byID: map[uint]*model.User{}}
	refreshRepo := &mockRefreshRepo{tokens: map[string]*model.RefreshToken{}}
	auth := service.NewAuthService(userRepo, refreshRepo)

	u, err := auth.Register("test@example.com", "pass123", "user")
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if u.Email != "test@example.com" || u.Role != "user" {
		t.Errorf("register wrong data: %+v", u)
	}

	_, err = auth.Register("test@example.com", "pass123", "user")
	if err == nil {
		t.Error("expected error for duplicate email")
	}


	access, refresh, user, err := auth.Authenticate("test@example.com", "pass123")
	if err != nil {
		t.Fatalf("auth failed: %v", err)
	}
	if access == "" || refresh == "" || user == nil {
		t.Error("expected tokens and user")
	}

	_, _, _, err = auth.Authenticate("test@example.com", "wrong")
	if err == nil {
		t.Error("expected error for wrong password")
	}
}

func TestRefreshTokenAndLogout(t *testing.T) {
	userRepo := &mockUserRepo{users: map[string]*model.User{}, byID: map[uint]*model.User{}}
	refreshRepo := &mockRefreshRepo{tokens: map[string]*model.RefreshToken{}}
	auth := service.NewAuthService(userRepo, refreshRepo)

	u, _ := auth.Register("a@b.com", "pw", "user")
	_, refresh, _, _ := auth.Authenticate("a@b.com", "pw")

	access2, refresh2, err := auth.RefreshToken(refresh)
	if err != nil {
		t.Fatalf("refresh failed: %v", err)
	}
	if access2 == "" || refresh2 == "" {
		t.Error("expected new tokens")
	}

	err = auth.Logout(u.ID)
	if err != nil {
		t.Errorf("logout failed: %v", err)
	}
	for _, rt := range refreshRepo.tokens {
		if rt.UserID == u.ID && !rt.Revoked {
			t.Error("token not revoked after logout")
		}
	}
}

func TestHashAndComparePassword(t *testing.T) {
	auth := service.NewAuthService(nil, nil)

	hash, err := auth.HashPassword("secret123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := auth.ComparePassword(hash, "secret123"); err != nil {
		t.Errorf("expected password to match, got error: %v", err)
	}

	if err := auth.ComparePassword(hash, "wrongpass"); err == nil {
		t.Errorf("expected error for wrong password, got nil")
	}
}

func TestAccessTokenTTL(t *testing.T) {
	auth := service.NewAuthService(nil, nil)
	if auth.AccessTokenTTL != time.Minute*60 {
		t.Errorf("expected default AccessTokenTTL = 60m, got %v", auth.AccessTokenTTL)
	}
}
