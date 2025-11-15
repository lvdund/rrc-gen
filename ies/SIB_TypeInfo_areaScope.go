package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_TypeInfo_areaScope_Enum_true uper.Enumerated = 0
)

type SIB_TypeInfo_areaScope struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SIB_TypeInfo_areaScope) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SIB_TypeInfo_areaScope", err)
	}
	return nil
}

func (ie *SIB_TypeInfo_areaScope) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SIB_TypeInfo_areaScope", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
