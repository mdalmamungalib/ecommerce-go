package repo

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productId int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
			return &product, nil
		}
	}
	return nil, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red, I Love Orange",
		Price:       100,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/4/43/Ambersweet_oranges.jpg",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red, I Love Apple",
		Price:       150,
		ImgUrl:      "https://cdn.britannica.com/22/187222-050-07B17FB6/apples-on-a-tree-branch.jpg",
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow, I Love Banana",
		Price:       80,
		ImgUrl:      "https://www.realsimple.com/thmb/3vUPZvQKPoN2G8Enwj12HTWYRac=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/banana-hack-GettyImages-1141226184-63241283ec5e4cd289290d40d0471c3c.jpg",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Grape",
		Description: "Grape is purple, I Love Grape",
		Price:       120,
		ImgUrl:      "https://png.pngtree.com/png-clipart/20250104/original/pngtree-delicious-black-grapes-png-image_20004046.png",
	}

	prd5 := &Product{
		ID:          5,
		Title:       "Strawberry",
		Description: "Strawberry is red, I Love Strawberry",
		Price:       200,
		ImgUrl:      "https://cdn.mos.cms.futurecdn.net/4wwQNKxhra9z9oUaPfwkP3.jpg",
	}

	prd6 := &Product{
		ID:          6,
		Title:       "Mango",
		Description: "Mango is yellow, I Love Mango",
		Price:       180,
		ImgUrl:      "https://www.melissas.com/cdn/shop/files/4-pounds-image-of-honey-mangos-fruit-1125637415_512x512.jpg?v=1738768090",
	}

	r.productList = append(r.productList, prd1, prd2, prd3, prd4, prd5, prd6)

}
