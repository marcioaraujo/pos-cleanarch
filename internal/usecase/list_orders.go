package usecase

import (
	"github.com/marcioaraujo/pos-cleanarch/internal/entity"
	"github.com/marcioaraujo/pos-cleanarch/pkg/events"
)

type ListOrdersOutputDTO struct {
	Quantity int
	Orders   []entity.Order
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {

	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	dto := ListOrdersOutputDTO{
		Quantity: len(orders),
		Orders:   orders,
	}

	c.OrderListed.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderListed)

	return dto, nil
}
