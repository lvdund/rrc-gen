package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1540 struct {
	overheatingAssistanceConfig *OverheatingAssistanceConfig `optional,setuprelease`
}

func (ie *OtherConfig_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.overheatingAssistanceConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.overheatingAssistanceConfig != nil {
		tmp_overheatingAssistanceConfig := utils.SetupRelease[*OverheatingAssistanceConfig]{
			Setup: ie.overheatingAssistanceConfig,
		}
		if err = tmp_overheatingAssistanceConfig.Encode(w); err != nil {
			return utils.WrapError("Encode overheatingAssistanceConfig", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1540) Decode(r *uper.UperReader) error {
	var err error
	var overheatingAssistanceConfigPresent bool
	if overheatingAssistanceConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if overheatingAssistanceConfigPresent {
		tmp_overheatingAssistanceConfig := utils.SetupRelease[*OverheatingAssistanceConfig]{}
		if err = tmp_overheatingAssistanceConfig.Decode(r); err != nil {
			return utils.WrapError("Decode overheatingAssistanceConfig", err)
		}
		ie.overheatingAssistanceConfig = tmp_overheatingAssistanceConfig.Setup
	}
	return nil
}
