package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest struct {
	rrcResumeRequest RRCResumeRequest_IEs `madatory`
}

func (ie *RRCResumeRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrcResumeRequest.Encode(w); err != nil {
		return utils.WrapError("Encode rrcResumeRequest", err)
	}
	return nil
}

func (ie *RRCResumeRequest) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrcResumeRequest.Decode(r); err != nil {
		return utils.WrapError("Decode rrcResumeRequest", err)
	}
	return nil
}
