package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_CCCH_Message struct {
	message UL_CCCH_MessageType `madatory`
}

func (ie *UL_CCCH_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.message.Encode(w); err != nil {
		return utils.WrapError("Encode message", err)
	}
	return nil
}

func (ie *UL_CCCH_Message) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.message.Decode(r); err != nil {
		return utils.WrapError("Decode message", err)
	}
	return nil
}
