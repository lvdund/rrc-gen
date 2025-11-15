package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_Resource_format_Choice_nothing uint64 = iota
	PUCCH_Resource_format_Choice_format0
	PUCCH_Resource_format_Choice_format1
	PUCCH_Resource_format_Choice_format2
	PUCCH_Resource_format_Choice_format3
	PUCCH_Resource_format_Choice_format4
)

type PUCCH_Resource_format struct {
	Choice  uint64
	format0 *PUCCH_format0
	format1 *PUCCH_format1
	format2 *PUCCH_format2
	format3 *PUCCH_format3
	format4 *PUCCH_format4
}

func (ie *PUCCH_Resource_format) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Resource_format_Choice_format0:
		if err = ie.format0.Encode(w); err != nil {
			err = utils.WrapError("Encode format0", err)
		}
	case PUCCH_Resource_format_Choice_format1:
		if err = ie.format1.Encode(w); err != nil {
			err = utils.WrapError("Encode format1", err)
		}
	case PUCCH_Resource_format_Choice_format2:
		if err = ie.format2.Encode(w); err != nil {
			err = utils.WrapError("Encode format2", err)
		}
	case PUCCH_Resource_format_Choice_format3:
		if err = ie.format3.Encode(w); err != nil {
			err = utils.WrapError("Encode format3", err)
		}
	case PUCCH_Resource_format_Choice_format4:
		if err = ie.format4.Encode(w); err != nil {
			err = utils.WrapError("Encode format4", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_Resource_format) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Resource_format_Choice_format0:
		ie.format0 = new(PUCCH_format0)
		if err = ie.format0.Decode(r); err != nil {
			return utils.WrapError("Decode format0", err)
		}
	case PUCCH_Resource_format_Choice_format1:
		ie.format1 = new(PUCCH_format1)
		if err = ie.format1.Decode(r); err != nil {
			return utils.WrapError("Decode format1", err)
		}
	case PUCCH_Resource_format_Choice_format2:
		ie.format2 = new(PUCCH_format2)
		if err = ie.format2.Decode(r); err != nil {
			return utils.WrapError("Decode format2", err)
		}
	case PUCCH_Resource_format_Choice_format3:
		ie.format3 = new(PUCCH_format3)
		if err = ie.format3.Decode(r); err != nil {
			return utils.WrapError("Decode format3", err)
		}
	case PUCCH_Resource_format_Choice_format4:
		ie.format4 = new(PUCCH_format4)
		if err = ie.format4.Decode(r); err != nil {
			return utils.WrapError("Decode format4", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
