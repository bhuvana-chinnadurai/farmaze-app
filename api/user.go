package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

var jwtKey = []byte("your_secret_key") // Keep this key secret and safe.

// Claims struct will add custom claims which extend jwt.StandardClaims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token for a authenticated user.
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Token is valid for 1 hour
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Users holds dependencies for user-related routes and functionalities.
type Users struct {
	repo *model.UserRepository
}

// NewUsers creates a new Users service with the necessary dependencies.
func NewUsers(db *sql.DB) *Users {
	return &Users{
		repo: model.NewUserRepository(db),
	}
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user with a username, password, and role
// @Tags users
// @Accept json
// @Produce json
// @Param request body registerRequest true "Registration Request"
// @Success 201 {object} map[string]string "Success: User successfully registered"
// @Failure 400 "Bad Request: Cannot parse request"
// @Failure 500 "Internal Server Error: Error while creating a new user"
// @Router /register [post]
// RegisterUser handles user registration.
func (u *Users) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := model.NewUser(req.Username, req.Password, req.Role)
	if err != nil {
		http.Error(w, "error while creating a new user", http.StatusInternalServerError)
		return
	}
	// Save the user using UserRepository through the Users struct
	err = u.repo.Save(user)
	if err != nil {
		http.Error(w, "Failed to save user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User successfully registered"})
}

// LoginUser godoc
// @Summary User login
// @Description Logs in a user with the provided username and password, returning a JWT token upon success
// @Tags users
// @Accept json
// @Produce json
// @Param request body loginRequest true "Login Request"
// @Success 200 {object} map[string]string "Success: Token generated and returned"
// @Failure 400 "Bad Request: Cannot parse request"
// @Failure 401 "Unauthorized: Invalid username or password"
// @Failure 500 "Internal Server Error: Failed to generate token"
// @Router /login [post]
// LoginUser handles user login, validating credentials and generating a session/token.
func (u *Users) LoginUser(w http.ResponseWriter, r *http.Request) {

	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user by username
	user, err := u.repo.FindByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// If the password does not match
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// In your LoginUser method, after validating the password:
	// If the password matches, generate a token
	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
