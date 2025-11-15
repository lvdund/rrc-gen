package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_formatExt_v1700 struct {
	nrofPRBs_r17 int64 `lb:1,ub:16,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_formatExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.nrofPRBs_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger nrofPRBs_r17", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_formatExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_nrofPRBs_r17 int64
	if tmp_int_nrofPRBs_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger nrofPRBs_r17", err)
	}
	ie.nrofPRBs_r17 = tmp_int_nrofPRBs_r17
	return nil
}
