package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFR2_2_r17_drx_Adaptation_r17 struct {
	non_SharedSpectrumChAccess_r17 *MinTimeGapFR2_2_r17 `optional`
	sharedSpectrumChAccess_r17     *MinTimeGapFR2_2_r17 `optional`
}

func (ie *MAC_ParametersFR2_2_r17_drx_Adaptation_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.non_SharedSpectrumChAccess_r17 != nil, ie.sharedSpectrumChAccess_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.non_SharedSpectrumChAccess_r17 != nil {
		if err = ie.non_SharedSpectrumChAccess_r17.Encode(w); err != nil {
			return utils.WrapError("Encode non_SharedSpectrumChAccess_r17", err)
		}
	}
	if ie.sharedSpectrumChAccess_r17 != nil {
		if err = ie.sharedSpectrumChAccess_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sharedSpectrumChAccess_r17", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFR2_2_r17_drx_Adaptation_r17) Decode(r *uper.UperReader) error {
	var err error
	var non_SharedSpectrumChAccess_r17Present bool
	if non_SharedSpectrumChAccess_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sharedSpectrumChAccess_r17Present bool
	if sharedSpectrumChAccess_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if non_SharedSpectrumChAccess_r17Present {
		ie.non_SharedSpectrumChAccess_r17 = new(MinTimeGapFR2_2_r17)
		if err = ie.non_SharedSpectrumChAccess_r17.Decode(r); err != nil {
			return utils.WrapError("Decode non_SharedSpectrumChAccess_r17", err)
		}
	}
	if sharedSpectrumChAccess_r17Present {
		ie.sharedSpectrumChAccess_r17 = new(MinTimeGapFR2_2_r17)
		if err = ie.sharedSpectrumChAccess_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sharedSpectrumChAccess_r17", err)
		}
	}
	return nil
}
