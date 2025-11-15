package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoListSCG_NR_r16 struct {
	Value []ServCellInfoXCG_NR_r16 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *ServCellInfoListSCG_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_NR_r16]([]*ServCellInfoXCG_NR_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellInfoListSCG_NR_r16", err)
	}
	return nil
}

func (ie *ServCellInfoListSCG_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_NR_r16]([]*ServCellInfoXCG_NR_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *ServCellInfoXCG_NR_r16 {
		return new(ServCellInfoXCG_NR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ServCellInfoListSCG_NR_r16", err)
	}
	ie.Value = []ServCellInfoXCG_NR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
