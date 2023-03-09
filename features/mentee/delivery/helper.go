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
		ClassID:         u.ClassID,
		Status:          u.Status,
	}
	if u.EducationGradDate != "" {
		education, err := time.Parse("2006-01-02", u.EducationGradDate)
		if err != nil {
			return mentee.MenteeCore{}, errors.New("invalid time format education_graduation_date")
		}
		m.EducationGradDate = &education
	}
	return m, nil
}

func convertToResponse(u *mentee.MenteeCore) mentee.MenteeResponse {
	data := mentee.MenteeResponse{
		ID:              u.ID,
		CreatedAt:       u.CreatedAt.Format("2006-01-02"),
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
		Status:          u.Status,
		ClassID:         u.ClassID,
	}

	if u.EducationGradDate != nil {
		data.EducationGradDate = u.EducationGradDate.Format("2006-01-02")
	}

	return data
}

func convertToResponseList(u []mentee.MenteeCore) []mentee.MenteeResponse {
	data := []mentee.MenteeResponse{}
	for _, v := range u {
		data = append(data, convertToResponse(&v))
	}
	return data
}
