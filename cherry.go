package cherry

import (
	"github.com/go2rust1/cherry/src/cherry"
	"github.com/go2rust1/cherry/src/mod/database"
	"github.com/go2rust1/cherry/src/trait"
)

type (
	Topic    = trait.Topic
	Response = trait.Response
)

func New() trait.Cherry {
	return cherry.New()
}

func DB2() trait.Database {
	return database.NewDB2()
}

func MySQL() trait.Database {
	return database.NewMySQL()
}

func Oracle() trait.Database {
	return database.NewOracle()
}
