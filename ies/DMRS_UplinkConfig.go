package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig struct {
	dmrs_Type                  *DMRS_UplinkConfig_dmrs_Type                  `optional`
	dmrs_AdditionalPosition    *DMRS_UplinkConfig_dmrs_AdditionalPosition    `optional`
	phaseTrackingRS            *PTRS_UplinkConfig                            `optional,setuprelease`
	maxLength                  *DMRS_UplinkConfig_maxLength                  `optional`
	transformPrecodingDisabled *DMRS_UplinkConfig_transformPrecodingDisabled `lb:0,ub:65535,optional`
	transformPrecodingEnabled  *DMRS_UplinkConfig_transformPrecodingEnabled  `lb:0,ub:1007,optional,ext`
}

func (ie *DMRS_UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dmrs_Type != nil, ie.dmrs_AdditionalPosition != nil, ie.phaseTrackingRS != nil, ie.maxLength != nil, ie.transformPrecodingDisabled != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dmrs_Type != nil {
		if err = ie.dmrs_Type.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_Type", err)
		}
	}
	if ie.dmrs_AdditionalPosition != nil {
		if err = ie.dmrs_AdditionalPosition.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_AdditionalPosition", err)
		}
	}
	if ie.phaseTrackingRS != nil {
		tmp_phaseTrackingRS := utils.SetupRelease[*PTRS_UplinkConfig]{
			Setup: ie.phaseTrackingRS,
		}
		if err = tmp_phaseTrackingRS.Encode(w); err != nil {
			return utils.WrapError("Encode phaseTrackingRS", err)
		}
	}
	if ie.maxLength != nil {
		if err = ie.maxLength.Encode(w); err != nil {
			return utils.WrapError("Encode maxLength", err)
		}
	}
	if ie.transformPrecodingDisabled != nil {
		if err = ie.transformPrecodingDisabled.Encode(w); err != nil {
			return utils.WrapError("Encode transformPrecodingDisabled", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var dmrs_TypePresent bool
	if dmrs_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_AdditionalPositionPresent bool
	if dmrs_AdditionalPositionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phaseTrackingRSPresent bool
	if phaseTrackingRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxLengthPresent bool
	if maxLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var transformPrecodingDisabledPresent bool
	if transformPrecodingDisabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dmrs_TypePresent {
		ie.dmrs_Type = new(DMRS_UplinkConfig_dmrs_Type)
		if err = ie.dmrs_Type.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_Type", err)
		}
	}
	if dmrs_AdditionalPositionPresent {
		ie.dmrs_AdditionalPosition = new(DMRS_UplinkConfig_dmrs_AdditionalPosition)
		if err = ie.dmrs_AdditionalPosition.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_AdditionalPosition", err)
		}
	}
	if phaseTrackingRSPresent {
		tmp_phaseTrackingRS := utils.SetupRelease[*PTRS_UplinkConfig]{}
		if err = tmp_phaseTrackingRS.Decode(r); err != nil {
			return utils.WrapError("Decode phaseTrackingRS", err)
		}
		ie.phaseTrackingRS = tmp_phaseTrackingRS.Setup
	}
	if maxLengthPresent {
		ie.maxLength = new(DMRS_UplinkConfig_maxLength)
		if err = ie.maxLength.Decode(r); err != nil {
			return utils.WrapError("Decode maxLength", err)
		}
	}
	if transformPrecodingDisabledPresent {
		ie.transformPrecodingDisabled = new(DMRS_UplinkConfig_transformPrecodingDisabled)
		if err = ie.transformPrecodingDisabled.Decode(r); err != nil {
			return utils.WrapError("Decode transformPrecodingDisabled", err)
		}
	}
	return nil
}
