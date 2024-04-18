package auth

import (
	"time"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	JWTSigningMethod = jwt.SigningMethodHS256
	LoginExpiration  = time.Duration(24) * time.Hour
	cfg              = config.Config{}
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

func createClaims(mc MyClaims, subject string) JwtClaims {
	j := JwtClaims{}

	expDate := LoginExpiration
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

func CreateAndSignToken(mc MyClaims, subject string) (string, error) {
	usedSecret := []byte(cfg.JWTKey)
	claims := createClaims(mc, subject)
	token := jwt.NewWithClaims(JWTSigningMethod, claims)
	signedToken, err := token.SignedString(usedSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

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
