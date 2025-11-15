package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_NeighbourCellList_r17 struct {
	Value []MBS_NeighbourCell_r17 `lb:0,ub:maxNeighCellMBS_r17,madatory`
}

func (ie *MBS_NeighbourCellList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MBS_NeighbourCell_r17]([]*MBS_NeighbourCell_r17{}, uper.Constraint{Lb: 0, Ub: maxNeighCellMBS_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MBS_NeighbourCellList_r17", err)
	}
	return nil
}

func (ie *MBS_NeighbourCellList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MBS_NeighbourCell_r17]([]*MBS_NeighbourCell_r17{}, uper.Constraint{Lb: 0, Ub: maxNeighCellMBS_r17}, false)
	fn := func() *MBS_NeighbourCell_r17 {
		return new(MBS_NeighbourCell_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MBS_NeighbourCellList_r17", err)
	}
	ie.Value = []MBS_NeighbourCell_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
