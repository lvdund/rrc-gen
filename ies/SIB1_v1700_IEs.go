package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1700_IEs struct {
	hsdn_Cell_r17                  *SIB1_v1700_IEs_hsdn_Cell_r17                  `optional`
	uac_BarringInfo_v1700          *SIB1_v1700_IEs_uac_BarringInfo_v1700          `optional`
	sdt_ConfigCommon_r17           *SDT_ConfigCommonSIB_r17                       `optional`
	redCap_ConfigCommon_r17        *RedCap_ConfigCommonSIB_r17                    `optional`
	featurePriorities_r17          *FeaturePriorities_r17                         `optional`
	si_SchedulingInfo_v1700        *SI_SchedulingInfo_v1700                       `optional`
	hyperSFN_r17                   *uper.BitString                                `lb:10,ub:10,optional`
	eDRX_AllowedIdle_r17           *SIB1_v1700_IEs_eDRX_AllowedIdle_r17           `optional`
	eDRX_AllowedInactive_r17       *SIB1_v1700_IEs_eDRX_AllowedInactive_r17       `optional`
	intraFreqReselectionRedCap_r17 *SIB1_v1700_IEs_intraFreqReselectionRedCap_r17 `optional`
	cellBarredNTN_r17              *SIB1_v1700_IEs_cellBarredNTN_r17              `optional`
	nonCriticalExtension           interface{}                                    `optional`
}

func (ie *SIB1_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.hsdn_Cell_r17 != nil, ie.uac_BarringInfo_v1700 != nil, ie.sdt_ConfigCommon_r17 != nil, ie.redCap_ConfigCommon_r17 != nil, ie.featurePriorities_r17 != nil, ie.si_SchedulingInfo_v1700 != nil, ie.hyperSFN_r17 != nil, ie.eDRX_AllowedIdle_r17 != nil, ie.eDRX_AllowedInactive_r17 != nil, ie.intraFreqReselectionRedCap_r17 != nil, ie.cellBarredNTN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.hsdn_Cell_r17 != nil {
		if err = ie.hsdn_Cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode hsdn_Cell_r17", err)
		}
	}
	if ie.uac_BarringInfo_v1700 != nil {
		if err = ie.uac_BarringInfo_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringInfo_v1700", err)
		}
	}
	if ie.sdt_ConfigCommon_r17 != nil {
		if err = ie.sdt_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_ConfigCommon_r17", err)
		}
	}
	if ie.redCap_ConfigCommon_r17 != nil {
		if err = ie.redCap_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode redCap_ConfigCommon_r17", err)
		}
	}
	if ie.featurePriorities_r17 != nil {
		if err = ie.featurePriorities_r17.Encode(w); err != nil {
			return utils.WrapError("Encode featurePriorities_r17", err)
		}
	}
	if ie.si_SchedulingInfo_v1700 != nil {
		if err = ie.si_SchedulingInfo_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode si_SchedulingInfo_v1700", err)
		}
	}
	if ie.hyperSFN_r17 != nil {
		if err = w.WriteBitString(ie.hyperSFN_r17.Bytes, uint(ie.hyperSFN_r17.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode hyperSFN_r17", err)
		}
	}
	if ie.eDRX_AllowedIdle_r17 != nil {
		if err = ie.eDRX_AllowedIdle_r17.Encode(w); err != nil {
			return utils.WrapError("Encode eDRX_AllowedIdle_r17", err)
		}
	}
	if ie.eDRX_AllowedInactive_r17 != nil {
		if err = ie.eDRX_AllowedInactive_r17.Encode(w); err != nil {
			return utils.WrapError("Encode eDRX_AllowedInactive_r17", err)
		}
	}
	if ie.intraFreqReselectionRedCap_r17 != nil {
		if err = ie.intraFreqReselectionRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqReselectionRedCap_r17", err)
		}
	}
	if ie.cellBarredNTN_r17 != nil {
		if err = ie.cellBarredNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cellBarredNTN_r17", err)
		}
	}
	return nil
}

