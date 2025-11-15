package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v15g0 struct {
	ca_ParametersNR_v15g0   *CA_ParametersNR_v15g0   `optional`
	ca_ParametersNRDC_v15g0 *CA_ParametersNRDC_v15g0 `optional`
	mrdc_Parameters_v15g0   *MRDC_Parameters_v15g0   `optional`
}

func (ie *BandCombination_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v15g0 != nil, ie.ca_ParametersNRDC_v15g0 != nil, ie.mrdc_Parameters_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v15g0 != nil {
		if err = ie.ca_ParametersNR_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v15g0", err)
		}
	}
	if ie.ca_ParametersNRDC_v15g0 != nil {
		if err = ie.ca_ParametersNRDC_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v15g0", err)
		}
	}
	if ie.mrdc_Parameters_v15g0 != nil {
		if err = ie.mrdc_Parameters_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Parameters_v15g0", err)
		}
	}
	return nil
}

func (ie *BandCombination_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v15g0Present bool
	if ca_ParametersNR_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v15g0Present bool
	if ca_ParametersNRDC_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_Parameters_v15g0Present bool
	if mrdc_Parameters_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v15g0Present {
		ie.ca_ParametersNR_v15g0 = new(CA_ParametersNR_v15g0)
		if err = ie.ca_ParametersNR_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v15g0", err)
		}
	}
	if ca_ParametersNRDC_v15g0Present {
		ie.ca_ParametersNRDC_v15g0 = new(CA_ParametersNRDC_v15g0)
		if err = ie.ca_ParametersNRDC_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v15g0", err)
		}
	}
	if mrdc_Parameters_v15g0Present {
		ie.mrdc_Parameters_v15g0 = new(MRDC_Parameters_v15g0)
		if err = ie.mrdc_Parameters_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Parameters_v15g0", err)
		}
	}
	return nil
}
