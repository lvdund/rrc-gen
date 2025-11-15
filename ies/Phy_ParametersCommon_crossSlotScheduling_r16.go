package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_crossSlotScheduling_r16 struct {
	non_SharedSpectrumChAccess_r16 *Phy_ParametersCommon_crossSlotScheduling_r16_non_SharedSpectrumChAccess_r16 `optional`
	sharedSpectrumChAccess_r16     *Phy_ParametersCommon_crossSlotScheduling_r16_sharedSpectrumChAccess_r16     `optional`
}

func (ie *Phy_ParametersCommon_crossSlotScheduling_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.non_SharedSpectrumChAccess_r16 != nil, ie.sharedSpectrumChAccess_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.non_SharedSpectrumChAccess_r16 != nil {
		if err = ie.non_SharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode non_SharedSpectrumChAccess_r16", err)
		}
	}
	if ie.sharedSpectrumChAccess_r16 != nil {
		if err = ie.sharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sharedSpectrumChAccess_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_crossSlotScheduling_r16) Decode(r *uper.UperReader) error {
	var err error
	var non_SharedSpectrumChAccess_r16Present bool
	if non_SharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sharedSpectrumChAccess_r16Present bool
	if sharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if non_SharedSpectrumChAccess_r16Present {
		ie.non_SharedSpectrumChAccess_r16 = new(Phy_ParametersCommon_crossSlotScheduling_r16_non_SharedSpectrumChAccess_r16)
		if err = ie.non_SharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode non_SharedSpectrumChAccess_r16", err)
		}
	}
	if sharedSpectrumChAccess_r16Present {
		ie.sharedSpectrumChAccess_r16 = new(Phy_ParametersCommon_crossSlotScheduling_r16_sharedSpectrumChAccess_r16)
		if err = ie.sharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sharedSpectrumChAccess_r16", err)
		}
	}
	return nil
}
