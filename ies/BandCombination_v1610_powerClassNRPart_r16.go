package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_v1610_powerClassNRPart_r16_Enum_pc1 uper.Enumerated = 0
	BandCombination_v1610_powerClassNRPart_r16_Enum_pc2 uper.Enumerated = 1
	BandCombination_v1610_powerClassNRPart_r16_Enum_pc3 uper.Enumerated = 2
	BandCombination_v1610_powerClassNRPart_r16_Enum_pc5 uper.Enumerated = 3
)

type BandCombination_v1610_powerClassNRPart_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BandCombination_v1610_powerClassNRPart_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BandCombination_v1610_powerClassNRPart_r16", err)
	}
	return nil
}

func (ie *BandCombination_v1610_powerClassNRPart_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BandCombination_v1610_powerClassNRPart_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
