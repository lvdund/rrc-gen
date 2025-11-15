package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3_Enum_short uper.Enumerated = 0
	FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3_Enum_long  uper.Enumerated = 1
)

type FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
