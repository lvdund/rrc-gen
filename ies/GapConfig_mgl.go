package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgl_Enum_ms1dot5 uper.Enumerated = 0
	GapConfig_mgl_Enum_ms3     uper.Enumerated = 1
	GapConfig_mgl_Enum_ms3dot5 uper.Enumerated = 2
	GapConfig_mgl_Enum_ms4     uper.Enumerated = 3
	GapConfig_mgl_Enum_ms5dot5 uper.Enumerated = 4
	GapConfig_mgl_Enum_ms6     uper.Enumerated = 5
)

type GapConfig_mgl struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *GapConfig_mgl) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgl", err)
	}
	return nil
}

func (ie *GapConfig_mgl) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgl", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
