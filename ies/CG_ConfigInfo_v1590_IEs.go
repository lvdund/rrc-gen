package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1590_IEs struct {
	servFrequenciesMN_NR []ARFCN_ValueNR          `lb:1,ub:maxNrofServingCells_1,optional`
	nonCriticalExtension *CG_ConfigInfo_v1610_IEs `optional`
}

func (ie *CG_ConfigInfo_v1590_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.servFrequenciesMN_NR) > 0, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.servFrequenciesMN_NR) > 0 {
		tmp_servFrequenciesMN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.servFrequenciesMN_NR {
			tmp_servFrequenciesMN_NR.Value = append(tmp_servFrequenciesMN_NR.Value, &i)
		}
		if err = tmp_servFrequenciesMN_NR.Encode(w); err != nil {
			return utils.WrapError("Encode servFrequenciesMN_NR", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1590_IEs) Decode(r *uper.UperReader) error {
	var err error
	var servFrequenciesMN_NRPresent bool
	if servFrequenciesMN_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if servFrequenciesMN_NRPresent {
		tmp_servFrequenciesMN_NR := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_servFrequenciesMN_NR := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_servFrequenciesMN_NR.Decode(r, fn_servFrequenciesMN_NR); err != nil {
			return utils.WrapError("Decode servFrequenciesMN_NR", err)
		}
		ie.servFrequenciesMN_NR = []ARFCN_ValueNR{}
		for _, i := range tmp_servFrequenciesMN_NR.Value {
			ie.servFrequenciesMN_NR = append(ie.servFrequenciesMN_NR, *i)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
