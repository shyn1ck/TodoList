package errs

import "errors"

var (
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrOperationNotFound           = errors.New("ErrOperationNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrTaskNotFound                = errors.New("TaskNotFound")
	ErrTaskCreationFailed          = errors.New("TaskCreationFailed")
	ErrTaskUpdateFailed            = errors.New("TaskUpdateFailed")
	ErrTaskDeleteFailed            = errors.New("TaskDeleteFailed")
	ErrInvalidTaskID               = errors.New("InvalidTaskID")
	ErrTaskAlreadyCompleted        = errors.New("TaskAlreadyCompleted")
	ErrTaskAlreadyExists           = errors.New("TaskAlreadyExists")
)
