package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceFR2_2_r17 struct {
	reducedMaxMIMO_LayersFR2_2_r17 *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17 `lb:1,ub:8,optional`
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxMIMO_LayersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxMIMO_LayersFR2_2_r17 != nil {
		if err = ie.reducedMaxMIMO_LayersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxMIMO_LayersFR2_2_r17Present bool
	if reducedMaxMIMO_LayersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxMIMO_LayersFR2_2_r17Present {
		ie.reducedMaxMIMO_LayersFR2_2_r17 = new(MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17)
		if err = ie.reducedMaxMIMO_LayersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR2_2_r17", err)
		}
	}
	return nil
}
