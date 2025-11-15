package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DefaultDC_Location_r17_Choice_nothing uint64 = iota
	DefaultDC_Location_r17_Choice_ul
	DefaultDC_Location_r17_Choice_dl
	DefaultDC_Location_r17_Choice_ulAndDL
)

type DefaultDC_Location_r17 struct {
	Choice  uint64
	ul      *FrequencyComponent_r17
	dl      *FrequencyComponent_r17
	ulAndDL *FrequencyComponent_r17
}

func (ie *DefaultDC_Location_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DefaultDC_Location_r17_Choice_ul:
		if err = ie.ul.Encode(w); err != nil {
			err = utils.WrapError("Encode ul", err)
		}
	case DefaultDC_Location_r17_Choice_dl:
		if err = ie.dl.Encode(w); err != nil {
			err = utils.WrapError("Encode dl", err)
		}
	case DefaultDC_Location_r17_Choice_ulAndDL:
		if err = ie.ulAndDL.Encode(w); err != nil {
			err = utils.WrapError("Encode ulAndDL", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DefaultDC_Location_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DefaultDC_Location_r17_Choice_ul:
		ie.ul = new(FrequencyComponent_r17)
		if err = ie.ul.Decode(r); err != nil {
			return utils.WrapError("Decode ul", err)
		}
	case DefaultDC_Location_r17_Choice_dl:
		ie.dl = new(FrequencyComponent_r17)
		if err = ie.dl.Decode(r); err != nil {
			return utils.WrapError("Decode dl", err)
		}
	case DefaultDC_Location_r17_Choice_ulAndDL:
		ie.ulAndDL = new(FrequencyComponent_r17)
		if err = ie.ulAndDL.Decode(r); err != nil {
			return utils.WrapError("Decode ulAndDL", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
