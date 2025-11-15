package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16_Enum_p0 uper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16_Enum_p1 uper.Enumerated = 1
)

type LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
