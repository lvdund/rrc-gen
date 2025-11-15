package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_reducedMaxMIMO_LayersFR2 struct {
	reducedMIMO_LayersFR2_DL MIMO_LayersDL `madatory`
	reducedMIMO_LayersFR2_UL MIMO_LayersUL `madatory`
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR2) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reducedMIMO_LayersFR2_DL.Encode(w); err != nil {
		return utils.WrapError("Encode reducedMIMO_LayersFR2_DL", err)
	}
	if err = ie.reducedMIMO_LayersFR2_UL.Encode(w); err != nil {
		return utils.WrapError("Encode reducedMIMO_LayersFR2_UL", err)
	}
	return nil
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR2) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reducedMIMO_LayersFR2_DL.Decode(r); err != nil {
		return utils.WrapError("Decode reducedMIMO_LayersFR2_DL", err)
	}
	if err = ie.reducedMIMO_LayersFR2_UL.Decode(r); err != nil {
		return utils.WrapError("Decode reducedMIMO_LayersFR2_UL", err)
	}
	return nil
}
