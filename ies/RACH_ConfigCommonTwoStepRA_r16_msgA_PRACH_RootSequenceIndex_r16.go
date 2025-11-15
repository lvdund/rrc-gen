package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_nothing uint64 = iota
	RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l839
	RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l139
	RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l571
	RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l1151
)

type RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16 struct {
	Choice uint64
	l839   int64 `lb:0,ub:837,madatory`
	l139   int64 `lb:0,ub:137,madatory`
	l571   int64 `lb:0,ub:569,madatory`
	l1151  int64 `lb:0,ub:1149,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l839:
		if err = w.WriteInteger(int64(ie.l839), &uper.Constraint{Lb: 0, Ub: 837}, false); err != nil {
			err = utils.WrapError("Encode l839", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l139:
		if err = w.WriteInteger(int64(ie.l139), &uper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			err = utils.WrapError("Encode l139", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l571:
		if err = w.WriteInteger(int64(ie.l571), &uper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			err = utils.WrapError("Encode l571", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l1151:
		if err = w.WriteInteger(int64(ie.l1151), &uper.Constraint{Lb: 0, Ub: 1149}, false); err != nil {
			err = utils.WrapError("Encode l1151", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l839:
		var tmp_int_l839 int64
		if tmp_int_l839, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 837}, false); err != nil {
			return utils.WrapError("Decode l839", err)
		}
		ie.l839 = tmp_int_l839
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l139:
		var tmp_int_l139 int64
		if tmp_int_l139, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Decode l139", err)
		}
		ie.l139 = tmp_int_l139
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l571:
		var tmp_int_l571 int64
		if tmp_int_l571, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			return utils.WrapError("Decode l571", err)
		}
		ie.l571 = tmp_int_l571
	case RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16_Choice_l1151:
		var tmp_int_l1151 int64
		if tmp_int_l1151, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1149}, false); err != nil {
			return utils.WrapError("Decode l1151", err)
		}
		ie.l1151 = tmp_int_l1151
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
