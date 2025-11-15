package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_PositionQCL_CellList_r17 struct {
	Value []SSB_PositionQCL_Cell_r17 `lb:1,ub:maxNrofCellMeas,madatory`
}

func (ie *SSB_PositionQCL_CellList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SSB_PositionQCL_Cell_r17]([]*SSB_PositionQCL_Cell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SSB_PositionQCL_CellList_r17", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_CellList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SSB_PositionQCL_Cell_r17]([]*SSB_PositionQCL_Cell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	fn := func() *SSB_PositionQCL_Cell_r17 {
		return new(SSB_PositionQCL_Cell_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SSB_PositionQCL_CellList_r17", err)
	}
	ie.Value = []SSB_PositionQCL_Cell_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
