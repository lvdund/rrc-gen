package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdToAddMod struct {
	measId         MeasId         `madatory`
	measObjectId   MeasObjectId   `madatory`
	reportConfigId ReportConfigId `madatory`
}

func (ie *MeasIdToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.measId.Encode(w); err != nil {
		return utils.WrapError("Encode measId", err)
	}
	if err = ie.measObjectId.Encode(w); err != nil {
		return utils.WrapError("Encode measObjectId", err)
	}
	if err = ie.reportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfigId", err)
	}
	return nil
}

func (ie *MeasIdToAddMod) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.measId.Decode(r); err != nil {
		return utils.WrapError("Decode measId", err)
	}
	if err = ie.measObjectId.Decode(r); err != nil {
		return utils.WrapError("Decode measObjectId", err)
	}
	if err = ie.reportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode reportConfigId", err)
	}
	return nil
}
