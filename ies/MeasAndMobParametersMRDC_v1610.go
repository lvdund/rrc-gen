package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1610 struct {
	measAndMobParametersMRDC_Common_v1610 *MeasAndMobParametersMRDC_Common_v1610                    `optional`
	interNR_MeasEUTRA_IAB_r16             *MeasAndMobParametersMRDC_v1610_interNR_MeasEUTRA_IAB_r16 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_Common_v1610 != nil, ie.interNR_MeasEUTRA_IAB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_Common_v1610 != nil {
		if err = ie.measAndMobParametersMRDC_Common_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_Common_v1610", err)
		}
	}
	if ie.interNR_MeasEUTRA_IAB_r16 != nil {
		if err = ie.interNR_MeasEUTRA_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interNR_MeasEUTRA_IAB_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_Common_v1610Present bool
	if measAndMobParametersMRDC_Common_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interNR_MeasEUTRA_IAB_r16Present bool
	if interNR_MeasEUTRA_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_Common_v1610Present {
		ie.measAndMobParametersMRDC_Common_v1610 = new(MeasAndMobParametersMRDC_Common_v1610)
		if err = ie.measAndMobParametersMRDC_Common_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_Common_v1610", err)
		}
	}
	if interNR_MeasEUTRA_IAB_r16Present {
		ie.interNR_MeasEUTRA_IAB_r16 = new(MeasAndMobParametersMRDC_v1610_interNR_MeasEUTRA_IAB_r16)
		if err = ie.interNR_MeasEUTRA_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interNR_MeasEUTRA_IAB_r16", err)
		}
	}
	return nil
}
