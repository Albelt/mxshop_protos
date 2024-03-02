// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: albelt/stock_srv/msg.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stock) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stock with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StockMultiError, or nil if none found.
func (m *Stock) ValidateAll() error {
	return m.validate(true)
}

func (m *Stock) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GoodsId

	// no validation rules for Number

	if len(errors) > 0 {
		return StockMultiError(errors)
	}

	return nil
}

// StockMultiError is an error wrapping multiple validation errors returned by
// Stock.ValidateAll() if the designated constraints aren't met.
type StockMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockMultiError) AllErrors() []error { return m }

// StockValidationError is the validation error returned by Stock.Validate if
// the designated constraints aren't met.
type StockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockValidationError) ErrorName() string { return "StockValidationError" }

// Error satisfies the builtin error interface
func (e StockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockValidationError{}

// Validate checks the field values on OrderDetails with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderDetails) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderDetails with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderDetailsMultiError, or
// nil if none found.
func (m *OrderDetails) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderDetails) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderSn

	if len(errors) > 0 {
		return OrderDetailsMultiError(errors)
	}

	return nil
}

// OrderDetailsMultiError is an error wrapping multiple validation errors
// returned by OrderDetails.ValidateAll() if the designated constraints aren't met.
type OrderDetailsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderDetailsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderDetailsMultiError) AllErrors() []error { return m }

// OrderDetailsValidationError is the validation error returned by
// OrderDetails.Validate if the designated constraints aren't met.
type OrderDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderDetailsValidationError) ErrorName() string { return "OrderDetailsValidationError" }

// Error satisfies the builtin error interface
func (e OrderDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderDetailsValidationError{}
