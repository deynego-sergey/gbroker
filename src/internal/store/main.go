package store

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/deynego-sergey/gbroker/internal/model"
)

type Queue struct {
	name string
	file string
	//
	buf    *bytes.Buffer
	tail   *os.File // указател для добавления сообщений
	head   *os.File //  читаем через этот
	update *os.File // обновляем  статусы

	// save timer
	timer *time.Timer
}

func NewQueue(name string) *Queue {

	return &Queue{}
}

func (q *Queue) init() error {
	// 1. chek files
	// 2.
}

func (q *Queue) Run(ctx context.Context) error {

}

func (q *Queue) Update(id uuid.UUID) error {
	return nil
}

func (q *Queue) Push(m model.QueMessage) (uuid.UUID, error) {

	q.buf.Write()

}
func (q *Queue) Pop() (model.QueMessage, error) {}

func (q *Queue) Close() error {
	return nil
}

func (q *Queue) startAppend() error {

	time.AfterFunc()
	go func() {

	}()
}
