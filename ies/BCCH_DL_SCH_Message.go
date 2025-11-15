package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BCCH_DL_SCH_Message struct {
	message BCCH_DL_SCH_MessageType `madatory`
}

func (ie *BCCH_DL_SCH_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.message.Encode(w); err != nil {
		return utils.WrapError("Encode message", err)
	}
	return nil
}

func (ie *BCCH_DL_SCH_Message) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.message.Decode(r); err != nil {
		return utils.WrapError("Decode message", err)
	}
	return nil
}
