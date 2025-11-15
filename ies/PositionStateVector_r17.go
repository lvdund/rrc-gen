package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PositionStateVector_r17 struct {
	Value uint64 `lb:-33554432,ub:33554431,madatory`
}

func (ie *PositionStateVector_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: -33554432, Ub: 33554431}, false); err != nil {
		return utils.WrapError("Encode PositionStateVector_r17", err)
	}
	return nil
}

func (ie *PositionStateVector_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: -33554432, Ub: 33554431}, false); err != nil {
		return utils.WrapError("Decode PositionStateVector_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
