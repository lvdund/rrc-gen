package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16 struct {
	maxNumberResWithinSlotAcrossCC_OneFR_r16 *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16 `optional`
	maxNumberResAcrossCC_OneFR_r16           *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16           `optional`
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberResWithinSlotAcrossCC_OneFR_r16 != nil, ie.maxNumberResAcrossCC_OneFR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberResWithinSlotAcrossCC_OneFR_r16 != nil {
		if err = ie.maxNumberResWithinSlotAcrossCC_OneFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberResWithinSlotAcrossCC_OneFR_r16", err)
		}
	}
	if ie.maxNumberResAcrossCC_OneFR_r16 != nil {
		if err = ie.maxNumberResAcrossCC_OneFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberResAcrossCC_OneFR_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberResWithinSlotAcrossCC_OneFR_r16Present bool
	if maxNumberResWithinSlotAcrossCC_OneFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberResAcrossCC_OneFR_r16Present bool
	if maxNumberResAcrossCC_OneFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberResWithinSlotAcrossCC_OneFR_r16Present {
		ie.maxNumberResWithinSlotAcrossCC_OneFR_r16 = new(Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16)
		if err = ie.maxNumberResWithinSlotAcrossCC_OneFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberResWithinSlotAcrossCC_OneFR_r16", err)
		}
	}
	if maxNumberResAcrossCC_OneFR_r16Present {
		ie.maxNumberResAcrossCC_OneFR_r16 = new(Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16)
		if err = ie.maxNumberResAcrossCC_OneFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberResAcrossCC_OneFR_r16", err)
		}
	}
	return nil
}
