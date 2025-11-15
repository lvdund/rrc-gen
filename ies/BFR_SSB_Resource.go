package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BFR_SSB_Resource struct {
	ssb              SSB_Index `madatory`
	ra_PreambleIndex int64     `lb:0,ub:63,madatory`
}

func (ie *BFR_SSB_Resource) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssb.Encode(w); err != nil {
		return utils.WrapError("Encode ssb", err)
	}
	if err = w.WriteInteger(ie.ra_PreambleIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger ra_PreambleIndex", err)
	}
	return nil
}

func (ie *BFR_SSB_Resource) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssb.Decode(r); err != nil {
		return utils.WrapError("Decode ssb", err)
	}
	var tmp_int_ra_PreambleIndex int64
	if tmp_int_ra_PreambleIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger ra_PreambleIndex", err)
	}
	ie.ra_PreambleIndex = tmp_int_ra_PreambleIndex
	return nil
}
