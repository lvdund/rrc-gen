package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigUC_Info_r17 struct {
	sl_DestinationIndex_r17 *SL_DestinationIndex_r16 `optional`
	sl_DRX_ConfigUC_r17     *SL_DRX_ConfigUC_r17     `optional`
}

func (ie *SL_DRX_ConfigUC_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_DestinationIndex_r17 != nil, ie.sl_DRX_ConfigUC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_DestinationIndex_r17 != nil {
		if err = ie.sl_DestinationIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DestinationIndex_r17", err)
		}
	}
	if ie.sl_DRX_ConfigUC_r17 != nil {
		if err = ie.sl_DRX_ConfigUC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DRX_ConfigUC_r17", err)
		}
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_DestinationIndex_r17Present bool
	if sl_DestinationIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DRX_ConfigUC_r17Present bool
	if sl_DRX_ConfigUC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DestinationIndex_r17Present {
		ie.sl_DestinationIndex_r17 = new(SL_DestinationIndex_r16)
		if err = ie.sl_DestinationIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DestinationIndex_r17", err)
		}
	}
	if sl_DRX_ConfigUC_r17Present {
		ie.sl_DRX_ConfigUC_r17 = new(SL_DRX_ConfigUC_r17)
		if err = ie.sl_DRX_ConfigUC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DRX_ConfigUC_r17", err)
		}
	}
	return nil
}
