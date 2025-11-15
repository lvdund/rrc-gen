package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17 struct {
	ppw_durationOfPRS_ProcessingSymbolsN2_r17 PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17_ppw_durationOfPRS_ProcessingSymbolsN2_r17 `madatory`
	ppw_durationOfPRS_ProcessingSymbolsT2_r17 PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17_ppw_durationOfPRS_ProcessingSymbolsT2_r17 `madatory`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ppw_durationOfPRS_ProcessingSymbolsN2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ppw_durationOfPRS_ProcessingSymbolsN2_r17", err)
	}
	if err = ie.ppw_durationOfPRS_ProcessingSymbolsT2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ppw_durationOfPRS_ProcessingSymbolsT2_r17", err)
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ppw_durationOfPRS_ProcessingSymbolsN2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ppw_durationOfPRS_ProcessingSymbolsN2_r17", err)
	}
	if err = ie.ppw_durationOfPRS_ProcessingSymbolsT2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ppw_durationOfPRS_ProcessingSymbolsT2_r17", err)
	}
	return nil
}
