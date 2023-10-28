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
	SQL := "INSERT INTO course(course_name, course_description, course_price, course_image, course_video, instructor_id, sneakpeek) VALUES(?,?,?,?,?,?, (SELECT JSON_ARRAY(?)))"
	result, err := tx.ExecContext(ctx, SQL, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorId, course.SneakPeak)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	course.Id = int(id)
	return course
}

func (repository *CourseRepositoryImpl) GetAllCourse(ctx context.Context, tx *sql.Tx) []domain.Course {
	SQL := "SELECT course_id, course_name, course_description, course_price, course_image, course_video, instructor_id FROM course"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		course := domain.Course{}
		err := rows.Scan(course.Id, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorId, course.SneakPeak)
		helper.PanicIfError(err)

		courses = append(courses, course)
	}

	return courses
}

func (repository *CourseRepositoryImpl) GetCourseById(ctx context.Context, tx *sql.Tx, courseId int) domain.Course {
	SQL := "SELECT course_id, course_name, course_description, course_price, course_image, course_video, instructor_id FROM course WHERE course_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, courseId)
	helper.PanicIfError(err)
	defer rows.Close()

	course := domain.Course{}
	for rows.Next() {
		err := rows.Scan(course.Id, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorId, course.SneakPeak)
		helper.PanicIfError(err)
	}

	return course
}

func (repository *CourseRepositoryImpl) GetCourseByInstructorId(ctx context.Context, tx *sql.Tx, instructorId int) []domain.Course {
	// query with join
	SQL := `SELECT
	course.course_id,
	course.course_name,
	course.course_description,
	course.course_price,
	course.course_image,
	course.course_video,
	course.instructor_id
FROM
	course
LEFT JOIN
	instructor ON course.instructor_id = instructor.instructor_id
WHERE
	instructor.instructor_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, instructorId)
	helper.PanicIfError(err)
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		course := domain.Course{}
		err := rows.Scan(course.Id, course.Name, course.Description, course.Price, course.Image, course.Video, course.InstructorId, course.SneakPeak)
		helper.PanicIfError(err)

		courses = append(courses, course)
	}

	return courses
}
