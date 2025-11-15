package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_BearerConfig_reestablishRLC_Enum_true uper.Enumerated = 0
)

type RLC_BearerConfig_reestablishRLC struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RLC_BearerConfig_reestablishRLC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RLC_BearerConfig_reestablishRLC", err)
	}
	return nil
}

func (ie *RLC_BearerConfig_reestablishRLC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RLC_BearerConfig_reestablishRLC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
