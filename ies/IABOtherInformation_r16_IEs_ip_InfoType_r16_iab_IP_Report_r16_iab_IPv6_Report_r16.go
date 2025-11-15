package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_nothing uint64 = iota
	IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_AddressReport_r16
	IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_PrefixReport_r16
)

type IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16 struct {
	Choice                     uint64
	iab_IPv6_AddressReport_r16 *IAB_IP_AddressAndTraffic_r16
	iab_IPv6_PrefixReport_r16  *IAB_IP_PrefixAndTraffic_r16
}

func (ie *IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_AddressReport_r16:
		if err = ie.iab_IPv6_AddressReport_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode iab_IPv6_AddressReport_r16", err)
		}
	case IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_PrefixReport_r16:
		if err = ie.iab_IPv6_PrefixReport_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode iab_IPv6_PrefixReport_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_AddressReport_r16:
		ie.iab_IPv6_AddressReport_r16 = new(IAB_IP_AddressAndTraffic_r16)
		if err = ie.iab_IPv6_AddressReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IPv6_AddressReport_r16", err)
		}
	case IABOtherInformation_r16_IEs_ip_InfoType_r16_iab_IP_Report_r16_iab_IPv6_Report_r16_Choice_iab_IPv6_PrefixReport_r16:
		ie.iab_IPv6_PrefixReport_r16 = new(IAB_IP_PrefixAndTraffic_r16)
		if err = ie.iab_IPv6_PrefixReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IPv6_PrefixReport_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
