package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1700 struct {
	f1c_OverNR_RRC_r17             *NRDC_Parameters_v1700_f1c_OverNR_RRC_r17 `optional`
	measAndMobParametersNRDC_v1700 MeasAndMobParametersMRDC_v1700            `madatory`
}

func (ie *NRDC_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.f1c_OverNR_RRC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.f1c_OverNR_RRC_r17 != nil {
		if err = ie.f1c_OverNR_RRC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode f1c_OverNR_RRC_r17", err)
		}
	}
	if err = ie.measAndMobParametersNRDC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode measAndMobParametersNRDC_v1700", err)
	}
	return nil
}

func (ie *NRDC_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var f1c_OverNR_RRC_r17Present bool
	if f1c_OverNR_RRC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if f1c_OverNR_RRC_r17Present {
		ie.f1c_OverNR_RRC_r17 = new(NRDC_Parameters_v1700_f1c_OverNR_RRC_r17)
		if err = ie.f1c_OverNR_RRC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode f1c_OverNR_RRC_r17", err)
		}
	}
	if err = ie.measAndMobParametersNRDC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode measAndMobParametersNRDC_v1700", err)
	}
	return nil
}
