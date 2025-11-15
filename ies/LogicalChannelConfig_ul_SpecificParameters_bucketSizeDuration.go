package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms5    uper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms10   uper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms20   uper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms50   uper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms100  uper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms150  uper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms300  uper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms500  uper.Enumerated = 7
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_ms1000 uper.Enumerated = 8
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare7 uper.Enumerated = 9
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare6 uper.Enumerated = 10
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare5 uper.Enumerated = 11
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare4 uper.Enumerated = 12
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare3 uper.Enumerated = 13
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare2 uper.Enumerated = 14
	LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration_Enum_spare1 uper.Enumerated = 15
)

type LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
