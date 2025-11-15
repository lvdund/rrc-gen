package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PRB_Id struct {
	Value uint64 `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
}

func (ie *PRB_Id) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("Encode PRB_Id", err)
	}
	return nil
}

func (ie *PRB_Id) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("Decode PRB_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
