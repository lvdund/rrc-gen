package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_nothing uint64 = iota
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs15
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs30
)

type PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16 struct {
	Choice uint64
	scs15  int64 `lb:0,ub:9,madatory`
	scs30  int64 `lb:0,ub:4,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs15:
		if err = w.WriteInteger(int64(ie.scs15), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode scs15", err)
		}
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs30:
		if err = w.WriteInteger(int64(ie.scs30), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode scs30", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs15:
		var tmp_int_scs15 int64
		if tmp_int_scs15, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode scs15", err)
		}
		ie.scs15 = tmp_int_scs15
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_scs30:
		var tmp_int_scs30 int64
		if tmp_int_scs30, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode scs30", err)
		}
		ie.scs30 = tmp_int_scs30
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
