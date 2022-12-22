package posts

type PostModelValidator struct {
	PostModel struct {
		Title       string `form:"title" json:"title" binding:"exists,min=4"`
		Description string `form:"description" json:"description" binding:"max=2048"`
		Body        string `form:"body" json:"body" binding:"max=2048"`
	}
}

func NewPostModelValidator() PostModelValidator {
	return PostModelValidator{}
}
