package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ScrapeLog struct {
	ID           string       `db:"id"`
	ScrapeStatus ScrapeStatus `db:"scrape_status"`
	TimeOfRun    time.Time    `db:"time_of_run"`
}

type ScrapeStatus string

const (
	ScrapeSuccess ScrapeStatus = "SUCCESS"
	ScrapeFailure ScrapeStatus = "FAILURE"
)

func GetLastSuccessfulScrapeDate(db *sqlx.DB) (time.Time, error) {
	var scrapeLog ScrapeLog

	query := `SELECT time_of_run FROM scrape_logs WHERE scrape_status = $1 ORDER BY time_of_run DESC LIMIT 1`

	err := db.Get(&scrapeLog, query, ScrapeSuccess)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, fmt.Errorf("no successful scrape found")
		}
		return time.Time{}, fmt.Errorf("failed to query last successful scrape date: %v", err)
	}

	return scrapeLog.TimeOfRun, nil
}
