package api

import (
	"api-server/internal/authentication"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignInRequest struct {
	WalletAddress string `json:"walletAddress"`
	Nonce         string `json:"nonce"`
}

type GetNonceResponse struct {
	Nonce string `json:"nonce"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func GetNonceHandler(w http.ResponseWriter, r *http.Request) {
	authChecker := authentication.NewAuthChecker()

	// Generate a random nonce
	nonce, err := authChecker.GenerateNonce()
	if err != nil {
		http.Error(w, "Failed to generate nonce", http.StatusInternalServerError)
		return
	}

	resp := GetNonceResponse{Nonce: nonce}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Verify the signature of the provided nonce
	authChecker := authentication.NewAuthChecker()
	valid, err := authChecker.VerifySignature(req.WalletAddress, req.Nonce, r.Header.Get("Signature"))
	if err != nil || !valid {
		http.Error(w, "Signature verification failed", http.StatusUnauthorized)
		return
	}

	// Generate a token
	claims := authentication.TokenClaims{
		Subject: req.WalletAddress,
		Nonce:   req.Nonce,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		},
	}
	token, err := authChecker.GenerateToken(claims)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	resp := SignInResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
