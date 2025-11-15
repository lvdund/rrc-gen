package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRB_Identity_v1700 struct {
	Value uint64 `lb:4,ub:4,madatory`
}

func (ie *SRB_Identity_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SRB_Identity_v1700", err)
	}
	return nil
}

func (ie *SRB_Identity_v1700) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SRB_Identity_v1700", err)
	}
	ie.Value = uint64(v)
	return nil
}
