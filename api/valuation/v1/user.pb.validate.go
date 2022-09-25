// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: valuation/v1/user.proto

package valuation

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

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserInfoMultiError, or nil
// if none found.
func (m *UserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() != 0 {

		if m.GetId() <= 0 {
			err := UserInfoValidationError{
				field:  "Id",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if l := utf8.RuneCountInString(m.GetUsername()); l < 1 || l > 16 {
		err := UserInfoValidationError{
			field:  "Username",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 1 || l > 20 {
		err := UserInfoValidationError{
			field:  "Password",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetAvatar()); l < 1 || l > 64 {
		err := UserInfoValidationError{
			field:  "Avatar",
			reason: "value length must be between 1 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExpiration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserInfoValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserInfoValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserInfoValidationError{
				field:  "Expiration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserInfoValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserInfoValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserInfoValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserInfoMultiError(errors)
	}

	return nil
}

// UserInfoMultiError is an error wrapping multiple validation errors returned
// by UserInfo.ValidateAll() if the designated constraints aren't met.
type UserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoMultiError) AllErrors() []error { return m }

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}

// Validate checks the field values on UserLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserLoginRequestMultiError, or nil if none found.
func (m *UserLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUsername()); l < 1 || l > 16 {
		err := UserLoginRequestValidationError{
			field:  "Username",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 1 || l > 16 {
		err := UserLoginRequestValidationError{
			field:  "Password",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserLoginRequestMultiError(errors)
	}

	return nil
}

// UserLoginRequestMultiError is an error wrapping multiple validation errors
// returned by UserLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type UserLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserLoginRequestMultiError) AllErrors() []error { return m }

// UserLoginRequestValidationError is the validation error returned by
// UserLoginRequest.Validate if the designated constraints aren't met.
type UserLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserLoginRequestValidationError) ErrorName() string { return "UserLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserLoginRequestValidationError{}

// Validate checks the field values on UserLoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserLoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserLoginReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserLoginReplyMultiError,
// or nil if none found.
func (m *UserLoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UserLoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return UserLoginReplyMultiError(errors)
	}

	return nil
}

// UserLoginReplyMultiError is an error wrapping multiple validation errors
// returned by UserLoginReply.ValidateAll() if the designated constraints
// aren't met.
type UserLoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserLoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserLoginReplyMultiError) AllErrors() []error { return m }

// UserLoginReplyValidationError is the validation error returned by
// UserLoginReply.Validate if the designated constraints aren't met.
type UserLoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserLoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserLoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserLoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserLoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserLoginReplyValidationError) ErrorName() string { return "UserLoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserLoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserLoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserLoginReplyValidationError{}
