package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_FSAI_r17 struct {
	Value []byte `lb:3,ub:3,madatory`
}

func (ie *MBS_FSAI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MBS_FSAI_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_r17) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MBS_FSAI_r17", err)
	}
	ie.Value = v
	return nil
}
