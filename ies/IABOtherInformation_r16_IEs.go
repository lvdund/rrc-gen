package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IABOtherInformation_r16_IEs struct {
	ip_InfoType_r16          *IABOtherInformation_r16_IEs_ip_InfoType_r16 `optional`
	lateNonCriticalExtension *[]byte                                      `optional,ext`
	nonCriticalExtension     interface{}                                  `optional,ext`
}

func (ie *IABOtherInformation_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ip_InfoType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ip_InfoType_r16 != nil {
		if err = ie.ip_InfoType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ip_InfoType_r16", err)
		}
	}
	return nil
}

func (ie *IABOtherInformation_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ip_InfoType_r16Present bool
	if ip_InfoType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ip_InfoType_r16Present {
		ie.ip_InfoType_r16 = new(IABOtherInformation_r16_IEs_ip_InfoType_r16)
		if err = ie.ip_InfoType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ip_InfoType_r16", err)
		}
	}
	return nil
}
