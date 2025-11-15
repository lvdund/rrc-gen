package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RadioLinkMonitoringRS struct {
	radioLinkMonitoringRS_Id RadioLinkMonitoringRS_Id                `madatory`
	purpose                  RadioLinkMonitoringRS_purpose           `madatory`
	detectionResource        RadioLinkMonitoringRS_detectionResource `madatory`
}

func (ie *RadioLinkMonitoringRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.radioLinkMonitoringRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode radioLinkMonitoringRS_Id", err)
	}
	if err = ie.purpose.Encode(w); err != nil {
		return utils.WrapError("Encode purpose", err)
	}
	if err = ie.detectionResource.Encode(w); err != nil {
		return utils.WrapError("Encode detectionResource", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.radioLinkMonitoringRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode radioLinkMonitoringRS_Id", err)
	}
	if err = ie.purpose.Decode(r); err != nil {
		return utils.WrapError("Decode purpose", err)
	}
	if err = ie.detectionResource.Decode(r); err != nil {
		return utils.WrapError("Decode detectionResource", err)
	}
	return nil
}
