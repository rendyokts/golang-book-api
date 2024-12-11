package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (*Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, BookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (*Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)
	newBook, err := s.repository.Update(book)
	return newBook, err
}
func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
