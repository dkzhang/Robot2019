package myDatabase

import (
	"fmt"
	"log"
)

func InitializeDB() error {
	schemaInspectionLogRecord := `
	CREATE TABLE INSPECTION_LOG_RECORD(
		INSPECTION_ID         BIGINT          NOT NULL,
		RECORD_ID             BIGINT          NOT NULL,
		RECORD_LEVEL          TEXT            NOT NULL,
		RECORD_DATETIME       TIMESTAMPTZ     NOT NULL,
		TEXT_CONTENT          TEXT            NOT NULL,
		IMAGE_URL             TEXT            NOT NULL,
		PRIMARY KEY(INSPECTION_ID, RECORD_ID)
	);`

	schemaInspectionLog := `
	CREATE TABLE INSPECTION_LOG(
		INSPECTION_ID         BIGINT          NOT NULL,
		DATETIME_BEGIN        TIMESTAMPTZ     NOT NULL,
		DATETIME_END          TIMESTAMPTZ     NOT NULL,
		INSPECTION_CONCLUSION TEXT            NOT NULL,
		PRIMARY KEY(INSPECTION_ID)
	);`

	db, err := ConnectToDatabase(GetPostgreSQLContext())
	if err != nil {
		log.Fatalf("database.ConnectToDatabase error: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(schemaInspectionLog)
	if err != nil {
		return fmt.Errorf("create table InpsecitonLog error: %v", err)
	}

	_, err = db.Exec(schemaInspectionLogRecord)
	if err != nil {
		return fmt.Errorf("create table InpsecitonLogRecord error: %v", err)
	}
	return nil
}
