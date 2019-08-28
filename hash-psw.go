package tools

import "golang.org/x/crypto/bcrypt"

const DefaultPasswordCost = 10

func GeneratePasswordHash(input []byte, costs ...int) ([]byte, error) {
	cost := DefaultPasswordCost
	if len(costs) > 0 {
		cost = costs[0]
	}
	return bcrypt.GenerateFromPassword(input, cost)
}

func ValidatePassword(input, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, input)
}
