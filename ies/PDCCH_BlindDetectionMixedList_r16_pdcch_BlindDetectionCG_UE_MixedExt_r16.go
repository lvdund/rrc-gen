package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16 struct {
	pdcch_BlindDetectionMCG_UE_Mixed_v16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16 `madatory`
	pdcch_BlindDetectionSCG_UE_Mixed_v16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16 `madatory`
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_v16a0.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_BlindDetectionMCG_UE_Mixed_v16a0", err)
	}
	if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_v16a0.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_BlindDetectionSCG_UE_Mixed_v16a0", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_v16a0.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_BlindDetectionMCG_UE_Mixed_v16a0", err)
	}
	if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_v16a0.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_BlindDetectionSCG_UE_Mixed_v16a0", err)
	}
	return nil
}
