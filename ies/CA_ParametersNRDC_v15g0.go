package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v15g0 struct {
	ca_ParametersNR_ForDC_v15g0 *CA_ParametersNR_v15g0 `optional`
}

func (ie *CA_ParametersNRDC_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_ForDC_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_ForDC_v15g0 != nil {
		if err = ie.ca_ParametersNR_ForDC_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v15g0", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_ForDC_v15g0Present bool
	if ca_ParametersNR_ForDC_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_ForDC_v15g0Present {
		ie.ca_ParametersNR_ForDC_v15g0 = new(CA_ParametersNR_v15g0)
		if err = ie.ca_ParametersNR_ForDC_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v15g0", err)
		}
	}
	return nil
}
