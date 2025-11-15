package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p02   uper.Enumerated = 0
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p04   uper.Enumerated = 1
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p0625 uper.Enumerated = 2
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p125  uper.Enumerated = 3
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p25   uper.Enumerated = 4
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_ms0p5    uper.Enumerated = 5
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_spare2   uper.Enumerated = 6
	SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16_Enum_spare1   uper.Enumerated = 7
)

type SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16", err)
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
