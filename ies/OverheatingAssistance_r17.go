package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_r17 struct {
	reducedMaxBW_FR2_2_r17     *OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17     `optional`
	reducedMaxMIMO_LayersFR2_2 *OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2 `optional`
}

func (ie *OverheatingAssistance_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxBW_FR2_2_r17 != nil, ie.reducedMaxMIMO_LayersFR2_2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxBW_FR2_2_r17 != nil {
		if err = ie.reducedMaxBW_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR2_2_r17", err)
		}
	}
	if ie.reducedMaxMIMO_LayersFR2_2 != nil {
		if err = ie.reducedMaxMIMO_LayersFR2_2.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR2_2", err)
		}
	}
	return nil
}

func (ie *OverheatingAssistance_r17) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxBW_FR2_2_r17Present bool
	if reducedMaxBW_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxMIMO_LayersFR2_2Present bool
	if reducedMaxMIMO_LayersFR2_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxBW_FR2_2_r17Present {
		ie.reducedMaxBW_FR2_2_r17 = new(OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17)
		if err = ie.reducedMaxBW_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR2_2_r17", err)
		}
	}
	if reducedMaxMIMO_LayersFR2_2Present {
		ie.reducedMaxMIMO_LayersFR2_2 = new(OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2)
		if err = ie.reducedMaxMIMO_LayersFR2_2.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR2_2", err)
		}
	}
	return nil
}
