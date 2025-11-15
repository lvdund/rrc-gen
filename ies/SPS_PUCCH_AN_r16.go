package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SPS_PUCCH_AN_r16 struct {
	sps_PUCCH_AN_ResourceID_r16 PUCCH_ResourceId `madatory`
	maxPayloadSize_r16          *int64           `lb:4,ub:256,optional`
}

func (ie *SPS_PUCCH_AN_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxPayloadSize_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sps_PUCCH_AN_ResourceID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sps_PUCCH_AN_ResourceID_r16", err)
	}
	if ie.maxPayloadSize_r16 != nil {
		if err = w.WriteInteger(*ie.maxPayloadSize_r16, &uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode maxPayloadSize_r16", err)
		}
	}
	return nil
}

func (ie *SPS_PUCCH_AN_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxPayloadSize_r16Present bool
	if maxPayloadSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sps_PUCCH_AN_ResourceID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sps_PUCCH_AN_ResourceID_r16", err)
	}
	if maxPayloadSize_r16Present {
		var tmp_int_maxPayloadSize_r16 int64
		if tmp_int_maxPayloadSize_r16, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode maxPayloadSize_r16", err)
		}
		ie.maxPayloadSize_r16 = &tmp_int_maxPayloadSize_r16
	}
	return nil
}
