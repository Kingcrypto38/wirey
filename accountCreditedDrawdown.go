// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// AccountCreditedDrawdown is the account which is credited in a drawdown
type AccountCreditedDrawdown struct {
	// tag
	tag string
	// DrawdownCreditAccountNumber  9 character ABA
	DrawdownCreditAccountNumber string `json:"drawdownCreditAccountNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAccountCreditedDrawdown returns a new AccountCreditedDrawdown
func NewAccountCreditedDrawdown() *AccountCreditedDrawdown {
	creditDD := &AccountCreditedDrawdown{
		tag: TagAccountCreditedDrawdown,
	}
	return creditDD
}

// Parse takes the input string and parses the AccountCreditedDrawdown values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (creditDD *AccountCreditedDrawdown) Parse(record string) error {
	if utf8.RuneCountInString(record) < 7 {
		return NewTagMinLengthErr(7, len(record))
	}

	creditDD.tag = record[:6]
	length := 6

	value, read, err := creditDD.parseFixedStringField(record[length:], 9)
	if err != nil {
		return fieldError("DrawdownCreditAccountNumber", err)
	}
	creditDD.DrawdownCreditAccountNumber = value
	length += read

	if err := creditDD.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
}

func (creditDD *AccountCreditedDrawdown) UnmarshalJSON(data []byte) error {
	type Alias AccountCreditedDrawdown
	aux := struct {
		*Alias
	}{
		(*Alias)(creditDD),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	creditDD.tag = TagAccountCreditedDrawdown
	return nil
}

// String returns a fixed-width AccountCreditedDrawdown record
func (creditDD *AccountCreditedDrawdown) String() string {
	return creditDD.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns an AccountCreditedDrawdown record formatted according to the FormatOptions
func (creditDD *AccountCreditedDrawdown) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(15)
	buf.WriteString(creditDD.tag)
	buf.WriteString(creditDD.DrawdownCreditAccountNumberField())
	return buf.String()
}

// Validate performs WIRE format rule checks on AccountCreditedDrawdown and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (creditDD *AccountCreditedDrawdown) Validate() error {
	if err := creditDD.fieldInclusion(); err != nil {
		return err
	}
	if creditDD.tag != TagAccountCreditedDrawdown {
		return fieldError("tag", ErrValidTagForType, creditDD.tag)
	}
	if err := creditDD.isNumeric(creditDD.DrawdownCreditAccountNumber); err != nil {
		return fieldError("DrawdownCreditAccountNumber", err, creditDD.DrawdownCreditAccountNumber)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (creditDD *AccountCreditedDrawdown) fieldInclusion() error {
	if creditDD.DrawdownCreditAccountNumber == "" {
		return fieldError("DrawdownCreditAccountNumber", ErrFieldRequired)
	}
	return nil
}

// DrawdownCreditAccountNumberField gets a string of the DrawdownCreditAccountNumber field
func (creditDD *AccountCreditedDrawdown) DrawdownCreditAccountNumberField() string {
	return creditDD.alphaField(creditDD.DrawdownCreditAccountNumber, 9)
}

// FormatCreditAccountNumber returns DrawdownCreditAccountNumber formatted according to the FormatOptions
func (creditDD *AccountCreditedDrawdown) FormatCreditAccountNumber(options FormatOptions) string {
	return creditDD.formatAlphaField(creditDD.DrawdownCreditAccountNumber, 9, options)
}
