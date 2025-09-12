package service_test

import (
	"testing"
	"time"

	"go-shop-api/internal/service"
)

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
