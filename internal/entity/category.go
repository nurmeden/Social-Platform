package entity

type Category struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts"`
}

func NewCategory(id int64, name string) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}

func (c *Category) AddPost(p Post) {
	c.Posts = append(c.Posts, p)
}
