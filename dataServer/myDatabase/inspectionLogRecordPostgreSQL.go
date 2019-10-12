package myDatabase

import (
	"Robot2019/dataServer/myDatabase/structure"
	"fmt"
	"log"
	"time"
)

func InsertInspectionLogRecordPostgreSQL(ilr structure.InspectionLogRecord) error {
	db, err := ConnectToDatabase(GetPostgreSQLContext())
	if err != nil {
		log.Fatalf("database.ConnectToDatabase error: %v", err)
	}
	defer db.Close()

	tx, err := db.Begin()

	insertRecord := `INSERT INTO INSPECTION_LOG_RECORD (INSPECTION_ID, RECORD_ID, RECORD_LEVEL, RECORD_DATETIME, TEXT_CONTENT, IMAGE_URL) VALUES ($1, $2, $3, $4, $5, $6)`
	//先插入record记录
	_, err = tx.Exec(insertRecord,
		ilr.InspectionID, ilr.RecordID, ilr.RecordLevel,
		ilr.RecordDateTime, ilr.TextContent, ilr.ImageUrl)

	insertInspection := `INSERT INTO INSPECTION_LOG (INSPECTION_ID, DATETIME_BEGIN, DATETIME_END, INSPECTION_CONCLUSION) VALUES ($1, $2, $3, $4)`
	if ilr.RecordID == 1 {
		//如果record记录ID为1，插入inspection记录
		_, err = tx.Exec(insertInspection,
			ilr.InspectionID, ilr.RecordDateTime, DefaultEndDatetime, "")
	}
	//如果record记录为-1，(统计巡检情况，更新inspection记录)，此处不处理，有上层应用另行处理

	if err == nil {
		err = tx.Commit()
		if err == nil {
			return nil
		}
	}

	if err != nil {
		err = fmt.Errorf("Insert InspectionLogRecordPostgreSQL error: %v", err)
		errRollback := tx.Rollback()
		if errRollback != nil {
			return fmt.Errorf("Insert InspectionLogRecordPostgreSQL error: %v , and Rollback error: %v", err, errRollback)
		} else {
			return fmt.Errorf("Insert InspectionLogRecordPostgreSQL error: %v , but Rollback successful", err)
		}
	}
	return err
}

var DefaultEndDatetime = time.Date(2099, 12, 31, 12, 13, 24, 0, time.UTC)
