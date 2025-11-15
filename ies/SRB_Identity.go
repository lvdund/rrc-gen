package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRB_Identity struct {
	Value uint64 `lb:1,ub:3,madatory`
}

func (ie *SRB_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SRB_Identity", err)
	}
	return nil
}

func (ie *SRB_Identity) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SRB_Identity", err)
	}
	ie.Value = uint64(v)
	return nil
}
