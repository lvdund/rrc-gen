package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIB_cellBarred_Enum_barred    uper.Enumerated = 0
	MIB_cellBarred_Enum_notBarred uper.Enumerated = 1
)

type MIB_cellBarred struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MIB_cellBarred) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MIB_cellBarred", err)
	}
	return nil
}

func (ie *MIB_cellBarred) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MIB_cellBarred", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
