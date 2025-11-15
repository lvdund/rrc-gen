package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_srs_PortReport_r17 struct {
	capVal1_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal1_r17 `optional`
	capVal2_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal2_r17 `optional`
	capVal3_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal3_r17 `optional`
	capVal4_r17 *MIMO_ParametersPerBand_srs_PortReport_r17_capVal4_r17 `optional`
}

func (ie *MIMO_ParametersPerBand_srs_PortReport_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.capVal1_r17 != nil, ie.capVal2_r17 != nil, ie.capVal3_r17 != nil, ie.capVal4_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.capVal1_r17 != nil {
		if err = ie.capVal1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode capVal1_r17", err)
		}
	}
	if ie.capVal2_r17 != nil {
		if err = ie.capVal2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode capVal2_r17", err)
		}
	}
	if ie.capVal3_r17 != nil {
		if err = ie.capVal3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode capVal3_r17", err)
		}
	}
	if ie.capVal4_r17 != nil {
		if err = ie.capVal4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode capVal4_r17", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_srs_PortReport_r17) Decode(r *uper.UperReader) error {
	var err error
	var capVal1_r17Present bool
	if capVal1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var capVal2_r17Present bool
	if capVal2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var capVal3_r17Present bool
	if capVal3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var capVal4_r17Present bool
	if capVal4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if capVal1_r17Present {
		ie.capVal1_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal1_r17)
		if err = ie.capVal1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode capVal1_r17", err)
		}
	}
	if capVal2_r17Present {
		ie.capVal2_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal2_r17)
		if err = ie.capVal2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode capVal2_r17", err)
		}
	}
	if capVal3_r17Present {
		ie.capVal3_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal3_r17)
		if err = ie.capVal3_r17.Decode(r); err != nil {
			return utils.WrapError("Decode capVal3_r17", err)
		}
	}
	if capVal4_r17Present {
		ie.capVal4_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17_capVal4_r17)
		if err = ie.capVal4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode capVal4_r17", err)
		}
	}
	return nil
}
