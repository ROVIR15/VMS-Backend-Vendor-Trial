// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/null/v8/convert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type VendorLogSection string

// Enum values for VendorLogSection
const (
	VendorLogSectionVendor        VendorLogSection = "vendor"
	VendorLogSectionVendorAddress VendorLogSection = "vendor_address"
	VendorLogSectionVendorPhone   VendorLogSection = "vendor_phone"
)

func AllVendorLogSection() []VendorLogSection {
	return []VendorLogSection{
		VendorLogSectionVendor,
		VendorLogSectionVendorAddress,
		VendorLogSectionVendorPhone,
	}
}

func (e VendorLogSection) IsValid() error {
	switch e {
	case VendorLogSectionVendor, VendorLogSectionVendorAddress, VendorLogSectionVendorPhone:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VendorLogSection) String() string {
	return string(e)
}

func (e VendorLogSection) Ordinal() int {
	switch e {
	case VendorLogSectionVendor:
		return 0
	case VendorLogSectionVendorAddress:
		return 1
	case VendorLogSectionVendorPhone:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}

// NullVendorLogSection is a nullable VendorLogSection enum type. It supports SQL and JSON serialization.
type NullVendorLogSection struct {
	Val   VendorLogSection
	Valid bool
}

// NullVendorLogSectionFrom creates a new VendorLogSection that will never be blank.
func NullVendorLogSectionFrom(v VendorLogSection) NullVendorLogSection {
	return NewNullVendorLogSection(v, true)
}

// NullVendorLogSectionFromPtr creates a new NullVendorLogSection that be null if s is nil.
func NullVendorLogSectionFromPtr(v *VendorLogSection) NullVendorLogSection {
	if v == nil {
		return NewNullVendorLogSection("", false)
	}
	return NewNullVendorLogSection(*v, true)
}

// NewNullVendorLogSection creates a new NullVendorLogSection
func NewNullVendorLogSection(v VendorLogSection, valid bool) NullVendorLogSection {
	return NullVendorLogSection{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullVendorLogSection) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullVendorLogSection) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullVendorLogSection) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullVendorLogSection) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = VendorLogSection(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullVendorLogSection value and also sets it to be non-null.
func (e *NullVendorLogSection) SetValid(v VendorLogSection) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullVendorLogSection value, or a nil pointer if this NullVendorLogSection is null.
func (e NullVendorLogSection) Ptr() *VendorLogSection {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullVendorLogSection) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullVendorLogSection) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullVendorLogSection) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type VendorPaymentPolicyCode string

// Enum values for VendorPaymentPolicyCode
const (
	VendorPaymentPolicyCodeAfterBookingConfirmed   VendorPaymentPolicyCode = "after_booking_confirmed"
	VendorPaymentPolicyCodeBeforeServiceDate       VendorPaymentPolicyCode = "before_service_date"
	VendorPaymentPolicyCodeAfterServiceDate        VendorPaymentPolicyCode = "after_service_date"
	VendorPaymentPolicyCodeAfterServiceIsFulfilled VendorPaymentPolicyCode = "after_service_is_fulfilled"
)

func AllVendorPaymentPolicyCode() []VendorPaymentPolicyCode {
	return []VendorPaymentPolicyCode{
		VendorPaymentPolicyCodeAfterBookingConfirmed,
		VendorPaymentPolicyCodeBeforeServiceDate,
		VendorPaymentPolicyCodeAfterServiceDate,
		VendorPaymentPolicyCodeAfterServiceIsFulfilled,
	}
}

