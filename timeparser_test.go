package utils

import (
	"testing"
	"time"
)

func TestTimeParser(t *testing.T) {
	_, err := Time().Parse("", 19800)
	Test().Contains(t, err.Error(), "empty value")

	p, err := Time().Parse("01/07/2020", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2020, p.Year)
	Test().Equals(t, time.July, p.Month)
	Test().Equals(t, 1, p.Date)
	Test().Equals(t, 0, p.Hour)
	Test().Equals(t, 0, p.Minute)
	Test().Equals(t, 0, p.Second)
	Test().Equals(t, int64(1593541800), p.Time.Unix())

	p, err = Time().Parse("01/01/2021 4:19 AM", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2021, p.Year)
	Test().Equals(t, time.January, p.Month)
	Test().Equals(t, 1, p.Date)
	Test().Equals(t, 4, p.Hour)
	Test().Equals(t, 19, p.Minute)
	Test().Equals(t, 0, p.Second)
	Test().Equals(t, int64(1609454940), p.Time.Unix())

	p, err = Time().Parse("01/07/2020 6:12 PM", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2020, p.Year)
	Test().Equals(t, time.July, p.Month)
	Test().Equals(t, 1, p.Date)
	Test().Equals(t, 18, p.Hour)
	Test().Equals(t, 12, p.Minute)
	Test().Equals(t, 0, p.Second)
	Test().Equals(t, int64(1593607320), p.Time.Unix())

	p, err = Time().Parse("2020-12-04", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2020, p.Year)
	Test().Equals(t, time.December, p.Month)
	Test().Equals(t, 4, p.Date)
	Test().Equals(t, 0, p.Hour)
	Test().Equals(t, 0, p.Minute)
	Test().Equals(t, 0, p.Second)
	Test().Equals(t, int64(1607020200), p.Time.Unix())

	p, err = Time().Parse("2019-02-04 09:00:07", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2019, p.Year)
	Test().Equals(t, time.February, p.Month)
	Test().Equals(t, 4, p.Date)
	Test().Equals(t, 9, p.Hour)
	Test().Equals(t, 0, p.Minute)
	Test().Equals(t, 7, p.Second)
	Test().Equals(t, int64(1549251007), p.Time.Unix())

	p, err = Time().Parse("2020-11-15 14:25:20", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2020, p.Year)
	Test().Equals(t, time.November, p.Month)
	Test().Equals(t, 15, p.Date)
	Test().Equals(t, 14, p.Hour)
	Test().Equals(t, 25, p.Minute)
	Test().Equals(t, 20, p.Second)
	Test().Equals(t, int64(1605430520), p.Time.Unix())

	p, err = Time().Parse("2021-01-21T12:09:09", 19800)
	Test().Nil(t, err)
	Test().Equals(t, 2021, p.Year)
	Test().Equals(t, time.January, p.Month)
	Test().Equals(t, 21, p.Date)
	Test().Equals(t, 12, p.Hour)
	Test().Equals(t, 9, p.Minute)
	Test().Equals(t, 9, p.Second)
	Test().Equals(t, int64(1611211149), p.Time.Unix())
}

// func TestTimeZoneDelta(t *testing.T) {
// 	k, err := getTimeZoneDelta("IST")
// 	Test().Nil(t, err)

// 	t530, err := time.ParseDuration("-5h30m")
// 	Test().Nil(t, err)

// 	Test().Equals(t, t530, k)
// }
