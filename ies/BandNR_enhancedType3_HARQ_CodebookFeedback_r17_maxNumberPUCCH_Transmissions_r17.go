package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n1 uper.Enumerated = 0
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n2 uper.Enumerated = 1
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n3 uper.Enumerated = 2
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n4 uper.Enumerated = 3
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n5 uper.Enumerated = 4
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n6 uper.Enumerated = 5
	BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17_Enum_n7 uper.Enumerated = 6
)

type BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17", err)
	}
	return nil
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
