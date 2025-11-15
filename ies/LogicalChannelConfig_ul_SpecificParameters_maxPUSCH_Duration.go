package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p02       uper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p04       uper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p0625     uper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p125      uper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p25       uper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p5        uper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p01_v1700 uper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_spare1       uper.Enumerated = 7
)

type LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
