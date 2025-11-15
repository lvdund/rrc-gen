package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList_v15g0 struct {
	Value []BandCombination_v15g0 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v15g0]([]*BandCombination_v15g0{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList_v15g0", err)
	}
	return nil
}

func (ie *BandCombinationList_v15g0) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v15g0]([]*BandCombination_v15g0{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination_v15g0 {
		return new(BandCombination_v15g0)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList_v15g0", err)
	}
	ie.Value = []BandCombination_v15g0{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
