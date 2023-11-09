package repository

import (
	"context"
	"courze-backend-app/helper"
	domain "courze-backend-app/model/domain"
	"database/sql"
	"errors"
)

type CourseRepositoryImpl struct {
}

func NewCourseRepository() CourseRepository {
	return &CourseRepositoryImpl{}
}

func (repository *CourseRepositoryImpl) GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course {
	SQL := "SELECT course_id, course_name, course_description, course_price, course_image, course_video, instructor_name FROM course"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		course := domain.Course{}
		err := rows.Scan(course.Id, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorName, course.SneakPeak)
		helper.PanicIfError(err)

		courses = append(courses, course)
	}

	return courses
}

func (repository *CourseRepositoryImpl) GetCourseById(ctx context.Context, tx *sql.Tx, courseId int) (domain.Course, error) {
	SQL := "SELECT course_id, course_name, course_description, course_price, course_image, course_video, instructor_name FROM course WHERE course_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, courseId)
	if err != nil {
		// Tangani kesalahan saat eksekusi query
		return domain.Course{}, err
	}
	defer rows.Close()

	course := domain.Course{}
	for rows.Next() {
		err := rows.Scan(&course.Id, &course.Name, &course.Description, &course.Price, &course.Image, &course.Video, &course.InstructorName, &course.SneakPeak)
		if err != nil {
			// Tangani kesalahan saat pembacaan baris
			return domain.Course{}, err
		}
	}

	// Periksa apakah course ditemukan
	if course.Id == 0 {
		// Course tidak ditemukan, kembalikan error sesuai kebutuhan
		return domain.Course{}, errors.New("course not found")
	}

	return course, nil
}

func (repository *CourseRepositoryImpl) GetCourseByInstructorName(ctx context.Context, tx *sql.Tx, instructorName string) []domain.Course {
	// query with join
	SQL := `SELECT course_id, course_name, course_description, course_price, course_image, course_video, instructor_name WHERE instructor_name = ?`
	rows, err := tx.QueryContext(ctx, SQL, instructorName)
	helper.PanicIfError(err)
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		course := domain.Course{}
		err := rows.Scan(course.Id, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorName, course.SneakPeak)
		helper.PanicIfError(err)

		courses = append(courses, course)
	}

	return courses
}
