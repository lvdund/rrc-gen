package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16 struct {
	supportNewDMRS_Port_r16   *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportNewDMRS_Port_r16   `optional`
	supportTwoPortDL_PTRS_r16 *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportTwoPortDL_PTRS_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportNewDMRS_Port_r16 != nil, ie.supportTwoPortDL_PTRS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportNewDMRS_Port_r16 != nil {
		if err = ie.supportNewDMRS_Port_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportNewDMRS_Port_r16", err)
		}
	}
	if ie.supportTwoPortDL_PTRS_r16 != nil {
		if err = ie.supportTwoPortDL_PTRS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportTwoPortDL_PTRS_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var supportNewDMRS_Port_r16Present bool
	if supportNewDMRS_Port_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportTwoPortDL_PTRS_r16Present bool
	if supportTwoPortDL_PTRS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportNewDMRS_Port_r16Present {
		ie.supportNewDMRS_Port_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportNewDMRS_Port_r16)
		if err = ie.supportNewDMRS_Port_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportNewDMRS_Port_r16", err)
		}
	}
	if supportTwoPortDL_PTRS_r16Present {
		ie.supportTwoPortDL_PTRS_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportTwoPortDL_PTRS_r16)
		if err = ie.supportTwoPortDL_PTRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportTwoPortDL_PTRS_r16", err)
		}
	}
	return nil
}
