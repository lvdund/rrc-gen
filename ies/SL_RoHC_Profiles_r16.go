package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RoHC_Profiles_r16 struct {
	profile0x0001_r16 bool `madatory`
	profile0x0002_r16 bool `madatory`
	profile0x0003_r16 bool `madatory`
	profile0x0004_r16 bool `madatory`
	profile0x0006_r16 bool `madatory`
	profile0x0101_r16 bool `madatory`
	profile0x0102_r16 bool `madatory`
	profile0x0103_r16 bool `madatory`
	profile0x0104_r16 bool `madatory`
}

func (ie *SL_RoHC_Profiles_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.profile0x0001_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0001_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0002_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0002_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0003_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0003_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0004_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0004_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0006_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0006_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0101_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0101_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0102_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0102_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0103_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0103_r16", err)
	}
	if err = w.WriteBoolean(ie.profile0x0104_r16); err != nil {
		return utils.WrapError("WriteBoolean profile0x0104_r16", err)
	}
	return nil
}

func (ie *SL_RoHC_Profiles_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_profile0x0001_r16 bool
	if tmp_bool_profile0x0001_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0001_r16", err)
	}
	ie.profile0x0001_r16 = tmp_bool_profile0x0001_r16
	var tmp_bool_profile0x0002_r16 bool
	if tmp_bool_profile0x0002_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0002_r16", err)
	}
	ie.profile0x0002_r16 = tmp_bool_profile0x0002_r16
	var tmp_bool_profile0x0003_r16 bool
	if tmp_bool_profile0x0003_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0003_r16", err)
	}
	ie.profile0x0003_r16 = tmp_bool_profile0x0003_r16
	var tmp_bool_profile0x0004_r16 bool
	if tmp_bool_profile0x0004_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0004_r16", err)
	}
	ie.profile0x0004_r16 = tmp_bool_profile0x0004_r16
	var tmp_bool_profile0x0006_r16 bool
	if tmp_bool_profile0x0006_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0006_r16", err)
	}
	ie.profile0x0006_r16 = tmp_bool_profile0x0006_r16
	var tmp_bool_profile0x0101_r16 bool
	if tmp_bool_profile0x0101_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0101_r16", err)
	}
	ie.profile0x0101_r16 = tmp_bool_profile0x0101_r16
	var tmp_bool_profile0x0102_r16 bool
	if tmp_bool_profile0x0102_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0102_r16", err)
	}
	ie.profile0x0102_r16 = tmp_bool_profile0x0102_r16
	var tmp_bool_profile0x0103_r16 bool
	if tmp_bool_profile0x0103_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0103_r16", err)
	}
	ie.profile0x0103_r16 = tmp_bool_profile0x0103_r16
	var tmp_bool_profile0x0104_r16 bool
	if tmp_bool_profile0x0104_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean profile0x0104_r16", err)
	}
	ie.profile0x0104_r16 = tmp_bool_profile0x0104_r16
	return nil
}
