package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCellActivationRS_ConfigId_r17 struct {
	Value uint64 `lb:1,ub:maxNrofSCellActRS_r17,madatory`
}

func (ie *SCellActivationRS_ConfigId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false); err != nil {
		return utils.WrapError("Encode SCellActivationRS_ConfigId_r17", err)
	}
	return nil
}

func (ie *SCellActivationRS_ConfigId_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false); err != nil {
		return utils.WrapError("Decode SCellActivationRS_ConfigId_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
