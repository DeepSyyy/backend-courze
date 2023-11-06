package repository

type Repository struct {
	AdminRepository  AdminRepository
	CourseRepository CourseRepository
}

func NewRepository(adminRepository AdminRepository, courseRepository CourseRepository) *Repository {
	return &Repository{
		AdminRepository:  adminRepository,
		CourseRepository: courseRepository,
	}
}
