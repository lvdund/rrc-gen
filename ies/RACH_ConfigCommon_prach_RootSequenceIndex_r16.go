package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_nothing uint64 = iota
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l571
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l1151
)

type RACH_ConfigCommon_prach_RootSequenceIndex_r16 struct {
	Choice uint64
	l571   int64 `lb:0,ub:569,madatory`
	l1151  int64 `lb:0,ub:1149,madatory`
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l571:
		if err = w.WriteInteger(int64(ie.l571), &uper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			err = utils.WrapError("Encode l571", err)
		}
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l1151:
		if err = w.WriteInteger(int64(ie.l1151), &uper.Constraint{Lb: 0, Ub: 1149}, false); err != nil {
			err = utils.WrapError("Encode l1151", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l571:
		var tmp_int_l571 int64
		if tmp_int_l571, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			return utils.WrapError("Decode l571", err)
		}
		ie.l571 = tmp_int_l571
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_l1151:
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
