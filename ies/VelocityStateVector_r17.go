package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VelocityStateVector_r17 struct {
	Value uint64 `lb:-131072,ub:131071,madatory`
}

func (ie *VelocityStateVector_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: -131072, Ub: 131071}, false); err != nil {
		return utils.WrapError("Encode VelocityStateVector_r17", err)
	}
	return nil
}

func (ie *VelocityStateVector_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: -131072, Ub: 131071}, false); err != nil {
		return utils.WrapError("Decode VelocityStateVector_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
