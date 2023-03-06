package service

import (
	"alta-dashboard-be/features/class"
	"errors"
	"regexp"
	"time"
)

func validate(c *class.ClassCore) error {
	//validate name only space and alphanumeric
	if !regexp.MustCompile(`/^[a-z\d\-_\s]+$/i`).MatchString(c.Name) {
		return errors.New("name must only contain space and alphanumeric")
	}

	//validate short name only space and alphanumeric
	if !regexp.MustCompile(`/^[a-z\d\-_\s]+$/i`).MatchString(c.ShortName) {
		return errors.New("short_name must only contain space and alphanumeric")
	}

	//validate empty
	if c.Name == "" && c.ShortName == "" {
		return errors.New("short_name and name cannot be empty")
	}

	//validate end date cannot earlier than start date
	if c.EndDate.Sub(c.StartDate).Hours() <= 0 {
		return errors.New("end_date cannot be earlier than start date")
	}

	//validate start date and end date cannot older than today
	if time.Since(c.StartDate).Hours() <= (1*time.Hour.Hours()) || time.Since(c.EndDate).Hours() <= (1*time.Hour.Hours()) {
		return errors.New("start_date and end_date cannot set in the a past")
	}

	return nil
}
