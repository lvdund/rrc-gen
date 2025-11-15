package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16 struct {
	condPSCellChangeFDD_TDD_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFDD_TDD_r16 `optional`
	condPSCellChangeFR1_FR2_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFR1_FR2_r16 `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condPSCellChangeFDD_TDD_r16 != nil, ie.condPSCellChangeFR1_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condPSCellChangeFDD_TDD_r16 != nil {
		if err = ie.condPSCellChangeFDD_TDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellChangeFDD_TDD_r16", err)
		}
	}
	if ie.condPSCellChangeFR1_FR2_r16 != nil {
		if err = ie.condPSCellChangeFR1_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellChangeFR1_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var condPSCellChangeFDD_TDD_r16Present bool
	if condPSCellChangeFDD_TDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condPSCellChangeFR1_FR2_r16Present bool
	if condPSCellChangeFR1_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if condPSCellChangeFDD_TDD_r16Present {
		ie.condPSCellChangeFDD_TDD_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFDD_TDD_r16)
		if err = ie.condPSCellChangeFDD_TDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellChangeFDD_TDD_r16", err)
		}
	}
	if condPSCellChangeFR1_FR2_r16Present {
		ie.condPSCellChangeFR1_FR2_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFR1_FR2_r16)
		if err = ie.condPSCellChangeFR1_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellChangeFR1_FR2_r16", err)
		}
	}
	return nil
}
