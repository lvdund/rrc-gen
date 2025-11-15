package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParameters_v1700 struct {
	measAndMobParametersFR2_2_r17 *MeasAndMobParametersFR2_2_r17 `optional`
}

func (ie *MeasAndMobParameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersFR2_2_r17 != nil {
		if err = ie.measAndMobParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersFR2_2_r17Present bool
	if measAndMobParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersFR2_2_r17Present {
		ie.measAndMobParametersFR2_2_r17 = new(MeasAndMobParametersFR2_2_r17)
		if err = ie.measAndMobParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersFR2_2_r17", err)
		}
	}
	return nil
}
