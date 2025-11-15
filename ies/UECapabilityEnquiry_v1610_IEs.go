package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_v1610_IEs struct {
	rrc_SegAllowed_r16   *UECapabilityEnquiry_v1610_IEs_rrc_SegAllowed_r16 `optional`
	nonCriticalExtension interface{}                                       `optional`
}

func (ie *UECapabilityEnquiry_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rrc_SegAllowed_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rrc_SegAllowed_r16 != nil {
		if err = ie.rrc_SegAllowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rrc_SegAllowed_r16", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var rrc_SegAllowed_r16Present bool
	if rrc_SegAllowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rrc_SegAllowed_r16Present {
		ie.rrc_SegAllowed_r16 = new(UECapabilityEnquiry_v1610_IEs_rrc_SegAllowed_r16)
		if err = ie.rrc_SegAllowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rrc_SegAllowed_r16", err)
		}
	}
	return nil
}
