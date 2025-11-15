package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfigInfo_r16 struct {
	sl_DestinationIndex_r16 SL_DestinationIndex_r16 `madatory`
	sl_MeasConfig_r16       SL_MeasConfig_r16       `madatory`
}

func (ie *SL_MeasConfigInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_DestinationIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DestinationIndex_r16", err)
	}
	if err = ie.sl_MeasConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_MeasConfig_r16", err)
	}
	return nil
}

func (ie *SL_MeasConfigInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_DestinationIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DestinationIndex_r16", err)
	}
	if err = ie.sl_MeasConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_MeasConfig_r16", err)
	}
	return nil
}
