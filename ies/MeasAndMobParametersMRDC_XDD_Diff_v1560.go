package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_XDD_Diff_v1560 struct {
	sftd_MeasPSCell_NEDC *MeasAndMobParametersMRDC_XDD_Diff_v1560_sftd_MeasPSCell_NEDC `optional`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sftd_MeasPSCell_NEDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sftd_MeasPSCell_NEDC != nil {
		if err = ie.sftd_MeasPSCell_NEDC.Encode(w); err != nil {
			return utils.WrapError("Encode sftd_MeasPSCell_NEDC", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Decode(r *uper.UperReader) error {
	var err error
	var sftd_MeasPSCell_NEDCPresent bool
	if sftd_MeasPSCell_NEDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sftd_MeasPSCell_NEDCPresent {
		ie.sftd_MeasPSCell_NEDC = new(MeasAndMobParametersMRDC_XDD_Diff_v1560_sftd_MeasPSCell_NEDC)
		if err = ie.sftd_MeasPSCell_NEDC.Decode(r); err != nil {
			return utils.WrapError("Decode sftd_MeasPSCell_NEDC", err)
		}
	}
	return nil
}
