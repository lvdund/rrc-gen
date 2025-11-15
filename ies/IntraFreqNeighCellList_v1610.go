package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqNeighCellList_v1610 struct {
	Value []IntraFreqNeighCellInfo_v1610 `lb:1,ub:maxCellIntra,madatory`
}

func (ie *IntraFreqNeighCellList_v1610) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo_v1610]([]*IntraFreqNeighCellInfo_v1610{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqNeighCellList_v1610", err)
	}
	return nil
}

func (ie *IntraFreqNeighCellList_v1610) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo_v1610]([]*IntraFreqNeighCellInfo_v1610{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	fn := func() *IntraFreqNeighCellInfo_v1610 {
		return new(IntraFreqNeighCellInfo_v1610)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraFreqNeighCellList_v1610", err)
	}
	ie.Value = []IntraFreqNeighCellInfo_v1610{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
