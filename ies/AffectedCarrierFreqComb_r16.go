package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqComb_r16 struct {
	affectedCarrierFreqComb_r16 []ARFCN_ValueNR      `lb:2,ub:maxNrofServingCells,optional`
	victimSystemType_r16        VictimSystemType_r16 `madatory`
}

func (ie *AffectedCarrierFreqComb_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.affectedCarrierFreqComb_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.affectedCarrierFreqComb_r16) > 0 {
		tmp_affectedCarrierFreqComb_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 2, Ub: maxNrofServingCells}, false)
		for _, i := range ie.affectedCarrierFreqComb_r16 {
			tmp_affectedCarrierFreqComb_r16.Value = append(tmp_affectedCarrierFreqComb_r16.Value, &i)
		}
		if err = tmp_affectedCarrierFreqComb_r16.Encode(w); err != nil {
			return utils.WrapError("Encode affectedCarrierFreqComb_r16", err)
		}
	}
	if err = ie.victimSystemType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode victimSystemType_r16", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqComb_r16) Decode(r *uper.UperReader) error {
	var err error
	var affectedCarrierFreqComb_r16Present bool
	if affectedCarrierFreqComb_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if affectedCarrierFreqComb_r16Present {
		tmp_affectedCarrierFreqComb_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 2, Ub: maxNrofServingCells}, false)
		fn_affectedCarrierFreqComb_r16 := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_affectedCarrierFreqComb_r16.Decode(r, fn_affectedCarrierFreqComb_r16); err != nil {
			return utils.WrapError("Decode affectedCarrierFreqComb_r16", err)
		}
		ie.affectedCarrierFreqComb_r16 = []ARFCN_ValueNR{}
		for _, i := range tmp_affectedCarrierFreqComb_r16.Value {
			ie.affectedCarrierFreqComb_r16 = append(ie.affectedCarrierFreqComb_r16, *i)
		}
	}
	if err = ie.victimSystemType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode victimSystemType_r16", err)
	}
	return nil
}
