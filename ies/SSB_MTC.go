package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC struct {
	periodicityAndOffset SSB_MTC_periodicityAndOffset `lb:0,ub:4,madatory`
	duration             SSB_MTC_duration             `madatory`
}

func (ie *SSB_MTC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.periodicityAndOffset.Encode(w); err != nil {
		return utils.WrapError("Encode periodicityAndOffset", err)
	}
	if err = ie.duration.Encode(w); err != nil {
		return utils.WrapError("Encode duration", err)
	}
	return nil
}

func (ie *SSB_MTC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.periodicityAndOffset.Decode(r); err != nil {
		return utils.WrapError("Decode periodicityAndOffset", err)
	}
	if err = ie.duration.Decode(r); err != nil {
		return utils.WrapError("Decode duration", err)
	}
	return nil
}
