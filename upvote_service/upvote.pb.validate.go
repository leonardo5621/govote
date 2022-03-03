// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protobuffers/upvote.proto

package upvote_service

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

// Validate checks the field values on VoteThreadRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *VoteThreadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VoteThreadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VoteThreadRequestMultiError, or nil if none found.
func (m *VoteThreadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VoteThreadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_VoteThreadRequest_UserId_Pattern.MatchString(m.GetUserId()) {
		err := VoteThreadRequestValidationError{
			field:  "UserId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _VoteThreadRequest_Votedir_InLookup[m.GetVotedir()]; !ok {
		err := VoteThreadRequestValidationError{
			field:  "Votedir",
			reason: "value must be in list [-1 1]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VoteThreadRequest_ThreadId_Pattern.MatchString(m.GetThreadId()) {
		err := VoteThreadRequestValidationError{
			field:  "ThreadId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return VoteThreadRequestMultiError(errors)
	}

	return nil
}

// VoteThreadRequestMultiError is an error wrapping multiple validation errors
// returned by VoteThreadRequest.ValidateAll() if the designated constraints
// aren't met.
type VoteThreadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VoteThreadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VoteThreadRequestMultiError) AllErrors() []error { return m }

// VoteThreadRequestValidationError is the validation error returned by
// VoteThreadRequest.Validate if the designated constraints aren't met.
type VoteThreadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VoteThreadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VoteThreadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VoteThreadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VoteThreadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VoteThreadRequestValidationError) ErrorName() string {
	return "VoteThreadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VoteThreadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVoteThreadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VoteThreadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VoteThreadRequestValidationError{}

var _VoteThreadRequest_UserId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

var _VoteThreadRequest_Votedir_InLookup = map[int32]struct{}{
	-1: {},
	1:  {},
}

var _VoteThreadRequest_ThreadId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

// Validate checks the field values on VoteThreadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VoteThreadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VoteThreadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VoteThreadResponseMultiError, or nil if none found.
func (m *VoteThreadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VoteThreadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Notification

	// no validation rules for Email

	if len(errors) > 0 {
		return VoteThreadResponseMultiError(errors)
	}

	return nil
}

// VoteThreadResponseMultiError is an error wrapping multiple validation errors
// returned by VoteThreadResponse.ValidateAll() if the designated constraints
// aren't met.
type VoteThreadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VoteThreadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VoteThreadResponseMultiError) AllErrors() []error { return m }

// VoteThreadResponseValidationError is the validation error returned by
// VoteThreadResponse.Validate if the designated constraints aren't met.
type VoteThreadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VoteThreadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VoteThreadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VoteThreadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VoteThreadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VoteThreadResponseValidationError) ErrorName() string {
	return "VoteThreadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VoteThreadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVoteThreadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VoteThreadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VoteThreadResponseValidationError{}

// Validate checks the field values on VoteCommentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VoteCommentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VoteCommentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VoteCommentRequestMultiError, or nil if none found.
func (m *VoteCommentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VoteCommentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_VoteCommentRequest_UserId_Pattern.MatchString(m.GetUserId()) {
		err := VoteCommentRequestValidationError{
			field:  "UserId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _VoteCommentRequest_Votedir_InLookup[m.GetVotedir()]; !ok {
		err := VoteCommentRequestValidationError{
			field:  "Votedir",
			reason: "value must be in list [-1 1]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_VoteCommentRequest_CommentId_Pattern.MatchString(m.GetCommentId()) {
		err := VoteCommentRequestValidationError{
			field:  "CommentId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return VoteCommentRequestMultiError(errors)
	}

	return nil
}

// VoteCommentRequestMultiError is an error wrapping multiple validation errors
// returned by VoteCommentRequest.ValidateAll() if the designated constraints
// aren't met.
type VoteCommentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VoteCommentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VoteCommentRequestMultiError) AllErrors() []error { return m }

// VoteCommentRequestValidationError is the validation error returned by
// VoteCommentRequest.Validate if the designated constraints aren't met.
type VoteCommentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VoteCommentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VoteCommentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VoteCommentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VoteCommentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VoteCommentRequestValidationError) ErrorName() string {
	return "VoteCommentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VoteCommentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVoteCommentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VoteCommentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VoteCommentRequestValidationError{}

var _VoteCommentRequest_UserId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

var _VoteCommentRequest_Votedir_InLookup = map[int32]struct{}{
	-1: {},
	1:  {},
}

var _VoteCommentRequest_CommentId_Pattern = regexp.MustCompile("^[A-Za-z0-9]*$")

// Validate checks the field values on VoteCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VoteCommentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VoteCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VoteCommentResponseMultiError, or nil if none found.
func (m *VoteCommentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VoteCommentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Notification

	if len(errors) > 0 {
		return VoteCommentResponseMultiError(errors)
	}

	return nil
}

// VoteCommentResponseMultiError is an error wrapping multiple validation
// errors returned by VoteCommentResponse.ValidateAll() if the designated
// constraints aren't met.
type VoteCommentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VoteCommentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VoteCommentResponseMultiError) AllErrors() []error { return m }

// VoteCommentResponseValidationError is the validation error returned by
// VoteCommentResponse.Validate if the designated constraints aren't met.
type VoteCommentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VoteCommentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VoteCommentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VoteCommentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VoteCommentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VoteCommentResponseValidationError) ErrorName() string {
	return "VoteCommentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VoteCommentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVoteCommentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VoteCommentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VoteCommentResponseValidationError{}
