package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_COT_Sharing_r16_Choice_nothing uint64 = iota
	CG_COT_Sharing_r16_Choice_noCOT_Sharing_r16
	CG_COT_Sharing_r16_Choice_cot_Sharing_r16
)

type CG_COT_Sharing_r16 struct {
	Choice            uint64
	noCOT_Sharing_r16 uper.NULL `madatory`
	cot_Sharing_r16   *CG_COT_Sharing_r16_cot_Sharing_r16
}

func (ie *CG_COT_Sharing_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r16_Choice_noCOT_Sharing_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode noCOT_Sharing_r16", err)
		}
	case CG_COT_Sharing_r16_Choice_cot_Sharing_r16:
		if err = ie.cot_Sharing_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode cot_Sharing_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_COT_Sharing_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r16_Choice_noCOT_Sharing_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode noCOT_Sharing_r16", err)
		}
	case CG_COT_Sharing_r16_Choice_cot_Sharing_r16:
		ie.cot_Sharing_r16 = new(CG_COT_Sharing_r16_cot_Sharing_r16)
		if err = ie.cot_Sharing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cot_Sharing_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
