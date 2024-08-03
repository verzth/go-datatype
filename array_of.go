package datatype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
)

type ArrayOf[T any] []T
type ArrayString ArrayOf[string]
type ArrayInt ArrayOf[int]
type ArrayInt8 ArrayOf[int8]
type ArrayInt16 ArrayOf[int16]
type ArrayInt32 ArrayOf[int32]
type ArrayInt64 ArrayOf[int64]
type ArrayUint ArrayOf[uint]
type ArrayUint8 ArrayOf[uint8]
type ArrayUint16 ArrayOf[uint16]
type ArrayUint32 ArrayOf[uint32]
type ArrayUint64 ArrayOf[uint64]
type ArrayFloat32 ArrayOf[float32]
type ArrayFloat64 ArrayOf[float64]
type ArrayBool ArrayOf[bool]
type ArrayBigInt ArrayOf[*big.Int]
type ArrayDecimal ArrayOf[decimal.Decimal]
type ArrayMutatedValue ArrayOf[MutatedValue]

func NewArrayOf[T comparable](values ...T) ArrayOf[T] {
	return values
}

func (t ArrayOf[T]) ToArray() []T {
	return []T(t)
}

func (t ArrayOf[T]) String() string {
	jsonStr, e := json.Marshal(t)
	if e != nil {
		return "[]"
	} else {
		return string(jsonStr)
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayOf[T]) Scan(value any) error {
	if value != nil {
		bytes, ok := value.([]byte)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
		}

		result := ArrayOf[T]{}
		err := json.Unmarshal(bytes, &result)
		if err != nil {
			return err
		}
		*t = result
		return err
	} else {
		*t = ArrayOf[T]{}
		return nil
	}
}

// Value return json value, implement driver.Valuer interface
func (t ArrayOf[T]) Value() (driver.Value, error) {
	return t.String(), nil
}

// ToArray convert to array
func (t ArrayString) ToArray() []string {
	return (ArrayOf[string])(t).ToArray()
}

// String convert to string
func (t ArrayString) String() string {
	return (ArrayOf[string])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayString) Scan(value any) error {
	return (*ArrayOf[string])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayString) Value() (driver.Value, error) {
	return (ArrayOf[string])(t).Value()
}

// ToArray convert to array
func (t ArrayInt) ToArray() []int {
	return (ArrayOf[int])(t).ToArray()
}

// String convert to string
func (t ArrayInt) String() string {
	return (ArrayOf[int])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayInt) Scan(value any) error {
	return (*ArrayOf[int])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayInt) Value() (driver.Value, error) {
	return (ArrayOf[int])(t).Value()
}

// ToArray convert to array
func (t ArrayInt8) ToArray() []int8 {
	return (ArrayOf[int8])(t).ToArray()
}

// String convert to string
func (t ArrayInt8) String() string {
	return (ArrayOf[int8])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayInt8) Scan(value any) error {
	return (*ArrayOf[int8])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayInt8) Value() (driver.Value, error) {
	return (ArrayOf[int8])(t).Value()
}

// ToArray convert to array
func (t ArrayInt16) ToArray() []int16 {
	return (ArrayOf[int16])(t).ToArray()
}

// String convert to string
func (t ArrayInt16) String() string {
	return (ArrayOf[int16])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayInt16) Scan(value any) error {
	return (*ArrayOf[int16])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayInt16) Value() (driver.Value, error) {
	return (ArrayOf[int16])(t).Value()
}

// ToArray convert to array
func (t ArrayInt32) ToArray() []int32 {
	return (ArrayOf[int32])(t).ToArray()
}

// String convert to string
func (t ArrayInt32) String() string {
	return (ArrayOf[int32])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayInt32) Scan(value any) error {
	return (*ArrayOf[int32])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayInt32) Value() (driver.Value, error) {
	return (ArrayOf[int32])(t).Value()
}

// ToArray convert to array
func (t ArrayInt64) ToArray() []int64 {
	return (ArrayOf[int64])(t).ToArray()
}

// String convert to string
func (t ArrayInt64) String() string {
	return (ArrayOf[int64])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayInt64) Scan(value any) error {
	return (*ArrayOf[int64])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayInt64) Value() (driver.Value, error) {
	return (ArrayOf[int64])(t).Value()
}

// ToArray convert to array
func (t ArrayUint) ToArray() []uint {
	return (ArrayOf[uint])(t).ToArray()
}

// String convert to string
func (t ArrayUint) String() string {
	return (ArrayOf[uint])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayUint) Scan(value any) error {
	return (*ArrayOf[uint])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayUint) Value() (driver.Value, error) {
	return (ArrayOf[uint])(t).Value()
}

