package calculation

import "errors"

// Ошибка для деления на ноль
var ErrDivideByZero = errors.New("divide by 0")

// Ошибка для недопустимого выражения
var ErrInvalidExpression = errors.New("expression is not valid")
