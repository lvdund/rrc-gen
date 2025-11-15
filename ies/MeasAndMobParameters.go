package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParameters struct {
	measAndMobParametersCommon   *MeasAndMobParametersCommon   `optional`
	measAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff `optional`
	measAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff `optional`
}

func (ie *MeasAndMobParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersCommon != nil, ie.measAndMobParametersXDD_Diff != nil, ie.measAndMobParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersCommon != nil {
		if err = ie.measAndMobParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersCommon", err)
		}
	}
	if ie.measAndMobParametersXDD_Diff != nil {
		if err = ie.measAndMobParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersXDD_Diff", err)
		}
	}
	if ie.measAndMobParametersFRX_Diff != nil {
		if err = ie.measAndMobParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParameters) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersCommonPresent bool
	if measAndMobParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersXDD_DiffPresent bool
	if measAndMobParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersFRX_DiffPresent bool
	if measAndMobParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersCommonPresent {
		ie.measAndMobParametersCommon = new(MeasAndMobParametersCommon)
		if err = ie.measAndMobParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersCommon", err)
		}
	}
	if measAndMobParametersXDD_DiffPresent {
		ie.measAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
		if err = ie.measAndMobParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersXDD_Diff", err)
		}
	}
	if measAndMobParametersFRX_DiffPresent {
		ie.measAndMobParametersFRX_Diff = new(MeasAndMobParametersFRX_Diff)
		if err = ie.measAndMobParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}
