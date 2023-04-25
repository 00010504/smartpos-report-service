package helper

import (
	"strings"
	"time"

	"github.com/Invan2/invan_report_service/config"
	"github.com/pkg/errors"
)

var (
	ErrInvalidArgument = errors.New("date is invalid. can't be parsed as date")
	ErrParsingDate     = func(err error) error {
		return errors.Wrap(err, "error parsing date")
	}
)

// date parameter formats: DD_MM_YYYY or DD_MM_YYYY_HH_MM_SS
// return Time.Zone - 5 hours of date
// ex. UTC+5 returned as UTC+0
func ParseDate(date string) (time.Time, error) {
	if len(date) < 10 {
		return time.Time{}, ErrInvalidArgument
	}

	var bDate time.Time
	var err error
	switch {
	case len(date) == 10:
		bDate, err = time.Parse(config.DD_MM_YYYY, date)
		if err != nil {
			return time.Time{}, ErrParsingDate(err)
		}
	default:
		date = strings.Replace(date, "+", " ", -1)
		bDate, err = time.Parse(config.DD_MM_YYYY_HH_MM_SS, date)
		if err != nil {
			return time.Time{}, ErrParsingDate(err)
		}
	}
	bDate = bDate.Add(-time.Hour * 5)
	return bDate, nil
}
