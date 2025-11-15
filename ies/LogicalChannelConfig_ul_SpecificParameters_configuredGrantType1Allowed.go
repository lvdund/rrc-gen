package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed_Enum_true uper.Enumerated = 0
)

type LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
