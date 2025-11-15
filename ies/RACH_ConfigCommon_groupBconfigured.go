package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommon_groupBconfigured struct {
	ra_Msg3SizeGroupA          RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA        `madatory`
	messagePowerOffsetGroupB   RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB `madatory`
	numberOfRA_PreamblesGroupA int64                                                       `lb:1,ub:64,madatory`
}

func (ie *RACH_ConfigCommon_groupBconfigured) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ra_Msg3SizeGroupA.Encode(w); err != nil {
		return utils.WrapError("Encode ra_Msg3SizeGroupA", err)
	}
	if err = ie.messagePowerOffsetGroupB.Encode(w); err != nil {
		return utils.WrapError("Encode messagePowerOffsetGroupB", err)
	}
	if err = w.WriteInteger(ie.numberOfRA_PreamblesGroupA, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfRA_PreamblesGroupA", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_groupBconfigured) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ra_Msg3SizeGroupA.Decode(r); err != nil {
		return utils.WrapError("Decode ra_Msg3SizeGroupA", err)
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
