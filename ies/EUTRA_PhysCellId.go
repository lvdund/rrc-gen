package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_PhysCellId struct {
	Value uint64 `lb:0,ub:503,madatory`
}

func (ie *EUTRA_PhysCellId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 503}, false); err != nil {
		return utils.WrapError("Encode EUTRA_PhysCellId", err)
	}
	return nil
}

func (ie *EUTRA_PhysCellId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 503}, false); err != nil {
		return utils.WrapError("Decode EUTRA_PhysCellId", err)
	}
	ie.Value = uint64(v)
	return nil
}
