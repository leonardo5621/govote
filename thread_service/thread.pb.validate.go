// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protobuffers/thread.proto

package thread_service

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

// Validate checks the field values on Thread with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Thread) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Thread with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ThreadMultiError, or nil if none found.
func (m *Thread) ValidateAll() error {
	return m.validate(true)
}

func (m *Thread) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetTitle()); l < 5 || l > 150 {
		err := ThreadValidationError{
			field:  "Title",
			reason: "value length must be between 5 and 150 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Thread_OwnerUserId_Pattern.MatchString(m.GetOwnerUserId()) {
		err := ThreadValidationError{
			field:  "OwnerUserId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Archived

	if utf8.RuneCountInString(m.GetDescription()) > 1000 {
		err := ThreadValidationError{
			field:  "Description",
			reason: "value length must be at most 1000 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for OwnerUserName

	if !_Thread_FirmId_Pattern.MatchString(m.GetFirmId()) {
		err := ThreadValidationError{
			field:  "FirmId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FirmName

	if len(errors) > 0 {
		return ThreadMultiError(errors)
	}

	return nil
}

// ThreadMultiError is an error wrapping multiple validation errors returned by
// Thread.ValidateAll() if the designated constraints aren't met.
type ThreadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThreadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThreadMultiError) AllErrors() []error { return m }

// ThreadValidationError is the validation error returned by Thread.Validate if
// the designated constraints aren't met.
type ThreadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThreadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThreadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThreadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThreadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThreadValidationError) ErrorName() string { return "ThreadValidationError" }

// Error satisfies the builtin error interface
func (e ThreadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThread.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThreadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThreadValidationError{}

var _Thread_OwnerUserId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

var _Thread_FirmId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

// Validate checks the field values on GetThreadRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetThreadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetThreadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetThreadRequestMultiError, or nil if none found.
func (m *GetThreadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetThreadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ThreadId

	if len(errors) > 0 {
		return GetThreadRequestMultiError(errors)
	}

	return nil
}

// GetThreadRequestMultiError is an error wrapping multiple validation errors
// returned by GetThreadRequest.ValidateAll() if the designated constraints
// aren't met.
type GetThreadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetThreadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetThreadRequestMultiError) AllErrors() []error { return m }

// GetThreadRequestValidationError is the validation error returned by
// GetThreadRequest.Validate if the designated constraints aren't met.
type GetThreadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetThreadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetThreadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetThreadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetThreadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetThreadRequestValidationError) ErrorName() string { return "GetThreadRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetThreadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetThreadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetThreadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetThreadRequestValidationError{}

// Validate checks the field values on GetThreadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetThreadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetThreadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetThreadResponseMultiError, or nil if none found.
func (m *GetThreadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetThreadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetThread()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetThreadResponseValidationError{
					field:  "Thread",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetThreadResponseValidationError{
					field:  "Thread",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetThread()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetThreadResponseValidationError{
				field:  "Thread",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetThreadResponseMultiError(errors)
	}

	return nil
}

// GetThreadResponseMultiError is an error wrapping multiple validation errors
// returned by GetThreadResponse.ValidateAll() if the designated constraints
// aren't met.
type GetThreadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetThreadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetThreadResponseMultiError) AllErrors() []error { return m }

// GetThreadResponseValidationError is the validation error returned by
// GetThreadResponse.Validate if the designated constraints aren't met.
type GetThreadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetThreadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetThreadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetThreadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetThreadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetThreadResponseValidationError) ErrorName() string {
	return "GetThreadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetThreadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetThreadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetThreadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetThreadResponseValidationError{}

// Validate checks the field values on CreateThreadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateThreadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateThreadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateThreadRequestMultiError, or nil if none found.
func (m *CreateThreadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateThreadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetThread()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateThreadRequestValidationError{
					field:  "Thread",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateThreadRequestValidationError{
					field:  "Thread",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetThread()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateThreadRequestValidationError{
				field:  "Thread",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateThreadRequestMultiError(errors)
	}

	return nil
}

// CreateThreadRequestMultiError is an error wrapping multiple validation
// errors returned by CreateThreadRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateThreadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateThreadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateThreadRequestMultiError) AllErrors() []error { return m }

// CreateThreadRequestValidationError is the validation error returned by
// CreateThreadRequest.Validate if the designated constraints aren't met.
type CreateThreadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateThreadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateThreadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateThreadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateThreadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateThreadRequestValidationError) ErrorName() string {
	return "CreateThreadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateThreadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateThreadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateThreadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateThreadRequestValidationError{}

// Validate checks the field values on CreateThreadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateThreadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateThreadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateThreadResponseMultiError, or nil if none found.
func (m *CreateThreadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateThreadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ThreadId

	if len(errors) > 0 {
		return CreateThreadResponseMultiError(errors)
	}

	return nil
}

// CreateThreadResponseMultiError is an error wrapping multiple validation
// errors returned by CreateThreadResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateThreadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateThreadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateThreadResponseMultiError) AllErrors() []error { return m }

// CreateThreadResponseValidationError is the validation error returned by
// CreateThreadResponse.Validate if the designated constraints aren't met.
type CreateThreadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateThreadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateThreadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateThreadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateThreadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateThreadResponseValidationError) ErrorName() string {
	return "CreateThreadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateThreadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateThreadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateThreadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateThreadResponseValidationError{}