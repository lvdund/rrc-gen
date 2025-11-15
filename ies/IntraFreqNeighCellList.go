package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqNeighCellList struct {
	Value []IntraFreqNeighCellInfo `lb:1,ub:maxCellIntra,madatory`
}

func (ie *IntraFreqNeighCellList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo]([]*IntraFreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqNeighCellList", err)
	}
	return nil
}

func (ie *IntraFreqNeighCellList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo]([]*IntraFreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	fn := func() *IntraFreqNeighCellInfo {
		return new(IntraFreqNeighCellInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraFreqNeighCellList", err)
	}
	ie.Value = []IntraFreqNeighCellInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
