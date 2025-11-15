package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1560 struct {
	measAndMobParametersMRDC_XDD_Diff_v1560 *MeasAndMobParametersMRDC_XDD_Diff_v1560 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_XDD_Diff_v1560 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_XDD_Diff_v1560 != nil {
		if err = ie.measAndMobParametersMRDC_XDD_Diff_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_XDD_Diff_v1560", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1560) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_XDD_Diff_v1560Present bool
	if measAndMobParametersMRDC_XDD_Diff_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_XDD_Diff_v1560Present {
		ie.measAndMobParametersMRDC_XDD_Diff_v1560 = new(MeasAndMobParametersMRDC_XDD_Diff_v1560)
		if err = ie.measAndMobParametersMRDC_XDD_Diff_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_XDD_Diff_v1560", err)
		}
	}
	return nil
}
