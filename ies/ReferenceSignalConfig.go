package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceSignalConfig struct {
	ssb_ConfigMobility            *SSB_ConfigMobility            `optional`
	csi_rs_ResourceConfigMobility *CSI_RS_ResourceConfigMobility `optional,setuprelease`
}

func (ie *ReferenceSignalConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_ConfigMobility != nil, ie.csi_rs_ResourceConfigMobility != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssb_ConfigMobility != nil {
		if err = ie.ssb_ConfigMobility.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ConfigMobility", err)
		}
	}
	if ie.csi_rs_ResourceConfigMobility != nil {
		tmp_csi_rs_ResourceConfigMobility := utils.SetupRelease[*CSI_RS_ResourceConfigMobility]{
			Setup: ie.csi_rs_ResourceConfigMobility,
		}
		if err = tmp_csi_rs_ResourceConfigMobility.Encode(w); err != nil {
			return utils.WrapError("Encode csi_rs_ResourceConfigMobility", err)
		}
	}
	return nil
}

func (ie *ReferenceSignalConfig) Decode(r *uper.UperReader) error {
	var err error
	var ssb_ConfigMobilityPresent bool
	if ssb_ConfigMobilityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_rs_ResourceConfigMobilityPresent bool
	if csi_rs_ResourceConfigMobilityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ssb_ConfigMobilityPresent {
		ie.ssb_ConfigMobility = new(SSB_ConfigMobility)
		if err = ie.ssb_ConfigMobility.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_ConfigMobility", err)
		}
	}
	if csi_rs_ResourceConfigMobilityPresent {
		tmp_csi_rs_ResourceConfigMobility := utils.SetupRelease[*CSI_RS_ResourceConfigMobility]{}
		if err = tmp_csi_rs_ResourceConfigMobility.Decode(r); err != nil {
			return utils.WrapError("Decode csi_rs_ResourceConfigMobility", err)
		}
		ie.csi_rs_ResourceConfigMobility = tmp_csi_rs_ResourceConfigMobility.Setup
	}
	return nil
}
