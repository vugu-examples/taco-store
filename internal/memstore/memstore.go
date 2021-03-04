package memstore

type TacoStore struct {
	TacoList []Taco
}

func NewTacoStore() *TacoStore {
	return &TacoStore{TacoList: []Taco{
		{
			ID:          1,
			Title:       "Fried Avocado",
			Description: "Fried, creamy, crunchy, avocado-y goodness, shall we? These yummy tacos are super easy to make and they’re perfect for Meatless Monday, Taco Tuesday, or Get In Mah Belly Day.",
			ImageUrl:    "https://i.imgur.com/4v5UYxZ.jpg",
			Price:       299,
		},
	}}
}

type Taco struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
}

func (t *TacoStore) SelectTacoList() []Taco {
	return t.TacoList
}
