package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited struct {
	differentTB_PerSlot_SCS_30kHz FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited_differentTB_PerSlot_SCS_30kHz `madatory`
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.differentTB_PerSlot_SCS_30kHz.Encode(w); err != nil {
		return utils.WrapError("Encode differentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.differentTB_PerSlot_SCS_30kHz.Decode(r); err != nil {
		return utils.WrapError("Decode differentTB_PerSlot_SCS_30kHz", err)
	}
	return nil
}
