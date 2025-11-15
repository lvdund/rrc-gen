package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_powerClass_v1530_Enum_pc2 uper.Enumerated = 0
)

type BandCombination_powerClass_v1530 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandCombination_powerClass_v1530) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandCombination_powerClass_v1530", err)
	}
	return nil
}

func (ie *BandCombination_powerClass_v1530) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandCombination_powerClass_v1530", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
