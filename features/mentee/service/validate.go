package service

import (
	"alta-dashboard-be/features/mentee"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

func validate(u *mentee.MenteeCore) error {
	//validate empty string
	fieldMap := map[string]string{
		"address":          u.Address,
		"full_name":        u.FullName,
		"email":            u.Email,
		"phone":            u.Phone,
		"telegram":         u.Telegram,
		"emergency_name":   u.EmergencyName,
		"emergency_phone":  u.EmergencyPhone,
		"emergency_status": u.EmergencyStatus,
		"education_type":   u.EducationType,
		"education_major":  u.EducationMajor,
	}
	err := validateEmptyString(fieldMap)
	if err != nil {
		return err
	}

	//validate phone no emergency phone
	err = validateNumericOnly(u.Phone, "phone")
	if err != nil {
		return err
	}
	err = validateNumericOnly(u.EmergencyPhone, "emergency_phone")
	if err != nil {
		return err
	}

	//validate address, fullname, emergency name, emergency status, education major
	err = validateAlphanumericSpaceOnly(u.Address, "address")
	if err != nil {
		return err
	}
	err = validateAlphabetSpaceOnly(u.FullName, "full_name")
	if err != nil {
		return err
	}
	err = validateAlphabetSpaceOnly(u.EmergencyName, "emergency_name")
	if err != nil {
		return err
	}
	err = validateAlphanumericSpaceOnly(u.EducationMajor, "emergency_status")
	if err != nil {
		return err
	}
	err = validateAlphanumericSpaceOnly(u.EmergencyStatus, "education_major")
	if err != nil {
		return err
	}

	//validate telegram
	err = validateAlphanumeric(u.Telegram, "telegram")
	if err != nil {
		return err
	}

	//validate education type
	if u.EducationType != "IT" && u.EducationType != "NON-IT" {
		return errors.New("education_type only IT or NON-IT")
	}

	//validate email
	_, err = mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("invalid email format")
	}

	//validate status if exists
	if u.Status != "" {
		err = validateStatus(u.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateStatus(status string) error {
	return validateENUM(status, "status", "Interview", "Continue Unit 1", "Continue Unit 2", "Continue Unit 3", "Graduated", "Eliminated", "Join Class", "Placement", "Repeat Unit 1", "Repeat Unit 2", "Repeat Unit 3")
}

func validateENUM(value string, name string, expected ...string) error {
	for _, v := range expected {
		if v == value {
			return nil
		}
	}
	return fmt.Errorf("%s only accept enum of %v", name, expected)
}

func validateEmptyString(val map[string]string) error {
	for k, v := range val {
		if v == "" {
			return errors.New(fmt.Sprintf("%s cannot be empty", k))
		}
	}
	return nil
}

func validateAlphabetSpaceOnly(val string, name string) error {
	if !regexp.MustCompile(`^[ a-zA-Z]+$`).MatchString(val) {
		return errors.New(fmt.Sprintf("%s must only contain space and alphanumeric", name))
	}
	return nil
}

func validateAlphanumericSpaceOnly(val string, name string) error {
	if !regexp.MustCompile(`^[-_ a-zA-Z0-9]+$`).MatchString(val) {
		return errors.New(fmt.Sprintf("%s must only contain space and alphanumeric", name))
	}
	return nil
}

func validateAlphanumeric(val string, name string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(val) {
		return errors.New(fmt.Sprintf("%s must only contain alphanumeric", name))
	}
	return nil
}

func validateNumericOnly(val string, name string) error {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(val) {
		return errors.New(fmt.Sprintf("%s must only contain numeric", name))
	}
	return nil
}
