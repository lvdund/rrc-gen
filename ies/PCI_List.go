package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_List struct {
	Value []PhysCellId `lb:1,ub:maxNrofCellMeas,madatory`
}

func (ie *PCI_List) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PCI_List", err)
	}
	return nil
}

func (ie *PCI_List) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	fn := func() *PhysCellId {
		return new(PhysCellId)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PCI_List", err)
	}
	ie.Value = []PhysCellId{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
