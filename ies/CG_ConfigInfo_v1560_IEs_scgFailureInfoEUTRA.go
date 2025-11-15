package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA struct {
	failureTypeEUTRA    CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA `madatory`
	measResultSCG_EUTRA []byte                                                       `madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.failureTypeEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode failureTypeEUTRA", err)
	}
	if err = w.WriteOctetString(ie.measResultSCG_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString measResultSCG_EUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.failureTypeEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode failureTypeEUTRA", err)
	}
	var tmp_os_measResultSCG_EUTRA []byte
	if tmp_os_measResultSCG_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString measResultSCG_EUTRA", err)
	}
	ie.measResultSCG_EUTRA = tmp_os_measResultSCG_EUTRA
	return nil
}
