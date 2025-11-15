package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkPC5_r16_sl_Reception_r16 struct {
	harq_RxProcessSidelink_r16   BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16    `madatory`
	pscch_RxSidelink_r16         BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16          `madatory`
	scs_CP_PatternRxSidelink_r16 *BandSidelinkPC5_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16 `optional`
	extendedCP_RxSidelink_r16    *BandSidelinkPC5_r16_sl_Reception_r16_extendedCP_RxSidelink_r16    `optional`
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_CP_PatternRxSidelink_r16 != nil, ie.extendedCP_RxSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.harq_RxProcessSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode harq_RxProcessSidelink_r16", err)
	}
	if err = ie.pscch_RxSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pscch_RxSidelink_r16", err)
	}
	if ie.scs_CP_PatternRxSidelink_r16 != nil {
		if err = ie.scs_CP_PatternRxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_CP_PatternRxSidelink_r16", err)
		}
	}
	if ie.extendedCP_RxSidelink_r16 != nil {
		if err = ie.extendedCP_RxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode extendedCP_RxSidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16) Decode(r *uper.UperReader) error {
	var err error
	var scs_CP_PatternRxSidelink_r16Present bool
	if scs_CP_PatternRxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extendedCP_RxSidelink_r16Present bool
	if extendedCP_RxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.harq_RxProcessSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode harq_RxProcessSidelink_r16", err)
	}
	if err = ie.pscch_RxSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pscch_RxSidelink_r16", err)
	}
	if scs_CP_PatternRxSidelink_r16Present {
		ie.scs_CP_PatternRxSidelink_r16 = new(BandSidelinkPC5_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16)
		if err = ie.scs_CP_PatternRxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_CP_PatternRxSidelink_r16", err)
		}
	}
	if extendedCP_RxSidelink_r16Present {
		ie.extendedCP_RxSidelink_r16 = new(BandSidelinkPC5_r16_sl_Reception_r16_extendedCP_RxSidelink_r16)
		if err = ie.extendedCP_RxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extendedCP_RxSidelink_r16", err)
		}
	}
	return nil
}
