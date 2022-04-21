package helper

import "golang.org/x/crypto/bcrypt"

func ResponseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ResponseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
	}
}

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
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
