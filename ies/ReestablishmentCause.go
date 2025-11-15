package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReestablishmentCause_Enum_reconfigurationFailure uper.Enumerated = 0
	ReestablishmentCause_Enum_handoverFailure        uper.Enumerated = 1
	ReestablishmentCause_Enum_otherFailure           uper.Enumerated = 2
	ReestablishmentCause_Enum_spare1                 uper.Enumerated = 3
)

type ReestablishmentCause struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ReestablishmentCause) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ReestablishmentCause", err)
	}
	return nil
}

func (ie *ReestablishmentCause) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ReestablishmentCause", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
