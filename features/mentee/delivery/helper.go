package delivery

import (
	"alta-dashboard-be/features/mentee"
	"errors"
	"time"
)

func convertToCore(u *mentee.MenteeRequest) (mentee.MenteeCore, error) {
	m := mentee.MenteeCore{
		FullName:        u.FullName,
		Email:           u.Email,
		Address:         u.Address,
		Phone:           u.Phone,
		Telegram:        u.Telegram,
		EmergencyName:   u.EmergencyName,
		EmergencyPhone:  u.EmergencyPhone,
		EmergencyStatus: u.EmergencyStatus,
		EducationType:   u.EducationType,
		EducationMajor:  u.EducationMajor,
	}
	if u.EducationGradDate != "" {
		education, err := time.Parse("", u.EducationGradDate)
		if err != nil {
			return mentee.MenteeCore{}, errors.New("invalid time format education_graduation_date")
		}
		m.EducationGradDate = &education
	}
	return m, nil
}

func convertToResponse(u *mentee.MenteeCore) mentee.MenteeResponse {
	return mentee.MenteeResponse{
		FullName:          u.FullName,
		Email:             u.Email,
		Address:           u.Address,
		Phone:             u.Phone,
		Telegram:          u.Telegram,
		EmergencyName:     u.EmergencyName,
		EmergencyPhone:    u.EmergencyPhone,
		EmergencyStatus:   u.EmergencyStatus,
		EducationType:     u.EducationType,
		EducationMajor:    u.EducationMajor,
		EducationGradDate: u.EducationGradDate.Format(""),
	}
}

func convertToResponseList(u []mentee.MenteeCore) []mentee.MenteeResponse {
	data := []mentee.MenteeResponse{}
	for _, v := range u {
		data = append(data, convertToResponse(&v))
	}
	return data
}
