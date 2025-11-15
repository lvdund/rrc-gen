package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SemiStaticChannelAccessConfig_r16_period_Enum_ms1     uper.Enumerated = 0
	SemiStaticChannelAccessConfig_r16_period_Enum_ms2     uper.Enumerated = 1
	SemiStaticChannelAccessConfig_r16_period_Enum_ms2dot5 uper.Enumerated = 2
	SemiStaticChannelAccessConfig_r16_period_Enum_ms4     uper.Enumerated = 3
	SemiStaticChannelAccessConfig_r16_period_Enum_ms5     uper.Enumerated = 4
	SemiStaticChannelAccessConfig_r16_period_Enum_ms10    uper.Enumerated = 5
)

type SemiStaticChannelAccessConfig_r16_period struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SemiStaticChannelAccessConfig_r16_period) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SemiStaticChannelAccessConfig_r16_period", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfig_r16_period) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SemiStaticChannelAccessConfig_r16_period", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
