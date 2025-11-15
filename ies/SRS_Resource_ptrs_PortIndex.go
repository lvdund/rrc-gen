package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_ptrs_PortIndex_Enum_n0 uper.Enumerated = 0
	SRS_Resource_ptrs_PortIndex_Enum_n1 uper.Enumerated = 1
)

type SRS_Resource_ptrs_PortIndex struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRS_Resource_ptrs_PortIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_ptrs_PortIndex", err)
	}
	return nil
}

func (ie *SRS_Resource_ptrs_PortIndex) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_ptrs_PortIndex", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
