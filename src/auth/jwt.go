package auth

import (
	"ApiCookMaster/src/models"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// Structure JWTManager
type JWTManager interface {
	Generate(user models.User) (string, error)
	GenerateToken(userID string, expiresIn int64) (string, error)
	VerifyToken(token string) (string, error)
}

// Structure JWT
type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  []byte
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyPath := "/app/keys/encrypted_private_key.pem"
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf("Error reading private key file: %v", err)
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	decryptedBytes, err := x509.DecryptPEMBlock(block, []byte("dana")) // Remplacez "YourPassphraseHere" par votre passphrase réelle
	if err != nil {
		log.Printf("Error decrypting PEM block: %v", err)
		return nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedBytes)
	if err != nil {
		log.Printf("Error parsing private key: %v", err)
		return nil, err
	}

	log.Println("Private key loaded successfully")

	return privateKey, nil
}

// Méthode pour charger la clé publique depuis le fichier
func loadPublicKey() ([]byte, error) {
	publicKeyPath := "/app/keys/public_key.pem"
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

// Méthode pour créer une instance de JWTManager
func NewJWT() JWTManager {
	privateKey, _ := loadPrivateKey()
	publicKey, _ := loadPublicKey()

	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// Méthode pour générer le token JWT
func (j *JWT) Generate(user models.User) (string, error) {
	return j.GenerateToken(strconv.FormatInt(int64(user.ID), 10), 3600)
}

// Méthode pour générer le token JWT avec un ID et une durée d'expiration
func (j *JWT) GenerateToken(userID string, expiresIn int64) (string, error) {
	// Construction du token
	builder := jwt.New()

	// Ajout des informations au token
	builder.Set("userID", userID)
	builder.Set(jwt.ExpirationKey, time.Now().Add(time.Duration(expiresIn)*time.Second))

	// Vérification de la clé privée
	if j.privateKey == nil {
		log.Println("Private key is nil")
		return "", errors.New("private key is nil")
	}

	// Signature du token avec la clé privée
	signedToken, err := jwt.Sign(builder, jwa.RS256, j.privateKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	log.Printf("Generated token: %v", string(signedToken))

	return string(signedToken), nil
}

// Méthode pour vérifier la validité du token et extraire l'ID utilisateur
func (j *JWT) VerifyToken(token string) (string, error) {
	// Construction de la clé publique pour vérifier la signature du token
	pubKeySet, err := jwk.Parse(j.publicKey)
	if err != nil {
		log.Printf("Error parsing public key: %v", err)
		return "", err
	}

	// Vérification de la validité du token et extraction des informations
	parsedToken, err := jwt.Parse([]byte(token), jwt.WithKeySet(pubKeySet))
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return "", err
	}

	log.Printf("Parsed token: %v", parsedToken)

	// Extraction de l'ID utilisateur du token
	val, ok := parsedToken.Get("userID")
	if !ok {
		log.Print("User ID not found in token")
		return "", errors.New("user ID not found in token")
	}

	userID, ok := val.(string)
	if !ok {
		log.Print("User ID is not a string")
		return "", errors.New("user ID is not a string")
	}

	return userID, nil
}
