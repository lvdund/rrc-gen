package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioPagingInformation_v1700_IEs struct {
	ue_RadioPagingInfo_r17            *[]byte                                                               `optional`
	inactiveStatePO_Determination_r17 *UERadioPagingInformation_v1700_IEs_inactiveStatePO_Determination_r17 `optional`
	numberOfRxRedCap_r17              *UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17              `optional`
	halfDuplexFDD_TypeA_RedCap_r17    []FreqBandIndicatorNR                                                 `lb:1,ub:maxBands,optional`
	nonCriticalExtension              interface{}                                                           `optional`
}

func (ie *UERadioPagingInformation_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_RadioPagingInfo_r17 != nil, ie.inactiveStatePO_Determination_r17 != nil, ie.numberOfRxRedCap_r17 != nil, len(ie.halfDuplexFDD_TypeA_RedCap_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_RadioPagingInfo_r17 != nil {
		if err = w.WriteOctetString(*ie.ue_RadioPagingInfo_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ue_RadioPagingInfo_r17", err)
		}
	}
	if ie.inactiveStatePO_Determination_r17 != nil {
		if err = ie.inactiveStatePO_Determination_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inactiveStatePO_Determination_r17", err)
		}
	}
	if ie.numberOfRxRedCap_r17 != nil {
		if err = ie.numberOfRxRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode numberOfRxRedCap_r17", err)
		}
	}
	if len(ie.halfDuplexFDD_TypeA_RedCap_r17) > 0 {
		tmp_halfDuplexFDD_TypeA_RedCap_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.halfDuplexFDD_TypeA_RedCap_r17 {
			tmp_halfDuplexFDD_TypeA_RedCap_r17.Value = append(tmp_halfDuplexFDD_TypeA_RedCap_r17.Value, &i)
		}
		if err = tmp_halfDuplexFDD_TypeA_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode halfDuplexFDD_TypeA_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *UERadioPagingInformation_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_RadioPagingInfo_r17Present bool
	if ue_RadioPagingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inactiveStatePO_Determination_r17Present bool
	if inactiveStatePO_Determination_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var numberOfRxRedCap_r17Present bool
	if numberOfRxRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var halfDuplexFDD_TypeA_RedCap_r17Present bool
	if halfDuplexFDD_TypeA_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_RadioPagingInfo_r17Present {
		var tmp_os_ue_RadioPagingInfo_r17 []byte
		if tmp_os_ue_RadioPagingInfo_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ue_RadioPagingInfo_r17", err)
		}
		ie.ue_RadioPagingInfo_r17 = &tmp_os_ue_RadioPagingInfo_r17
	}
	if inactiveStatePO_Determination_r17Present {
		ie.inactiveStatePO_Determination_r17 = new(UERadioPagingInformation_v1700_IEs_inactiveStatePO_Determination_r17)
		if err = ie.inactiveStatePO_Determination_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inactiveStatePO_Determination_r17", err)
		}
	}
	if numberOfRxRedCap_r17Present {
		ie.numberOfRxRedCap_r17 = new(UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17)
		if err = ie.numberOfRxRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode numberOfRxRedCap_r17", err)
		}
	}
	if halfDuplexFDD_TypeA_RedCap_r17Present {
		tmp_halfDuplexFDD_TypeA_RedCap_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_halfDuplexFDD_TypeA_RedCap_r17 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_halfDuplexFDD_TypeA_RedCap_r17.Decode(r, fn_halfDuplexFDD_TypeA_RedCap_r17); err != nil {
			return utils.WrapError("Decode halfDuplexFDD_TypeA_RedCap_r17", err)
		}
		ie.halfDuplexFDD_TypeA_RedCap_r17 = []FreqBandIndicatorNR{}
		for _, i := range tmp_halfDuplexFDD_TypeA_RedCap_r17.Value {
			ie.halfDuplexFDD_TypeA_RedCap_r17 = append(ie.halfDuplexFDD_TypeA_RedCap_r17, *i)
		}
	}
	return nil
}
