package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationIndex struct {
	Value uint64 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
		return utils.WrapError("Encode BandCombinationIndex", err)
	}
	return nil
}

func (ie *BandCombinationIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
		return utils.WrapError("Decode BandCombinationIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
