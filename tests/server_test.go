package tests

import (
	//"github.com/stretchr/testify/assert"
	//"net/http"
	"server/ip2country/server"
	"testing"
)

func test(t *testing.T) {
	s := server.NewServer()
	s.Run()

	//req, _ := http.NewRequest(http.MethodGet, "/v1/find-country?ip=192.168.1.31", nil)
	//assert.Equal(t, Location{}, 123, "they should be equal")

}
