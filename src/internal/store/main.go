package store

import (
	"context"
	"os"

	"github.com/google/uuid"

	"github.com/deynego-sergey/gbroker/internal/model"
)

type Queue struct {
	name string
	file string
	//
	tail   *os.File // указател для добавления сообщений
	head   *os.File //  читаем через этот
	update *os.File // обновляем  статусы
}

func NewQueue(name string) *Queue {

	return &Queue{}
}

func (q *Queue) init() error {

}

func (q *Queue) Run(ctx context.Context) error {

}

func (q *Queue) Update(id uuid.UUID) error {
	return nil
}

func (q *Queue) Push(m model.QueMessage) (uuid.UUID, error) {}
func (q *Queue) Pop() (model.QueMessage, error)             {}

func (q *Queue) Close() error {
	return nil
}
