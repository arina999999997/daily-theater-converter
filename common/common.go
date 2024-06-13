package common

import (
	"log"
)

func AddFormat(name string, parser Parser, serializer Serializer) {
	if parser == nil {
		log.Printf(`Warning: format "%s" doesn't have a parser`, name)
	}
	if serializer == nil {
		log.Printf(`Warning: format "%s" doesn't have a serializer`, name)
	}
	_, exist := parsers[name]
	if exist {
		log.Fatalf(`Error: format "%s" already exists`, name)
	}
	parsers[name] = parser
	serializers[name] = serializer
}
