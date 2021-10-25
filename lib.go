package libTask2

import (
	"errors"
	"github.com/aidarkhanov/nanoid/v2"
	"math"
	"unicode"
)

func ChangeRegistCharacter(character rune) rune {
	if unicode.IsLower(character) {
		return unicode.ToUpper(character)
	} else {
		return unicode.ToLower(character)
	}
}

func QuadraticRoot(a, b, c float64) (float64, float64, error) {
	d := b*b - 4*a*c
	if d < 0 {
		return -1, -1, errors.New("Minus discriminant")
	}
	d = math.Sqrt(d)
	r1 := (-b + d) / (2 * a)
	r2 := (-b - d) / (2 * a)
	return r1, r2, nil
}

func GenerateUUID() string {
	uuid, _ := nanoid.New()
	return uuid
}