package services

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"github.com/yudha-Dlesmana/fakeAPI/types"
)

func GenerateFakeData(jsonString string, count int) ([]map[string]any, error) {

	var parseType map[string]any

	err := json.Unmarshal([]byte(jsonString), &parseType)

	if err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	fakes := make([]map[string]any, 0, count)

	for i := 1; i <= count; i++ {
		fake := make(map[string]any)
		for key, typ := range parseType {
			if typStr, ok := typ.(string); ok {
				if typStr == "autoIncrement" {
					fake[key] = i

				} else {
					fake[key] = GenerateFakeValue(typStr)
				}
			}
		}
		fakes = append(fakes, fake)
	}
	return fakes, nil
}

func GenerateFakeValue(typeName string) any {
	fakerMap := map[string]func() any{
		"timestamp":   func() any { return faker.Timestamp() },
		"email":       func() any { return faker.Email() },
		"username":    func() any { return faker.Username() },
		"name":        func() any { return faker.Name() },
		"phoneNumber": func() any { return faker.E164PhoneNumber() },
		"uuid":        func() any { return faker.UUIDHyphenated() },
		"word":        func() any { return faker.Word() },
		"sentence":    func() any { return faker.Sentence() },
		"paragraf":    func() any { return faker.Paragraph() },
		"landscape":   func() any { return types.Landscape },
		"potret":      func() any { return types.Potret },
		"square":      func() any { return types.Square },
		"boolean":     func() any { return rand.Intn(2) == 1 },
	}

	if fn, ok := fakerMap[typeName]; ok {
		return fn()
	}
	return "unknownType"
}
