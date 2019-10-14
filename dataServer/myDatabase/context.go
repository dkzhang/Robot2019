package myDatabase

func GetPostgreSQLContext() (driverName, dataSourceName string) {
	//return "postgres","user=postgres dbname=team_building password=Jim980911 sslmode=disable"
	return "postgres", "host=myPostgreSQL001 user=postgres dbname=postgres password=xxjs2019 sslmode=disable"
}
