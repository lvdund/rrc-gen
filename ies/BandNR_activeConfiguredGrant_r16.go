package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_activeConfiguredGrant_r16 struct {
	maxNumberConfigsPerBWP_r16 BandNR_activeConfiguredGrant_r16_maxNumberConfigsPerBWP_r16 `madatory`
	maxNumberConfigsAllCC_r16  int64                                                       `lb:2,ub:32,madatory`
}

func (ie *BandNR_activeConfiguredGrant_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberConfigsPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberConfigsPerBWP_r16", err)
	}
	if err = w.WriteInteger(ie.maxNumberConfigsAllCC_r16, &uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberConfigsAllCC_r16", err)
	}
	return nil
}

func (ie *BandNR_activeConfiguredGrant_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberConfigsPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberConfigsPerBWP_r16", err)
	}
	var tmp_int_maxNumberConfigsAllCC_r16 int64
	if tmp_int_maxNumberConfigsAllCC_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberConfigsAllCC_r16", err)
	}
	ie.maxNumberConfigsAllCC_r16 = tmp_int_maxNumberConfigsAllCC_r16
	return nil
}
