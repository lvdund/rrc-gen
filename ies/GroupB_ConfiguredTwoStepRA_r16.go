package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GroupB_ConfiguredTwoStepRA_r16 struct {
	ra_MsgA_SizeGroupA         GroupB_ConfiguredTwoStepRA_r16_ra_MsgA_SizeGroupA       `madatory`
	messagePowerOffsetGroupB   GroupB_ConfiguredTwoStepRA_r16_messagePowerOffsetGroupB `madatory`
	numberOfRA_PreamblesGroupA int64                                                   `lb:1,ub:64,madatory`
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ra_MsgA_SizeGroupA.Encode(w); err != nil {
		return utils.WrapError("Encode ra_MsgA_SizeGroupA", err)
	}
	if err = ie.messagePowerOffsetGroupB.Encode(w); err != nil {
		return utils.WrapError("Encode messagePowerOffsetGroupB", err)
	}
	if err = w.WriteInteger(ie.numberOfRA_PreamblesGroupA, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfRA_PreamblesGroupA", err)
	}
	return nil
}

func (ie *GroupB_ConfiguredTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ra_MsgA_SizeGroupA.Decode(r); err != nil {
		return utils.WrapError("Decode ra_MsgA_SizeGroupA", err)
	}
	if err = ie.messagePowerOffsetGroupB.Decode(r); err != nil {
		return utils.WrapError("Decode messagePowerOffsetGroupB", err)
	}
	var tmp_int_numberOfRA_PreamblesGroupA int64
	if tmp_int_numberOfRA_PreamblesGroupA, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfRA_PreamblesGroupA", err)
	}
	ie.numberOfRA_PreamblesGroupA = tmp_int_numberOfRA_PreamblesGroupA
	return nil
}
