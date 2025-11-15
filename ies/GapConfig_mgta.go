package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgta_Enum_ms0      uper.Enumerated = 0
	GapConfig_mgta_Enum_ms0dot25 uper.Enumerated = 1
	GapConfig_mgta_Enum_ms0dot5  uper.Enumerated = 2
)

type GapConfig_mgta struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *GapConfig_mgta) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgta", err)
	}
	return nil
}

func (ie *GapConfig_mgta) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgta", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
