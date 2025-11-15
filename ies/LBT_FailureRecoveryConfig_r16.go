package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LBT_FailureRecoveryConfig_r16 struct {
	lbt_FailureInstanceMaxCount_r16 LBT_FailureRecoveryConfig_r16_lbt_FailureInstanceMaxCount_r16 `madatory`
	lbt_FailureDetectionTimer_r16   LBT_FailureRecoveryConfig_r16_lbt_FailureDetectionTimer_r16   `madatory`
}

func (ie *LBT_FailureRecoveryConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.lbt_FailureInstanceMaxCount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode lbt_FailureInstanceMaxCount_r16", err)
	}
	if err = ie.lbt_FailureDetectionTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode lbt_FailureDetectionTimer_r16", err)
	}
	return nil
}

func (ie *LBT_FailureRecoveryConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.lbt_FailureInstanceMaxCount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode lbt_FailureInstanceMaxCount_r16", err)
	}
	if err = ie.lbt_FailureDetectionTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode lbt_FailureDetectionTimer_r16", err)
	}
	return nil
}
