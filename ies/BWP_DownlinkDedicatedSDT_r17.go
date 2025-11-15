package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkDedicatedSDT_r17 struct {
	pdcch_Config_r17 *PDCCH_Config `optional,setuprelease`
	pdsch_Config_r17 *PDSCH_Config `optional,setuprelease`
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_Config_r17 != nil, ie.pdsch_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_Config_r17 != nil {
		tmp_pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{
			Setup: ie.pdcch_Config_r17,
		}
		if err = tmp_pdcch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_Config_r17", err)
		}
	}
	if ie.pdsch_Config_r17 != nil {
		tmp_pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{
			Setup: ie.pdsch_Config_r17,
		}
		if err = tmp_pdsch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_Config_r17", err)
		}
	}
	return nil
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_Config_r17Present bool
	if pdcch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_Config_r17Present bool
	if pdsch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_Config_r17Present {
		tmp_pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{}
		if err = tmp_pdcch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_Config_r17", err)
		}
		ie.pdcch_Config_r17 = tmp_pdcch_Config_r17.Setup
	}
	if pdsch_Config_r17Present {
		tmp_pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{}
		if err = tmp_pdsch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_Config_r17", err)
		}
		ie.pdsch_Config_r17 = tmp_pdsch_Config_r17.Setup
	}
	return nil
}
