package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_dummy_Enum_dynamic    uper.Enumerated = 0
	RateMatchPattern_dummy_Enum_semiStatic uper.Enumerated = 1
)

type RateMatchPattern_dummy struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RateMatchPattern_dummy) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RateMatchPattern_dummy", err)
	}
	return nil
}

func (ie *RateMatchPattern_dummy) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RateMatchPattern_dummy", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
