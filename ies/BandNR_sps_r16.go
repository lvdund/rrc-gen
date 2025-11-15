package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_sps_r16 struct {
	maxNumberConfigsPerBWP_r16 int64 `lb:1,ub:8,madatory`
	maxNumberConfigsAllCC_r16  int64 `lb:2,ub:32,madatory`
}

func (ie *BandNR_sps_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberConfigsPerBWP_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberConfigsPerBWP_r16", err)
	}
	if err = w.WriteInteger(ie.maxNumberConfigsAllCC_r16, &uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberConfigsAllCC_r16", err)
	}
	return nil
}

func (ie *BandNR_sps_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberConfigsPerBWP_r16 int64
	if tmp_int_maxNumberConfigsPerBWP_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberConfigsPerBWP_r16", err)
	}
	ie.maxNumberConfigsPerBWP_r16 = tmp_int_maxNumberConfigsPerBWP_r16
	var tmp_int_maxNumberConfigsAllCC_r16 int64
	if tmp_int_maxNumberConfigsAllCC_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberConfigsAllCC_r16", err)
	}
	ie.maxNumberConfigsAllCC_r16 = tmp_int_maxNumberConfigsAllCC_r16
	return nil
}
