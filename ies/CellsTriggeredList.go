package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsTriggeredList struct {
	Value []CellsTriggeredItem `lb:1,ub:maxNrofCellMeas,madatory`
}

func (ie *CellsTriggeredList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CellsTriggeredItem]([]*CellsTriggeredItem{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellsTriggeredList", err)
	}
	return nil
}

func (ie *CellsTriggeredList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CellsTriggeredItem]([]*CellsTriggeredItem{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	fn := func() *CellsTriggeredItem {
		return new(CellsTriggeredItem)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellsTriggeredList", err)
	}
	ie.Value = []CellsTriggeredItem{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
