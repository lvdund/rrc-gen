package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCCH_Message struct {
	message SCCH_MessageType `madatory`
}

func (ie *SCCH_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.message.Encode(w); err != nil {
		return utils.WrapError("Encode message", err)
	}
	return nil
}

func (ie *SCCH_Message) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.message.Decode(r); err != nil {
		return utils.WrapError("Decode message", err)
	}
	return nil
}
