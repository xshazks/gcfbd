package gcfbd

import (
	"encoding/json"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllGeoData(mconn, collectionname)
	jsondatagedung, _ := json.Marshal(datagedung)
	return string(jsondatagedung)
}