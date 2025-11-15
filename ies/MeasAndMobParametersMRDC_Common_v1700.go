package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1700 struct {
	condPSCellChangeParameters_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17 `optional`
	condHandoverWithSCG_ENDC_r17   *MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_ENDC_r17   `optional`
	condHandoverWithSCG_NEDC_r17   *MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_NEDC_r17   `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condPSCellChangeParameters_r17 != nil, ie.condHandoverWithSCG_ENDC_r17 != nil, ie.condHandoverWithSCG_NEDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condPSCellChangeParameters_r17 != nil {
		if err = ie.condPSCellChangeParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellChangeParameters_r17", err)
		}
	}
	if ie.condHandoverWithSCG_ENDC_r17 != nil {
		if err = ie.condHandoverWithSCG_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode condHandoverWithSCG_ENDC_r17", err)
		}
	}
	if ie.condHandoverWithSCG_NEDC_r17 != nil {
		if err = ie.condHandoverWithSCG_NEDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode condHandoverWithSCG_NEDC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Decode(r *uper.UperReader) error {
	var err error
	var condPSCellChangeParameters_r17Present bool
	if condPSCellChangeParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condHandoverWithSCG_ENDC_r17Present bool
	if condHandoverWithSCG_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condHandoverWithSCG_NEDC_r17Present bool
	if condHandoverWithSCG_NEDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if condPSCellChangeParameters_r17Present {
		ie.condPSCellChangeParameters_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17)
		if err = ie.condPSCellChangeParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellChangeParameters_r17", err)
		}
	}
	if condHandoverWithSCG_ENDC_r17Present {
		ie.condHandoverWithSCG_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_ENDC_r17)
		if err = ie.condHandoverWithSCG_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condHandoverWithSCG_ENDC_r17", err)
		}
	}
	if condHandoverWithSCG_NEDC_r17Present {
		ie.condHandoverWithSCG_NEDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_NEDC_r17)
		if err = ie.condHandoverWithSCG_NEDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condHandoverWithSCG_NEDC_r17", err)
		}
	}
	return nil
}
