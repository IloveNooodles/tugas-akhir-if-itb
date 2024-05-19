package auth

import (
	"net/http"
	"time"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	JWTSigningMethod  = jwt.SigningMethodHS256
	LoginExpiration   = time.Duration(1) * time.Hour
	RefreshExpiration = time.Duration(24) * time.Hour
	cfg               = config.Config{}
)

const (
	Authentication string = "Authentication"
)

type MyClaims struct {
	UserID    uuid.UUID `json:"id"`
	CompanyID uuid.UUID `json:"company_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	MyClaims
}

func init() {
	cfg = config.New()
}

func createClaims(mc MyClaims, subject string, expDate time.Duration) JwtClaims {
	j := JwtClaims{}

	j.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "Deployment Manager",
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDate)),
	}

	j.UserID = mc.UserID
	j.Email = mc.Email
	j.Name = mc.Name
	j.CompanyID = mc.CompanyID
	return j
}

func jwtSecretHelper(t *jwt.Token) (interface{}, error) {
	return []byte(cfg.JWTKey), nil
}

// Given claims and subject it will create the access and refresh token
func CreateAndSignToken(mc MyClaims, subject string) (string, string, error) {
	usedSecret := []byte(cfg.JWTKey)
	claims := createClaims(mc, subject, LoginExpiration)
	token := jwt.NewWithClaims(JWTSigningMethod, claims)
	signedToken, err := token.SignedString(usedSecret)
	if err != nil {
		return "", "", err
	}

	rtClaims := createClaims(mc, subject, RefreshExpiration)
	rtToken := jwt.NewWithClaims(JWTSigningMethod, rtClaims)
	refreshToken, err := rtToken.SignedString(usedSecret)
	if err != nil {
		return "", "", err
	}

	return signedToken, refreshToken, nil
}

// This function will be called by the create and sign token.
func GeneratePairToken(claims JwtClaims) (string, string, error) {
	subject := claims.Subject
	mc := MyClaims{
		UserID:    claims.UserID,
		CompanyID: claims.CompanyID,
		Email:     claims.Email,
		Name:      claims.Name,
	}

	return CreateAndSignToken(mc, subject)
}

// Validate the given JWTStringToken and return the parsed version of it.
func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{}, jwtSecretHelper)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errx.ErrInvalidToken
	}

	return token, nil
}

// Create http cookie based on key, val and age
func CreateCookie(key, val string, age int) *http.Cookie {
	return &http.Cookie{
		Name:     key,
		Value:    val,
		MaxAge:   age,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
