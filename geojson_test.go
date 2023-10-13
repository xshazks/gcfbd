package gcfbd

import (
	"fmt"
	"testing"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gisaw")
	datasaw := GetAllGeoData(mconn, "petasaw")
	fmt.Println(datasaw)
}