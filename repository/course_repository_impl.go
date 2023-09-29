package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
)

type CourseRepositoryImpl struct {
}

func NewCourseRepository() CourseRepository {
	return &CourseRepositoryImpl{}
}

func (repository *CourseRepositoryImpl) CreateCourse(ctx context.Context, tx *sql.Tx, course domain.Course) domain.Course {
	SQL := "INSERT INTO course (course_name, course_desc, course_price, course_image, course_video, instructor_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, course.CourseName, course.CourseDescription, course.CoursePrice, course.CourseImage, course.CourseVideo, course.InstructorId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	course.CourseId = int(id)
	return course
}

func (repository *CourseRepositoryImpl) GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course {
	SQL := "SELECT course_id, course_name, course_desc, course_price, course_image, course_video, intructor_id FROM course"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		course := domain.Course{}
		err := rows.Scan(&course.CourseId, &course.CourseName, &course.CourseDescription, &course.CoursePrice, &course.CourseImage, &course.CourseVideo, &course.InstructorId)
		helper.PanicIfError(err)

		courses = append(courses, course)
	}

	return courses
}
