package adapter

import "context"

type DataHandler interface {
	Create(ctx context.Context, value interface{}) error
	Update(ctx context.Context, value interface{}) error
	Delete(ctx context.Context, value interface{}) error
	Query(value interface{}, sql string, params ...interface{}) error
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
