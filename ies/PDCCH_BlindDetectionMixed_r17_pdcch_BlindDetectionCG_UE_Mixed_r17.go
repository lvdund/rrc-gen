package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17 struct {
	pdcch_BlindDetectionMCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17 `madatory`
	pdcch_BlindDetectionSCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17 `madatory`
}

func (ie *PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_v17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_BlindDetectionMCG_UE_Mixed_v17", err)
	}
	if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_v17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_BlindDetectionSCG_UE_Mixed_v17", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_v17.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_BlindDetectionMCG_UE_Mixed_v17", err)
	}
	if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_v17.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_BlindDetectionSCG_UE_Mixed_v17", err)
	}
	return nil
}
