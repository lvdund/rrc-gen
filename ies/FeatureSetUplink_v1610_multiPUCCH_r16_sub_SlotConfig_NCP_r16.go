package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16_Enum_set1 uper.Enumerated = 0
	FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16_Enum_set2 uper.Enumerated = 1
)

type FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1610_multiPUCCH_r16_sub_SlotConfig_NCP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
