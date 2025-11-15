package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_t313_Expiry         uper.Enumerated = 0
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_randomAccessProblem uper.Enumerated = 1
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_rlc_MaxNumRetx      uper.Enumerated = 2
	CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA_Enum_scg_ChangeFailure   uper.Enumerated = 3
)

type CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA_failureTypeEUTRA", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
