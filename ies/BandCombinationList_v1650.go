package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList_v1650 struct {
	Value []BandCombination_v1650 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList_v1650) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v1650]([]*BandCombination_v1650{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList_v1650", err)
	}
	return nil
}

func (ie *BandCombinationList_v1650) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v1650]([]*BandCombination_v1650{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination_v1650 {
		return new(BandCombination_v1650)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList_v1650", err)
	}
	ie.Value = []BandCombination_v1650{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
