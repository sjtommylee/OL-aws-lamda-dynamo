package posts

import "github.com/gin-gonic/gin"

type PostSerializer struct {
	C *gin.Context
	PostModel
}

type PostResponse struct {
	ID          uint   `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	CreatedAt   string `json:"createdAt"`
	Author      string
}

type PostsSerializer struct {
	C     *gin.Context
	Posts []PostModel
}

func (s *PostSerializer) Response() PostResponse {
	// myPostModel := s.C.MustGet("my_post_model")
	response := PostResponse{
		// ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		Body:        s.Body,
		Author:      s.Author,
	}
	return response
}
