package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nikita-Smirnov-idk/gOZON/order-service/internal/domain"
	"github.com/Nikita-Smirnov-idk/gOZON/order-service/internal/infrastructure/repositories"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrdersRepo struct {
	pool *pgxpool.Pool
}

func NewOrdersRepo(pool *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{pool: pool}
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, order *domain.Order) error {
	query := `
        INSERT INTO orders (id, user_id, amount, description, status)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.pool.Exec(ctx, query,
		order.GetId(),
		order.GetUserId(),
		order.GetAmount(),
		order.GetDescription(),
		order.GetStatus(),
	)

	return err
}

func (r *OrdersRepo) ListOrdersByUser(ctx context.Context, userID uuid.UUID) ([]domain.Order, error) {
	query := `
        SELECT id, user_id, amount, description, status
        FROM orders 
        WHERE user_id = $1
    `

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders by user: %w", err)
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		order, rowErr := r.scanOrder(rows)

		if rowErr != nil {
			if errors.Is(rowErr, repositories.ErrNotFound) {
				return nil, repositories.ErrNotFound
			}
			return nil, fmt.Errorf("failed to scan row")
		}

		orders = append(orders, *order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return orders, nil
}

func (r *OrdersRepo) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	query := `
        SELECT id, user_id, amount, description, status
        FROM orders
        WHERE id = $1
    `

	row := r.pool.QueryRow(ctx, query, orderID)
	return r.scanOrder(row)
}

func (r *OrdersRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status domain.OrderStatus) error {
	query := `
        UPDATE orders 
        SET status = $1
        WHERE id = $2
    `

	result, err := r.pool.Exec(ctx, query, status, orderID)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return repositories.ErrNotFound
	}

	return nil
}

func (r *OrdersRepo) scanOrder(row pgx.Row) (*domain.Order, error) {
	var (
		id          uuid.UUID
		userId      uuid.UUID
		amount      int
		description string
		statusInt   int
	)

	err := row.Scan(&id, &userId, &amount, &description, &statusInt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repositories.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan order: %w", err)
	}

	order := domain.NewOrder(id, userId, amount, description, domain.OrderStatus(statusInt))

	return order, nil
}
