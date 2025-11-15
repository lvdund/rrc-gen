package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInfoDAPS_r16 struct {
	failureType_r16 FailureInfoDAPS_r16_failureType_r16 `madatory`
}

func (ie *FailureInfoDAPS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.failureType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode failureType_r16", err)
	}
	return nil
}

func (ie *FailureInfoDAPS_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.failureType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode failureType_r16", err)
	}
	return nil
}
