package helper

import (
	domain "courze-backend-app/model/domain"
	web "courze-backend-app/model/web"
)

func ToInstructorResponse(instructor domain.Instructor) web.InstructorResponse {
	return web.InstructorResponse{
		Id:    instructor.Id,
		Name:  instructor.Name,
		Email: instructor.Email,
	}
}

func ToInstructorResponses(instructors []domain.Instructor) []web.InstructorResponse {
	instructorResponses := []web.InstructorResponse{}
	for _, instructor := range instructors {
		instructorResponses = append(instructorResponses, ToInstructorResponse(instructor))
	}
	return instructorResponses
}
