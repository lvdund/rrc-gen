package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n2   uper.Enumerated = 0
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n4   uper.Enumerated = 1
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n8   uper.Enumerated = 2
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n12  uper.Enumerated = 3
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n16  uper.Enumerated = 4
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n32  uper.Enumerated = 5
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n40  uper.Enumerated = 6
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n48  uper.Enumerated = 7
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n64  uper.Enumerated = 8
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n72  uper.Enumerated = 9
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n80  uper.Enumerated = 10
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n96  uper.Enumerated = 11
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n128 uper.Enumerated = 12
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16_Enum_n256 uper.Enumerated = 13
)

type Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16 struct {
	Value uper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResAcrossCC_OneFR_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
