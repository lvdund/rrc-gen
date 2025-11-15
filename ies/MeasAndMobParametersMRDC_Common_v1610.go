package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1610 struct {
	condPSCellChangeParametersCommon_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16 `optional`
	pscellT312_r16                       *MeasAndMobParametersMRDC_Common_v1610_pscellT312_r16                       `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condPSCellChangeParametersCommon_r16 != nil, ie.pscellT312_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condPSCellChangeParametersCommon_r16 != nil {
		if err = ie.condPSCellChangeParametersCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellChangeParametersCommon_r16", err)
		}
	}
	if ie.pscellT312_r16 != nil {
		if err = ie.pscellT312_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pscellT312_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Decode(r *uper.UperReader) error {
	var err error
	var condPSCellChangeParametersCommon_r16Present bool
	if condPSCellChangeParametersCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pscellT312_r16Present bool
	if pscellT312_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if condPSCellChangeParametersCommon_r16Present {
		ie.condPSCellChangeParametersCommon_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16)
		if err = ie.condPSCellChangeParametersCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellChangeParametersCommon_r16", err)
		}
	}
	if pscellT312_r16Present {
		ie.pscellT312_r16 = new(MeasAndMobParametersMRDC_Common_v1610_pscellT312_r16)
		if err = ie.pscellT312_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pscellT312_r16", err)
		}
	}
	return nil
}
