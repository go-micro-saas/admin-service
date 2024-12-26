// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: app/admin-service/internal/conf/config.conf.proto

package conf

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

// Validate checks the field values on ServiceConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServiceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServiceConfigMultiError, or
// nil if none found.
func (m *ServiceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccountService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceConfigValidationError{
					field:  "AccountService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceConfigValidationError{
					field:  "AccountService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccountService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfigValidationError{
				field:  "AccountService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceConfigMultiError(errors)
	}

	return nil
}

// ServiceConfigMultiError is an error wrapping multiple validation errors
// returned by ServiceConfig.ValidateAll() if the designated constraints
// aren't met.
type ServiceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfigMultiError) AllErrors() []error { return m }

// ServiceConfigValidationError is the validation error returned by
// ServiceConfig.Validate if the designated constraints aren't met.
type ServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfigValidationError) ErrorName() string { return "ServiceConfigValidationError" }

// Error satisfies the builtin error interface
func (e ServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfigValidationError{}

// Validate checks the field values on ServiceConfig_AccountService with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServiceConfig_AccountService) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceConfig_AccountService with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceConfig_AccountServiceMultiError, or nil if none found.
func (m *ServiceConfig_AccountService) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig_AccountService) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSnowflake()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceConfig_AccountServiceValidationError{
					field:  "Snowflake",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceConfig_AccountServiceValidationError{
					field:  "Snowflake",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSnowflake()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfig_AccountServiceValidationError{
				field:  "Snowflake",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSendEmailCode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceConfig_AccountServiceValidationError{
					field:  "SendEmailCode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceConfig_AccountServiceValidationError{
					field:  "SendEmailCode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSendEmailCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfig_AccountServiceValidationError{
				field:  "SendEmailCode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceConfig_AccountServiceMultiError(errors)
	}

	return nil
}

// ServiceConfig_AccountServiceMultiError is an error wrapping multiple
// validation errors returned by ServiceConfig_AccountService.ValidateAll() if
// the designated constraints aren't met.
type ServiceConfig_AccountServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfig_AccountServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfig_AccountServiceMultiError) AllErrors() []error { return m }

// ServiceConfig_AccountServiceValidationError is the validation error returned
// by ServiceConfig_AccountService.Validate if the designated constraints
// aren't met.
type ServiceConfig_AccountServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfig_AccountServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfig_AccountServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfig_AccountServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfig_AccountServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfig_AccountServiceValidationError) ErrorName() string {
	return "ServiceConfig_AccountServiceValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceConfig_AccountServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig_AccountService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfig_AccountServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfig_AccountServiceValidationError{}

// Validate checks the field values on ServiceConfig_AccountService_Snowflake
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *ServiceConfig_AccountService_Snowflake) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// ServiceConfig_AccountService_Snowflake with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// ServiceConfig_AccountService_SnowflakeMultiError, or nil if none found.
func (m *ServiceConfig_AccountService_Snowflake) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig_AccountService_Snowflake) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetInstanceId()) < 1 {
		err := ServiceConfig_AccountService_SnowflakeValidationError{
			field:  "InstanceId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for InstanceName

	// no validation rules for Metadata

	if len(errors) > 0 {
		return ServiceConfig_AccountService_SnowflakeMultiError(errors)
	}

	return nil
}

// ServiceConfig_AccountService_SnowflakeMultiError is an error wrapping
// multiple validation errors returned by
// ServiceConfig_AccountService_Snowflake.ValidateAll() if the designated
// constraints aren't met.
type ServiceConfig_AccountService_SnowflakeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfig_AccountService_SnowflakeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfig_AccountService_SnowflakeMultiError) AllErrors() []error { return m }

// ServiceConfig_AccountService_SnowflakeValidationError is the validation
// error returned by ServiceConfig_AccountService_Snowflake.Validate if the
// designated constraints aren't met.
type ServiceConfig_AccountService_SnowflakeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfig_AccountService_SnowflakeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfig_AccountService_SnowflakeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfig_AccountService_SnowflakeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfig_AccountService_SnowflakeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfig_AccountService_SnowflakeValidationError) ErrorName() string {
	return "ServiceConfig_AccountService_SnowflakeValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceConfig_AccountService_SnowflakeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig_AccountService_Snowflake.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfig_AccountService_SnowflakeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfig_AccountService_SnowflakeValidationError{}

// Validate checks the field values on
// ServiceConfig_AccountService_SendEmailCode with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServiceConfig_AccountService_SendEmailCode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// ServiceConfig_AccountService_SendEmailCode with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// ServiceConfig_AccountService_SendEmailCodeMultiError, or nil if none found.
func (m *ServiceConfig_AccountService_SendEmailCode) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig_AccountService_SendEmailCode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	if utf8.RuneCountInString(m.GetIssuer()) < 1 {
		err := ServiceConfig_AccountService_SendEmailCodeValidationError{
			field:  "Issuer",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSubject()) < 1 {
		err := ServiceConfig_AccountService_SendEmailCodeValidationError{
			field:  "Subject",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHost()) < 1 {
		err := ServiceConfig_AccountService_SendEmailCodeValidationError{
			field:  "Host",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPort(); val < 1 || val > 65535 {
		err := ServiceConfig_AccountService_SendEmailCodeValidationError{
			field:  "Port",
			reason: "value must be inside range [1, 65535]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Username

	// no validation rules for Password

	if utf8.RuneCountInString(m.GetFrom()) < 1 {
		err := ServiceConfig_AccountService_SendEmailCodeValidationError{
			field:  "From",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ServiceConfig_AccountService_SendEmailCodeMultiError(errors)
	}

	return nil
}

// ServiceConfig_AccountService_SendEmailCodeMultiError is an error wrapping
// multiple validation errors returned by
// ServiceConfig_AccountService_SendEmailCode.ValidateAll() if the designated
// constraints aren't met.
type ServiceConfig_AccountService_SendEmailCodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfig_AccountService_SendEmailCodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfig_AccountService_SendEmailCodeMultiError) AllErrors() []error { return m }

// ServiceConfig_AccountService_SendEmailCodeValidationError is the validation
// error returned by ServiceConfig_AccountService_SendEmailCode.Validate if
// the designated constraints aren't met.
type ServiceConfig_AccountService_SendEmailCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) ErrorName() string {
	return "ServiceConfig_AccountService_SendEmailCodeValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceConfig_AccountService_SendEmailCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig_AccountService_SendEmailCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfig_AccountService_SendEmailCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfig_AccountService_SendEmailCodeValidationError{}