package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BetaOffsetsCrossPri_r17 struct {
	Value []int64 `lb:3,ub:3,madatory`
}

func (ie *BetaOffsetsCrossPri_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BetaOffsetsCrossPri_r17", err)
	}
	return nil
}

func (ie *BetaOffsetsCrossPri_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	fn := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BetaOffsetsCrossPri_r17", err)
	}
	ie.Value = []int64{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, int64(i.Value))
	}
	return err
}
