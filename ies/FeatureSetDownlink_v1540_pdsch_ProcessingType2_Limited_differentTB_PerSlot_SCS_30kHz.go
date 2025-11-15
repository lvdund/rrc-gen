package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto1 uper.Enumerated = 0
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto2 uper.Enumerated = 1
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto4 uper.Enumerated = 2
	FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz_Enum_upto7 uper.Enumerated = 3
)

type FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
