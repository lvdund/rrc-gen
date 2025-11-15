package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoListMCG_EUTRA_r16 struct {
	Value []ServCellInfoXCG_EUTRA_r16 `lb:1,ub:maxNrofServingCellsEUTRA,madatory`
}

func (ie *ServCellInfoListMCG_EUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_EUTRA_r16]([]*ServCellInfoXCG_EUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsEUTRA}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellInfoListMCG_EUTRA_r16", err)
	}
	return nil
}

func (ie *ServCellInfoListMCG_EUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_EUTRA_r16]([]*ServCellInfoXCG_EUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsEUTRA}, false)
	fn := func() *ServCellInfoXCG_EUTRA_r16 {
		return new(ServCellInfoXCG_EUTRA_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ServCellInfoListMCG_EUTRA_r16", err)
	}
	ie.Value = []ServCellInfoXCG_EUTRA_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
