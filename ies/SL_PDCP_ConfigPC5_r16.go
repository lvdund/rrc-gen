package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PDCP_ConfigPC5_r16 struct {
	sl_PDCP_SN_Size_r16       *SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16       `optional`
	sl_OutOfOrderDelivery_r16 *SL_PDCP_ConfigPC5_r16_sl_OutOfOrderDelivery_r16 `optional`
}

func (ie *SL_PDCP_ConfigPC5_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PDCP_SN_Size_r16 != nil, ie.sl_OutOfOrderDelivery_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PDCP_SN_Size_r16 != nil {
		if err = ie.sl_PDCP_SN_Size_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PDCP_SN_Size_r16", err)
		}
	}
	if ie.sl_OutOfOrderDelivery_r16 != nil {
		if err = ie.sl_OutOfOrderDelivery_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_OutOfOrderDelivery_r16", err)
		}
	}
	return nil
}

func (ie *SL_PDCP_ConfigPC5_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PDCP_SN_Size_r16Present bool
	if sl_PDCP_SN_Size_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_OutOfOrderDelivery_r16Present bool
	if sl_OutOfOrderDelivery_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PDCP_SN_Size_r16Present {
		ie.sl_PDCP_SN_Size_r16 = new(SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16)
		if err = ie.sl_PDCP_SN_Size_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PDCP_SN_Size_r16", err)
		}
	}
	if sl_OutOfOrderDelivery_r16Present {
		ie.sl_OutOfOrderDelivery_r16 = new(SL_PDCP_ConfigPC5_r16_sl_OutOfOrderDelivery_r16)
		if err = ie.sl_OutOfOrderDelivery_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_OutOfOrderDelivery_r16", err)
		}
	}
	return nil
}
