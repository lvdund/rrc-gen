package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_r17_mgl_r17_Enum_ms1     uper.Enumerated = 0
	GapConfig_r17_mgl_r17_Enum_ms1dot5 uper.Enumerated = 1
	GapConfig_r17_mgl_r17_Enum_ms2     uper.Enumerated = 2
	GapConfig_r17_mgl_r17_Enum_ms3     uper.Enumerated = 3
	GapConfig_r17_mgl_r17_Enum_ms3dot5 uper.Enumerated = 4
	GapConfig_r17_mgl_r17_Enum_ms4     uper.Enumerated = 5
	GapConfig_r17_mgl_r17_Enum_ms5     uper.Enumerated = 6
	GapConfig_r17_mgl_r17_Enum_ms5dot5 uper.Enumerated = 7
	GapConfig_r17_mgl_r17_Enum_ms6     uper.Enumerated = 8
	GapConfig_r17_mgl_r17_Enum_ms10    uper.Enumerated = 9
	GapConfig_r17_mgl_r17_Enum_ms20    uper.Enumerated = 10
)

type GapConfig_r17_mgl_r17 struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *GapConfig_r17_mgl_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode GapConfig_r17_mgl_r17", err)
	}
	return nil
}

func (ie *GapConfig_r17_mgl_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode GapConfig_r17_mgl_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
