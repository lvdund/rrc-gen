package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1590 struct {
	interBandContiguousMRDC *MRDC_Parameters_v1590_interBandContiguousMRDC `optional`
}

func (ie *MRDC_Parameters_v1590) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.interBandContiguousMRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interBandContiguousMRDC != nil {
		if err = ie.interBandContiguousMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode interBandContiguousMRDC", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1590) Decode(r *uper.UperReader) error {
	var err error
	var interBandContiguousMRDCPresent bool
	if interBandContiguousMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if interBandContiguousMRDCPresent {
		ie.interBandContiguousMRDC = new(MRDC_Parameters_v1590_interBandContiguousMRDC)
		if err = ie.interBandContiguousMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode interBandContiguousMRDC", err)
		}
	}
	return nil
}
