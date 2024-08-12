package simple

type Database struct {
	Username string
}

type DatabaseMysql Database
type DatabasePostgres Database

func NewDatabasePostgresql() *DatabasePostgres {
	return (*DatabasePostgres)(&Database{Username: "postgressql"})
}

func NewDatabaseMysql() *DatabaseMysql {
	return (*DatabaseMysql)(&Database{Username: "mysql"})
}

type DatabaseRepository struct {
	DatabaseMysql    *DatabaseMysql
	DatabasePostgres *DatabasePostgres
}

func NewDatabaseRepository(postgresql *DatabasePostgres, mysql *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMysql:    mysql,
		DatabasePostgres: postgresql,
	}
}
