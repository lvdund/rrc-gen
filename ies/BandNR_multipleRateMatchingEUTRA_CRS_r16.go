package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_multipleRateMatchingEUTRA_CRS_r16 struct {
	maxNumberPatterns_r16            int64 `lb:2,ub:6,madatory`
	maxNumberNon_OverlapPatterns_r16 int64 `lb:1,ub:3,madatory`
}

func (ie *BandNR_multipleRateMatchingEUTRA_CRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberPatterns_r16, &uper.Constraint{Lb: 2, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPatterns_r16", err)
	}
	if err = w.WriteInteger(ie.maxNumberNon_OverlapPatterns_r16, &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberNon_OverlapPatterns_r16", err)
	}
	return nil
}

func (ie *BandNR_multipleRateMatchingEUTRA_CRS_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberPatterns_r16 int64
	if tmp_int_maxNumberPatterns_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPatterns_r16", err)
	}
	ie.maxNumberPatterns_r16 = tmp_int_maxNumberPatterns_r16
	var tmp_int_maxNumberNon_OverlapPatterns_r16 int64
	if tmp_int_maxNumberNon_OverlapPatterns_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberNon_OverlapPatterns_r16", err)
	}
	ie.maxNumberNon_OverlapPatterns_r16 = tmp_int_maxNumberNon_OverlapPatterns_r16
	return nil
}
