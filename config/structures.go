package config

type Conf struct {
	Port           string `json:"server_port"`
	MongoURI       string `json:"mongo_uri"`
	MongoDBName    string `json:"mongodb_name"`
	MongoDBTimeout int    `json:"mongodb_timeout"`
	JWTKey         string `json:"jwt_key"`
}
