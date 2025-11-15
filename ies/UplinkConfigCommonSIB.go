package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommonSIB struct {
	frequencyInfoUL          FrequencyInfoUL_SIB `madatory`
	initialUplinkBWP         BWP_UplinkCommon    `madatory`
	timeAlignmentTimerCommon TimeAlignmentTimer  `madatory`
}

func (ie *UplinkConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.frequencyInfoUL.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyInfoUL", err)
	}
	if err = ie.initialUplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode initialUplinkBWP", err)
	}
	if err = ie.timeAlignmentTimerCommon.Encode(w); err != nil {
		return utils.WrapError("Encode timeAlignmentTimerCommon", err)
	}
	return nil
}

func (ie *UplinkConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.frequencyInfoUL.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyInfoUL", err)
	}
	if err = ie.initialUplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode initialUplinkBWP", err)
	}
	if err = ie.timeAlignmentTimerCommon.Decode(r); err != nil {
		return utils.WrapError("Decode timeAlignmentTimerCommon", err)
	}
	return nil
}
