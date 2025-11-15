package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_t310_Expiry              uper.Enumerated = 0
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_randomAccessProblem      uper.Enumerated = 1
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_rlc_MaxNumRetx           uper.Enumerated = 2
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_synchReconfigFailure_SCG uper.Enumerated = 3
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_scg_reconfigFailure      uper.Enumerated = 4
	CG_ConfigInfo_IEs_scgFailureInfo_failureType_Enum_srb3_IntegrityFailure    uper.Enumerated = 5
)

type CG_ConfigInfo_IEs_scgFailureInfo_failureType struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo_failureType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_IEs_scgFailureInfo_failureType", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs_scgFailureInfo_failureType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_IEs_scgFailureInfo_failureType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
