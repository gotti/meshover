// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/ip.proto

package spec

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

// Validate checks the field values on AddressIPv4 with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddressIPv4) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressIPv4 with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddressIPv4MultiError, or
// nil if none found.
func (m *AddressIPv4) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressIPv4) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetIpaddress()); ip == nil || ip.To4() == nil {
		err := AddressIPv4ValidationError{
			field:  "Ipaddress",
			reason: "value must be a valid IPv4 address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddressIPv4MultiError(errors)
	}

	return nil
}

// AddressIPv4MultiError is an error wrapping multiple validation errors
// returned by AddressIPv4.ValidateAll() if the designated constraints aren't met.
type AddressIPv4MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressIPv4MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressIPv4MultiError) AllErrors() []error { return m }

// AddressIPv4ValidationError is the validation error returned by
// AddressIPv4.Validate if the designated constraints aren't met.
type AddressIPv4ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressIPv4ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressIPv4ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressIPv4ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressIPv4ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressIPv4ValidationError) ErrorName() string { return "AddressIPv4ValidationError" }

// Error satisfies the builtin error interface
func (e AddressIPv4ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressIPv4.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressIPv4ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressIPv4ValidationError{}

// Validate checks the field values on AddressIPv6 with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddressIPv6) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressIPv6 with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddressIPv6MultiError, or
// nil if none found.
func (m *AddressIPv6) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressIPv6) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if ip := net.ParseIP(m.GetIpaddress()); ip == nil || ip.To4() != nil {
		err := AddressIPv6ValidationError{
			field:  "Ipaddress",
			reason: "value must be a valid IPv6 address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddressIPv6MultiError(errors)
	}

	return nil
}

// AddressIPv6MultiError is an error wrapping multiple validation errors
// returned by AddressIPv6.ValidateAll() if the designated constraints aren't met.
type AddressIPv6MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressIPv6MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressIPv6MultiError) AllErrors() []error { return m }

// AddressIPv6ValidationError is the validation error returned by
// AddressIPv6.Validate if the designated constraints aren't met.
type AddressIPv6ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressIPv6ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressIPv6ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressIPv6ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressIPv6ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressIPv6ValidationError) ErrorName() string { return "AddressIPv6ValidationError" }

