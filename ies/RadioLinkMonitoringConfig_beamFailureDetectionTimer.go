package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd1  uper.Enumerated = 0
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd2  uper.Enumerated = 1
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd3  uper.Enumerated = 2
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd4  uper.Enumerated = 3
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd5  uper.Enumerated = 4
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd6  uper.Enumerated = 5
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd8  uper.Enumerated = 6
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd10 uper.Enumerated = 7
)

type RadioLinkMonitoringConfig_beamFailureDetectionTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RadioLinkMonitoringConfig_beamFailureDetectionTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringConfig_beamFailureDetectionTimer", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringConfig_beamFailureDetectionTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringConfig_beamFailureDetectionTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
