package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC struct {
	ca_ParametersNR_ForDC       *CA_ParametersNR         `optional`
	ca_ParametersNR_ForDC_v1540 *CA_ParametersNR_v1540   `optional`
	ca_ParametersNR_ForDC_v1550 *CA_ParametersNR_v1550   `optional`
	ca_ParametersNR_ForDC_v1560 *CA_ParametersNR_v1560   `optional`
	featureSetCombinationDC     *FeatureSetCombinationId `optional`
}

func (ie *CA_ParametersNRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_ForDC != nil, ie.ca_ParametersNR_ForDC_v1540 != nil, ie.ca_ParametersNR_ForDC_v1550 != nil, ie.ca_ParametersNR_ForDC_v1560 != nil, ie.featureSetCombinationDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_ForDC != nil {
		if err = ie.ca_ParametersNR_ForDC.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC", err)
		}
	}
	if ie.ca_ParametersNR_ForDC_v1540 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1540", err)
		}
	}
	if ie.ca_ParametersNR_ForDC_v1550 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1550.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1550", err)
		}
	}
	if ie.ca_ParametersNR_ForDC_v1560 != nil {
		if err = ie.ca_ParametersNR_ForDC_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_ForDC_v1560", err)
		}
	}
	if ie.featureSetCombinationDC != nil {
		if err = ie.featureSetCombinationDC.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetCombinationDC", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_ForDCPresent bool
	if ca_ParametersNR_ForDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_ForDC_v1540Present bool
	if ca_ParametersNR_ForDC_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_ForDC_v1550Present bool
	if ca_ParametersNR_ForDC_v1550Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_ForDC_v1560Present bool
	if ca_ParametersNR_ForDC_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetCombinationDCPresent bool
	if featureSetCombinationDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_ForDCPresent {
		ie.ca_ParametersNR_ForDC = new(CA_ParametersNR)
		if err = ie.ca_ParametersNR_ForDC.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC", err)
		}
	}
	if ca_ParametersNR_ForDC_v1540Present {
		ie.ca_ParametersNR_ForDC_v1540 = new(CA_ParametersNR_v1540)
		if err = ie.ca_ParametersNR_ForDC_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1540", err)
		}
	}
	if ca_ParametersNR_ForDC_v1550Present {
		ie.ca_ParametersNR_ForDC_v1550 = new(CA_ParametersNR_v1550)
		if err = ie.ca_ParametersNR_ForDC_v1550.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1550", err)
		}
	}
	if ca_ParametersNR_ForDC_v1560Present {
		ie.ca_ParametersNR_ForDC_v1560 = new(CA_ParametersNR_v1560)
		if err = ie.ca_ParametersNR_ForDC_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_ForDC_v1560", err)
		}
	}
	if featureSetCombinationDCPresent {
		ie.featureSetCombinationDC = new(FeatureSetCombinationId)
		if err = ie.featureSetCombinationDC.Decode(r); err != nil {
			return utils.WrapError("Decode featureSetCombinationDC", err)
		}
	}
	return nil
}
