package initial

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

var conf Config
var db *sql.DB

// Config ...
type Config struct {
	Database database
}

type database struct {
	Server   string
	Port     string
	Name     string
	User     string
	Password string
	Provider string
}

func init() {

	if _, err := toml.DecodeFile("./configs/db.toml", &conf); err != nil {
		panic(err)
	}

	dbConn, err := sql.Open(conf.Database.Provider,
		fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Server,
			conf.Database.Port,
		),
	)

	if err != nil {
		panic(err)
	}

	db = dbConn

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database un reachable ...")
		panic(errPing)
	}

	fmt.Println("database is running ....")
}

// GetDB ...
func GetDB() *sql.DB {
	return db
}
