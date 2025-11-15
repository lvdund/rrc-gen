package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1630 struct {
	ca_ParametersNR_ForDC_v1610 *CA_ParametersNR_v1610 `optional`
	ca_ParametersNR_ForDC_v1630 *CA_ParametersNR_v1630 `optional`
}

func (ie *CA_ParametersNRDC_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_ForDC_v1610 != nil, ie.ca_ParametersNR_ForDC_v1630 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_ForDC_v1610 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1610", err)
		}
	}
	if ie.ca_ParametersNR_ForDC_v1630 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1630", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1630) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_ForDC_v1610Present bool
	if ca_ParametersNR_ForDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_ForDC_v1630Present bool
	if ca_ParametersNR_ForDC_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_ForDC_v1610Present {
		ie.ca_ParametersNR_ForDC_v1610 = new(CA_ParametersNR_v1610)
		if err = ie.ca_ParametersNR_ForDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1610", err)
		}
	}
	if ca_ParametersNR_ForDC_v1630Present {
		ie.ca_ParametersNR_ForDC_v1630 = new(CA_ParametersNR_v1630)
		if err = ie.ca_ParametersNR_ForDC_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1630", err)
		}
	}
	return nil
}
