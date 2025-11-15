package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MPE_Config_FR2_r17 struct {
	mpe_ProhibitTimer_r17 MPE_Config_FR2_r17_mpe_ProhibitTimer_r17 `madatory`
	mpe_Threshold_r17     MPE_Config_FR2_r17_mpe_Threshold_r17     `madatory`
	numberOfN_r17         int64                                    `lb:1,ub:4,madatory`
}

func (ie *MPE_Config_FR2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.mpe_ProhibitTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mpe_ProhibitTimer_r17", err)
	}
	if err = ie.mpe_Threshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mpe_Threshold_r17", err)
	}
	if err = w.WriteInteger(ie.numberOfN_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfN_r17", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.mpe_ProhibitTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mpe_ProhibitTimer_r17", err)
	}
	if err = ie.mpe_Threshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mpe_Threshold_r17", err)
	}
	var tmp_int_numberOfN_r17 int64
	if tmp_int_numberOfN_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfN_r17", err)
	}
	ie.numberOfN_r17 = tmp_int_numberOfN_r17
	return nil
}
