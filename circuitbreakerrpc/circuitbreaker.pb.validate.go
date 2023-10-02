// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: circuitbreaker.proto

package circuitbreakerrpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetInfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetInfoRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetInfoRequestValidationError is the validation error returned by
// GetInfoRequest.Validate if the designated constraints aren't met.
type GetInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoRequestValidationError) ErrorName() string { return "GetInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoRequestValidationError{}

// Validate checks the field values on GetInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetInfoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeKey

	// no validation rules for NodeAlias

	// no validation rules for NodeVersion

	// no validation rules for Version

	return nil
}

// GetInfoResponseValidationError is the validation error returned by
// GetInfoResponse.Validate if the designated constraints aren't met.
type GetInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoResponseValidationError) ErrorName() string { return "GetInfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoResponseValidationError{}

// Validate checks the field values on ClearLimitsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClearLimitsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ClearLimitsRequestValidationError is the validation error returned by
// ClearLimitsRequest.Validate if the designated constraints aren't met.
type ClearLimitsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClearLimitsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClearLimitsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClearLimitsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClearLimitsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClearLimitsRequestValidationError) ErrorName() string {
	return "ClearLimitsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ClearLimitsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClearLimitsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClearLimitsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClearLimitsRequestValidationError{}

// Validate checks the field values on ClearLimitsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClearLimitsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ClearLimitsResponseValidationError is the validation error returned by
// ClearLimitsResponse.Validate if the designated constraints aren't met.
type ClearLimitsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClearLimitsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClearLimitsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClearLimitsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClearLimitsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClearLimitsResponseValidationError) ErrorName() string {
	return "ClearLimitsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ClearLimitsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClearLimitsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClearLimitsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClearLimitsResponseValidationError{}

// Validate checks the field values on UpdateLimitsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateLimitsRequest) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetLimits() {
		_ = val

		// no validation rules for Limits[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateLimitsRequestValidationError{
					field:  fmt.Sprintf("Limits[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UpdateLimitsRequestValidationError is the validation error returned by
// UpdateLimitsRequest.Validate if the designated constraints aren't met.
type UpdateLimitsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitsRequestValidationError) ErrorName() string {
	return "UpdateLimitsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitsRequestValidationError{}

// Validate checks the field values on UpdateLimitsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateLimitsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateLimitsResponseValidationError is the validation error returned by
// UpdateLimitsResponse.Validate if the designated constraints aren't met.
type UpdateLimitsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitsResponseValidationError) ErrorName() string {
	return "UpdateLimitsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitsResponseValidationError{}

// Validate checks the field values on UpdateDefaultLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDefaultLimitRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateDefaultLimitRequestValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateDefaultLimitRequestValidationError is the validation error returned by
// UpdateDefaultLimitRequest.Validate if the designated constraints aren't met.
type UpdateDefaultLimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDefaultLimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDefaultLimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDefaultLimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDefaultLimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDefaultLimitRequestValidationError) ErrorName() string {
	return "UpdateDefaultLimitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDefaultLimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDefaultLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDefaultLimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDefaultLimitRequestValidationError{}

// Validate checks the field values on UpdateDefaultLimitResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDefaultLimitResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateDefaultLimitResponseValidationError is the validation error returned
// by UpdateDefaultLimitResponse.Validate if the designated constraints aren't met.
type UpdateDefaultLimitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDefaultLimitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDefaultLimitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDefaultLimitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDefaultLimitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDefaultLimitResponseValidationError) ErrorName() string {
	return "UpdateDefaultLimitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDefaultLimitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDefaultLimitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDefaultLimitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDefaultLimitResponseValidationError{}

// Validate checks the field values on ListLimitsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListLimitsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListLimitsRequestValidationError is the validation error returned by
// ListLimitsRequest.Validate if the designated constraints aren't met.
type ListLimitsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitsRequestValidationError) ErrorName() string {
	return "ListLimitsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitsRequestValidationError{}

// Validate checks the field values on ListLimitsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListLimitsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetLimits() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListLimitsResponseValidationError{
					field:  fmt.Sprintf("Limits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetDefaultLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListLimitsResponseValidationError{
				field:  "DefaultLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListLimitsResponseValidationError is the validation error returned by
// ListLimitsResponse.Validate if the designated constraints aren't met.
type ListLimitsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitsResponseValidationError) ErrorName() string {
	return "ListLimitsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitsResponseValidationError{}

// Validate checks the field values on NodeLimit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NodeLimit) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Node

	// no validation rules for Alias

	if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeLimitValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCounter_1H()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeLimitValidationError{
				field:  "Counter_1H",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCounter_24H()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeLimitValidationError{
				field:  "Counter_24H",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for QueueLen

	// no validation rules for PendingHtlcCount

	return nil
}

// NodeLimitValidationError is the validation error returned by
// NodeLimit.Validate if the designated constraints aren't met.
type NodeLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeLimitValidationError) ErrorName() string { return "NodeLimitValidationError" }

// Error satisfies the builtin error interface
func (e NodeLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeLimitValidationError{}

// Validate checks the field values on Limit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Limit) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MaxHourlyRate

	// no validation rules for MaxPending

	// no validation rules for Mode

	return nil
}

// LimitValidationError is the validation error returned by Limit.Validate if
// the designated constraints aren't met.
type LimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimitValidationError) ErrorName() string { return "LimitValidationError" }

// Error satisfies the builtin error interface
func (e LimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimitValidationError{}

// Validate checks the field values on Counter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Counter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Fail

	// no validation rules for Success

	// no validation rules for Reject

	return nil
}

// CounterValidationError is the validation error returned by Counter.Validate
// if the designated constraints aren't met.
type CounterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CounterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CounterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CounterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CounterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CounterValidationError) ErrorName() string { return "CounterValidationError" }

// Error satisfies the builtin error interface
func (e CounterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCounter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CounterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CounterValidationError{}

// Validate checks the field values on ListForwardingHistoryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListForwardingHistoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AddStartTimeNs

	// no validation rules for AddEndTimeNs

	return nil
}

// ListForwardingHistoryRequestValidationError is the validation error returned
// by ListForwardingHistoryRequest.Validate if the designated constraints
// aren't met.
type ListForwardingHistoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListForwardingHistoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListForwardingHistoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListForwardingHistoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListForwardingHistoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListForwardingHistoryRequestValidationError) ErrorName() string {
	return "ListForwardingHistoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListForwardingHistoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListForwardingHistoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListForwardingHistoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListForwardingHistoryRequestValidationError{}

// Validate checks the field values on ListForwardingHistoryResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListForwardingHistoryResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetForwards() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListForwardingHistoryResponseValidationError{
					field:  fmt.Sprintf("Forwards[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListForwardingHistoryResponseValidationError is the validation error
// returned by ListForwardingHistoryResponse.Validate if the designated
// constraints aren't met.
type ListForwardingHistoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListForwardingHistoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListForwardingHistoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListForwardingHistoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListForwardingHistoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListForwardingHistoryResponseValidationError) ErrorName() string {
	return "ListForwardingHistoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListForwardingHistoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListForwardingHistoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListForwardingHistoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListForwardingHistoryResponseValidationError{}

// Validate checks the field values on CircuitKey with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CircuitKey) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ShortChannelId

	// no validation rules for HtlcIndex

	return nil
}

// CircuitKeyValidationError is the validation error returned by
// CircuitKey.Validate if the designated constraints aren't met.
type CircuitKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitKeyValidationError) ErrorName() string { return "CircuitKeyValidationError" }

// Error satisfies the builtin error interface
func (e CircuitKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitKeyValidationError{}

// Validate checks the field values on Forward with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Forward) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AddTimeNs

	// no validation rules for ResolveTimeNs

	// no validation rules for Settled

	// no validation rules for IncomingAmount

	// no validation rules for OutgoingAmount

	// no validation rules for IncomingPeer

	if v, ok := interface{}(m.GetIncomingCircuit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForwardValidationError{
				field:  "IncomingCircuit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OutgoingPeer

	if v, ok := interface{}(m.GetOutgoingCircuit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForwardValidationError{
				field:  "OutgoingCircuit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ForwardValidationError is the validation error returned by Forward.Validate
// if the designated constraints aren't met.
type ForwardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForwardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForwardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForwardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForwardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForwardValidationError) ErrorName() string { return "ForwardValidationError" }

// Error satisfies the builtin error interface
func (e ForwardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForward.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForwardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForwardValidationError{}
