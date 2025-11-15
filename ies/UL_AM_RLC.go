package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_AM_RLC struct {
	sn_FieldLength   *SN_FieldLengthAM          `optional`
	t_PollRetransmit T_PollRetransmit           `madatory`
	pollPDU          PollPDU                    `madatory`
	pollByte         PollByte                   `madatory`
	maxRetxThreshold UL_AM_RLC_maxRetxThreshold `madatory`
}

func (ie *UL_AM_RLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sn_FieldLength != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sn_FieldLength != nil {
		if err = ie.sn_FieldLength.Encode(w); err != nil {
			return utils.WrapError("Encode sn_FieldLength", err)
		}
	}
	if err = ie.t_PollRetransmit.Encode(w); err != nil {
		return utils.WrapError("Encode t_PollRetransmit", err)
	}
	if err = ie.pollPDU.Encode(w); err != nil {
		return utils.WrapError("Encode pollPDU", err)
	}
	if err = ie.pollByte.Encode(w); err != nil {
		return utils.WrapError("Encode pollByte", err)
	}
	if err = ie.maxRetxThreshold.Encode(w); err != nil {
		return utils.WrapError("Encode maxRetxThreshold", err)
	}
	return nil
}

func (ie *UL_AM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var sn_FieldLengthPresent bool
	if sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sn_FieldLengthPresent {
		ie.sn_FieldLength = new(SN_FieldLengthAM)
		if err = ie.sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode sn_FieldLength", err)
		}
	}
	if err = ie.t_PollRetransmit.Decode(r); err != nil {
		return utils.WrapError("Decode t_PollRetransmit", err)
	}
	if err = ie.pollPDU.Decode(r); err != nil {
		return utils.WrapError("Decode pollPDU", err)
	}
	if err = ie.pollByte.Decode(r); err != nil {
		return utils.WrapError("Decode pollByte", err)
	}
	if err = ie.maxRetxThreshold.Decode(r); err != nil {
		return utils.WrapError("Decode maxRetxThreshold", err)
	}
	return nil
}
