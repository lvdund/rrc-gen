package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles struct {
	profile0x0006 bool `madatory`
}

func (ie *PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.profile0x0006); err != nil {
		return utils.WrapError("WriteBoolean profile0x0006", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_profile0x0006 bool
	if tmp_bool_profile0x0006, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0006", err)
	}
	ie.profile0x0006 = tmp_bool_profile0x0006
	return nil
}
