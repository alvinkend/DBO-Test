package usecase

import (
	"log"
	"time"

	"pustaka-api/model/mysqlmodels"
	"pustaka-api/repository"
	"pustaka-api/request"
	"pustaka-api/responses"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	LoginUsecase(req request.LoginRequest) (interface{}, error)
}

type authUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewAuthUsecase(customerRepository repository.CustomerRepository) AuthUsecase {
	return &authUsecase{customerRepository: customerRepository}
}

func (srv *authUsecase) LoginUsecase(req request.LoginRequest) (interface{}, error) {
	var customer *mysqlmodels.Customer
	criteria := make(map[string]interface{})
	criteria["email"] = req.Email
	customer, err := srv.customerRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	if !srv.comparePasswords(customer.Password, req.Password) {
		return nil, err
	}

	// return JWT token
	return srv.getToken(customer)
}

func (srv *authUsecase) comparePasswords(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	
	return err == nil
}

func (srv *authUsecase) getToken(customer *mysqlmodels.Customer) (responses.TokenResponse, error) {
	role := "customer"

	structure := request.TokenStructure{
		CustomerID: customer.ID,
		Email:  customer.Email,
		Role:   role,
	}
	signature := "DBO123"

	resp, err := srv.generateToken(structure, signature)
	return resp, err
}

func (srv *authUsecase) generateToken(data request.TokenStructure, signature string) (response responses.TokenResponse, err error) {
	
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	log.Println(signature)
	expiredIn := time.Hour * (24 * 7)
	expiredAt := time.Now().Add(time.Hour * (24 * 7))

	myCrypt, err := bcrypt.GenerateFromPassword([]byte(signature), 8)
	if err != nil {
		return
	}

	claims["user_id"] = data.CustomerID
	claims["email"] = data.Email
	claims["role"] = data.Role
	claims["hash"] = string(myCrypt)
	claims["exp"] = expiredIn

	byteSig := []byte(signature)
	tokenString, err := token.SignedString(byteSig)
	if err != nil {
		log.Println(err)
	}

	response.AccessToken = tokenString
	response.TokenType = "Bearer"
	response.Role = data.Role
	response.ExpiredAt = expiredAt.Unix()
	response.ExpiredIn = expiredIn.Seconds()

	return 
}