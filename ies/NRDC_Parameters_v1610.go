package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v1610 struct {
	measAndMobParametersNRDC_v1610 *MeasAndMobParametersMRDC_v1610 `optional`
}

func (ie *NRDC_Parameters_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersNRDC_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersNRDC_v1610 != nil {
		if err = ie.measAndMobParametersNRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersNRDC_v1610", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters_v1610) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersNRDC_v1610Present bool
	if measAndMobParametersNRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersNRDC_v1610Present {
		ie.measAndMobParametersNRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
		if err = ie.measAndMobParametersNRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersNRDC_v1610", err)
		}
	}
	return nil
}
