package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MainTesting(t *testing.T){
	assert := assert.New(t)
	assert.Equal("A","A", "is Trie")

}