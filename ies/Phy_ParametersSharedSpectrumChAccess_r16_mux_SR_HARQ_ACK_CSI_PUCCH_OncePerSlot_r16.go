package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 struct {
	sameSymbol_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_sameSymbol_r16 `optional`
	diffSymbol_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_diffSymbol_r16 `optional`
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sameSymbol_r16 != nil, ie.diffSymbol_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sameSymbol_r16 != nil {
		if err = ie.sameSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sameSymbol_r16", err)
		}
	}
	if ie.diffSymbol_r16 != nil {
		if err = ie.diffSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode diffSymbol_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16) Decode(r *uper.UperReader) error {
	var err error
	var sameSymbol_r16Present bool
	if sameSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var diffSymbol_r16Present bool
	if diffSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sameSymbol_r16Present {
		ie.sameSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_sameSymbol_r16)
		if err = ie.sameSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sameSymbol_r16", err)
		}
	}
	if diffSymbol_r16Present {
		ie.diffSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_diffSymbol_r16)
		if err = ie.diffSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode diffSymbol_r16", err)
		}
	}
	return nil
}
