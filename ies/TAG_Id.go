package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TAG_Id struct {
	Value uint64 `lb:0,ub:maxNrofTAGs_1,madatory`
}

func (ie *TAG_Id) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofTAGs_1}, false); err != nil {
		return utils.WrapError("Encode TAG_Id", err)
	}
	return nil
}

func (ie *TAG_Id) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofTAGs_1}, false); err != nil {
		return utils.WrapError("Decode TAG_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
