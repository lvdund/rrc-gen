package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupRequest struct {
	rrcSetupRequest RRCSetupRequest_IEs `madatory`
}

func (ie *RRCSetupRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrcSetupRequest.Encode(w); err != nil {
		return utils.WrapError("Encode rrcSetupRequest", err)
	}
	return nil
}

func (ie *RRCSetupRequest) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrcSetupRequest.Decode(r); err != nil {
		return utils.WrapError("Decode rrcSetupRequest", err)
	}
	return nil
}
