package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_schedulingCellInfo_own struct {
	cif_Presence bool `madatory`
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_own) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.cif_Presence); err != nil {
		return utils.WrapError("WriteBoolean cif_Presence", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo_own) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_cif_Presence bool
	if tmp_bool_cif_Presence, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean cif_Presence", err)
	}
	ie.cif_Presence = tmp_bool_cif_Presence
	return nil
}
