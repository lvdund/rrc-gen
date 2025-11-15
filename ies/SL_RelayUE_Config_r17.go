package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RelayUE_Config_r17 struct {
	threshHighRelay_r17 *RSRP_Range `optional`
	threshLowRelay_r17  *RSRP_Range `optional`
	hystMaxRelay_r17    *Hysteresis `optional`
	hystMinRelay_r17    *Hysteresis `optional`
}

func (ie *SL_RelayUE_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.threshHighRelay_r17 != nil, ie.threshLowRelay_r17 != nil, ie.hystMaxRelay_r17 != nil, ie.hystMinRelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.threshHighRelay_r17 != nil {
		if err = ie.threshHighRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode threshHighRelay_r17", err)
		}
	}
	if ie.threshLowRelay_r17 != nil {
		if err = ie.threshLowRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode threshLowRelay_r17", err)
		}
	}
	if ie.hystMaxRelay_r17 != nil {
		if err = ie.hystMaxRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode hystMaxRelay_r17", err)
		}
	}
	if ie.hystMinRelay_r17 != nil {
		if err = ie.hystMinRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode hystMinRelay_r17", err)
		}
	}
	return nil
}

func (ie *SL_RelayUE_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var threshHighRelay_r17Present bool
	if threshHighRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var threshLowRelay_r17Present bool
	if threshLowRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var hystMaxRelay_r17Present bool
	if hystMaxRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var hystMinRelay_r17Present bool
	if hystMinRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if threshHighRelay_r17Present {
		ie.threshHighRelay_r17 = new(RSRP_Range)
		if err = ie.threshHighRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode threshHighRelay_r17", err)
		}
	}
	if threshLowRelay_r17Present {
		ie.threshLowRelay_r17 = new(RSRP_Range)
		if err = ie.threshLowRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode threshLowRelay_r17", err)
		}
	}
	if hystMaxRelay_r17Present {
		ie.hystMaxRelay_r17 = new(Hysteresis)
		if err = ie.hystMaxRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode hystMaxRelay_r17", err)
		}
	}
	if hystMinRelay_r17Present {
		ie.hystMinRelay_r17 = new(Hysteresis)
		if err = ie.hystMinRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode hystMinRelay_r17", err)
		}
	}
	return nil
}