// Error satisfies the builtin error interface
func (e AddressIPv6ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressIPv6.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressIPv6ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressIPv6ValidationError{}

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AddressMultiError, or nil if none found.
func (m *Address) ValidateAll() error {
	return m.validate(true)
}

func (m *Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Ipaddress.(type) {

	case *Address_AddressIPv4:

		if all {
			switch v := interface{}(m.GetAddressIPv4()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "AddressIPv4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "AddressIPv4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddressIPv4()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressValidationError{
					field:  "AddressIPv4",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Address_AddressIPv6:

		if all {
			switch v := interface{}(m.GetAddressIPv6()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "AddressIPv6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "AddressIPv6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddressIPv6()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressValidationError{
					field:  "AddressIPv6",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := AddressValidationError{
			field:  "Ipaddress",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return AddressMultiError(errors)
	}

	return nil
}

// AddressMultiError is an error wrapping multiple validation errors returned
// by Address.ValidateAll() if the designated constraints aren't met.
type AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressMultiError) AllErrors() []error { return m }

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}

// Validate checks the field values on AddressCIDRIPv4 with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddressCIDRIPv4) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressCIDRIPv4 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddressCIDRIPv4MultiError, or nil if none found.
func (m *AddressCIDRIPv4) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressCIDRIPv4) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetIpaddress() == nil {
		err := AddressCIDRIPv4ValidationError{
			field:  "Ipaddress",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetIpaddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddressCIDRIPv4ValidationError{
					field:  "Ipaddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddressCIDRIPv4ValidationError{
					field:  "Ipaddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIpaddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddressCIDRIPv4ValidationError{
				field:  "Ipaddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if val := m.GetMask(); val < 0 || val > 32 {
		err := AddressCIDRIPv4ValidationError{
			field:  "Mask",
			reason: "value must be inside range [0, 32]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddressCIDRIPv4MultiError(errors)
	}

	return nil
}

// AddressCIDRIPv4MultiError is an error wrapping multiple validation errors
// returned by AddressCIDRIPv4.ValidateAll() if the designated constraints
// aren't met.
type AddressCIDRIPv4MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressCIDRIPv4MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressCIDRIPv4MultiError) AllErrors() []error { return m }

// AddressCIDRIPv4ValidationError is the validation error returned by
// AddressCIDRIPv4.Validate if the designated constraints aren't met.
type AddressCIDRIPv4ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressCIDRIPv4ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressCIDRIPv4ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressCIDRIPv4ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressCIDRIPv4ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressCIDRIPv4ValidationError) ErrorName() string { return "AddressCIDRIPv4ValidationError" }

// Error satisfies the builtin error interface
func (e AddressCIDRIPv4ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressCIDRIPv4.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressCIDRIPv4ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressCIDRIPv4ValidationError{}

// Validate checks the field values on AddressCIDRIPv6 with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddressCIDRIPv6) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressCIDRIPv6 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddressCIDRIPv6MultiError, or nil if none found.
func (m *AddressCIDRIPv6) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressCIDRIPv6) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetIpaddress() == nil {
		err := AddressCIDRIPv6ValidationError{
			field:  "Ipaddress",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetIpaddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddressCIDRIPv6ValidationError{
					field:  "Ipaddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddressCIDRIPv6ValidationError{
					field:  "Ipaddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIpaddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddressCIDRIPv6ValidationError{
				field:  "Ipaddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if val := m.GetMask(); val < 0 || val > 128 {
		err := AddressCIDRIPv6ValidationError{
			field:  "Mask",
			reason: "value must be inside range [0, 128]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddressCIDRIPv6MultiError(errors)
	}

	return nil
}

// AddressCIDRIPv6MultiError is an error wrapping multiple validation errors
// returned by AddressCIDRIPv6.ValidateAll() if the designated constraints
// aren't met.
type AddressCIDRIPv6MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressCIDRIPv6MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressCIDRIPv6MultiError) AllErrors() []error { return m }

// AddressCIDRIPv6ValidationError is the validation error returned by
// AddressCIDRIPv6.Validate if the designated constraints aren't met.
type AddressCIDRIPv6ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressCIDRIPv6ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressCIDRIPv6ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressCIDRIPv6ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressCIDRIPv6ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressCIDRIPv6ValidationError) ErrorName() string { return "AddressCIDRIPv6ValidationError" }

// Error satisfies the builtin error interface
func (e AddressCIDRIPv6ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressCIDRIPv6.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressCIDRIPv6ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressCIDRIPv6ValidationError{}

// Validate checks the field values on AddressCIDR with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddressCIDR) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressCIDR with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddressCIDRMultiError, or
// nil if none found.
func (m *AddressCIDR) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressCIDR) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Addresscidr.(type) {

	case *AddressCIDR_AddressCIDRIPv4:

		if all {
			switch v := interface{}(m.GetAddressCIDRIPv4()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressCIDRValidationError{
						field:  "AddressCIDRIPv4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressCIDRValidationError{
						field:  "AddressCIDRIPv4",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddressCIDRIPv4()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressCIDRValidationError{
					field:  "AddressCIDRIPv4",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AddressCIDR_AddressCIDRIPv6:

		if all {
			switch v := interface{}(m.GetAddressCIDRIPv6()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressCIDRValidationError{
						field:  "AddressCIDRIPv6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressCIDRValidationError{
						field:  "AddressCIDRIPv6",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddressCIDRIPv6()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressCIDRValidationError{
					field:  "AddressCIDRIPv6",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := AddressCIDRValidationError{
			field:  "Addresscidr",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return AddressCIDRMultiError(errors)
	}

	return nil
}

// AddressCIDRMultiError is an error wrapping multiple validation errors
// returned by AddressCIDR.ValidateAll() if the designated constraints aren't met.
type AddressCIDRMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressCIDRMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressCIDRMultiError) AllErrors() []error { return m }

// AddressCIDRValidationError is the validation error returned by
// AddressCIDR.Validate if the designated constraints aren't met.
type AddressCIDRValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressCIDRValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressCIDRValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressCIDRValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressCIDRValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressCIDRValidationError) ErrorName() string { return "AddressCIDRValidationError" }

// Error satisfies the builtin error interface
func (e AddressCIDRValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressCIDR.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressCIDRValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressCIDRValidationError{}

// Validate checks the field values on ASN with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *ASN) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ASN with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ASNMultiError, or nil if none found.
func (m *ASN) ValidateAll() error {
	return m.validate(true)
}

func (m *ASN) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Number

	if len(errors) > 0 {
		return ASNMultiError(errors)
	}

	return nil
}

// ASNMultiError is an error wrapping multiple validation errors returned by
// ASN.ValidateAll() if the designated constraints aren't met.
type ASNMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ASNMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ASNMultiError) AllErrors() []error { return m }

// ASNValidationError is the validation error returned by ASN.Validate if the
// designated constraints aren't met.
type ASNValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ASNValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ASNValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ASNValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ASNValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ASNValidationError) ErrorName() string { return "ASNValidationError" }

// Error satisfies the builtin error interface
func (e ASNValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sASN.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ASNValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ASNValidationError{}
