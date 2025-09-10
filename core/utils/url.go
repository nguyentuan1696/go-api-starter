package utils

import (
	"strings"

	"github.com/gosimple/slug"
)

func GenerateIDAndSlug(s string) (string, string) {
	id := GenerateID()
	if s == "" {
		return id, id
	}

	var b strings.Builder
	b.WriteString(slug.Make(s))
	b.WriteByte('-')
	b.WriteString(id)
	return id, b.String()
}

func GenerateSlugWithID(s, id string) string {
	var b strings.Builder
	b.WriteString(slug.Make(s))
	b.WriteByte('-')
	b.WriteString(id)
	return b.String()
}

func GenerateSlugWithName(name string) string {
	id := GenerateID()
	if name == "" {
		return id
	}

	var finalSlug strings.Builder
	finalSlug.WriteString(slug.Make(name))
	finalSlug.WriteByte('-')
	finalSlug.WriteString(id)
	return finalSlug.String()

}
