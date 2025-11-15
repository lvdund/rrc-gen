package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedNAS_Message struct {
	Value []byte `madatory`
}

func (ie *DedicatedNAS_Message) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DedicatedNAS_Message", err)
	}
	return nil
}

func (ie *DedicatedNAS_Message) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DedicatedNAS_Message", err)
	}
	ie.Value = v
	return nil
}
