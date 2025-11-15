package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_RangeIndex struct {
	Value uint64 `lb:1,ub:maxNrofPCI_Ranges,madatory`
}

func (ie *PCI_RangeIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false); err != nil {
		return utils.WrapError("Encode PCI_RangeIndex", err)
	}
	return nil
}

func (ie *PCI_RangeIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false); err != nil {
		return utils.WrapError("Decode PCI_RangeIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
