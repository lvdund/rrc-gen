package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TwoPUCCH_Grp_ConfigParams_r16 struct {
	pucch_GroupMapping_r16 PUCCH_Grp_CarrierTypes_r16 `madatory`
	pucch_TX_r16           PUCCH_Grp_CarrierTypes_r16 `madatory`
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_GroupMapping_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_GroupMapping_r16", err)
	}
	if err = ie.pucch_TX_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_TX_r16", err)
	}
	return nil
}

func (ie *TwoPUCCH_Grp_ConfigParams_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_GroupMapping_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_GroupMapping_r16", err)
	}
	if err = ie.pucch_TX_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_TX_r16", err)
	}
	return nil
}
