package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuspendConfig struct {
	fullI_RNTI                  I_RNTI_Value              `madatory`
	shortI_RNTI                 ShortI_RNTI_Value         `madatory`
	ran_PagingCycle             PagingCycle               `madatory`
	ran_NotificationAreaInfo    *RAN_NotificationAreaInfo `optional`
	t380                        *PeriodicRNAU_TimerValue  `optional`
	nextHopChainingCount        NextHopChainingCount      `madatory`
	sl_UEIdentityRemote_r17     *RNTI_Value               `optional,ext-1`
	sdt_Config_r17              *SDT_Config_r17           `optional,ext-1,setuprelease`
	srs_PosRRC_Inactive_r17     *SRS_PosRRC_Inactive_r17  `optional,ext-1,setuprelease`
	ran_ExtendedPagingCycle_r17 *ExtendedPagingCycle_r17  `optional,ext-1`
}

func (ie *SuspendConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_UEIdentityRemote_r17 != nil || ie.sdt_Config_r17 != nil || ie.srs_PosRRC_Inactive_r17 != nil || ie.ran_ExtendedPagingCycle_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ran_NotificationAreaInfo != nil, ie.t380 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.fullI_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode fullI_RNTI", err)
	}
	if err = ie.shortI_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode shortI_RNTI", err)
	}
	if err = ie.ran_PagingCycle.Encode(w); err != nil {
		return utils.WrapError("Encode ran_PagingCycle", err)
	}
	if ie.ran_NotificationAreaInfo != nil {
		if err = ie.ran_NotificationAreaInfo.Encode(w); err != nil {
			return utils.WrapError("Encode ran_NotificationAreaInfo", err)
		}
	}
	if ie.t380 != nil {
		if err = ie.t380.Encode(w); err != nil {
			return utils.WrapError("Encode t380", err)
		}
	}
	if err = ie.nextHopChainingCount.Encode(w); err != nil {
		return utils.WrapError("Encode nextHopChainingCount", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_UEIdentityRemote_r17 != nil || ie.sdt_Config_r17 != nil || ie.srs_PosRRC_Inactive_r17 != nil || ie.ran_ExtendedPagingCycle_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SuspendConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_UEIdentityRemote_r17 != nil, ie.sdt_Config_r17 != nil, ie.srs_PosRRC_Inactive_r17 != nil, ie.ran_ExtendedPagingCycle_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_UEIdentityRemote_r17 optional
			if ie.sl_UEIdentityRemote_r17 != nil {
				if err = ie.sl_UEIdentityRemote_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_UEIdentityRemote_r17", err)
				}
			}
			// encode sdt_Config_r17 optional
			if ie.sdt_Config_r17 != nil {
				tmp_sdt_Config_r17 := utils.SetupRelease[*SDT_Config_r17]{
					Setup: ie.sdt_Config_r17,
				}
				if err = tmp_sdt_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sdt_Config_r17", err)
				}
			}
			// encode srs_PosRRC_Inactive_r17 optional
			if ie.srs_PosRRC_Inactive_r17 != nil {
				tmp_srs_PosRRC_Inactive_r17 := utils.SetupRelease[*SRS_PosRRC_Inactive_r17]{
					Setup: ie.srs_PosRRC_Inactive_r17,
				}
				if err = tmp_srs_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PosRRC_Inactive_r17", err)
				}
			}
			// encode ran_ExtendedPagingCycle_r17 optional
			if ie.ran_ExtendedPagingCycle_r17 != nil {
				if err = ie.ran_ExtendedPagingCycle_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ran_ExtendedPagingCycle_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SuspendConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ran_NotificationAreaInfoPresent bool
	if ran_NotificationAreaInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var t380Present bool
	if t380Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.fullI_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode fullI_RNTI", err)
	}
	if err = ie.shortI_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode shortI_RNTI", err)
	}
	if err = ie.ran_PagingCycle.Decode(r); err != nil {
		return utils.WrapError("Decode ran_PagingCycle", err)
	}
	if ran_NotificationAreaInfoPresent {
		ie.ran_NotificationAreaInfo = new(RAN_NotificationAreaInfo)
		if err = ie.ran_NotificationAreaInfo.Decode(r); err != nil {
			return utils.WrapError("Decode ran_NotificationAreaInfo", err)
		}
	}
	if t380Present {
		ie.t380 = new(PeriodicRNAU_TimerValue)
		if err = ie.t380.Decode(r); err != nil {
			return utils.WrapError("Decode t380", err)
		}
	}
	if err = ie.nextHopChainingCount.Decode(r); err != nil {
		return utils.WrapError("Decode nextHopChainingCount", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_UEIdentityRemote_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sdt_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ran_ExtendedPagingCycle_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_UEIdentityRemote_r17 optional
			if sl_UEIdentityRemote_r17Present {
				ie.sl_UEIdentityRemote_r17 = new(RNTI_Value)
				if err = ie.sl_UEIdentityRemote_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_UEIdentityRemote_r17", err)
				}
			}
			// decode sdt_Config_r17 optional
			if sdt_Config_r17Present {
				tmp_sdt_Config_r17 := utils.SetupRelease[*SDT_Config_r17]{}
				if err = tmp_sdt_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sdt_Config_r17", err)
				}
				ie.sdt_Config_r17 = tmp_sdt_Config_r17.Setup
			}
			// decode srs_PosRRC_Inactive_r17 optional
			if srs_PosRRC_Inactive_r17Present {
				tmp_srs_PosRRC_Inactive_r17 := utils.SetupRelease[*SRS_PosRRC_Inactive_r17]{}
				if err = tmp_srs_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_PosRRC_Inactive_r17", err)
				}
				ie.srs_PosRRC_Inactive_r17 = tmp_srs_PosRRC_Inactive_r17.Setup
			}
			// decode ran_ExtendedPagingCycle_r17 optional
			if ran_ExtendedPagingCycle_r17Present {
				ie.ran_ExtendedPagingCycle_r17 = new(ExtendedPagingCycle_r17)
				if err = ie.ran_ExtendedPagingCycle_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ran_ExtendedPagingCycle_r17", err)
				}
			}
		}
	}
	return nil
}
