package payment

// import (
// 	"github.com/MCPutro/goCoffeShop/internal/domain"
// 	"gorm.io/gorm"
// )

// type paymentRepository struct {
// }

// func NewPaymentRepository(db *gorm.DB) Repository {
// 	return &paymentRepository{db: db}
// }

// func (r *paymentRepository) Create(payment *domain.Payment) error {
// 	return r.db.Create(payment).Error
// }

// func (r *paymentRepository) FindById(id string) (*domain.Payment, error) {
// 	var payment domain.Payment
// 	err := r.db.First(&payment, "id = ?", id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &payment, nil
// }

// func (r *paymentRepository) FindAll() ([]domain.Payment, error) {
// 	var payments []domain.Payment
// 	err := r.db.Find(&payments).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return payments, nil
// }

// func (r *paymentRepository) Update(payment *domain.Payment) error {
// 	return r.db.Save(payment).Error
// }

// func (r *paymentRepository) Delete(id string) error {
// 	return r.db.Delete(&domain.Payment{}, id).Error
// }
