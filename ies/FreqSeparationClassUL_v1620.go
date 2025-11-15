package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClassUL_v1620_Enum_mhz1000 uper.Enumerated = 0
)

type FreqSeparationClassUL_v1620 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FreqSeparationClassUL_v1620) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClassUL_v1620", err)
	}
	return nil
}

func (ie *FreqSeparationClassUL_v1620) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClassUL_v1620", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
