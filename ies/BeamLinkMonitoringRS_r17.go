package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamLinkMonitoringRS_r17 struct {
	beamLinkMonitoringRS_Id_r17 BeamLinkMonitoringRS_Id_r17                    `madatory`
	detectionResource_r17       BeamLinkMonitoringRS_r17_detectionResource_r17 `madatory`
}

func (ie *BeamLinkMonitoringRS_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.beamLinkMonitoringRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode beamLinkMonitoringRS_Id_r17", err)
	}
	if err = ie.detectionResource_r17.Encode(w); err != nil {
		return utils.WrapError("Encode detectionResource_r17", err)
	}
	return nil
}

func (ie *BeamLinkMonitoringRS_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.beamLinkMonitoringRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode beamLinkMonitoringRS_Id_r17", err)
	}
	if err = ie.detectionResource_r17.Decode(r); err != nil {
		return utils.WrapError("Decode detectionResource_r17", err)
	}
	return nil
}
