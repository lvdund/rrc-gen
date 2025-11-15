package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_Index struct {
	Value uint64 `lb:0,ub:maxNrofSSBs_1,madatory`
}

func (ie *SSB_Index) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSSBs_1}, false); err != nil {
		return utils.WrapError("Encode SSB_Index", err)
	}
	return nil
}

func (ie *SSB_Index) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSSBs_1}, false); err != nil {
		return utils.WrapError("Decode SSB_Index", err)
	}
	ie.Value = uint64(v)
	return nil
}
