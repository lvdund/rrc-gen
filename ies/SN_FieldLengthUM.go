package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SN_FieldLengthUM_Enum_size6  uper.Enumerated = 0
	SN_FieldLengthUM_Enum_size12 uper.Enumerated = 1
)

type SN_FieldLengthUM struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SN_FieldLengthUM) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SN_FieldLengthUM", err)
	}
	return nil
}

func (ie *SN_FieldLengthUM) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SN_FieldLengthUM", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
