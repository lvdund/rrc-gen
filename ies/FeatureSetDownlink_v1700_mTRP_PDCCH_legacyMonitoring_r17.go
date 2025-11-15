package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17 struct {
	scs_15kHz_r17 *PDCCH_RepetitionParameters_r17 `optional`
	scs_30kHz_r17 *PDCCH_RepetitionParameters_r17 `optional`
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz_r17 != nil, ie.scs_30kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz_r17 != nil {
		if err = ie.scs_15kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_15kHz_r17", err)
		}
	}
	if ie.scs_30kHz_r17 != nil {
		if err = ie.scs_30kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_30kHz_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHz_r17Present bool
	if scs_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHz_r17Present bool
	if scs_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHz_r17Present {
		ie.scs_15kHz_r17 = new(PDCCH_RepetitionParameters_r17)
		if err = ie.scs_15kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_15kHz_r17", err)
		}
	}
	if scs_30kHz_r17Present {
		ie.scs_30kHz_r17 = new(PDCCH_RepetitionParameters_r17)
		if err = ie.scs_30kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_30kHz_r17", err)
		}
	}
	return nil
}
