package main

import (
	crand "crypto/rand"
	"strings"

	"github.com/google/uuid"
	oklogULID "github.com/oklog/ulid/v2"
)

// UuidSorted generación de IDs (ordenados por time) de 26 caracteres para claves de la base de datos
func UuidSorted() string {
	return strings.ToLower(oklogULID.MustNew(oklogULID.Now(), crand.Reader).String())
}

// Uuid returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx  (36 caracteres)
func Uuid() string {
	return uuid.New().String()
}

// UuidPlain returns the string form of uuid, xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (32 caracteres)
func UuidPlain() string {
	return strings.Replace(Uuid(), "-", "", -1)
}
func ReplaceAll(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
