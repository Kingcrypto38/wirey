// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := &FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return fibfia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfia *FIBeneficiaryFIAdvice) Parse(record string) error {
	if utf8.RuneCountInString(record) < 9 {
		return NewTagMinLengthErr(9, len(record))
	}

	fibfia.tag = record[:6]
	fibfia.Advice.AdviceCode = fibfia.parseStringField(record[6:9])
	length := 9

	value, read, err := fibfia.parseVariableStringField(record[length:], 26)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fibfia.Advice.LineOne = value
	length += read

	value, read, err = fibfia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fibfia.Advice.LineTwo = value
	length += read

	value, read, err = fibfia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fibfia.Advice.LineThree = value
	length += read

	value, read, err = fibfia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fibfia.Advice.LineFour = value
	length += read

	value, read, err = fibfia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fibfia.Advice.LineFive = value
	length += read

	value, read, err = fibfia.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fibfia.Advice.LineSix = value
	length += read

	if err := fibfia.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
}

func (fibfia *FIBeneficiaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfia.tag = TagFIBeneficiaryFIAdvice
	return nil
}

// String returns a fixed-width FIBeneficiaryFIAdvice record
func (fibfia *FIBeneficiaryFIAdvice) String() string {
	return fibfia.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a FIBeneficiaryFIAdvice record formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fibfia.tag)
	buf.WriteString(fibfia.AdviceCodeField())
	buf.WriteString(fibfia.FormatLineOne(options) + Delimiter)
	buf.WriteString(fibfia.FormatLineTwo(options) + Delimiter)
	buf.WriteString(fibfia.FormatLineThree(options) + Delimiter)
	buf.WriteString(fibfia.FormatLineFour(options) + Delimiter)
	buf.WriteString(fibfia.FormatLineFive(options) + Delimiter)
	buf.WriteString(fibfia.FormatLineSix(options) + Delimiter)

	if options.VariableLengthFields {
		return fibfia.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfia *FIBeneficiaryFIAdvice) Validate() error {
	if fibfia.tag != TagFIBeneficiaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fibfia.tag)
	}
	if err := fibfia.isAdviceCode(fibfia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fibfia.Advice.AdviceCode)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fibfia.Advice.LineOne)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfia.Advice.LineTwo)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fibfia.Advice.LineThree)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fibfia.Advice.LineFour)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fibfia.Advice.LineFive)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fibfia.Advice.LineSix)
	}
	return nil
}

// AdviceCodeField gets a string of the AdviceCode field
func (fibfia *FIBeneficiaryFIAdvice) AdviceCodeField() string {
	return fibfia.alphaField(fibfia.Advice.AdviceCode, 3)
}

// LineOneField gets a string of the LineOne field
func (fibfia *FIBeneficiaryFIAdvice) LineOneField() string {
	return fibfia.alphaField(fibfia.Advice.LineOne, 26)
}

// LineTwoField gets a string of the LineTwo field
func (fibfia *FIBeneficiaryFIAdvice) LineTwoField() string {
	return fibfia.alphaField(fibfia.Advice.LineTwo, 33)
}

// LineThreeField gets a string of the LineThree field
func (fibfia *FIBeneficiaryFIAdvice) LineThreeField() string {
	return fibfia.alphaField(fibfia.Advice.LineThree, 33)
}

// LineFourField gets a string of the LineFour field
func (fibfia *FIBeneficiaryFIAdvice) LineFourField() string {
	return fibfia.alphaField(fibfia.Advice.LineFour, 33)
}

// LineFiveField gets a string of the LineFive field
func (fibfia *FIBeneficiaryFIAdvice) LineFiveField() string {
	return fibfia.alphaField(fibfia.Advice.LineFive, 33)
}

// LineSixField gets a string of the LineSix field
func (fibfia *FIBeneficiaryFIAdvice) LineSixField() string {
	return fibfia.alphaField(fibfia.Advice.LineSix, 33)
}

// FormatLineOne returns Advice.LineOne formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineOne(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineOne, 26, options)
}

// FormatLineTwo returns Advice.LineTwo formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineTwo(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineTwo, 33, options)
}

// FormatLineThree returns Advice.LineThree formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineThree(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineThree, 33, options)
}

// FormatLineFour returns Advice.LineFour formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineFour(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineFour, 33, options)
}

// FormatLineFive returns Advice.LineFive formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineFive(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineFive, 33, options)
}

// FormatLineSix returns Advice.LineSix formatted according to the FormatOptions
func (fibfia *FIBeneficiaryFIAdvice) FormatLineSix(options FormatOptions) string {
	return fibfia.formatAlphaField(fibfia.Advice.LineSix, 33, options)
}
