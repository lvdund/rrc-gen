package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest1 struct {
	rrcResumeRequest1 RRCResumeRequest1_IEs `madatory`
}

func (ie *RRCResumeRequest1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrcResumeRequest1.Encode(w); err != nil {
		return utils.WrapError("Encode rrcResumeRequest1", err)
	}
	return nil
}

func (ie *RRCResumeRequest1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrcResumeRequest1.Decode(r); err != nil {
		return utils.WrapError("Decode rrcResumeRequest1", err)
	}
	return nil
}
