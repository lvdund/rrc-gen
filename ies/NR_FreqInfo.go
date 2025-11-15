package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_FreqInfo struct {
	measuredFrequency *ARFCN_ValueNR `optional`
}

func (ie *NR_FreqInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measuredFrequency != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measuredFrequency != nil {
		if err = ie.measuredFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode measuredFrequency", err)
		}
	}
	return nil
}

func (ie *NR_FreqInfo) Decode(r *uper.UperReader) error {
	var err error
	var measuredFrequencyPresent bool
	if measuredFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measuredFrequencyPresent {
		ie.measuredFrequency = new(ARFCN_ValueNR)
		if err = ie.measuredFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode measuredFrequency", err)
		}
	}
	return nil
}
