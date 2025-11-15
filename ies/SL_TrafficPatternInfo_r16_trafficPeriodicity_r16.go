package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms20   uper.Enumerated = 0
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms50   uper.Enumerated = 1
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms100  uper.Enumerated = 2
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms200  uper.Enumerated = 3
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms300  uper.Enumerated = 4
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms400  uper.Enumerated = 5
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms500  uper.Enumerated = 6
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms600  uper.Enumerated = 7
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms700  uper.Enumerated = 8
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms800  uper.Enumerated = 9
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms900  uper.Enumerated = 10
	SL_TrafficPatternInfo_r16_trafficPeriodicity_r16_Enum_ms1000 uper.Enumerated = 11
)

type SL_TrafficPatternInfo_r16_trafficPeriodicity_r16 struct {
	Value uper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *SL_TrafficPatternInfo_r16_trafficPeriodicity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode SL_TrafficPatternInfo_r16_trafficPeriodicity_r16", err)
	}
	return nil
}

func (ie *SL_TrafficPatternInfo_r16_trafficPeriodicity_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode SL_TrafficPatternInfo_r16_trafficPeriodicity_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
