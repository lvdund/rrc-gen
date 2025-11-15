package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetection struct {
	Value uint64 `lb:1,ub:15,madatory`
}

func (ie *PDCCH_BlindDetection) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDCCH_BlindDetection", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetection) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDCCH_BlindDetection", err)
	}
	ie.Value = uint64(v)
	return nil
}