func (ie *SIB1_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var hsdn_Cell_r17Present bool
	if hsdn_Cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uac_BarringInfo_v1700Present bool
	if uac_BarringInfo_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_ConfigCommon_r17Present bool
	if sdt_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var redCap_ConfigCommon_r17Present bool
	if redCap_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var featurePriorities_r17Present bool
	if featurePriorities_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var si_SchedulingInfo_v1700Present bool
	if si_SchedulingInfo_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var hyperSFN_r17Present bool
	if hyperSFN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eDRX_AllowedIdle_r17Present bool
	if eDRX_AllowedIdle_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eDRX_AllowedInactive_r17Present bool
	if eDRX_AllowedInactive_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqReselectionRedCap_r17Present bool
	if intraFreqReselectionRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellBarredNTN_r17Present bool
	if cellBarredNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if hsdn_Cell_r17Present {
		ie.hsdn_Cell_r17 = new(SIB1_v1700_IEs_hsdn_Cell_r17)
		if err = ie.hsdn_Cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode hsdn_Cell_r17", err)
		}
	}
	if uac_BarringInfo_v1700Present {
		ie.uac_BarringInfo_v1700 = new(SIB1_v1700_IEs_uac_BarringInfo_v1700)
		if err = ie.uac_BarringInfo_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringInfo_v1700", err)
		}
	}
	if sdt_ConfigCommon_r17Present {
		ie.sdt_ConfigCommon_r17 = new(SDT_ConfigCommonSIB_r17)
		if err = ie.sdt_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_ConfigCommon_r17", err)
		}
	}
	if redCap_ConfigCommon_r17Present {
		ie.redCap_ConfigCommon_r17 = new(RedCap_ConfigCommonSIB_r17)
		if err = ie.redCap_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode redCap_ConfigCommon_r17", err)
		}
	}
	if featurePriorities_r17Present {
		ie.featurePriorities_r17 = new(FeaturePriorities_r17)
		if err = ie.featurePriorities_r17.Decode(r); err != nil {
			return utils.WrapError("Decode featurePriorities_r17", err)
		}
	}
	if si_SchedulingInfo_v1700Present {
		ie.si_SchedulingInfo_v1700 = new(SI_SchedulingInfo_v1700)
		if err = ie.si_SchedulingInfo_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode si_SchedulingInfo_v1700", err)
		}
	}
	if hyperSFN_r17Present {
		var tmp_bs_hyperSFN_r17 []byte
		var n_hyperSFN_r17 uint
		if tmp_bs_hyperSFN_r17, n_hyperSFN_r17, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode hyperSFN_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_hyperSFN_r17,
			NumBits: uint64(n_hyperSFN_r17),
		}
		ie.hyperSFN_r17 = &tmp_bitstring
	}
	if eDRX_AllowedIdle_r17Present {
		ie.eDRX_AllowedIdle_r17 = new(SIB1_v1700_IEs_eDRX_AllowedIdle_r17)
		if err = ie.eDRX_AllowedIdle_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eDRX_AllowedIdle_r17", err)
		}
	}
	if eDRX_AllowedInactive_r17Present {
		ie.eDRX_AllowedInactive_r17 = new(SIB1_v1700_IEs_eDRX_AllowedInactive_r17)
		if err = ie.eDRX_AllowedInactive_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eDRX_AllowedInactive_r17", err)
		}
	}
	if intraFreqReselectionRedCap_r17Present {
		ie.intraFreqReselectionRedCap_r17 = new(SIB1_v1700_IEs_intraFreqReselectionRedCap_r17)
		if err = ie.intraFreqReselectionRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqReselectionRedCap_r17", err)
		}
	}
	if cellBarredNTN_r17Present {
		ie.cellBarredNTN_r17 = new(SIB1_v1700_IEs_cellBarredNTN_r17)
		if err = ie.cellBarredNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cellBarredNTN_r17", err)
		}
	}
	return nil
}
