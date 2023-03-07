package delivery

import (
	"alta-dashboard-be/features/class"
	"errors"
	"time"
)

func convertToCore(u *class.ClassRequest) (class.ClassCore, error) {
	//parse time
	startDate, err := time.Parse("", u.StartDate)
	if err != nil {
		return class.ClassCore{}, errors.New("invalid start Date format")
	}
	endDate, err := time.Parse("", u.EndDate)
	if err != nil {
		return class.ClassCore{}, errors.New("invalid end Date format")
	}
	return class.ClassCore{
		Name:      u.Name,
		ShortName: u.ShortName,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}

func convertToResponse(u *class.ClassCore) (class.ClassResponse, error) {
	startDate := u.StartDate.Format("")
	endDate := u.EndDate.Format("")
	createdAt := u.CreatedAt.Format("")
	return class.ClassResponse{
		ID:        u.ID,
		Name:      u.Name,
		ShortName: u.ShortName,
		CreatedAt: createdAt,
		EndDate:   endDate,
		StartDate: startDate,
		UserID:    u.UserID,
	}, nil
}

func convertToResponseList(u []class.ClassCore) ([]class.ClassResponse, error) {
	data := []class.ClassResponse{}

	for _, v := range u {
		tmp, err := convertToResponse(&v)
		if err != nil {
			return nil, err
		}
		data = append(data, tmp)
	}

	return data, nil
}
