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

	err := db.Model(&dbSheets).Select()

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

func InsertSheet(name string, userID int) (interface{}, error) {
	if userID == 0 || name == "" {
		return nil, errors.New("Missing arguments")
	}

	_, err := users.QueryUserById(userID)

	if err != nil {
		return nil, errors.New("Tried to attach sheet to a user. User did not exist")
	}

	currTime := time.Now().Local()
	sheet := database.Sheet{
		Name:      name,
		UserID:    userID,
		CreatedAt: currTime.String(),
	}

	if _, err := db.Model(&sheet).Returning("id").Insert(); err != nil {
		return nil, err
	}

	return model.Sheet{
		ID:        sheet.ID,
		UserID:    sheet.UserID,
		Name:      sheet.Name,
		CreatedAt: sheet.CreatedAt,
	}, nil
}
