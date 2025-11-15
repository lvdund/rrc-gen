package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_CapabilityAddXDD_Mode struct {
	measAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff `optional`
	generalParametersMRDC_XDD_Diff    *GeneralParametersMRDC_XDD_Diff    `optional`
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_XDD_Diff != nil, ie.generalParametersMRDC_XDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_XDD_Diff != nil {
		if err = ie.measAndMobParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if ie.generalParametersMRDC_XDD_Diff != nil {
		if err = ie.generalParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode generalParametersMRDC_XDD_Diff", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_XDD_DiffPresent bool
	if measAndMobParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var generalParametersMRDC_XDD_DiffPresent bool
	if generalParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_XDD_DiffPresent {
		ie.measAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
		if err = ie.measAndMobParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if generalParametersMRDC_XDD_DiffPresent {
		ie.generalParametersMRDC_XDD_Diff = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.generalParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode generalParametersMRDC_XDD_Diff", err)
		}
	}
	return nil
}
