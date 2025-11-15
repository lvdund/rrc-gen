package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList struct {
	Value []BandCombination `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination]([]*BandCombination{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList", err)
	}
	return nil
}

func (ie *BandCombinationList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination]([]*BandCombination{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination {
		return new(BandCombination)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList", err)
	}
	ie.Value = []BandCombination{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
