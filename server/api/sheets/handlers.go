package sheets

import (
	"errors"
	"time"

	"github.com/raunofreiberg/kyrene/server/api/users"
	"github.com/raunofreiberg/kyrene/server/database"
	"github.com/raunofreiberg/kyrene/server/model"
)

var db = database.Database()

func QuerySheets() (interface{}, error) {
	var sheets []model.Sheet
	var dbSheets []database.Sheet

	err := db.Model(&dbSheets).OrderExpr("created_at DESC").Select()

	if err != nil {
		return nil, err
	}

	for _, sheet := range dbSheets {
		sheets = append(sheets, model.Sheet{
			ID:        sheet.ID,
			UserID:    sheet.UserID,
			Name:      sheet.Name,
			CreatedAt: sheet.CreatedAt,
		})
	}

	return sheets, nil
}

func QuerySheet(sheetID int) (interface{}, error) {
	if sheetID == 0 {
		return nil, errors.New("Missing sheetID")
	}

	sheet := database.Sheet{}

	_, err := db.QueryOne(
		&sheet,
		"SELECT id, user_id, name, created_at FROM sheets WHERE id = ?", sheetID,
	)

	if err != nil {
		return nil, errors.New("Sheet not found")
	}

	return model.Sheet{
		ID:        sheet.ID,
		Name:      sheet.Name,
		UserID:    sheet.UserID,
		CreatedAt: sheet.CreatedAt,
	}, nil
}

func InsertSheet(name string, userID int, segments []interface{}) (interface{}, error) {
	var segmentsJSON []model.Segment
	var dbSegments []database.Segment

	if userID == 0 || name == "" || segments == nil {
		return nil, errors.New("Missing arguments")
	}

	_, err := users.QueryUserById(userID)

	if err != nil {
		return nil, errors.New("Tried to attach sheet to a user. User did not exist")
	}

	currTime := time.Now().Local().String()
	sheet := database.Sheet{
		Name:      name,
		UserID:    userID,
		CreatedAt: currTime,
	}

	// Insert sheet
	if _, err := db.Model(&sheet).Returning("id").Insert(); err != nil {
		return nil, err
	}

	// Bulk insert related segments
	for _, data := range segments {
		dbSegments = append(dbSegments, database.Segment{
			SheetID:   sheet.ID,
			Content:   data.(map[string]interface{})["content"].(string),
			Label:     data.(map[string]interface{})["label"].(string),
			CreatedAt: currTime,
		})
	}

	if err = db.Insert(&dbSegments); err != nil {
		return nil, err
	}

	for _, segment := range dbSegments {
		segmentsJSON = append(segmentsJSON, model.Segment{
			ID:        segment.ID,
			SheetID:   segment.SheetID,
			Label:     segment.Label,
			Content:   segment.Content,
			CreatedAt: segment.CreatedAt,
		})
	}

	return model.Sheet{
		ID:        sheet.ID,
		UserID:    sheet.UserID,
		Name:      sheet.Name,
		CreatedAt: sheet.CreatedAt,
		Segments:  segmentsJSON,
	}, nil
}