// ToArray convert to array
func (t ArrayUint8) ToArray() []uint8 {
	return (ArrayOf[uint8])(t).ToArray()
}

// String convert to string
func (t ArrayUint8) String() string {
	return (ArrayOf[uint8])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayUint8) Scan(value any) error {
	return (*ArrayOf[uint8])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayUint8) Value() (driver.Value, error) {
	return (ArrayOf[uint8])(t).Value()
}

// ToArray convert to array
func (t ArrayUint16) ToArray() []uint16 {
	return (ArrayOf[uint16])(t).ToArray()
}

// String convert to string
func (t ArrayUint16) String() string {
	return (ArrayOf[uint16])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayUint16) Scan(value any) error {
	return (*ArrayOf[uint16])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayUint16) Value() (driver.Value, error) {
	return (ArrayOf[uint16])(t).Value()
}

// ToArray convert to array
func (t ArrayUint32) ToArray() []uint32 {
	return (ArrayOf[uint32])(t).ToArray()
}

// String convert to string
func (t ArrayUint32) String() string {
	return (ArrayOf[uint32])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayUint32) Scan(value any) error {
	return (*ArrayOf[uint32])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayUint32) Value() (driver.Value, error) {
	return (ArrayOf[uint32])(t).Value()
}

// ToArray convert to array
func (t ArrayUint64) ToArray() []uint64 {
	return (ArrayOf[uint64])(t).ToArray()
}

// String convert to string
func (t ArrayUint64) String() string {
	return (ArrayOf[uint64])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayUint64) Scan(value any) error {
	return (*ArrayOf[uint64])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayUint64) Value() (driver.Value, error) {
	return (ArrayOf[uint64])(t).Value()
}


// ToArray convert to array
func (t ArrayFloat32) ToArray() []float32 {
	return (ArrayOf[float32])(t).ToArray()
}

// String convert to string
func (t ArrayFloat32) String() string {
	return (ArrayOf[float32])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayFloat32) Scan(value any) error {
	return (*ArrayOf[float32])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayFloat32) Value() (driver.Value, error) {
	return (ArrayOf[float32])(t).Value()
}

// ToArray convert to array
func (t ArrayFloat64) ToArray() []float64 {
	return (ArrayOf[float64])(t).ToArray()
}

// String convert to string
func (t ArrayFloat64) String() string {
	return (ArrayOf[float64])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayFloat64) Scan(value any) error {
	return (*ArrayOf[float64])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayFloat64) Value() (driver.Value, error) {
	return (ArrayOf[float64])(t).Value()
}

// ToArray convert to array
func (t ArrayBool) ToArray() []bool {
	return (ArrayOf[bool])(t).ToArray()
}

// String convert to string
func (t ArrayBool) String() string {
	return (ArrayOf[bool])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayBool) Scan(value any) error {
	return (*ArrayOf[bool])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayBool) Value() (driver.Value, error) {
	return (ArrayOf[bool])(t).Value()
}

// ToArray convert to array
func (t ArrayBigInt) ToArray() []*big.Int {
	return (ArrayOf[*big.Int])(t).ToArray()
}

// String convert to string
func (t ArrayBigInt) String() string {
	return (ArrayOf[*big.Int])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayBigInt) Scan(value any) error {
	return (*ArrayOf[*big.Int])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayBigInt) Value() (driver.Value, error) {
	return (ArrayOf[*big.Int])(t).Value()
}

// ToArray convert to array
func (t ArrayDecimal) ToArray() []decimal.Decimal {
	return (ArrayOf[decimal.Decimal])(t).ToArray()
}

// String convert to string
func (t ArrayDecimal) String() string {
	return (ArrayOf[decimal.Decimal])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayDecimal) Scan(value any) error {
	return (*ArrayOf[decimal.Decimal])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayDecimal) Value() (driver.Value, error) {
	return (ArrayOf[decimal.Decimal])(t).Value()
}

// ToArray convert to array
func (t ArrayMutatedValue) ToArray() []MutatedValue {
	return (ArrayOf[MutatedValue])(t).ToArray()
}

// String convert to string
func (t ArrayMutatedValue) String() string {
	return (ArrayOf[MutatedValue])(t).String()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *ArrayMutatedValue) Scan(value any) error {
	return (*ArrayOf[MutatedValue])(t).Scan(value)
}

// Value return json value, implement driver.Valuer interface
func (t ArrayMutatedValue) Value() (driver.Value, error) {
	return (ArrayOf[MutatedValue])(t).Value()
}