package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common struct {
	independentGapConfig *MeasAndMobParametersMRDC_Common_independentGapConfig `optional`
}

func (ie *MeasAndMobParametersMRDC_Common) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.independentGapConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.independentGapConfig != nil {
		if err = ie.independentGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode independentGapConfig", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common) Decode(r *uper.UperReader) error {
	var err error
	var independentGapConfigPresent bool
	if independentGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if independentGapConfigPresent {
		ie.independentGapConfig = new(MeasAndMobParametersMRDC_Common_independentGapConfig)
		if err = ie.independentGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode independentGapConfig", err)
		}
	}
	return nil
}
