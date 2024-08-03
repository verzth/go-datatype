package datatype

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

type MutatedValue []byte

func (t MutatedValue) String() string {
	if t != nil {
		return string(t)
	}
	return ""
}

// Scan value into Jsonb, implements sql.Scanner interface
func (t *MutatedValue) Scan(value any) error {
	fmt.Println(value)
	if value != nil {
		bytes, ok := value.([]byte)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
		}

		result := string(bytes)
		*t = MutatedValue(result)
	}
	return nil
}

func (t MutatedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *MutatedValue) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = MutatedValue(value)
	return nil
}

func (t MutatedValue) Slice() []MutatedValue {
	var result []MutatedValue
	if err := json.Unmarshal(t, &result); err != nil {
		return nil
	}
	return result
}

func (t MutatedValue) Float64() (float64, error) {
	if f, e := strconv.ParseFloat(t.String(), 64); e == nil {
		return f, nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Float32() (float32, error) {
	if f, e := strconv.ParseFloat(t.String(), 32); e == nil {
		return float32(f), nil
	} else {
		return 0, e
	}
}

// Float64IE : Float64 with ignore error
func (t MutatedValue) Float64IE() float64 {
	v, _ := t.Float64()
	return v
}

// Float32IE : Float32 with ignore error
func (t MutatedValue) Float32IE() float32 {
	v, _ := t.Float32()
	return v
}

func (t MutatedValue) Int() (int, error) {
	if f, e := strconv.ParseInt(t.String(), 10, 32); e == nil {
		return int(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Int8() (int8, error) {
	if f, e := strconv.ParseInt(t.String(), 10, 8); e == nil {
		return int8(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Int16() (int16, error) {
	if f, e := strconv.ParseInt(t.String(), 10, 16); e == nil {
		return int16(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Int32() (int32, error) {
	if f, e := strconv.ParseInt(t.String(), 10, 32); e == nil {
		return int32(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Int64() (int64, error) {
	if f, e := strconv.ParseInt(t.String(), 10, 64); e == nil {
		return f, nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) IntIE() int {
	v, _ := t.Int()
	return v
}

func (t MutatedValue) Int8IE() int8 {
	v, _ := t.Int8()
	return v
}

func (t MutatedValue) Int16IE() int16 {
	v, _ := t.Int16()
	return v
}

func (t MutatedValue) Int32IE() int32 {
	v, _ := t.Int32()
	return v
}

func (t MutatedValue) Int64IE() int64 {
	v, _ := t.Int64()
	return v
}

func (t MutatedValue) Uint() (uint, error) {
	if f, e := strconv.ParseUint(t.String(), 10, 32); e == nil {
		return uint(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Uint8() (uint8, error) {
	if f, e := strconv.ParseUint(t.String(), 10, 8); e == nil {
		return uint8(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Uint16() (uint16, error) {
	if f, e := strconv.ParseUint(t.String(), 10, 16); e == nil {
		return uint16(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Uint32() (uint32, error) {
	if f, e := strconv.ParseUint(t.String(), 10, 32); e == nil {
		return uint32(f), nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) Uint64() (uint64, error) {
	if f, e := strconv.ParseUint(t.String(), 10, 64); e == nil {
		return f, nil
	} else {
		return 0, e
	}
}

func (t MutatedValue) UintIE() uint {
	v, _ := t.Uint()
	return v
}

func (t MutatedValue) Uint8IE() uint8 {
	v, _ := t.Uint8()
	return v
}

func (t MutatedValue) Uint16IE() uint16 {
	v, _ := t.Uint16()
	return v
}

func (t MutatedValue) Uint32IE() uint32 {
	v, _ := t.Uint32()
	return v
}

func (t MutatedValue) Uint64IE() uint64 {
	v, _ := t.Uint64()
	return v
}

func (t MutatedValue) Bytes() []byte {
	return []byte(t.String())
}

func (t MutatedValue) HexToFloat64() (float64, error) {
	if len(t) >= 2 && (t).String()[:2] == "0x" {
		if f, err := strconv.ParseFloat(string(t)[2:], 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	} else {
		if f, err := strconv.ParseFloat(t.String(), 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	}
}

func (t MutatedValue) HexToInt64() (int64, error) {
	if len(t) >= 2 && (t).String()[:2] == "0x" {
		if f, err := strconv.ParseInt(string(t)[2:], 16, 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	} else {
		if f, err := strconv.ParseInt(t.String(), 16, 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	}
}

func (t MutatedValue) HexToInt64IE() int64 {
	v, _ := t.HexToInt64()
	return v
}

func (t MutatedValue) HexToUint64() (uint64, error) {
	if len(t) >= 2 && (t).String()[:2] == "0x" {
		if f, err := strconv.ParseUint(string(t)[2:], 16, 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	} else {
		if f, err := strconv.ParseUint(t.String(), 16, 64); err == nil {
			return f, nil
		} else {
			return 0, err
		}
	}
}

func (t MutatedValue) HexToUint64IE() uint64 {
	v, _ := t.HexToUint64()
	return v
}

func (t MutatedValue) BigInt() *big.Int {
	b := new(big.Int)
	if len(t) >= 2 && (t).String()[:2] == "0x" {
		b.SetString(string(t)[2:], 10)
	} else {
		b.SetString(t.String(), 10)
	}
	return b
}

func (t MutatedValue) HexToBigInt() *big.Int {
	b := new(big.Int)
	if len(t) >= 2 && (t).String()[:2] == "0x" {
		b.SetString(string(t)[2:], 16)
	} else {
		b.SetString(t.String(), 16)
	}
	return b
}

func (t MutatedValue) Decimal() (decimal.Decimal, error) {
	d, err := decimal.NewFromString(t.String())
	if err != nil {
		return decimal.Zero, err
	}
	return d, nil
}

func (t MutatedValue) DecimalIE() decimal.Decimal {
	v, _ := t.Decimal()
	return v
}
