package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosGapConfig_r17_mgl_r17_Enum_ms1dot5 uper.Enumerated = 0
	PosGapConfig_r17_mgl_r17_Enum_ms3     uper.Enumerated = 1
	PosGapConfig_r17_mgl_r17_Enum_ms3dot5 uper.Enumerated = 2
	PosGapConfig_r17_mgl_r17_Enum_ms4     uper.Enumerated = 3
	PosGapConfig_r17_mgl_r17_Enum_ms5dot5 uper.Enumerated = 4
	PosGapConfig_r17_mgl_r17_Enum_ms6     uper.Enumerated = 5
	PosGapConfig_r17_mgl_r17_Enum_ms10    uper.Enumerated = 6
	PosGapConfig_r17_mgl_r17_Enum_ms20    uper.Enumerated = 7
)

type PosGapConfig_r17_mgl_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PosGapConfig_r17_mgl_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PosGapConfig_r17_mgl_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17_mgl_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PosGapConfig_r17_mgl_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
