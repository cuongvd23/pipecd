// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/model/deployment_trace.proto

package model

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

// Validate checks the field values on DeploymentTrace with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeploymentTrace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentTrace with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentTraceMultiError, or nil if none found.
func (m *DeploymentTrace) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentTrace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := DeploymentTraceValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetProjectId()) < 1 {
		err := DeploymentTraceValidationError{
			field:  "ProjectId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Title

	// no validation rules for CommitMessage

	if utf8.RuneCountInString(m.GetCommitHash()) < 1 {
		err := DeploymentTraceValidationError{
			field:  "CommitHash",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCommitUrl()) < 1 {
		err := DeploymentTraceValidationError{
			field:  "CommitUrl",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCommitTimestamp() < 0 {
		err := DeploymentTraceValidationError{
			field:  "CommitTimestamp",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Author

	if m.GetCreatedAt() < 0 {
		err := DeploymentTraceValidationError{
			field:  "CreatedAt",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() < 0 {
		err := DeploymentTraceValidationError{
			field:  "UpdatedAt",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeploymentTraceMultiError(errors)
	}

	return nil
}

// DeploymentTraceMultiError is an error wrapping multiple validation errors
// returned by DeploymentTrace.ValidateAll() if the designated constraints
// aren't met.
type DeploymentTraceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentTraceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentTraceMultiError) AllErrors() []error { return m }

// DeploymentTraceValidationError is the validation error returned by
// DeploymentTrace.Validate if the designated constraints aren't met.
type DeploymentTraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentTraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentTraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentTraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentTraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentTraceValidationError) ErrorName() string { return "DeploymentTraceValidationError" }

// Error satisfies the builtin error interface
func (e DeploymentTraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentTraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentTraceValidationError{}
