package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize_Enum_n2 uper.Enumerated = 0
	ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize_Enum_n3 uper.Enumerated = 1
	ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize_Enum_n6 uper.Enumerated = 2
)

type ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize", err)
	}
	return nil
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
