package segment

import (
	"bytes"
	"fmt"

	"github.com/mitch000001/go-hbci/element"
)

func (c *CommonUserParameterDataSegment) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	header := &element.SegmentHeader{}
	err = header.UnmarshalHBCI(elements[0])
	if err != nil {
		return err
	}
	var segment commonUserParameterDataSegment
	switch header.Version.Val() {
	case 2:
		segment = &CommonUserParameterDataV2{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	case 3:
		segment = &CommonUserParameterDataV3{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	case 4:
		segment = &CommonUserParameterDataV4{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown segment version: %d", header.Version.Val())
	}
	c.commonUserParameterDataSegment = segment
	return nil
}

func (c *CommonUserParameterDataV2) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], c)
	if err != nil {
		return err
	}
	c.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		c.UserID = &element.IdentificationDataElement{}
		err = c.UserID.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		c.UPDVersion = &element.NumberDataElement{}
		err = c.UPDVersion.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		c.UPDUsage = &element.NumberDataElement{}
		if len(elements)+1 > 3 {
			err = c.UPDUsage.UnmarshalHBCI(bytes.Join(elements[3:], []byte("+")))
		} else {
			err = c.UPDUsage.UnmarshalHBCI(elements[3])
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CommonUserParameterDataV3) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], c)
	if err != nil {
		return err
	}
	c.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		c.UserID = &element.IdentificationDataElement{}
		err = c.UserID.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		c.UPDVersion = &element.NumberDataElement{}
		err = c.UPDVersion.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		c.UPDUsage = &element.NumberDataElement{}
		err = c.UPDUsage.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		c.UserName = &element.AlphaNumericDataElement{}
		err = c.UserName.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		c.CommonExtensions = &element.AlphaNumericDataElement{}
		if len(elements)+1 > 5 {
			err = c.CommonExtensions.UnmarshalHBCI(bytes.Join(elements[5:], []byte("+")))
		} else {
			err = c.CommonExtensions.UnmarshalHBCI(elements[5])
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CommonUserParameterDataV4) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], c)
	if err != nil {
		return err
	}
	c.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		c.UserID = &element.IdentificationDataElement{}
		err = c.UserID.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		c.UPDVersion = &element.NumberDataElement{}
		err = c.UPDVersion.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		c.UPDUsage = &element.NumberDataElement{}
		err = c.UPDUsage.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		c.UserName = &element.AlphaNumericDataElement{}
		err = c.UserName.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		c.CommonExtensions = &element.AlphaNumericDataElement{}
		if len(elements)+1 > 5 {
			err = c.CommonExtensions.UnmarshalHBCI(bytes.Join(elements[5:], []byte("+")))
		} else {
			err = c.CommonExtensions.UnmarshalHBCI(elements[5])
		}
		if err != nil {
			return err
		}
	}
	return nil
}
