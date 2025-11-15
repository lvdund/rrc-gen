package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps0     uper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps8     uper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps16    uper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps32    uper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps64    uper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps128   uper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps256   uper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps512   uper.Enumerated = 7
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps1024  uper.Enumerated = 8
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps2048  uper.Enumerated = 9
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps4096  uper.Enumerated = 10
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps8192  uper.Enumerated = 11
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps16384 uper.Enumerated = 12
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps32768 uper.Enumerated = 13
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps65536 uper.Enumerated = 14
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_infinity  uper.Enumerated = 15
)

type LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
