package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps0     uper.Enumerated = 0
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps8     uper.Enumerated = 1
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps16    uper.Enumerated = 2
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps32    uper.Enumerated = 3
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps64    uper.Enumerated = 4
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps128   uper.Enumerated = 5
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps256   uper.Enumerated = 6
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps512   uper.Enumerated = 7
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps1024  uper.Enumerated = 8
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps2048  uper.Enumerated = 9
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps4096  uper.Enumerated = 10
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps8192  uper.Enumerated = 11
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps16384 uper.Enumerated = 12
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps32768 uper.Enumerated = 13
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps65536 uper.Enumerated = 14
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_infinity  uper.Enumerated = 15
)

type SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16", err)
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
