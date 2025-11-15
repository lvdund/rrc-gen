package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SN_FieldLengthAM_Enum_size12 uper.Enumerated = 0
	SN_FieldLengthAM_Enum_size18 uper.Enumerated = 1
)

type SN_FieldLengthAM struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SN_FieldLengthAM) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SN_FieldLengthAM", err)
	}
	return nil
}

func (ie *SN_FieldLengthAM) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SN_FieldLengthAM", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
