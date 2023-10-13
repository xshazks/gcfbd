package gcfbd

import "encoding/json"

func GCHandlerFunc(Mongostring, dbname, colname string) []byte {
	koneksyen := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllGeoData(koneksyen, colname)

	jsonsaw, _ := json.Marshal(datageo)

	return jsonsaw
}