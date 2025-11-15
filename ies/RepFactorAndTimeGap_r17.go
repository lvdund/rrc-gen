package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RepFactorAndTimeGap_r17 struct {
	repetitionFactor_r17 RepFactorAndTimeGap_r17_repetitionFactor_r17 `madatory`
	timeGap_r17          RepFactorAndTimeGap_r17_timeGap_r17          `madatory`
}

func (ie *RepFactorAndTimeGap_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.repetitionFactor_r17.Encode(w); err != nil {
		return utils.WrapError("Encode repetitionFactor_r17", err)
	}
	if err = ie.timeGap_r17.Encode(w); err != nil {
		return utils.WrapError("Encode timeGap_r17", err)
	}
	return nil
}

func (ie *RepFactorAndTimeGap_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.repetitionFactor_r17.Decode(r); err != nil {
		return utils.WrapError("Decode repetitionFactor_r17", err)
	}
	if err = ie.timeGap_r17.Decode(r); err != nil {
		return utils.WrapError("Decode timeGap_r17", err)
	}
	return nil
}