func (e VendorPaymentPolicyCode) IsValid() error {
	switch e {
	case VendorPaymentPolicyCodeAfterBookingConfirmed, VendorPaymentPolicyCodeBeforeServiceDate, VendorPaymentPolicyCodeAfterServiceDate, VendorPaymentPolicyCodeAfterServiceIsFulfilled:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VendorPaymentPolicyCode) String() string {
	return string(e)
}

func (e VendorPaymentPolicyCode) Ordinal() int {
	switch e {
	case VendorPaymentPolicyCodeAfterBookingConfirmed:
		return 0
	case VendorPaymentPolicyCodeBeforeServiceDate:
		return 1
	case VendorPaymentPolicyCodeAfterServiceDate:
		return 2
	case VendorPaymentPolicyCodeAfterServiceIsFulfilled:
		return 3

	default:
		panic(errors.New("enum is not valid"))
	}
}

type VendorPhoneType string

// Enum values for VendorPhoneType
const (
	VendorPhoneTypeWork     VendorPhoneType = "work"
	VendorPhoneTypePersonal VendorPhoneType = "personal"
)

func AllVendorPhoneType() []VendorPhoneType {
	return []VendorPhoneType{
		VendorPhoneTypeWork,
		VendorPhoneTypePersonal,
	}
}

func (e VendorPhoneType) IsValid() error {
	switch e {
	case VendorPhoneTypeWork, VendorPhoneTypePersonal:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VendorPhoneType) String() string {
	return string(e)
}

func (e VendorPhoneType) Ordinal() int {
	switch e {
	case VendorPhoneTypeWork:
		return 0
	case VendorPhoneTypePersonal:
		return 1

	default:
		panic(errors.New("enum is not valid"))
	}
}

// NullVendorPhoneType is a nullable VendorPhoneType enum type. It supports SQL and JSON serialization.
type NullVendorPhoneType struct {
	Val   VendorPhoneType
	Valid bool
}

// NullVendorPhoneTypeFrom creates a new VendorPhoneType that will never be blank.
func NullVendorPhoneTypeFrom(v VendorPhoneType) NullVendorPhoneType {
	return NewNullVendorPhoneType(v, true)
}

// NullVendorPhoneTypeFromPtr creates a new NullVendorPhoneType that be null if s is nil.
func NullVendorPhoneTypeFromPtr(v *VendorPhoneType) NullVendorPhoneType {
	if v == nil {
		return NewNullVendorPhoneType("", false)
	}
	return NewNullVendorPhoneType(*v, true)
}

// NewNullVendorPhoneType creates a new NullVendorPhoneType
func NewNullVendorPhoneType(v VendorPhoneType, valid bool) NullVendorPhoneType {
	return NullVendorPhoneType{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullVendorPhoneType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullVendorPhoneType) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullVendorPhoneType) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullVendorPhoneType) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = VendorPhoneType(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullVendorPhoneType value and also sets it to be non-null.
func (e *NullVendorPhoneType) SetValid(v VendorPhoneType) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullVendorPhoneType value, or a nil pointer if this NullVendorPhoneType is null.
func (e NullVendorPhoneType) Ptr() *VendorPhoneType {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullVendorPhoneType) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullVendorPhoneType) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullVendorPhoneType) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type VendorVisibility string

// Enum values for VendorVisibility
const (
	VendorVisibilityAnyone      VendorVisibility = "anyone"
	VendorVisibilityInvitedHost VendorVisibility = "invited_host"
)

func AllVendorVisibility() []VendorVisibility {
	return []VendorVisibility{
		VendorVisibilityAnyone,
		VendorVisibilityInvitedHost,
	}
}

func (e VendorVisibility) IsValid() error {
	switch e {
	case VendorVisibilityAnyone, VendorVisibilityInvitedHost:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VendorVisibility) String() string {
	return string(e)
}

func (e VendorVisibility) Ordinal() int {
	switch e {
	case VendorVisibilityAnyone:
		return 0
	case VendorVisibilityInvitedHost:
		return 1

	default:
		panic(errors.New("enum is not valid"))
	}
}

// NullVendorVisibility is a nullable VendorVisibility enum type. It supports SQL and JSON serialization.
type NullVendorVisibility struct {
	Val   VendorVisibility
	Valid bool
}

// NullVendorVisibilityFrom creates a new VendorVisibility that will never be blank.
func NullVendorVisibilityFrom(v VendorVisibility) NullVendorVisibility {
	return NewNullVendorVisibility(v, true)
}

// NullVendorVisibilityFromPtr creates a new NullVendorVisibility that be null if s is nil.
func NullVendorVisibilityFromPtr(v *VendorVisibility) NullVendorVisibility {
	if v == nil {
		return NewNullVendorVisibility("", false)
	}
	return NewNullVendorVisibility(*v, true)
}

// NewNullVendorVisibility creates a new NullVendorVisibility
func NewNullVendorVisibility(v VendorVisibility, valid bool) NullVendorVisibility {
	return NullVendorVisibility{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullVendorVisibility) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullVendorVisibility) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullVendorVisibility) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullVendorVisibility) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = VendorVisibility(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullVendorVisibility value and also sets it to be non-null.
func (e *NullVendorVisibility) SetValid(v VendorVisibility) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullVendorVisibility value, or a nil pointer if this NullVendorVisibility is null.
func (e NullVendorVisibility) Ptr() *VendorVisibility {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullVendorVisibility) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullVendorVisibility) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullVendorVisibility) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}
