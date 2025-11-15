package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_InfoEUTRA_r17 struct {
	needForNCSG_EUTRA_r17 []NeedForNCSG_EUTRA_r17 `lb:1,ub:maxBandsEUTRA,madatory`
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_needForNCSG_EUTRA_r17 := utils.NewSequence[*NeedForNCSG_EUTRA_r17]([]*NeedForNCSG_EUTRA_r17{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	for _, i := range ie.needForNCSG_EUTRA_r17 {
		tmp_needForNCSG_EUTRA_r17.Value = append(tmp_needForNCSG_EUTRA_r17.Value, &i)
	}
	if err = tmp_needForNCSG_EUTRA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode needForNCSG_EUTRA_r17", err)
	}
	return nil
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_needForNCSG_EUTRA_r17 := utils.NewSequence[*NeedForNCSG_EUTRA_r17]([]*NeedForNCSG_EUTRA_r17{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	fn_needForNCSG_EUTRA_r17 := func() *NeedForNCSG_EUTRA_r17 {
		return new(NeedForNCSG_EUTRA_r17)
	}
	if err = tmp_needForNCSG_EUTRA_r17.Decode(r, fn_needForNCSG_EUTRA_r17); err != nil {
		return utils.WrapError("Decode needForNCSG_EUTRA_r17", err)
	}
	ie.needForNCSG_EUTRA_r17 = []NeedForNCSG_EUTRA_r17{}
	for _, i := range tmp_needForNCSG_EUTRA_r17.Value {
		ie.needForNCSG_EUTRA_r17 = append(ie.needForNCSG_EUTRA_r17, *i)
	}
	return nil
}
