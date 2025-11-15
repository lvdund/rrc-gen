package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Request_r16 struct {
	iab_IPv4_AddressNumReq_r16 *IAB_IP_AddressNumReq_r16                                                               `optional`
	iab_IPv6_AddressReq_r16    *IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Request_r16_iab_IPv6_AddressReq_r16 `optional`
}

func (ie *IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Request_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.iab_IPv4_AddressNumReq_r16 != nil, ie.iab_IPv6_AddressReq_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.iab_IPv4_AddressNumReq_r16 != nil {
		if err = ie.iab_IPv4_AddressNumReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IPv4_AddressNumReq_r16", err)
		}
	}
	if ie.iab_IPv6_AddressReq_r16 != nil {
		if err = ie.iab_IPv6_AddressReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IPv6_AddressReq_r16", err)
		}
	}
	return nil
}

func (ie *IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Request_r16) Decode(r *uper.UperReader) error {
	var err error
	var iab_IPv4_AddressNumReq_r16Present bool
	if iab_IPv4_AddressNumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_IPv6_AddressReq_r16Present bool
	if iab_IPv6_AddressReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if iab_IPv4_AddressNumReq_r16Present {
		ie.iab_IPv4_AddressNumReq_r16 = new(IAB_IP_AddressNumReq_r16)
		if err = ie.iab_IPv4_AddressNumReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IPv4_AddressNumReq_r16", err)
		}
	}
	if iab_IPv6_AddressReq_r16Present {
		ie.iab_IPv6_AddressReq_r16 = new(IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Request_r16_iab_IPv6_AddressReq_r16)
		if err = ie.iab_IPv6_AddressReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IPv6_AddressReq_r16", err)
		}
	}
	return nil
}
