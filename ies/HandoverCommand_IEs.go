package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HandoverCommand_IEs struct {
	handoverCommandMessage []byte      `madatory`
	nonCriticalExtension   interface{} `optional`
}

func (ie *HandoverCommand_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.handoverCommandMessage, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString handoverCommandMessage", err)
	}
	return nil
}

func (ie *HandoverCommand_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_os_handoverCommandMessage []byte
	if tmp_os_handoverCommandMessage, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString handoverCommandMessage", err)
	}
	ie.handoverCommandMessage = tmp_os_handoverCommandMessage
	return nil
}
