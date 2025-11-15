package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigSN struct {
	measuredFrequenciesSN []NR_FreqInfo `lb:1,ub:maxMeasFreqsSN,optional`
}

func (ie *MeasConfigSN) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.measuredFrequenciesSN) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.measuredFrequenciesSN) > 0 {
		tmp_measuredFrequenciesSN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsSN}, false)
		for _, i := range ie.measuredFrequenciesSN {
			tmp_measuredFrequenciesSN.Value = append(tmp_measuredFrequenciesSN.Value, &i)
		}
		if err = tmp_measuredFrequenciesSN.Encode(w); err != nil {
			return utils.WrapError("Encode measuredFrequenciesSN", err)
		}
	}
	return nil
}

func (ie *MeasConfigSN) Decode(r *uper.UperReader) error {
	var err error
	var measuredFrequenciesSNPresent bool
	if measuredFrequenciesSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measuredFrequenciesSNPresent {
		tmp_measuredFrequenciesSN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsSN}, false)
		fn_measuredFrequenciesSN := func() *NR_FreqInfo {
			return new(NR_FreqInfo)
		}
		if err = tmp_measuredFrequenciesSN.Decode(r, fn_measuredFrequenciesSN); err != nil {
			return utils.WrapError("Decode measuredFrequenciesSN", err)
		}
		ie.measuredFrequenciesSN = []NR_FreqInfo{}
		for _, i := range tmp_measuredFrequenciesSN.Value {
			ie.measuredFrequenciesSN = append(ie.measuredFrequenciesSN, *i)
		}
	}
	return nil
}
