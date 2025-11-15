package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL struct {
	scs_15kHz  *PTRS_DensityRecommendationDL `optional`
	scs_30kHz  *PTRS_DensityRecommendationDL `optional`
	scs_60kHz  *PTRS_DensityRecommendationDL `optional`
	scs_120kHz *PTRS_DensityRecommendationDL `optional`
}

func (ie *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL) Encode(w *uper.UperWriter) error {
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

func (ie *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL) Decode(r *uper.UperReader) error {
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
		ie.scs_15kHz = new(PTRS_DensityRecommendationDL)
		if err = ie.scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_15kHz", err)
		}
	}
	if scs_30kHzPresent {
		ie.scs_30kHz = new(PTRS_DensityRecommendationDL)
		if err = ie.scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_30kHz", err)
		}
	}
	if scs_60kHzPresent {
		ie.scs_60kHz = new(PTRS_DensityRecommendationDL)
		if err = ie.scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
	}
	if scs_120kHzPresent {
		ie.scs_120kHz = new(PTRS_DensityRecommendationDL)
		if err = ie.scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_120kHz", err)
		}
	}
	return nil
}
