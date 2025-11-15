package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_enhancedType3_HARQ_CodebookFeedback_r17 struct {
	enhancedType3_HARQ_Codebooks_r17 BandNR_enhancedType3_HARQ_CodebookFeedback_r17_enhancedType3_HARQ_Codebooks_r17 `madatory`
	maxNumberPUCCH_Transmissions_r17 BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17 `madatory`
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.enhancedType3_HARQ_Codebooks_r17.Encode(w); err != nil {
		return utils.WrapError("Encode enhancedType3_HARQ_Codebooks_r17", err)
	}
	if err = ie.maxNumberPUCCH_Transmissions_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPUCCH_Transmissions_r17", err)
	}
	return nil
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.enhancedType3_HARQ_Codebooks_r17.Decode(r); err != nil {
		return utils.WrapError("Decode enhancedType3_HARQ_Codebooks_r17", err)
	}
	if err = ie.maxNumberPUCCH_Transmissions_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPUCCH_Transmissions_r17", err)
	}
	return nil
}
