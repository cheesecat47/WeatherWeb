package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cheesecat47/webpractice/api/model"
	"github.com/stretchr/testify/assert"
)

// TestClosedDB func
func TestClosedDB(t *testing.T) {
	err = model.CloseDB()
	assert.Equal(t, nil, err, fmt.Errorf("TestClosedDB: Error: %v", err))
	log.Println("TestClosedDB: DB closed")
}
