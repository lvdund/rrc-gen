package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ChoCandidateCellList_r17 struct {
	Value []ChoCandidateCell_r17 `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *ChoCandidateCellList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ChoCandidateCell_r17]([]*ChoCandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ChoCandidateCellList_r17", err)
	}
	return nil
}

func (ie *ChoCandidateCellList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ChoCandidateCell_r17]([]*ChoCandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	fn := func() *ChoCandidateCell_r17 {
		return new(ChoCandidateCell_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ChoCandidateCellList_r17", err)
	}
	ie.Value = []ChoCandidateCell_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
