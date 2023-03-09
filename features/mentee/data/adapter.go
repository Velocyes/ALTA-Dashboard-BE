package data

import (
	"alta-dashboard-be/features/mentee"
	"alta-dashboard-be/features/mentee/models"
	"database/sql"
)

func convertToModels(m *mentee.MenteeCore) (models.Mentee, models.Emergency, models.Education) {
	mentee := models.Mentee{
		FullName: m.FullName,
		Email:    m.Email,
		Address:  m.Address,
		Phone:    m.Phone,
		Telegram: m.Telegram,
		ClassID:  m.ClassID,
		Status:   m.Status,
	}

	emergency := models.Emergency{
		Name:   m.EmergencyName,
		Phone:  m.EmergencyPhone,
		Status: m.EmergencyStatus,
	}

	education := models.Education{
		Type:  m.EducationType,
		Major: m.EducationMajor,
	}

	//set graddate if not nil
	if m.EducationGradDate != nil {
		education.GraduationDate = sql.NullTime{
			Valid: true,
			Time:  *m.EducationGradDate,
		}
	}

	return mentee, emergency, education
}

func convertToCore(u *models.Mentee) mentee.MenteeCore {
	m := mentee.MenteeCore{
		ID:              int(u.ID),
		ClassID:         u.ClassID,
		CreatedAt:       u.CreatedAt,
		FullName:        u.FullName,
		Email:           u.Email,
		Telegram:        u.Telegram,
		Address:         u.Address,
		EmergencyName:   u.Emergency.Name,
		EmergencyPhone:  u.Emergency.Phone,
		EmergencyStatus: u.Emergency.Status,
		EducationType:   u.Education.Type,
		EducationMajor:  u.Education.Major,
		Status:          u.Status,
		Phone:		u.Phone,
	}

	//set graduation date
	if u.Education.GraduationDate.Valid {
		m.EducationGradDate = &u.Education.GraduationDate.Time
	}

	return m
}

func convertToCoreList(u []models.Mentee) []mentee.MenteeCore {
	data := []mentee.MenteeCore{}
	for _, v := range u {
		data = append(data, convertToCore(&v))
	}
	return data
}
