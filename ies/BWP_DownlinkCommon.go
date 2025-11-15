package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkCommon struct {
	genericParameters  BWP                 `madatory`
	pdcch_ConfigCommon *PDCCH_ConfigCommon `optional,setuprelease`
	pdsch_ConfigCommon *PDSCH_ConfigCommon `optional,setuprelease`
}

func (ie *BWP_DownlinkCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_ConfigCommon != nil, ie.pdsch_ConfigCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.genericParameters.Encode(w); err != nil {
		return utils.WrapError("Encode genericParameters", err)
	}
	if ie.pdcch_ConfigCommon != nil {
		tmp_pdcch_ConfigCommon := utils.SetupRelease[*PDCCH_ConfigCommon]{
			Setup: ie.pdcch_ConfigCommon,
		}
		if err = tmp_pdcch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_ConfigCommon", err)
		}
	}
	if ie.pdsch_ConfigCommon != nil {
		tmp_pdsch_ConfigCommon := utils.SetupRelease[*PDSCH_ConfigCommon]{
			Setup: ie.pdsch_ConfigCommon,
		}
		if err = tmp_pdsch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ConfigCommon", err)
		}
	}
	return nil
}

func (ie *BWP_DownlinkCommon) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_ConfigCommonPresent bool
	if pdcch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigCommonPresent bool
	if pdsch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.genericParameters.Decode(r); err != nil {
		return utils.WrapError("Decode genericParameters", err)
	}
	if pdcch_ConfigCommonPresent {
		tmp_pdcch_ConfigCommon := utils.SetupRelease[*PDCCH_ConfigCommon]{}
		if err = tmp_pdcch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_ConfigCommon", err)
		}
		ie.pdcch_ConfigCommon = tmp_pdcch_ConfigCommon.Setup
	}
	if pdsch_ConfigCommonPresent {
		tmp_pdsch_ConfigCommon := utils.SetupRelease[*PDSCH_ConfigCommon]{}
		if err = tmp_pdsch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ConfigCommon", err)
		}
		ie.pdsch_ConfigCommon = tmp_pdsch_ConfigCommon.Setup
	}
	return nil
}
