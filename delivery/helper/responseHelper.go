package helper

import "golang.org/x/crypto/bcrypt"

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": message,
		"data":    data,
	}
}

func ResponseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": message,
	}
}

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    400,
		"message": message,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
