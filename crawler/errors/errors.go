package errors

import (
	"errors"
	"fmt"
)

type CrawlerError struct {
	Code    int
	Message string
	Err     error
}

func (e *CrawlerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *CrawlerError) Unwrap() error {
	return e.Err
}

func NewCrawlerError(code int, message string) *CrawlerError {
	return &CrawlerError{
		Code:    code,
		Message: message,
	}
}

func NewCrawlerErrorWithCause(code int, message string, err error) *CrawlerError {
	return &CrawlerError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

const (
	ErrCodeFetchFailed       = 1001
	ErrCodeParseFailed       = 1002
	ErrCodeSaveFailed        = 1003
	ErrCodeInvalidURL        = 1004
	ErrCodeRateLimitExceeded = 1005
	ErrCodeTimeout           = 1006
)

func IsCrawlerError(err error) bool {
	var ce *CrawlerError
	return errors.As(err, &ce)
}

func GetCrawlerError(err error) (*CrawlerError, bool) {
	var ce *CrawlerError
	if errors.As(err, &ce) {
		return ce, true
	}
	return nil, false
}

func IsFetchError(err error) bool {
	ce, ok := GetCrawlerError(err)
	return ok && ce.Code == ErrCodeFetchFailed
}

func IsParseError(err error) bool {
	ce, ok := GetCrawlerError(err)
	return ok && ce.Code == ErrCodeParseFailed
}
