package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance struct {
	reducedMaxCCs            *ReducedMaxCCs_r16                              `optional`
	reducedMaxBW_FR1         *ReducedMaxBW_FRx_r16                           `optional`
	reducedMaxBW_FR2         *ReducedMaxBW_FRx_r16                           `optional`
	reducedMaxMIMO_LayersFR1 *OverheatingAssistance_reducedMaxMIMO_LayersFR1 `optional`
	reducedMaxMIMO_LayersFR2 *OverheatingAssistance_reducedMaxMIMO_LayersFR2 `optional`
}

func (ie *OverheatingAssistance) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxCCs != nil, ie.reducedMaxBW_FR1 != nil, ie.reducedMaxBW_FR2 != nil, ie.reducedMaxMIMO_LayersFR1 != nil, ie.reducedMaxMIMO_LayersFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxCCs != nil {
		if err = ie.reducedMaxCCs.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxCCs", err)
		}
	}
	if ie.reducedMaxBW_FR1 != nil {
		if err = ie.reducedMaxBW_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR1", err)
		}
	}
	if ie.reducedMaxBW_FR2 != nil {
		if err = ie.reducedMaxBW_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR2", err)
		}
	}
	if ie.reducedMaxMIMO_LayersFR1 != nil {
		if err = ie.reducedMaxMIMO_LayersFR1.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR1", err)
		}
	}
	if ie.reducedMaxMIMO_LayersFR2 != nil {
		if err = ie.reducedMaxMIMO_LayersFR2.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxMIMO_LayersFR2", err)
		}
	}
	return nil
}

func (ie *OverheatingAssistance) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxCCsPresent bool
	if reducedMaxCCsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxBW_FR1Present bool
	if reducedMaxBW_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxBW_FR2Present bool
	if reducedMaxBW_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxMIMO_LayersFR1Present bool
	if reducedMaxMIMO_LayersFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxMIMO_LayersFR2Present bool
	if reducedMaxMIMO_LayersFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxCCsPresent {
		ie.reducedMaxCCs = new(ReducedMaxCCs_r16)
		if err = ie.reducedMaxCCs.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxCCs", err)
		}
	}
	if reducedMaxBW_FR1Present {
		ie.reducedMaxBW_FR1 = new(ReducedMaxBW_FRx_r16)
		if err = ie.reducedMaxBW_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR1", err)
		}
	}
	if reducedMaxBW_FR2Present {
		ie.reducedMaxBW_FR2 = new(ReducedMaxBW_FRx_r16)
		if err = ie.reducedMaxBW_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR2", err)
		}
	}
	if reducedMaxMIMO_LayersFR1Present {
		ie.reducedMaxMIMO_LayersFR1 = new(OverheatingAssistance_reducedMaxMIMO_LayersFR1)
		if err = ie.reducedMaxMIMO_LayersFR1.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR1", err)
		}
	}
	if reducedMaxMIMO_LayersFR2Present {
		ie.reducedMaxMIMO_LayersFR2 = new(OverheatingAssistance_reducedMaxMIMO_LayersFR2)
		if err = ie.reducedMaxMIMO_LayersFR2.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxMIMO_LayersFR2", err)
		}
	}
	return nil
}
