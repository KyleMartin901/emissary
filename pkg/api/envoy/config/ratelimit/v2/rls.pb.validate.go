// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/ratelimit/v2/rls.proto

package envoy_config_ratelimit_v2

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

// define the regex for a UUID once up-front
var _rls_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RateLimitServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return RateLimitServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetGrpcService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RateLimitServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RateLimitServiceConfigValidationError is the validation error returned by
// RateLimitServiceConfig.Validate if the designated constraints aren't met.
type RateLimitServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitServiceConfigValidationError) ErrorName() string {
	return "RateLimitServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitServiceConfigValidationError{}
