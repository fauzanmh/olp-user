package course

type CourseDetailRequest struct {
	ID int64 `param:"id" json:"-" validate:"required"`
}
