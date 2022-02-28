package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/betas-in/logger"
)

type TimeFunctions struct{}

type TimeParser struct {
	UserValue string
	Parsed    bool
	Year      int
	Month     time.Month
	Date      int
	Hour      int
	Minute    int
	Second    int
	Time      time.Time
	Offset    int
}

func Time() TimeFunctions {
	return TimeFunctions{}
}

func (t TimeFunctions) Elapsed(log *logger.Logger, what string, logging bool) func() {
	start := time.Now()
	return func() {
		if logging {
			log.Debug(what).Msgf("%s\n", time.Since(start))
		}
	}
}

func (t TimeFunctions) Parse(userValue string, offset int) (*TimeParser, error) {
	// Examples
	// ""
	// 01/07/2020
	// 01/01/2021 4:19 AM
	// 01/07/2020 6:12 PM
	// 2020-12-04
	// 2019-02-04 09:00:07
	// 2020-11-15 14:25:20
	// 2021-01-21T12:09:09

	tp := TimeParser{UserValue: userValue, Offset: offset}

	if tp.UserValue == "" {
		return &tp, fmt.Errorf("cannot parse empty value")
	}

	ampm := []string{"am", "pm"}

	v := tp.UserValue
	// Remove T
	v = strings.Replace(v, "T", " ", -1)
	// Split on spaces
	vs := strings.Split(v, " ")

	for _, item := range vs {
		if strings.Contains(item, "/") || strings.Contains(item, "-") {
			err := tp.parseDate(item)
			if err != nil {
				return &tp, err
			}
		} else if strings.Contains(item, ":") {
			err := tp.parseTime(item)
			if err != nil {
				return &tp, err
			}
		} else if Array().Contains(ampm, item, false) >= 0 {
			tp.parseAMPM(strings.ToLower(item))
		}
	}

	tp.Time = time.Date(tp.Year, tp.Month, tp.Date, tp.Hour, tp.Minute, tp.Second, 0, time.UTC)

	delta := tp.Offset * -1
	tp.Time = tp.Time.Add(time.Second * time.Duration(delta))

	return &tp, nil
}

func (t *TimeParser) parseDate(item string) error {
	v := strings.Replace(item, "/", " ", -1)
	v = strings.Replace(v, "-", " ", -1)

	vs := strings.Split(v, " ")

	var err error
	s1, err := strconv.Atoi(vs[0])
	if err != nil {
		return fmt.Errorf("could not parse date in %s", item)
	}
	s2, err := strconv.Atoi(vs[1])
	if err != nil {
		return fmt.Errorf("could not parse date in %s", item)
	}
	s3, err := strconv.Atoi(vs[2])
	if err != nil {
		return fmt.Errorf("could not parse date in %s", item)
	}

	if s1 > 1970 {
		t.Year = s1
		t.Month = time.Month(s2)
		t.Date = s3
	} else if s3 > 1970 {
		t.Year = s3
		t.Month = time.Month(s2)
		t.Date = s1
	}
	return nil
}

func (t *TimeParser) parseTime(item string) error {
	v := strings.Replace(item, ":", " ", -1)
	vs := strings.Split(v, " ")
	var nms []int

	for _, vsi := range vs {
		s, err := strconv.Atoi(vsi)
		if err != nil {
			return fmt.Errorf("could not parse time in %s", item)
		}
		nms = append(nms, s)
	}

	t.Hour = nms[0]
	t.Minute = nms[1]
	if len(nms) == 3 {
		t.Second = nms[2]
	} else {
		t.Second = 0
	}

	return nil
}

func (t *TimeParser) parseAMPM(item string) {
	if item == "pm" {
		t.Hour += 12
	}
}
