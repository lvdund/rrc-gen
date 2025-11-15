package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqNeighCellList_v1710 struct {
	Value []IntraFreqNeighCellInfo_v1710 `lb:1,ub:maxCellIntra,madatory`
}

func (ie *IntraFreqNeighCellList_v1710) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo_v1710]([]*IntraFreqNeighCellInfo_v1710{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqNeighCellList_v1710", err)
	}
	return nil
}

func (ie *IntraFreqNeighCellList_v1710) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*IntraFreqNeighCellInfo_v1710]([]*IntraFreqNeighCellInfo_v1710{}, uper.Constraint{Lb: 1, Ub: maxCellIntra}, false)
	fn := func() *IntraFreqNeighCellInfo_v1710 {
		return new(IntraFreqNeighCellInfo_v1710)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraFreqNeighCellList_v1710", err)
	}
	ie.Value = []IntraFreqNeighCellInfo_v1710{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
