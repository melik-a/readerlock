package limit

import "errors"

// Необходимо предоставить четыре реализации интерфейса, используя разные возможности go.
// Для каждой реализации должны быть написаны тесты.
// Логика следующая: для Reader должно задаваться значение и число, которое устанавливает количество возможных чтений этого значения.
// Если число чтений превышено, то необходимо возвращать ошибку.

type Reader[T any] interface {
	Read() (v T, err error)
}

var ReadLimitExceededError = errors.New("reading limit exceeded")
