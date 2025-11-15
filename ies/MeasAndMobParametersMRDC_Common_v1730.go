package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1730 struct {
	independentGapConfig_maxCC_r17 *MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17 `lb:1,ub:32,optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.independentGapConfig_maxCC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.independentGapConfig_maxCC_r17 != nil {
		if err = ie.independentGapConfig_maxCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode independentGapConfig_maxCC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Decode(r *uper.UperReader) error {
	var err error
	var independentGapConfig_maxCC_r17Present bool
	if independentGapConfig_maxCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if independentGapConfig_maxCC_r17Present {
		ie.independentGapConfig_maxCC_r17 = new(MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17)
		if err = ie.independentGapConfig_maxCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode independentGapConfig_maxCC_r17", err)
		}
	}
	return nil
}
