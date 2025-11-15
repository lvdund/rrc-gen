package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1640 struct {
	ca_ParametersNR_v1640   *CA_ParametersNR_v1640   `optional`
	ca_ParametersNRDC_v1640 *CA_ParametersNRDC_v1640 `optional`
}

func (ie *BandCombination_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1640 != nil, ie.ca_ParametersNRDC_v1640 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v1640 != nil {
		if err = ie.ca_ParametersNR_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1640", err)
		}
	}
	if ie.ca_ParametersNRDC_v1640 != nil {
		if err = ie.ca_ParametersNRDC_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1640", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1640) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1640Present bool
	if ca_ParametersNR_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v1640Present bool
	if ca_ParametersNRDC_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v1640Present {
		ie.ca_ParametersNR_v1640 = new(CA_ParametersNR_v1640)
		if err = ie.ca_ParametersNR_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1640", err)
		}
	}
	if ca_ParametersNRDC_v1640Present {
		ie.ca_ParametersNRDC_v1640 = new(CA_ParametersNRDC_v1640)
		if err = ie.ca_ParametersNRDC_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1640", err)
		}
	}
	return nil
}
