package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_InfoBroadcast_r17 struct {
	pdcp_Config_r17 MRB_PDCP_ConfigBroadcast_r17 `madatory`
	rlc_Config_r17  MRB_RLC_ConfigBroadcast_r17  `madatory`
}

func (ie *MRB_InfoBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pdcp_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcp_Config_r17", err)
	}
	if err = ie.rlc_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode rlc_Config_r17", err)
	}
	return nil
}

func (ie *MRB_InfoBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pdcp_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pdcp_Config_r17", err)
	}
	if err = ie.rlc_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode rlc_Config_r17", err)
	}
	return nil
}
