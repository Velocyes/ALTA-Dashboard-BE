package data

import (
	"alta-dashboard-be/features/class"
	"alta-dashboard-be/features/class/models"
)

func convertToCore(u *models.Class) class.ClassCore {
	return class.ClassCore{
		ID:        int(u.ID),
		CreatedAt: u.CreatedAt,
		Name:      u.Name,
		ShortName: u.ShortName,
		StartDate: u.StartDate,
		EndDate:   u.EndDate,
		UserID:    u.UserID,
	}
}

func convertToCoreList(u []models.Class) []class.ClassCore {
	data := []class.ClassCore{}
	for _, v := range u {
		data = append(data, convertToCore(&v))
	}
	return data
}

func convertToModel(u *class.ClassCore) models.Class {
	x := models.Class{
		Name:      u.Name,
		ShortName: u.ShortName,
		StartDate: u.StartDate,
		EndDate:   u.EndDate,
		UserID:    u.UserID,
	}

	return x
}
