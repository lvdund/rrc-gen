package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_triggeredHARQ_CodebookRetx_r17 struct {
	minHARQ_Retx_Offset_r17 BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17 `madatory`
	maxHARQ_Retx_Offset_r17 BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17 `madatory`
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.minHARQ_Retx_Offset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode minHARQ_Retx_Offset_r17", err)
	}
	if err = ie.maxHARQ_Retx_Offset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxHARQ_Retx_Offset_r17", err)
	}
	return nil
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.minHARQ_Retx_Offset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode minHARQ_Retx_Offset_r17", err)
	}
	if err = ie.maxHARQ_Retx_Offset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxHARQ_Retx_Offset_r17", err)
	}
	return nil
}
