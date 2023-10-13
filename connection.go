package gcfbd

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllGeoData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}