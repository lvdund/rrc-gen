package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1730 struct {
	ca_ParametersNR_ForDC_v1730 *CA_ParametersNR_v1730 `optional`
}

func (ie *CA_ParametersNRDC_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_ForDC_v1730 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_ForDC_v1730 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1730", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1730) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_ForDC_v1730Present bool
	if ca_ParametersNR_ForDC_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_ForDC_v1730Present {
		ie.ca_ParametersNR_ForDC_v1730 = new(CA_ParametersNR_v1730)
		if err = ie.ca_ParametersNR_ForDC_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1730", err)
		}
	}
	return nil
}
