package authservice

import (
	"context"
	"fmt"
	"time"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/dgrijalva/jwt-go"
)

func (s *AuthService) createAndSaveRefreshToken(
	ctx context.Context,
	userID int,
	email string,
	identifyData entities.UserIdentifyData,
) (string, error) {
	errorBase := fmt.Sprintf(
		"userservice.GenerateAndSaveRefreshToken(user_id: %d, email: %s, user_agent: %s, finger_print: %s, ip: %s)",
		userID, email, identifyData.UserAgent, identifyData.FingerPrint, identifyData.IP)

	signedToken, err := s.createAndSignedToken(userID, email, jwtExpireRefreshTime)
	if err != nil {
		return "", fmt.Errorf("%s.createAndSignedToken.error: %w", errorBase, err)
	}

	refreshToken := entities.RefreshToken{
		UserID:      userID,
		Token:       signedToken,
		UserAgent:   identifyData.UserAgent,
		IP:          identifyData.IP,
		FingerPrint: identifyData.FingerPrint,
	}

	if err := s.tokenStorage.CreateRefreshToken(ctx, refreshToken); err != nil {
		return "", fmt.Errorf("%s.CreateRefreshToken.error: %w", errorBase, err)
	}

	return signedToken, nil
}

func (s *AuthService) createAndSignedToken(userID int, email string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:  email,
		UserID: userID,
	})

	signedToken, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token error: %w", err)
	}

	return signedToken, nil
}