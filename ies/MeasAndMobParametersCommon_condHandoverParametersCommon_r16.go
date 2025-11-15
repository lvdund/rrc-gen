package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersCommon_condHandoverParametersCommon_r16 struct {
	condHandoverFDD_TDD_r16 *MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFDD_TDD_r16 `optional`
	condHandoverFR1_FR2_r16 *MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFR1_FR2_r16 `optional`
}

func (ie *MeasAndMobParametersCommon_condHandoverParametersCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condHandoverFDD_TDD_r16 != nil, ie.condHandoverFR1_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condHandoverFDD_TDD_r16 != nil {
		if err = ie.condHandoverFDD_TDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condHandoverFDD_TDD_r16", err)
		}
	}
	if ie.condHandoverFR1_FR2_r16 != nil {
		if err = ie.condHandoverFR1_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condHandoverFR1_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_condHandoverParametersCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var condHandoverFDD_TDD_r16Present bool
	if condHandoverFDD_TDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condHandoverFR1_FR2_r16Present bool
	if condHandoverFR1_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if condHandoverFDD_TDD_r16Present {
		ie.condHandoverFDD_TDD_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFDD_TDD_r16)
		if err = ie.condHandoverFDD_TDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condHandoverFDD_TDD_r16", err)
		}
	}
	if condHandoverFR1_FR2_r16Present {
		ie.condHandoverFR1_FR2_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFR1_FR2_r16)
		if err = ie.condHandoverFR1_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condHandoverFR1_FR2_r16", err)
		}
	}
	return nil
}
