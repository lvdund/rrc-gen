package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_ResourceConfig_r16 struct {
	srs_ResourceConfig_r16  *SRS_ResourceListConfigCLI_r16  `optional,setuprelease`
	rssi_ResourceConfig_r16 *RSSI_ResourceListConfigCLI_r16 `optional,setuprelease`
}

func (ie *CLI_ResourceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_ResourceConfig_r16 != nil, ie.rssi_ResourceConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_ResourceConfig_r16 != nil {
		tmp_srs_ResourceConfig_r16 := utils.SetupRelease[*SRS_ResourceListConfigCLI_r16]{
			Setup: ie.srs_ResourceConfig_r16,
		}
		if err = tmp_srs_ResourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ResourceConfig_r16", err)
		}
	}
	if ie.rssi_ResourceConfig_r16 != nil {
		tmp_rssi_ResourceConfig_r16 := utils.SetupRelease[*RSSI_ResourceListConfigCLI_r16]{
			Setup: ie.rssi_ResourceConfig_r16,
		}
		if err = tmp_rssi_ResourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rssi_ResourceConfig_r16", err)
		}
	}
	return nil
}

func (ie *CLI_ResourceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var srs_ResourceConfig_r16Present bool
	if srs_ResourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rssi_ResourceConfig_r16Present bool
	if rssi_ResourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_ResourceConfig_r16Present {
		tmp_srs_ResourceConfig_r16 := utils.SetupRelease[*SRS_ResourceListConfigCLI_r16]{}
		if err = tmp_srs_ResourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode srs_ResourceConfig_r16", err)
		}
		ie.srs_ResourceConfig_r16 = tmp_srs_ResourceConfig_r16.Setup
	}
	if rssi_ResourceConfig_r16Present {
		tmp_rssi_ResourceConfig_r16 := utils.SetupRelease[*RSSI_ResourceListConfigCLI_r16]{}
		if err = tmp_rssi_ResourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rssi_ResourceConfig_r16", err)
		}
		ie.rssi_ResourceConfig_r16 = tmp_rssi_ResourceConfig_r16.Setup
	}
	return nil
}
