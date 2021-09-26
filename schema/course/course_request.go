package course

type CourseDetailRequest struct {
	ID int64 `param:"id" json:"-" validate:"required"`
}

type CourseGetRequest struct {
	Search string `query:"search" json:"-"`
}
