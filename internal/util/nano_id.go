package util

import gonanoid "github.com/matoous/go-nanoid"

func NewNanoID(prefix ...string) string {
	id, err := gonanoid.Generate("2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ", 21)
	if err != nil {
		panic(err)
	}
	if len(prefix) > 0 {
		return prefix[0] + id
	}
	return id
}
