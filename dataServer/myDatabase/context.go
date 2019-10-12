package myDatabase

func GetPostgreSQLContext() (driverName, dataSourceName string) {
	//return "postgres","user=postgres dbname=team_building password=Jim980911 sslmode=disable"
	return "postgres", "user=postgres dbname=team_building password=Jim980911 sslmode=disable"
}
