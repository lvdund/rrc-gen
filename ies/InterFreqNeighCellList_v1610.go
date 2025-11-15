package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellList_v1610 struct {
	Value []InterFreqNeighCellInfo_v1610 `lb:1,ub:maxCellInter,madatory`
}

func (ie *InterFreqNeighCellList_v1610) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo_v1610]([]*InterFreqNeighCellInfo_v1610{}, uper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqNeighCellList_v1610", err)
	}
	return nil
}

func (ie *InterFreqNeighCellList_v1610) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo_v1610]([]*InterFreqNeighCellInfo_v1610{}, uper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	fn := func() *InterFreqNeighCellInfo_v1610 {
		return new(InterFreqNeighCellInfo_v1610)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqNeighCellList_v1610", err)
	}
	ie.Value = []InterFreqNeighCellInfo_v1610{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
