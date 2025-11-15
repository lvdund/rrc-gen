package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceId struct {
	Value uint64 `lb:0,ub:maxNrofSearchSpaces_1,madatory`
}

func (ie *SearchSpaceId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSearchSpaces_1}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceId", err)
	}
	return nil
}

func (ie *SearchSpaceId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSearchSpaces_1}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
