package memstore

type MemStore struct {
	TacoList []Taco
	Cart     []Taco
}

func NewMemStore() *MemStore {
	return &MemStore{TacoList: []Taco{
		{
			ID:          1,
			Title:       "Fried Avocado",
			Description: "Fried, creamy, crunchy, avocado-y goodness, shall we? These yummy tacos are super easy to make and they’re perfect for Meatless Monday, Taco Tuesday, or Get In Mah Belly Day.",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       13.99,
		},
		{
			ID:          2,
			Title:       "Blackberry Chicken",
			Description: "Shredded Hoisin-Blackberry Chicken Tacos with Crunchy Slaw. The chicken is cooked in the slow cooker until it’s so tender that it basically falls apart with no effort.",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       14.99,
		},
		{
			ID:          3,
			Title:       "Cilantro Lime Shrimp",
			Description: "Although the cilantro lime shrimp would be good in just about everything, they are particularly nice in these fresh summery tacos! ",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       11.99,
		},
		{
			ID:          4,
			Title:       "Mango Mahi-Mahi",
			Description: "Mahi-Mahi Fish Tacos with Mango and Avocado Salsa. If your not too familiar with mahi-mahi it’s a white fish that is a little firmer and perfect for fish tacos!",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       9.99,
		},
		{
			ID:          5,
			Title:       "Wild Caught Lobster",
			Description: "Fresh lobster with sweet mango, refreshing cucumber, and creamy avocado with a hint of lime, topped with bright cilantro—it's the ultimate summer taco!",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       15.99,
		},
		{
			ID:          6,
			Title:       "The Hungry Vegan ",
			Description: "These Meatless wonders will have you feeling great about your new years resolution. Mushroom and vegan beef crumbles with avocado and plenty of cilantro. Yummy!",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       12.99,
		},
		{
			ID:          7,
			Title:       "Carne Asada",
			Description: "The undisputed king of the taco world. Classic with a twist... of lime! Flank steak marinated in our signature Mojo with homemade pico de gallo. Enough saID.",
			ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
			Price:       13.99,
		},
	},
		Cart: []Taco{
			{
				ID:          7,
				Title:       "Carne Asada",
				Description: "The undisputed king of the taco world. Classic with a twist... of lime! Flank steak marinated in our signature Mojo with homemade pico de gallo. Enough saID.",
				ImageUrl:    "https://i.ibb.co/FhCBtQW/taco-Image.png",
				Price:       13.99,
			},
		}}
}

func (t *MemStore) SelectTacoList() []Taco {
	return t.TacoList
}

func (t *MemStore) SelectCart() []Taco {
	return t.Cart
}

func (t *MemStore) PostCartItem(taco Taco) {
	t.Cart = append(t.Cart, taco)
}

func (t *MemStore) PatchCart(newTacoList []Taco) {
	t.Cart = newTacoList
}
