package graphql

import (
	"fmt"
	"io"
	"time"

	"github.com/volatiletech/null"
)

type Date string

func (d *Date) UnmarshalGQL(v interface{}) error {
	date, ok := v.(string)
	if !ok {
		return fmt.Errorf("%v", v)
	}
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return err
	}
	*d = Date(date)
	return nil
}

func (d Date) MarshalGQL(w io.Writer) {
	w.Write([]byte("\"" + d + "\""))
}

func (d *Date) Time() (*time.Time, error) {
	t, err := time.Parse("2006-01-02", string(*d))
	if err != nil {
		fmt.Println("date error")
		return nil, err
	}
	return &t, nil
}

func (d *Date) LocationTime(loc *time.Location) (*time.Time, error) {
	t, err := time.ParseInLocation("2006-01-02", string(*d), loc)
	if err != nil {
		fmt.Println("date error")
		return nil, err
	}
	return &t, nil
}

func (d *Date) NullTime() (null.Time, error) {
	t := null.NewTime(time.Time{}, false)
	if d == nil {
		return t, nil
	}
	tm, err := time.Parse("2006-01-02", string(*d))
	if err != nil {
		return t, err
	}
	t.Time = tm
	t.Valid = true
	return t, nil
}

func ConvertNullTime(t null.Time) *Date {
	if !t.Valid {
		return nil
	}
	var date Date = Date(t.Time.Format("2006-01-02"))
	return &date
}

func ConvertTime(t time.Time) *Date {
	var date Date = Date(t.Format("2006-01-02"))
	return &date
}
