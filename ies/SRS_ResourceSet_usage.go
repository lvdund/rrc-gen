package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_ResourceSet_usage_Enum_beamManagement   uper.Enumerated = 0
	SRS_ResourceSet_usage_Enum_codebook         uper.Enumerated = 1
	SRS_ResourceSet_usage_Enum_nonCodebook      uper.Enumerated = 2
	SRS_ResourceSet_usage_Enum_antennaSwitching uper.Enumerated = 3
)

type SRS_ResourceSet_usage struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SRS_ResourceSet_usage) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SRS_ResourceSet_usage", err)
	}
	return nil
}

func (ie *SRS_ResourceSet_usage) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SRS_ResourceSet_usage", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
