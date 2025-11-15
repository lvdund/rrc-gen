package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n1  uper.Enumerated = 0
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n2  uper.Enumerated = 1
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n3  uper.Enumerated = 2
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n4  uper.Enumerated = 3
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n5  uper.Enumerated = 4
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n6  uper.Enumerated = 5
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n8  uper.Enumerated = 6
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n10 uper.Enumerated = 7
)

type RadioLinkMonitoringConfig_beamFailureInstanceMaxCount struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringConfig_beamFailureInstanceMaxCount", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringConfig_beamFailureInstanceMaxCount", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
