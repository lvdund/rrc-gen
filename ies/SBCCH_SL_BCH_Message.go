package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SBCCH_SL_BCH_Message struct {
	message SBCCH_SL_BCH_MessageType `madatory`
}

func (ie *SBCCH_SL_BCH_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.message.Encode(w); err != nil {
		return utils.WrapError("Encode message", err)
	}
	return nil
}

func (ie *SBCCH_SL_BCH_Message) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.message.Decode(r); err != nil {
		return utils.WrapError("Decode message", err)
	}
	return nil
}
