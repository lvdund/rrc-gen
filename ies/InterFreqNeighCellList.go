package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellList struct {
	Value []InterFreqNeighCellInfo `lb:1,ub:maxCellInter,madatory`
}

func (ie *InterFreqNeighCellList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo]([]*InterFreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqNeighCellList", err)
	}
	return nil
}

func (ie *InterFreqNeighCellList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqNeighCellInfo]([]*InterFreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellInter}, false)
	fn := func() *InterFreqNeighCellInfo {
		return new(InterFreqNeighCellInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqNeighCellList", err)
	}
	ie.Value = []InterFreqNeighCellInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
