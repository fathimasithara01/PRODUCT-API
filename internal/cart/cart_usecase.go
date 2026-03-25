package cart

type Usecase interface {
	AddToCart(userID, productID uint, price float64) error
	GetCart(userID uint) ([]Cart, float64, error)
}

type usecase struct {
	repo Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{r}
}

func (u *usecase) AddToCart(userID, productID uint, price float64) error {

	item, _ := u.repo.GetItem(userID, productID)

	if item != nil && item.ID != 0 {
		item.Quantity += 1
		return u.repo.Update(item)
	}

	newItem := &Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  1,
		Price:     price,
	}

	return u.repo.Add(newItem)
}

func (u *usecase) GetCart(userID uint) ([]Cart, float64, error) {

	items, err := u.repo.GetByUser(userID)
	if err != nil {
		return nil, 0, err
	}

	var total float64

	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return items, total, nil
}
