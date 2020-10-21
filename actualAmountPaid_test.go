package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"
)

// ActualAmountPaid creates a ActualAmountPaid
func mockActualAmountPaid() *ActualAmountPaid {
	aap := NewActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = "USD"
	aap.RemittanceAmount.Amount = "1234.56"
	return aap
}

// TestMockActualAmountPaid validates mockActualAmountPaid
func TestMockActualAmountPaid(t *testing.T) {
	aap := mockActualAmountPaid()
	if err := aap.Validate(); err != nil {
		t.Error("mockActualAmountPaid does not validate and will break other tests")
	}
}

// TestActualAmountPaidAmountRequired validates ActualAmountPaid Amount is required
func TestActualAmountPaidAmountRequired(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.Amount = ""
	if err := aap.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestActualAmountPaidCurrencyCodeRequired validates ActualAmountPaid CurrencyCode is required
func TestCurrencyCodeRequired(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = ""
	if err := aap.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestActualAmountPaidAmountValid validates Amount
func TestActualAmountPaidAmountValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.Amount = "X,"
	if err := aap.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestActualAmountPaidCurrencyCodeValid validates Amount
func TestActualAmountPaidCurrencyCodeValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = "XZP"
	if err := aap.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseActualAmountPaidWrongLength parses a wrong ActualAmountPaid record length
func TestParseActualAmountPaidWrongLength(t *testing.T) {
	var line = "{8450}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)
	err := r.parseActualAmountPaid()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(28, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseActualAmountPaidReaderParseError parses a wrong ActualAmountPaid reader parse error
func TestParseActualAmountPaidReaderParseError(t *testing.T) {
	var line = "{8450}USD1234.56Z           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseActualAmountPaid()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	fwm.LocalInstrument = mockLocalInstrument()
	fwm.LocalInstrument.LocalInstrumentCode = RemittanceInformationStructured
	fwm.ActualAmountPaid = &ActualAmountPaid{}
	require.NoError(t, fwm.ActualAmountPaid.Parse(line))
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	fwm.RemittanceOriginator = mockRemittanceOriginator()
	fwm.RemittanceBeneficiary = mockRemittanceBeneficiary()
	fwm.DateRemittanceDocument = mockDateRemittanceDocument()
	fwm.PrimaryRemittanceDocument = mockPrimaryRemittanceDocument()
	fwm.GrossAmountRemittanceDocument = mockGrossAmountRemittanceDocument()
	fwm.RemittanceFreeText = mockRemittanceFreeText()
	fwm.Adjustment = mockAdjustment()
	r.currentFEDWireMessage = fwm

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestActualAmountPaidTagError validates ActualAmountPaid tag
func TestActualAmountPaidTagError(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.tag = "{9999}"
	if err := aap.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
