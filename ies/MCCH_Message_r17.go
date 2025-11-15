package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MCCH_Message_r17 struct {
	message MCCH_MessageType_r17 `madatory`
}

func (ie *MCCH_Message_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.message.Encode(w); err != nil {
		return utils.WrapError("Encode message", err)
	}
	return nil
}

func (ie *MCCH_Message_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.message.Decode(r); err != nil {
		return utils.WrapError("Decode message", err)
	}
	return nil
}
