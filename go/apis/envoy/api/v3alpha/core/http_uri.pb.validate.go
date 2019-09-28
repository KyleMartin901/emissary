// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/core/http_uri.proto

package envoy_api_v3alpha_core

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on HttpUri with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpUri) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetUri()) < 1 {
		return HttpUriValidationError{
			field:  "Uri",
			reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetTimeout() == nil {
		return HttpUriValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
	}

	switch m.HttpUpstreamType.(type) {

	case *HttpUri_Cluster:

		if len(m.GetCluster()) < 1 {
			return HttpUriValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 bytes",
			}
		}

	default:
		return HttpUriValidationError{
			field:  "HttpUpstreamType",
			reason: "value is required",
		}

	}

	return nil
}

// HttpUriValidationError is the validation error returned by HttpUri.Validate
// if the designated constraints aren't met.
type HttpUriValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpUriValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpUriValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpUriValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpUriValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpUriValidationError) ErrorName() string { return "HttpUriValidationError" }

// Error satisfies the builtin error interface
func (e HttpUriValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpUri.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpUriValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpUriValidationError{}
