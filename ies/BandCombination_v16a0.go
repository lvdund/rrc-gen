package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v16a0 struct {
	ca_ParametersNR_v16a0   *CA_ParametersNR_v16a0   `optional`
	ca_ParametersNRDC_v16a0 *CA_ParametersNRDC_v16a0 `optional`
}

func (ie *BandCombination_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v16a0 != nil, ie.ca_ParametersNRDC_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v16a0 != nil {
		if err = ie.ca_ParametersNR_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v16a0", err)
		}
	}
	if ie.ca_ParametersNRDC_v16a0 != nil {
		if err = ie.ca_ParametersNRDC_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v16a0", err)
		}
	}
	return nil
}

func (ie *BandCombination_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v16a0Present bool
	if ca_ParametersNR_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v16a0Present bool
	if ca_ParametersNRDC_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v16a0Present {
		ie.ca_ParametersNR_v16a0 = new(CA_ParametersNR_v16a0)
		if err = ie.ca_ParametersNR_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v16a0", err)
		}
	}
	if ca_ParametersNRDC_v16a0Present {
		ie.ca_ParametersNRDC_v16a0 = new(CA_ParametersNRDC_v16a0)
		if err = ie.ca_ParametersNRDC_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v16a0", err)
		}
	}
	return nil
}
