package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1700 struct {
	supportedMinBandwidthUL_r17    *SupportedBandwidth_v1700                                   `optional`
	mTRP_PUSCH_RepetitionTypeB_r17 *FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_RepetitionTypeB_r17 `optional`
	mTRP_PUSCH_TypeB_CB_r17        *FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_TypeB_CB_r17        `optional`
	supportedBandwidthUL_v1710     *SupportedBandwidth_v1700                                   `optional`
}

func (ie *FeatureSetUplinkPerCC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedMinBandwidthUL_r17 != nil, ie.mTRP_PUSCH_RepetitionTypeB_r17 != nil, ie.mTRP_PUSCH_TypeB_CB_r17 != nil, ie.supportedBandwidthUL_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedMinBandwidthUL_r17 != nil {
		if err = ie.supportedMinBandwidthUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportedMinBandwidthUL_r17", err)
		}
	}
	if ie.mTRP_PUSCH_RepetitionTypeB_r17 != nil {
		if err = ie.mTRP_PUSCH_RepetitionTypeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PUSCH_RepetitionTypeB_r17", err)
		}
	}
	if ie.mTRP_PUSCH_TypeB_CB_r17 != nil {
		if err = ie.mTRP_PUSCH_TypeB_CB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PUSCH_TypeB_CB_r17", err)
		}
	}
	if ie.supportedBandwidthUL_v1710 != nil {
		if err = ie.supportedBandwidthUL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandwidthUL_v1710", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var supportedMinBandwidthUL_r17Present bool
	if supportedMinBandwidthUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PUSCH_RepetitionTypeB_r17Present bool
	if mTRP_PUSCH_RepetitionTypeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PUSCH_TypeB_CB_r17Present bool
	if mTRP_PUSCH_TypeB_CB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandwidthUL_v1710Present bool
	if supportedBandwidthUL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedMinBandwidthUL_r17Present {
		ie.supportedMinBandwidthUL_r17 = new(SupportedBandwidth_v1700)
		if err = ie.supportedMinBandwidthUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportedMinBandwidthUL_r17", err)
		}
	}
	if mTRP_PUSCH_RepetitionTypeB_r17Present {
		ie.mTRP_PUSCH_RepetitionTypeB_r17 = new(FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_RepetitionTypeB_r17)
		if err = ie.mTRP_PUSCH_RepetitionTypeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PUSCH_RepetitionTypeB_r17", err)
		}
	}
	if mTRP_PUSCH_TypeB_CB_r17Present {
		ie.mTRP_PUSCH_TypeB_CB_r17 = new(FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_TypeB_CB_r17)
		if err = ie.mTRP_PUSCH_TypeB_CB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PUSCH_TypeB_CB_r17", err)
		}
	}
	if supportedBandwidthUL_v1710Present {
		ie.supportedBandwidthUL_v1710 = new(SupportedBandwidth_v1700)
		if err = ie.supportedBandwidthUL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandwidthUL_v1710", err)
		}
	}
	return nil
}
