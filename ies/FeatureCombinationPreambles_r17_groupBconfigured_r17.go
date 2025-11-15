package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17 struct {
	ra_SizeGroupA_r17              FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17            `madatory`
	messagePowerOffsetGroupB_r17   FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17 `madatory`
	numberOfRA_PreamblesGroupA_r17 int64                                                                             `lb:1,ub:64,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ra_SizeGroupA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ra_SizeGroupA_r17", err)
	}
	if err = ie.messagePowerOffsetGroupB_r17.Encode(w); err != nil {
		return utils.WrapError("Encode messagePowerOffsetGroupB_r17", err)
	}
	if err = w.WriteInteger(ie.numberOfRA_PreamblesGroupA_r17, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfRA_PreamblesGroupA_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ra_SizeGroupA_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ra_SizeGroupA_r17", err)
	}
	if err = ie.messagePowerOffsetGroupB_r17.Decode(r); err != nil {
		return utils.WrapError("Decode messagePowerOffsetGroupB_r17", err)
	}
	var tmp_int_numberOfRA_PreamblesGroupA_r17 int64
	if tmp_int_numberOfRA_PreamblesGroupA_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfRA_PreamblesGroupA_r17", err)
	}
	ie.numberOfRA_PreamblesGroupA_r17 = tmp_int_numberOfRA_PreamblesGroupA_r17
	return nil
}
