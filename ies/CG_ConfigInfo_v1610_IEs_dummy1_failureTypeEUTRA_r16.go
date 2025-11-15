package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_scg_lbtFailure_r16             uper.Enumerated = 0
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_beamFailureRecoveryFailure_r16 uper.Enumerated = 1
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_t312_Expiry_r16                uper.Enumerated = 2
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_spare5                         uper.Enumerated = 3
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_spare4                         uper.Enumerated = 4
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_spare3                         uper.Enumerated = 5
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_spare2                         uper.Enumerated = 6
	CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16_Enum_spare1                         uper.Enumerated = 7
)

type CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_v1610_IEs_dummy1_failureTypeEUTRA_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
