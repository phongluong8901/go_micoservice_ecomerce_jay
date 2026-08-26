package helper

import (
	"errors"
	"fmt"
	"proj_1/internal/domain"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) < 6 {
		return "", errors.New("password length should be a least 6 charaters long")
	}

	//dùng để mã hóa (hash) một mật khẩu dạng văn bản thuần túy (plaintext) thành một chuỗi mã hóa an toàn bằng thuật toán Bcrypt.
	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		//log actual error and report to logging tool
		return "", errors.New("password hash failed")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs are missing to generate token")
	}

	//khởi tạo một JSON Web Token (JWT) mới kèm theo phương thức mã hóa và phần dữ liệu payload (claims) bên trong.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//dùng để ký (sign) đối tượng JWT đã tạo trước đó bằng một khóa bí mật (Secret Key) để biến nó thành một chuỗi token hoàn chỉnh (string) có thể gửi về cho client.
	tokenStr, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("unable to signed the token")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {
	if len(pP) < 6 {
		return errors.New("password length should be a least 6 charaters long")
	}

	//dùng để so sánh mật khẩu người dùng vừa nhập với mật khẩu đã mã hóa (hash) trong cơ sở dữ liệu xem có khớp nhau hay không
	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))
	if err != nil {
		return errors.New("Password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	//Bearer t4564t1asd...
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]
	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})
	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//dùng để kiểm tra xem token JWT đã hết hạn hay chưa.
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	//dùng để lấy giá trị của header Authorization từ HTTP request gửi lên bởi client. Dòng này thường được đặt trong các middleware xác thực (Authentication) để kiểm tra xem client có đính kèm token khi gọi API hay không.
	authHeader := ctx.GetReqHeaders()["Authorization"]
	user, err := a.VerifyToken(authHeader[0])
	if err == nil && user.ID > 0 {
		//dùng để lưu trữ dữ liệu vào không gian bộ nhớ tạm thời của request hiện tại trong framework Fiber (Go).
		ctx.Locals("user", user)

		return ctx.Next()

	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}

}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")

	return user.(domain.User)

}
