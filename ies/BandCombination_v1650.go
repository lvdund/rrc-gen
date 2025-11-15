package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1650 struct {
	ca_ParametersNRDC_v1650 *CA_ParametersNRDC_v1650 `optional`
}

func (ie *BandCombination_v1650) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNRDC_v1650 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNRDC_v1650 != nil {
		if err = ie.ca_ParametersNRDC_v1650.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1650", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1650) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNRDC_v1650Present bool
	if ca_ParametersNRDC_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNRDC_v1650Present {
		ie.ca_ParametersNRDC_v1650 = new(CA_ParametersNRDC_v1650)
		if err = ie.ca_ParametersNRDC_v1650.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1650", err)
		}
	}
	return nil
}
