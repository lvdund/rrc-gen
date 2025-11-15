package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17 struct {
	profile0x0000_r17 bool `madatory`
	profile0x0001_r17 bool `madatory`
	profile0x0002_r17 bool `madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.profile0x0000_r17); err != nil {
		return utils.WrapError("WriteBoolean profile0x0000_r17", err)
	}
	if err = w.WriteBoolean(ie.profile0x0001_r17); err != nil {
		return utils.WrapError("WriteBoolean profile0x0001_r17", err)
	}
	if err = w.WriteBoolean(ie.profile0x0002_r17); err != nil {
		return utils.WrapError("WriteBoolean profile0x0002_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_profile0x0000_r17 bool
	if tmp_bool_profile0x0000_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0000_r17", err)
	}
	ie.profile0x0000_r17 = tmp_bool_profile0x0000_r17
	var tmp_bool_profile0x0001_r17 bool
	if tmp_bool_profile0x0001_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0001_r17", err)
	}
	ie.profile0x0001_r17 = tmp_bool_profile0x0001_r17
	var tmp_bool_profile0x0002_r17 bool
	if tmp_bool_profile0x0002_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0002_r17", err)
	}
	ie.profile0x0002_r17 = tmp_bool_profile0x0002_r17
	return nil
}
