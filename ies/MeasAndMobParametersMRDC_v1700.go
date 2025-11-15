package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1700 struct {
	measAndMobParametersMRDC_Common_v1700 *MeasAndMobParametersMRDC_Common_v1700 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_Common_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_Common_v1700 != nil {
		if err = ie.measAndMobParametersMRDC_Common_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_Common_v1700", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_Common_v1700Present bool
	if measAndMobParametersMRDC_Common_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_Common_v1700Present {
		ie.measAndMobParametersMRDC_Common_v1700 = new(MeasAndMobParametersMRDC_Common_v1700)
		if err = ie.measAndMobParametersMRDC_Common_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_Common_v1700", err)
		}
	}
	return nil
}
