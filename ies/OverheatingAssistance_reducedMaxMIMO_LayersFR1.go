package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_reducedMaxMIMO_LayersFR1 struct {
	reducedMIMO_LayersFR1_DL MIMO_LayersDL `madatory`
	reducedMIMO_LayersFR1_UL MIMO_LayersUL `madatory`
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reducedMIMO_LayersFR1_DL.Encode(w); err != nil {
		return utils.WrapError("Encode reducedMIMO_LayersFR1_DL", err)
	}
	if err = ie.reducedMIMO_LayersFR1_UL.Encode(w); err != nil {
		return utils.WrapError("Encode reducedMIMO_LayersFR1_UL", err)
	}
	return nil
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reducedMIMO_LayersFR1_DL.Decode(r); err != nil {
		return utils.WrapError("Decode reducedMIMO_LayersFR1_DL", err)
	}
	if err = ie.reducedMIMO_LayersFR1_UL.Decode(r); err != nil {
		return utils.WrapError("Decode reducedMIMO_LayersFR1_UL", err)
	}
	return nil
}
