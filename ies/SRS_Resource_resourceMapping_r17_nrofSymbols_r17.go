package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n1  uper.Enumerated = 0
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n2  uper.Enumerated = 1
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n4  uper.Enumerated = 2
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n8  uper.Enumerated = 3
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n10 uper.Enumerated = 4
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n12 uper.Enumerated = 5
	SRS_Resource_resourceMapping_r17_nrofSymbols_r17_Enum_n14 uper.Enumerated = 6
)

type SRS_Resource_resourceMapping_r17_nrofSymbols_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SRS_Resource_resourceMapping_r17_nrofSymbols_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_resourceMapping_r17_nrofSymbols_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_r17_nrofSymbols_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_resourceMapping_r17_nrofSymbols_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
