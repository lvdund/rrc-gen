package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1590_IEs struct {
	scellFrequenciesSN_NR    []ARFCN_ValueNR      `lb:1,ub:maxNrofServingCells_1,optional`
	scellFrequenciesSN_EUTRA []ARFCN_ValueEUTRA   `lb:1,ub:maxNrofServingCells_1,optional`
	nonCriticalExtension     *CG_Config_v1610_IEs `optional`
}

func (ie *CG_Config_v1590_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.scellFrequenciesSN_NR) > 0, len(ie.scellFrequenciesSN_EUTRA) > 0, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.scellFrequenciesSN_NR) > 0 {
		tmp_scellFrequenciesSN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.scellFrequenciesSN_NR {
			tmp_scellFrequenciesSN_NR.Value = append(tmp_scellFrequenciesSN_NR.Value, &i)
		}
		if err = tmp_scellFrequenciesSN_NR.Encode(w); err != nil {
			return utils.WrapError("Encode scellFrequenciesSN_NR", err)
		}
	}
	if len(ie.scellFrequenciesSN_EUTRA) > 0 {
		tmp_scellFrequenciesSN_EUTRA := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.scellFrequenciesSN_EUTRA {
			tmp_scellFrequenciesSN_EUTRA.Value = append(tmp_scellFrequenciesSN_EUTRA.Value, &i)
		}
		if err = tmp_scellFrequenciesSN_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode scellFrequenciesSN_EUTRA", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1590_IEs) Decode(r *uper.UperReader) error {
	var err error
	var scellFrequenciesSN_NRPresent bool
	if scellFrequenciesSN_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scellFrequenciesSN_EUTRAPresent bool
	if scellFrequenciesSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scellFrequenciesSN_NRPresent {
		tmp_scellFrequenciesSN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_scellFrequenciesSN_NR := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_scellFrequenciesSN_NR.Decode(r, fn_scellFrequenciesSN_NR); err != nil {
			return utils.WrapError("Decode scellFrequenciesSN_NR", err)
		}
		ie.scellFrequenciesSN_NR = []ARFCN_ValueNR{}
		for _, i := range tmp_scellFrequenciesSN_NR.Value {
			ie.scellFrequenciesSN_NR = append(ie.scellFrequenciesSN_NR, *i)
		}
	}
	if scellFrequenciesSN_EUTRAPresent {
		tmp_scellFrequenciesSN_EUTRA := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_scellFrequenciesSN_EUTRA := func() *ARFCN_ValueEUTRA {
			return new(ARFCN_ValueEUTRA)
		}
		if err = tmp_scellFrequenciesSN_EUTRA.Decode(r, fn_scellFrequenciesSN_EUTRA); err != nil {
			return utils.WrapError("Decode scellFrequenciesSN_EUTRA", err)
		}
		ie.scellFrequenciesSN_EUTRA = []ARFCN_ValueEUTRA{}
		for _, i := range tmp_scellFrequenciesSN_EUTRA.Value {
			ie.scellFrequenciesSN_EUTRA = append(ie.scellFrequenciesSN_EUTRA, *i)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
