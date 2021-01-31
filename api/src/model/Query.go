package model

import (
	"fmt"
	"log"
)

func makeSelectQuery(params map[string]string) string {
	query := fmt.Sprintf("SELECT *")
	paramOrder := [...]string{"attr", "from", "where", "order by", "limit"}

	for _, p := range paramOrder {
		switch val, exists := params[p]; p {
		case "from":
			if !exists {
				return ""
			}
			query += fmt.Sprintf(" %s %s", p, val)
		case "where", "order by", "limit":
			if !exists {
				continue
			}
			query += fmt.Sprintf(" %s %s", p, val)
		default:
			continue
		}
	}
	log.Println("Query: makeQueryString: query:", query)
	return query
}
