package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n4  uper.Enumerated = 0
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n6  uper.Enumerated = 1
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n8  uper.Enumerated = 2
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n10 uper.Enumerated = 3
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n12 uper.Enumerated = 4
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n14 uper.Enumerated = 5
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n16 uper.Enumerated = 6
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n18 uper.Enumerated = 7
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n20 uper.Enumerated = 8
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n22 uper.Enumerated = 9
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n24 uper.Enumerated = 10
)

type BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17 struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17", err)
	}
	return nil
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
