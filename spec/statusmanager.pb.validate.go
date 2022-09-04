// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/statusmanager.proto

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

// Validate checks the field values on MinimumPeerStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MinimumPeerStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MinimumPeerStatus with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MinimumPeerStatusMultiError, or nil if none found.
func (m *MinimumPeerStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *MinimumPeerStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocalAS()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MinimumPeerStatusValidationError{
					field:  "LocalAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MinimumPeerStatusValidationError{
					field:  "LocalAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocalAS()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MinimumPeerStatusValidationError{
				field:  "LocalAS",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAddresses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MinimumPeerStatusValidationError{
						field:  fmt.Sprintf("Addresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MinimumPeerStatusValidationError{
						field:  fmt.Sprintf("Addresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MinimumPeerStatusValidationError{
					field:  fmt.Sprintf("Addresses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetEndpoint() == nil {
		err := MinimumPeerStatusValidationError{
			field:  "Endpoint",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEndpoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MinimumPeerStatusValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MinimumPeerStatusValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MinimumPeerStatusValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MinimumPeerStatusMultiError(errors)
	}

	return nil
}

// MinimumPeerStatusMultiError is an error wrapping multiple validation errors
// returned by MinimumPeerStatus.ValidateAll() if the designated constraints
// aren't met.
type MinimumPeerStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MinimumPeerStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MinimumPeerStatusMultiError) AllErrors() []error { return m }

// MinimumPeerStatusValidationError is the validation error returned by
// MinimumPeerStatus.Validate if the designated constraints aren't met.
type MinimumPeerStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MinimumPeerStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MinimumPeerStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MinimumPeerStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MinimumPeerStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MinimumPeerStatusValidationError) ErrorName() string {
	return "MinimumPeerStatusValidationError"
}

// Error satisfies the builtin error interface
func (e MinimumPeerStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMinimumPeerStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MinimumPeerStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MinimumPeerStatusValidationError{}

// Validate checks the field values on WireguardStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WireguardStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WireguardStatus with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WireguardStatusMultiError, or nil if none found.
func (m *WireguardStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *WireguardStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LatestHandshake

	// no validation rules for TxBytes

	// no validation rules for RxBytes

	if len(errors) > 0 {
		return WireguardStatusMultiError(errors)
	}

	return nil
}

// WireguardStatusMultiError is an error wrapping multiple validation errors
// returned by WireguardStatus.ValidateAll() if the designated constraints
// aren't met.
type WireguardStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WireguardStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WireguardStatusMultiError) AllErrors() []error { return m }

// WireguardStatusValidationError is the validation error returned by
// WireguardStatus.Validate if the designated constraints aren't met.
type WireguardStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WireguardStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WireguardStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WireguardStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WireguardStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WireguardStatusValidationError) ErrorName() string { return "WireguardStatusValidationError" }

// Error satisfies the builtin error interface
func (e WireguardStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWireguardStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WireguardStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WireguardStatusValidationError{}

// Validate checks the field values on BGPPeerStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BGPPeerStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BGPPeerStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BGPPeerStatusMultiError, or
// nil if none found.
func (m *BGPPeerStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *BGPPeerStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocalAS()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "LocalAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "LocalAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocalAS()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BGPPeerStatusValidationError{
				field:  "LocalAS",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRemoteAS()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "RemoteAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "RemoteAS",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRemoteAS()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BGPPeerStatusValidationError{
				field:  "RemoteAS",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBgpNeighborAddr()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "BgpNeighborAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BGPPeerStatusValidationError{
					field:  "BgpNeighborAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBgpNeighborAddr()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BGPPeerStatusValidationError{
				field:  "BgpNeighborAddr",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RemoteHostName

	// no validation rules for BGPState

	if len(errors) > 0 {
		return BGPPeerStatusMultiError(errors)
	}

	return nil
}

// BGPPeerStatusMultiError is an error wrapping multiple validation errors
// returned by BGPPeerStatus.ValidateAll() if the designated constraints
// aren't met.
type BGPPeerStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BGPPeerStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BGPPeerStatusMultiError) AllErrors() []error { return m }

// BGPPeerStatusValidationError is the validation error returned by
// BGPPeerStatus.Validate if the designated constraints aren't met.
type BGPPeerStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BGPPeerStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BGPPeerStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BGPPeerStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BGPPeerStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BGPPeerStatusValidationError) ErrorName() string { return "BGPPeerStatusValidationError" }

// Error satisfies the builtin error interface
func (e BGPPeerStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBGPPeerStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BGPPeerStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BGPPeerStatusValidationError{}

// Validate checks the field values on BGPStatus with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BGPStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BGPStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BGPStatusMultiError, or nil
// if none found.
func (m *BGPStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *BGPStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBGPPeers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BGPStatusValidationError{
						field:  fmt.Sprintf("BGPPeers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BGPStatusValidationError{
						field:  fmt.Sprintf("BGPPeers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BGPStatusValidationError{
					field:  fmt.Sprintf("BGPPeers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BGPStatusMultiError(errors)
	}

	return nil
}

// BGPStatusMultiError is an error wrapping multiple validation errors returned
// by BGPStatus.ValidateAll() if the designated constraints aren't met.
type BGPStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BGPStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BGPStatusMultiError) AllErrors() []error { return m }

// BGPStatusValidationError is the validation error returned by
// BGPStatus.Validate if the designated constraints aren't met.
type BGPStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BGPStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BGPStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BGPStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BGPStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BGPStatusValidationError) ErrorName() string { return "BGPStatusValidationError" }

// Error satisfies the builtin error interface
func (e BGPStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBGPStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BGPStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BGPStatusValidationError{}

// Validate checks the field values on StatusManagerPeerStatus with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StatusManagerPeerStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatusManagerPeerStatus with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatusManagerPeerStatusMultiError, or nil if none found.
func (m *StatusManagerPeerStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *StatusManagerPeerStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_StatusManagerPeerStatus_Hostname_Pattern.MatchString(m.GetHostname()) {
		err := StatusManagerPeerStatusValidationError{
			field:  "Hostname",
			reason: "value does not match regex pattern \"^[0-9a-fA-Z-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPeerStatus() == nil {
		err := StatusManagerPeerStatusValidationError{
			field:  "PeerStatus",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPeerStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "PeerStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "PeerStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPeerStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusManagerPeerStatusValidationError{
				field:  "PeerStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetWireguardStatus() == nil {
		err := StatusManagerPeerStatusValidationError{
			field:  "WireguardStatus",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetWireguardStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "WireguardStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "WireguardStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWireguardStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusManagerPeerStatusValidationError{
				field:  "WireguardStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetBgpStatus() == nil {
		err := StatusManagerPeerStatusValidationError{
			field:  "BgpStatus",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetBgpStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "BgpStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatusManagerPeerStatusValidationError{
					field:  "BgpStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBgpStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusManagerPeerStatusValidationError{
				field:  "BgpStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StatusManagerPeerStatusMultiError(errors)
	}

	return nil
}

// StatusManagerPeerStatusMultiError is an error wrapping multiple validation
// errors returned by StatusManagerPeerStatus.ValidateAll() if the designated
// constraints aren't met.
type StatusManagerPeerStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatusManagerPeerStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatusManagerPeerStatusMultiError) AllErrors() []error { return m }

// StatusManagerPeerStatusValidationError is the validation error returned by
// StatusManagerPeerStatus.Validate if the designated constraints aren't met.
type StatusManagerPeerStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusManagerPeerStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusManagerPeerStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusManagerPeerStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusManagerPeerStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusManagerPeerStatusValidationError) ErrorName() string {
	return "StatusManagerPeerStatusValidationError"
}

// Error satisfies the builtin error interface
func (e StatusManagerPeerStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusManagerPeerStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusManagerPeerStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusManagerPeerStatusValidationError{}

var _StatusManagerPeerStatus_Hostname_Pattern = regexp.MustCompile("^[0-9a-fA-Z-]+$")

// Validate checks the field values on RegisterStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterStatusRequestMultiError, or nil if none found.
func (m *RegisterStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStatus() == nil {
		err := RegisterStatusRequestValidationError{
			field:  "Status",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterStatusRequestValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterStatusRequestValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterStatusRequestValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterStatusRequestMultiError(errors)
	}

	return nil
}

// RegisterStatusRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterStatusRequestMultiError) AllErrors() []error { return m }

// RegisterStatusRequestValidationError is the validation error returned by
// RegisterStatusRequest.Validate if the designated constraints aren't met.
type RegisterStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterStatusRequestValidationError) ErrorName() string {
	return "RegisterStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterStatusRequestValidationError{}

// Validate checks the field values on RegisterStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterStatusResponseMultiError, or nil if none found.
func (m *RegisterStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterStatusResponseMultiError(errors)
	}

	return nil
}

// RegisterStatusResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterStatusResponseMultiError) AllErrors() []error { return m }

// RegisterStatusResponseValidationError is the validation error returned by
// RegisterStatusResponse.Validate if the designated constraints aren't met.
type RegisterStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterStatusResponseValidationError) ErrorName() string {
	return "RegisterStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterStatusResponseValidationError{}