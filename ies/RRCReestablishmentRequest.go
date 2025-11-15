package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentRequest struct {
	rrcReestablishmentRequest RRCReestablishmentRequest_IEs `madatory`
}

func (ie *RRCReestablishmentRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrcReestablishmentRequest.Encode(w); err != nil {
		return utils.WrapError("Encode rrcReestablishmentRequest", err)
	}
	return nil
}

func (ie *RRCReestablishmentRequest) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrcReestablishmentRequest.Decode(r); err != nil {
		return utils.WrapError("Decode rrcReestablishmentRequest", err)
	}
	return nil
}
