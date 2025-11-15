package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16 struct {
	maxNumberResWithinSlotAcrossCC_AcrossFR_r16 *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResWithinSlotAcrossCC_AcrossFR_r16 `optional`
	maxNumberResAcrossCC_AcrossFR_r16           *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16           `optional`
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil, ie.maxNumberResAcrossCC_AcrossFR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil {
		if err = ie.maxNumberResWithinSlotAcrossCC_AcrossFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberResWithinSlotAcrossCC_AcrossFR_r16", err)
		}
	}
	if ie.maxNumberResAcrossCC_AcrossFR_r16 != nil {
		if err = ie.maxNumberResAcrossCC_AcrossFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberResAcrossCC_AcrossFR_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberResWithinSlotAcrossCC_AcrossFR_r16Present bool
	if maxNumberResWithinSlotAcrossCC_AcrossFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberResAcrossCC_AcrossFR_r16Present bool
	if maxNumberResAcrossCC_AcrossFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberResWithinSlotAcrossCC_AcrossFR_r16Present {
		ie.maxNumberResWithinSlotAcrossCC_AcrossFR_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResWithinSlotAcrossCC_AcrossFR_r16)
		if err = ie.maxNumberResWithinSlotAcrossCC_AcrossFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberResWithinSlotAcrossCC_AcrossFR_r16", err)
		}
	}
	if maxNumberResAcrossCC_AcrossFR_r16Present {
		ie.maxNumberResAcrossCC_AcrossFR_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16)
		if err = ie.maxNumberResAcrossCC_AcrossFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberResAcrossCC_AcrossFR_r16", err)
		}
	}
	return nil
}
