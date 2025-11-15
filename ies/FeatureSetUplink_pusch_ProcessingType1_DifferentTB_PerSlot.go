package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot struct {
	scs_15kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_15kHz  `optional`
	scs_30kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_30kHz  `optional`
	scs_60kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_60kHz  `optional`
	scs_120kHz *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_120kHz `optional`
}

func (ie *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz != nil, ie.scs_30kHz != nil, ie.scs_60kHz != nil, ie.scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz != nil {
		if err = ie.scs_15kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_15kHz", err)
		}
	}
	if ie.scs_30kHz != nil {
		if err = ie.scs_30kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_30kHz", err)
		}
	}
	if ie.scs_60kHz != nil {
		if err = ie.scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_60kHz", err)
		}
	}
	if ie.scs_120kHz != nil {
		if err = ie.scs_120kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_120kHz", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHzPresent bool
	if scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHzPresent bool
	if scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_60kHzPresent bool
	if scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_120kHzPresent bool
	if scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHzPresent {
		ie.scs_15kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_15kHz)
		if err = ie.scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_15kHz", err)
		}
	}
	if scs_30kHzPresent {
		ie.scs_30kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_30kHz)
		if err = ie.scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_30kHz", err)
		}
	}
	if scs_60kHzPresent {
		ie.scs_60kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_60kHz)
		if err = ie.scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
	}
	if scs_120kHzPresent {
		ie.scs_120kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_120kHz)
		if err = ie.scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_120kHz", err)
		}
	}
	return nil
}
