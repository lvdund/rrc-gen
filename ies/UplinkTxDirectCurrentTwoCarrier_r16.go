package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentTwoCarrier_r16 struct {
	carrierOneInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16     `madatory`
	carrierTwoInfo_r16           UplinkTxDirectCurrentCarrierInfo_r16     `madatory`
	singlePA_TxDirectCurrent_r16 UplinkTxDirectCurrentTwoCarrierInfo_r16  `madatory`
	secondPA_TxDirectCurrent_r16 *UplinkTxDirectCurrentTwoCarrierInfo_r16 `optional`
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.secondPA_TxDirectCurrent_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierOneInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierOneInfo_r16", err)
	}
	if err = ie.carrierTwoInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierTwoInfo_r16", err)
	}
	if err = ie.singlePA_TxDirectCurrent_r16.Encode(w); err != nil {
		return utils.WrapError("Encode singlePA_TxDirectCurrent_r16", err)
	}
	if ie.secondPA_TxDirectCurrent_r16 != nil {
		if err = ie.secondPA_TxDirectCurrent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode secondPA_TxDirectCurrent_r16", err)
		}
	}
	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrier_r16) Decode(r *uper.UperReader) error {
	var err error
	var secondPA_TxDirectCurrent_r16Present bool
	if secondPA_TxDirectCurrent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierOneInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierOneInfo_r16", err)
	}
	if err = ie.carrierTwoInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierTwoInfo_r16", err)
	}
	if err = ie.singlePA_TxDirectCurrent_r16.Decode(r); err != nil {
		return utils.WrapError("Decode singlePA_TxDirectCurrent_r16", err)
	}
	if secondPA_TxDirectCurrent_r16Present {
		ie.secondPA_TxDirectCurrent_r16 = new(UplinkTxDirectCurrentTwoCarrierInfo_r16)
		if err = ie.secondPA_TxDirectCurrent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode secondPA_TxDirectCurrent_r16", err)
		}
	}
	return nil
}
