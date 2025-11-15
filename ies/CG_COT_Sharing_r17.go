package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_COT_Sharing_r17_Choice_nothing uint64 = iota
	CG_COT_Sharing_r17_Choice_noCOT_Sharing_r17
	CG_COT_Sharing_r17_Choice_cot_Sharing_r17
)

type CG_COT_Sharing_r17 struct {
	Choice            uint64
	noCOT_Sharing_r17 uper.NULL `madatory`
	cot_Sharing_r17   *CG_COT_Sharing_r17_cot_Sharing_r17
}

func (ie *CG_COT_Sharing_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r17_Choice_noCOT_Sharing_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode noCOT_Sharing_r17", err)
		}
	case CG_COT_Sharing_r17_Choice_cot_Sharing_r17:
		if err = ie.cot_Sharing_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode cot_Sharing_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_COT_Sharing_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r17_Choice_noCOT_Sharing_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode noCOT_Sharing_r17", err)
		}
	case CG_COT_Sharing_r17_Choice_cot_Sharing_r17:
		ie.cot_Sharing_r17 = new(CG_COT_Sharing_r17_cot_Sharing_r17)
		if err = ie.cot_Sharing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cot_Sharing_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
