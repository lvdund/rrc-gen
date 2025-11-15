package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectToAddMod struct {
	measObjectId MeasObjectId                  `madatory`
	measObject   MeasObjectToAddMod_measObject `madatory`
}

func (ie *MeasObjectToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.measObjectId.Encode(w); err != nil {
		return utils.WrapError("Encode measObjectId", err)
	}
	if err = ie.measObject.Encode(w); err != nil {
		return utils.WrapError("Encode measObject", err)
	}
	return nil
}

func (ie *MeasObjectToAddMod) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.measObjectId.Decode(r); err != nil {
		return utils.WrapError("Decode measObjectId", err)
	}
	if err = ie.measObject.Decode(r); err != nil {
		return utils.WrapError("Decode measObject", err)
	}
	return nil
}
