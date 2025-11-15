package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17 struct {
	maxNumBeamGroups_r17    int64                                                                     `lb:1,ub:4,madatory`
	maxNumRS_WithinSlot_r17 MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17_maxNumRS_WithinSlot_r17 `madatory`
	maxNumRS_AcrossSlot_r17 MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17_maxNumRS_AcrossSlot_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumBeamGroups_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumBeamGroups_r17", err)
	}
	if err = ie.maxNumRS_WithinSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumRS_WithinSlot_r17", err)
	}
	if err = ie.maxNumRS_AcrossSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumRS_AcrossSlot_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumBeamGroups_r17 int64
	if tmp_int_maxNumBeamGroups_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumBeamGroups_r17", err)
	}
	ie.maxNumBeamGroups_r17 = tmp_int_maxNumBeamGroups_r17
	if err = ie.maxNumRS_WithinSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumRS_WithinSlot_r17", err)
	}
	if err = ie.maxNumRS_AcrossSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumRS_AcrossSlot_r17", err)
	}
	return nil
}
