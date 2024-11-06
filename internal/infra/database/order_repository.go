package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marcioaraujo/pos-cleanarch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	var orders []entity.Order

	res, err := r.Db.Query("Select * from orders")

	if err != nil {
		return []entity.Order{}, err
	}

	defer res.Close()

	for res.Next() {

		var order entity.Order
		err := res.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", order)

		orders = append(orders, order)
	}

	return orders, nil
}
