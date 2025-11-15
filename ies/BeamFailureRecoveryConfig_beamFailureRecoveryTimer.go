package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms10  uper.Enumerated = 0
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms20  uper.Enumerated = 1
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms40  uper.Enumerated = 2
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms60  uper.Enumerated = 3
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms80  uper.Enumerated = 4
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms100 uper.Enumerated = 5
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms150 uper.Enumerated = 6
	BeamFailureRecoveryConfig_beamFailureRecoveryTimer_Enum_ms200 uper.Enumerated = 7
)

type BeamFailureRecoveryConfig_beamFailureRecoveryTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BeamFailureRecoveryConfig_beamFailureRecoveryTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BeamFailureRecoveryConfig_beamFailureRecoveryTimer", err)
	}
	return nil
}

func (ie *BeamFailureRecoveryConfig_beamFailureRecoveryTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BeamFailureRecoveryConfig_beamFailureRecoveryTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
