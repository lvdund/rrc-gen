package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17 struct {
	numBD_twoPDCCH_r17 int64                                                                 `lb:2,ub:3,madatory`
	maxNumOverlaps_r17 FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17 `madatory`
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.numBD_twoPDCCH_r17, &uper.Constraint{Lb: 2, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger numBD_twoPDCCH_r17", err)
	}
	if err = ie.maxNumOverlaps_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumOverlaps_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_numBD_twoPDCCH_r17 int64
	if tmp_int_numBD_twoPDCCH_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger numBD_twoPDCCH_r17", err)
	}
	ie.numBD_twoPDCCH_r17 = tmp_int_numBD_twoPDCCH_r17
	if err = ie.maxNumOverlaps_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumOverlaps_r17", err)
	}
	return nil
}
