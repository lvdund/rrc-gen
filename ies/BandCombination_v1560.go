package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1560 struct {
	ne_DC_BC                 *BandCombination_v1560_ne_DC_BC `optional`
	ca_ParametersNRDC        *CA_ParametersNRDC              `optional`
	ca_ParametersEUTRA_v1560 *CA_ParametersEUTRA_v1560       `optional`
	ca_ParametersNR_v1560    *CA_ParametersNR_v1560          `optional`
}

func (ie *BandCombination_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ne_DC_BC != nil, ie.ca_ParametersNRDC != nil, ie.ca_ParametersEUTRA_v1560 != nil, ie.ca_ParametersNR_v1560 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ne_DC_BC != nil {
		if err = ie.ne_DC_BC.Encode(w); err != nil {
			return utils.WrapError("Encode ne_DC_BC", err)
		}
	}
	if ie.ca_ParametersNRDC != nil {
		if err = ie.ca_ParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC", err)
		}
	}
	if ie.ca_ParametersEUTRA_v1560 != nil {
		if err = ie.ca_ParametersEUTRA_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersEUTRA_v1560", err)
		}
	}
	if ie.ca_ParametersNR_v1560 != nil {
		if err = ie.ca_ParametersNR_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1560", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1560) Decode(r *uper.UperReader) error {
	var err error
	var ne_DC_BCPresent bool
	if ne_DC_BCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDCPresent bool
	if ca_ParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersEUTRA_v1560Present bool
	if ca_ParametersEUTRA_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_v1560Present bool
	if ca_ParametersNR_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ne_DC_BCPresent {
		ie.ne_DC_BC = new(BandCombination_v1560_ne_DC_BC)
		if err = ie.ne_DC_BC.Decode(r); err != nil {
			return utils.WrapError("Decode ne_DC_BC", err)
		}
	}
	if ca_ParametersNRDCPresent {
		ie.ca_ParametersNRDC = new(CA_ParametersNRDC)
		if err = ie.ca_ParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC", err)
		}
	}
	if ca_ParametersEUTRA_v1560Present {
		ie.ca_ParametersEUTRA_v1560 = new(CA_ParametersEUTRA_v1560)
		if err = ie.ca_ParametersEUTRA_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersEUTRA_v1560", err)
		}
	}
	if ca_ParametersNR_v1560Present {
		ie.ca_ParametersNR_v1560 = new(CA_ParametersNR_v1560)
		if err = ie.ca_ParametersNR_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1560", err)
		}
	}
	return nil
}
