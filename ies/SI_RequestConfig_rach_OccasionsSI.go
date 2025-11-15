package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestConfig_rach_OccasionsSI struct {
	rach_ConfigSI        RACH_ConfigGeneric                                     `madatory`
	ssb_perRACH_Occasion SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion `madatory`
}

func (ie *SI_RequestConfig_rach_OccasionsSI) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rach_ConfigSI.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigSI", err)
	}
	if err = ie.ssb_perRACH_Occasion.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_perRACH_Occasion", err)
	}
	return nil
}

func (ie *SI_RequestConfig_rach_OccasionsSI) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rach_ConfigSI.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigSI", err)
	}
	if err = ie.ssb_perRACH_Occasion.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_perRACH_Occasion", err)
	}
	return nil
}
