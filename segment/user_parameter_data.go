package segment

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mitch000001/go-hbci/element"
)

type CommonUserParameterDataSegment struct {
	Segment
	UserId     *element.IdentificationDataElement
	UPDVersion *element.NumberDataElement
	// Status |￼Beschreibung
	// -----------------------------------------------------------------
	// 0	  | Die nicht aufgeführten Geschäftsvorfälle sind gesperrt
	//		  | (die aufgeführten Geschäftsvorfälle sind zugelassen).
	// 1 ￼ ￼  | Bei den nicht aufgeführten Geschäftsvorfällen ist anhand
	//        | der UPD keine Aussage darüber möglich, ob diese erlaubt
	//        | oder gesperrt sind. Diese Prüfung kann nur online vom
	//        | Kreditinstitutssystem vorgenommen werden.
	UPDUsage *element.NumberDataElement
}

func (c *CommonUserParameterDataSegment) init() {
	*c.UserId = *new(element.IdentificationDataElement)
	*c.UPDVersion = *new(element.NumberDataElement)
	*c.UPDUsage = *new(element.NumberDataElement)
}
func (c *CommonUserParameterDataSegment) version() int         { return 2 }
func (c *CommonUserParameterDataSegment) id() string           { return "HIUPA" }
func (c *CommonUserParameterDataSegment) referencedId() string { return "HKVVB" }
func (c *CommonUserParameterDataSegment) sender() string       { return senderBank }

func (c *CommonUserParameterDataSegment) elements() []element.DataElement {
	return []element.DataElement{
		c.UserId,
		c.UPDVersion,
		c.UPDUsage,
	}
}

type AccountInformationSegment struct {
	Segment
	AccountConnection           *element.AccountConnectionDataElement
	UserID                      *element.IdentificationDataElement
	AccountCurrency             *element.CurrencyDataElement
	Name1                       *element.AlphaNumericDataElement
	Name2                       *element.AlphaNumericDataElement
	AccountProductID            *element.AlphaNumericDataElement
	AccountLimit                *element.AccountLimitDataElement
	AllowedBusinessTransactions *element.AllowedBusinessTransactionsDataElement
}

func (a *AccountInformationSegment) init() {
	*a.AccountConnection = *new(element.AccountConnectionDataElement)
	*a.UserID = *new(element.IdentificationDataElement)
	*a.AccountCurrency = *new(element.CurrencyDataElement)
	*a.Name1 = *new(element.AlphaNumericDataElement)
	*a.Name2 = *new(element.AlphaNumericDataElement)
	*a.AccountProductID = *new(element.AlphaNumericDataElement)
	*a.AccountLimit = *new(element.AccountLimitDataElement)
	*a.AllowedBusinessTransactions = *new(element.AllowedBusinessTransactionsDataElement)
}
func (a *AccountInformationSegment) version() int         { return 4 }
func (a *AccountInformationSegment) id() string           { return "HIUPD" }
func (a *AccountInformationSegment) referencedId() string { return "HKVVB" }
func (a *AccountInformationSegment) sender() string       { return senderBank }

func (a *AccountInformationSegment) UnmarshalHBCI(value []byte) error {
	elements := bytes.Split(value, []byte("+"))
	header := elements[0]
	headerElems := bytes.Split(header, []byte(":"))
	num, err := strconv.Atoi(string(headerElems[1]))
	if err != nil {
		return fmt.Errorf("Malformed segment header")
	}
	if len(headerElems) == 4 {
		ref, err := strconv.Atoi(string(headerElems[3]))
		if err != nil {
			return fmt.Errorf("Malformed segment header reference: %v", err)
		}
		a.Segment = NewReferencingBasicSegment(num, ref, a)
	} else {
		a.Segment = NewBasicSegment(num, a)
	}
	elements = elements[1:]
	a.AccountConnection = &element.AccountConnectionDataElement{}
	err = a.AccountConnection.UnmarshalHBCI(elements[0])
	if err != nil {
		return fmt.Errorf("%T: Unmarshaling AccountConnection failed: %T:%v", a, err, err)
	}
	a.UserID = element.NewIdentification(string(elements[1]))
	a.AccountCurrency = element.NewCurrency(string(elements[2]))
	a.Name1 = element.NewAlphaNumeric(string(elements[3]), 27)
	a.Name2 = element.NewAlphaNumeric(string(elements[4]), 27)
	a.AccountProductID = element.NewAlphaNumeric(string(elements[5]), 30)
	accountLimit := elements[6]
	if len(accountLimit) > 0 {
		a.AccountLimit = &element.AccountLimitDataElement{}
		err = a.AccountLimit.UnmarshalHBCI(accountLimit)
		if err != nil {
			return fmt.Errorf("%T: Unmarshaling AccountLimit failed: %T:%v", a, err, err)
		}
	}
	allowedBusinessTransactions := elements[7]
	if len(allowedBusinessTransactions) > 0 {
		a.AllowedBusinessTransactions = &element.AllowedBusinessTransactionsDataElement{}
		err = a.AllowedBusinessTransactions.UnmarshalHBCI(allowedBusinessTransactions)
		if err != nil {
			return fmt.Errorf("%T: Unmarshaling AllowedBusinessTransactions failed: %T:%v", a, err, err)
		}
	}
	return nil
}

func (a *AccountInformationSegment) elements() []element.DataElement {
	return []element.DataElement{
		a.AccountConnection,
		a.UserID,
		a.AccountCurrency,
		a.Name1,
		a.Name2,
		a.AccountProductID,
		a.AccountLimit,
		a.AllowedBusinessTransactions,
	}
}
