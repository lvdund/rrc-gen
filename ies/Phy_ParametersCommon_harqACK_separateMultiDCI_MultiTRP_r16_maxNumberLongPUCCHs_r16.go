package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_longAndLong   uper.Enumerated = 0
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_longAndShort  uper.Enumerated = 1
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_shortAndShort uper.Enumerated = 2
)

type Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
