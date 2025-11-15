package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1730 struct {
	measAndMobParametersMRDC_Common_v1730 *MeasAndMobParametersMRDC_Common_v1730 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_Common_v1730 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_Common_v1730 != nil {
		if err = ie.measAndMobParametersMRDC_Common_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_Common_v1730", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1730) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_Common_v1730Present bool
	if measAndMobParametersMRDC_Common_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_Common_v1730Present {
		ie.measAndMobParametersMRDC_Common_v1730 = new(MeasAndMobParametersMRDC_Common_v1730)
		if err = ie.measAndMobParametersMRDC_Common_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_Common_v1730", err)
		}
	}
	return nil
}
