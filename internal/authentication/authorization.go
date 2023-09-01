package authentication

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type AuthChecker struct {
	// You can add fields here if needed
}

func NewAuthChecker() *AuthChecker {
	// Initialize and return a new instance of AuthChecker
	return &AuthChecker{}
}

type TokenClaims struct {
	Subject string `json:"sub"`
	Nonce   string `json:"nonce"`
	jwt.StandardClaims
}

func (ac *AuthChecker) VerifySignature(walletAddress, nonce, signatureHeader string) (bool, error) {
	addressTy, _ := abi.NewType("address", "address", nil)
	stringTy, _ := abi.NewType("string", "string", nil)

	arguments := abi.Arguments{
		{
			Type: addressTy,
		},
		{
			Type: stringTy,
		},
	}

	// Equivalent to abi.encode
	bytes, err := arguments.Pack(
		common.HexToAddress(walletAddress),
		nonce,
	)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// Hash the bytes
	hashedBytes := crypto.Keccak256(bytes)

	// Cut the '0x' and return bytes
	signatureBytes, err := hex.DecodeString(signatureHeader[2:])
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	// fmt.Println(signatureBytes)

	// recovery id support
	if signatureBytes[len(signatureBytes)-1] == 27 || signatureBytes[len(signatureBytes)-1] == 28 {
		signatureBytes[len(signatureBytes)-1] -= 27
	}

	// Recover public key
	sigPublicKey, err := crypto.Ecrecover(hashedBytes, signatureBytes)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	pubKey, _ := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	// REcover ethereum address
	publicEtherAddress := crypto.PubkeyToAddress(*pubKey)

	if strings.ToLower(walletAddress) == strings.ToLower(publicEtherAddress.String()) {

		return true, nil
	}

	return false, errors.New("Invalid sign in request")
}

func (ac *AuthChecker) GenerateToken(claims TokenClaims) (string, error) {
	// Create a new JWT token with the provided claims
	// Sign the token with your secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	// Replace "your-secret-key" with your actual secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ac *AuthChecker) GenerateNonce() (string, error) {
	// // Generate a random nonce
	// nonceBytes := make([]byte, 32) // Adjust the nonce size as needed
	// _, err := rand.Read(nonceBytes)
	// if err != nil {
	// 	return "", err
	// }

	// // Encode the nonce as base64
	// nonce := base64.StdEncoding.EncodeToString(nonceBytes)
	// return nonce, nil

	// For easy testability just return this:
	return "xHL5/YhoDO18iq7MUlhKmocUlY8QXciMhOAp1K2RIJU=", nil

}

func (ac *AuthChecker) ValidateToken(tokenString string, nonce string) error {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // Replace with your actual secret key
	})
	if err != nil {
		return err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the nonce matches
		if nonceClaim, ok := claims["nonce"].(string); ok {
			if nonceClaim != nonce {
				return errors.New("Nonce mismatch")
			}
			return nil // Token and nonce are valid
		}
		return errors.New("Nonce not found in token claims")
	}

	return errors.New("Invalid token")
}
