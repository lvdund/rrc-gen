package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellReselectionSubPriority_Enum_oDot2 uper.Enumerated = 0
	CellReselectionSubPriority_Enum_oDot4 uper.Enumerated = 1
	CellReselectionSubPriority_Enum_oDot6 uper.Enumerated = 2
	CellReselectionSubPriority_Enum_oDot8 uper.Enumerated = 3
)

type CellReselectionSubPriority struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CellReselectionSubPriority) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CellReselectionSubPriority", err)
	}
	return nil
}

func (ie *CellReselectionSubPriority) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CellReselectionSubPriority", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
