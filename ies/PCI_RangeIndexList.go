package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_RangeIndexList struct {
	Value []PCI_RangeIndex `lb:1,ub:maxNrofPCI_Ranges,madatory`
}

func (ie *PCI_RangeIndexList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PCI_RangeIndex]([]*PCI_RangeIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PCI_RangeIndexList", err)
	}
	return nil
}

func (ie *PCI_RangeIndexList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PCI_RangeIndex]([]*PCI_RangeIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
	fn := func() *PCI_RangeIndex {
		return new(PCI_RangeIndex)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PCI_RangeIndexList", err)
	}
	ie.Value = []PCI_RangeIndex{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
