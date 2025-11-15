package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17_Enum_pc1dot5 uper.Enumerated = 0
	FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17_Enum_pc2     uper.Enumerated = 1
	FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17_Enum_pc3     uper.Enumerated = 2
)

type FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
