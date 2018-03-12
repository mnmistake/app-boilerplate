package segments

import (
	"errors"
	"time"

	"github.com/raunofreiberg/kyrene/server/api/sheets"
	"github.com/raunofreiberg/kyrene/server/database"
	"github.com/raunofreiberg/kyrene/server/model"
)

var db = database.Database()

func QuerySegments(sheetID int) (interface{}, error) {
	if sheetID == 0 {
		return nil, errors.New("This query requires a SheetID param")
	}

	var segments []model.Segment
	var dbSegments []database.Segment

	_, err := db.Query(
		&dbSegments,
		"SELECT id, sheet_id, label, content, created_at FROM segments WHERE sheet_id = ?",
		sheetID,
	)

	if err != nil {
		return nil, err
	}

	for _, segment := range dbSegments {
		segments = append(segments, model.Segment{
			ID:        segment.ID,
			SheetID:   segment.SheetID,
			Label:     segment.Label,
			Content:   segment.Content,
			CreatedAt: segment.CreatedAt,
		})
	}

	return segments, nil
}

func InsertSegment(sheetID int, label string, content string) (interface{}, error) {
	if sheetID == 0 || label == "" || content == "" {
		return nil, errors.New("Missing arguments")
	}

	// Make sure the related sheet exists
	_, err := sheets.QuerySheet(sheetID)

	if err != nil {
		return nil, err
	}

	currTime := time.Now().Local()
	segment := database.Segment{
		SheetID:   sheetID,
		Label:     label,
		Content:   content,
		CreatedAt: currTime.String(),
	}

	if _, err := db.Model(&segment).Returning("id").Insert(); err != nil {
		return nil, err
	}

	return model.Segment{
		ID:        segment.ID,
		SheetID:   segment.SheetID,
		Label:     segment.Label,
		Content:   segment.Content,
		CreatedAt: segment.CreatedAt,
	}, nil
}
