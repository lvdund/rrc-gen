package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC struct {
	measAndMobParametersMRDC_Common   *MeasAndMobParametersMRDC_Common   `optional`
	measAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff `optional`
	measAndMobParametersMRDC_FRX_Diff *MeasAndMobParametersMRDC_FRX_Diff `optional`
}

func (ie *MeasAndMobParametersMRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_Common != nil, ie.measAndMobParametersMRDC_XDD_Diff != nil, ie.measAndMobParametersMRDC_FRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_Common != nil {
		if err = ie.measAndMobParametersMRDC_Common.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_Common", err)
		}
	}
	if ie.measAndMobParametersMRDC_XDD_Diff != nil {
		if err = ie.measAndMobParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if ie.measAndMobParametersMRDC_FRX_Diff != nil {
		if err = ie.measAndMobParametersMRDC_FRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_FRX_Diff", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_CommonPresent bool
	if measAndMobParametersMRDC_CommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersMRDC_XDD_DiffPresent bool
	if measAndMobParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersMRDC_FRX_DiffPresent bool
	if measAndMobParametersMRDC_FRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_CommonPresent {
		ie.measAndMobParametersMRDC_Common = new(MeasAndMobParametersMRDC_Common)
		if err = ie.measAndMobParametersMRDC_Common.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_Common", err)
		}
	}
	if measAndMobParametersMRDC_XDD_DiffPresent {
		ie.measAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
		if err = ie.measAndMobParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if measAndMobParametersMRDC_FRX_DiffPresent {
		ie.measAndMobParametersMRDC_FRX_Diff = new(MeasAndMobParametersMRDC_FRX_Diff)
		if err = ie.measAndMobParametersMRDC_FRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_FRX_Diff", err)
		}
	}
	return nil
}
