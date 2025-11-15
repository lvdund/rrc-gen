package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16_Enum_one   uper.Enumerated = 0
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16_Enum_upto2 uper.Enumerated = 1
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16_Enum_upto4 uper.Enumerated = 2
	FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16_Enum_upto7 uper.Enumerated = 3
)

type FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_scs_60kHz_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
