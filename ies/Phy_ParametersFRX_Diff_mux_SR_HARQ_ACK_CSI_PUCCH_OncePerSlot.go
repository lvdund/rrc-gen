package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot struct {
	sameSymbol *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_sameSymbol `optional`
	diffSymbol *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_diffSymbol `optional`
}

func (ie *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sameSymbol != nil, ie.diffSymbol != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sameSymbol != nil {
		if err = ie.sameSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode sameSymbol", err)
		}
	}
	if ie.diffSymbol != nil {
		if err = ie.diffSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode diffSymbol", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot) Decode(r *uper.UperReader) error {
	var err error
	var sameSymbolPresent bool
	if sameSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var diffSymbolPresent bool
	if diffSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sameSymbolPresent {
		ie.sameSymbol = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_sameSymbol)
		if err = ie.sameSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode sameSymbol", err)
		}
	}
	if diffSymbolPresent {
		ie.diffSymbol = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_diffSymbol)
		if err = ie.diffSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode diffSymbol", err)
		}
	}
	return nil
}
