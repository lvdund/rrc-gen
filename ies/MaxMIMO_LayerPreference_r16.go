package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreference_r16 struct {
	reducedMaxMIMO_LayersFR1_r16 *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16 `lb:1,ub:8,optional`
	reducedMaxMIMO_LayersFR2_r16 *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR2_r16 `lb:1,ub:8,optional`
}

func (ie *MaxMIMO_LayerPreference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxMIMO_LayersFR1_r16 != nil, ie.reducedMaxMIMO_LayersFR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxMIMO_LayersFR1_r16 != nil {
		if err = ie.reducedMaxMIMO_LayersFR1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR1_r16", err)
		}
	}
	if ie.reducedMaxMIMO_LayersFR2_r16 != nil {
		if err = ie.reducedMaxMIMO_LayersFR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR2_r16", err)
		}
	}
	return nil
}

func (ie *MaxMIMO_LayerPreference_r16) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxMIMO_LayersFR1_r16Present bool
	if reducedMaxMIMO_LayersFR1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxMIMO_LayersFR2_r16Present bool
	if reducedMaxMIMO_LayersFR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxMIMO_LayersFR1_r16Present {
		ie.reducedMaxMIMO_LayersFR1_r16 = new(MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16)
		if err = ie.reducedMaxMIMO_LayersFR1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR1_r16", err)
		}
	}
	if reducedMaxMIMO_LayersFR2_r16Present {
		ie.reducedMaxMIMO_LayersFR2_r16 = new(MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR2_r16)
		if err = ie.reducedMaxMIMO_LayersFR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR2_r16", err)
		}
	}
	return nil
}
