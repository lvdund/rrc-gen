package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_resourceMapping_nrofSymbols_Enum_n1 uper.Enumerated = 0
	SRS_Resource_resourceMapping_nrofSymbols_Enum_n2 uper.Enumerated = 1
	SRS_Resource_resourceMapping_nrofSymbols_Enum_n4 uper.Enumerated = 2
)

type SRS_Resource_resourceMapping_nrofSymbols struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SRS_Resource_resourceMapping_nrofSymbols) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_resourceMapping_nrofSymbols", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_nrofSymbols) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_resourceMapping_nrofSymbols", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
