package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_RSTD_Info struct {
	carrierFreq    ARFCN_ValueEUTRA `madatory`
	measPRS_Offset int64            `lb:0,ub:39,madatory`
}

func (ie *EUTRA_RSTD_Info) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if err = w.WriteInteger(ie.measPRS_Offset, &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger measPRS_Offset", err)
	}
	return nil
}

func (ie *EUTRA_RSTD_Info) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	var tmp_int_measPRS_Offset int64
	if tmp_int_measPRS_Offset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger measPRS_Offset", err)
	}
	ie.measPRS_Offset = tmp_int_measPRS_Offset
	return nil
}
