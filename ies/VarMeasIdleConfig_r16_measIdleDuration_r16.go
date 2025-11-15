package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec10  uper.Enumerated = 0
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec30  uper.Enumerated = 1
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec60  uper.Enumerated = 2
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec120 uper.Enumerated = 3
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec180 uper.Enumerated = 4
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec240 uper.Enumerated = 5
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec300 uper.Enumerated = 6
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_spare  uper.Enumerated = 7
)

type VarMeasIdleConfig_r16_measIdleDuration_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *VarMeasIdleConfig_r16_measIdleDuration_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode VarMeasIdleConfig_r16_measIdleDuration_r16", err)
	}
	return nil
}

func (ie *VarMeasIdleConfig_r16_measIdleDuration_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode VarMeasIdleConfig_r16_measIdleDuration_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
