package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ScramblingId struct {
	Value uint64 `lb:0,ub:1023,madatory`
}

func (ie *ScramblingId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("Encode ScramblingId", err)
	}
	return nil
}

func (ie *ScramblingId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("Decode ScramblingId", err)
	}
	ie.Value = uint64(v)
	return nil
}
