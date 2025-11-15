package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceType_periodic struct {
	periodicityAndOffset_p SRS_PeriodicityAndOffset `madatory`
}

func (ie *SRS_Resource_resourceType_periodic) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.periodicityAndOffset_p.Encode(w); err != nil {
		return utils.WrapError("Encode periodicityAndOffset_p", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceType_periodic) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.periodicityAndOffset_p.Decode(r); err != nil {
		return utils.WrapError("Decode periodicityAndOffset_p", err)
	}
	return nil
}
