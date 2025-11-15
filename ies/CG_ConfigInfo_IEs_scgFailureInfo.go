package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_IEs_scgFailureInfo struct {
	failureType   CG_ConfigInfo_IEs_scgFailureInfo_failureType `madatory`
	measResultSCG []byte                                       `madatory`
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.failureType.Encode(w); err != nil {
		return utils.WrapError("Encode failureType", err)
	}
	if err = w.WriteOctetString(ie.measResultSCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString measResultSCG", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.failureType.Decode(r); err != nil {
		return utils.WrapError("Decode failureType", err)
	}
	var tmp_os_measResultSCG []byte
	if tmp_os_measResultSCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString measResultSCG", err)
	}
	ie.measResultSCG = tmp_os_measResultSCG
	return nil
}
