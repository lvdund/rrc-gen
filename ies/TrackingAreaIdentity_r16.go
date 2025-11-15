package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TrackingAreaIdentity_r16 struct {
	plmn_Identity_r16    PLMN_Identity    `madatory`
	trackingAreaCode_r16 TrackingAreaCode `madatory`
}

func (ie *TrackingAreaIdentity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r16", err)
	}
	if err = ie.trackingAreaCode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode trackingAreaCode_r16", err)
	}
	return nil
}

func (ie *TrackingAreaIdentity_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r16", err)
	}
	if err = ie.trackingAreaCode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode trackingAreaCode_r16", err)
	}
	return nil
}
