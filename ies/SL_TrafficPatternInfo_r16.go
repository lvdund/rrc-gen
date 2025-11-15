package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TrafficPatternInfo_r16 struct {
	trafficPeriodicity_r16  SL_TrafficPatternInfo_r16_trafficPeriodicity_r16 `madatory`
	timingOffset_r16        int64                                            `lb:0,ub:10239,madatory`
	messageSize_r16         uper.BitString                                   `lb:8,ub:8,madatory`
	sl_QoS_FlowIdentity_r16 SL_QoS_FlowIdentity_r16                          `madatory`
}

func (ie *SL_TrafficPatternInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.trafficPeriodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode trafficPeriodicity_r16", err)
	}
	if err = w.WriteInteger(ie.timingOffset_r16, &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
		return utils.WrapError("WriteInteger timingOffset_r16", err)
	}
	if err = w.WriteBitString(ie.messageSize_r16.Bytes, uint(ie.messageSize_r16.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString messageSize_r16", err)
	}
	if err = ie.sl_QoS_FlowIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_QoS_FlowIdentity_r16", err)
	}
	return nil
}

func (ie *SL_TrafficPatternInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.trafficPeriodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode trafficPeriodicity_r16", err)
	}
	var tmp_int_timingOffset_r16 int64
	if tmp_int_timingOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
		return utils.WrapError("ReadInteger timingOffset_r16", err)
	}
	ie.timingOffset_r16 = tmp_int_timingOffset_r16
	var tmp_bs_messageSize_r16 []byte
	var n_messageSize_r16 uint
	if tmp_bs_messageSize_r16, n_messageSize_r16, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString messageSize_r16", err)
	}
	ie.messageSize_r16 = uper.BitString{
		Bytes:   tmp_bs_messageSize_r16,
		NumBits: uint64(n_messageSize_r16),
	}
	if err = ie.sl_QoS_FlowIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_QoS_FlowIdentity_r16", err)
	}
	return nil
}
