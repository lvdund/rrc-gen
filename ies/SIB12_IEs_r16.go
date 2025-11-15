package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB12_IEs_r16 struct {
	sl_ConfigCommonNR_r16             SL_ConfigCommonNR_r16                      `madatory`
	lateNonCriticalExtension          *[]byte                                    `optional`
	sl_DRX_ConfigCommonGC_BC_r17      *SL_DRX_ConfigGC_BC_r17                    `optional,ext-1`
	sl_DiscConfigCommon_r17           *SL_DiscConfigCommon_r17                   `optional,ext-1`
	sl_L2U2N_Relay_r17                *SIB12_IEs_r16_sl_L2U2N_Relay_r17          `optional,ext-1`
	sl_NonRelayDiscovery_r17          *SIB12_IEs_r16_sl_NonRelayDiscovery_r17    `optional,ext-1`
	sl_L3U2N_RelayDiscovery_r17       *SIB12_IEs_r16_sl_L3U2N_RelayDiscovery_r17 `optional,ext-1`
	sl_TimersAndConstantsRemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17         `optional,ext-1`
}

func (ie *SIB12_IEs_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_DRX_ConfigCommonGC_BC_r17 != nil || ie.sl_DiscConfigCommon_r17 != nil || ie.sl_L2U2N_Relay_r17 != nil || ie.sl_NonRelayDiscovery_r17 != nil || ie.sl_L3U2N_RelayDiscovery_r17 != nil || ie.sl_TimersAndConstantsRemoteUE_r17 != nil
	preambleBits := []bool{hasExtensions, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_ConfigCommonNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ConfigCommonNR_r16", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_DRX_ConfigCommonGC_BC_r17 != nil || ie.sl_DiscConfigCommon_r17 != nil || ie.sl_L2U2N_Relay_r17 != nil || ie.sl_NonRelayDiscovery_r17 != nil || ie.sl_L3U2N_RelayDiscovery_r17 != nil || ie.sl_TimersAndConstantsRemoteUE_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB12_IEs_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_DRX_ConfigCommonGC_BC_r17 != nil, ie.sl_DiscConfigCommon_r17 != nil, ie.sl_L2U2N_Relay_r17 != nil, ie.sl_NonRelayDiscovery_r17 != nil, ie.sl_L3U2N_RelayDiscovery_r17 != nil, ie.sl_TimersAndConstantsRemoteUE_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_DRX_ConfigCommonGC_BC_r17 optional
			if ie.sl_DRX_ConfigCommonGC_BC_r17 != nil {
				if err = ie.sl_DRX_ConfigCommonGC_BC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_DRX_ConfigCommonGC_BC_r17", err)
				}
			}
			// encode sl_DiscConfigCommon_r17 optional
			if ie.sl_DiscConfigCommon_r17 != nil {
				if err = ie.sl_DiscConfigCommon_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_DiscConfigCommon_r17", err)
				}
			}
			// encode sl_L2U2N_Relay_r17 optional
			if ie.sl_L2U2N_Relay_r17 != nil {
				if err = ie.sl_L2U2N_Relay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_L2U2N_Relay_r17", err)
				}
			}
			// encode sl_NonRelayDiscovery_r17 optional
			if ie.sl_NonRelayDiscovery_r17 != nil {
				if err = ie.sl_NonRelayDiscovery_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_NonRelayDiscovery_r17", err)
				}
			}
			// encode sl_L3U2N_RelayDiscovery_r17 optional
			if ie.sl_L3U2N_RelayDiscovery_r17 != nil {
				if err = ie.sl_L3U2N_RelayDiscovery_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_L3U2N_RelayDiscovery_r17", err)
				}
			}
			// encode sl_TimersAndConstantsRemoteUE_r17 optional
			if ie.sl_TimersAndConstantsRemoteUE_r17 != nil {
				if err = ie.sl_TimersAndConstantsRemoteUE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_TimersAndConstantsRemoteUE_r17", err)
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

func (ie *SIB12_IEs_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_ConfigCommonNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ConfigCommonNR_r16", err)
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
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

			sl_DRX_ConfigCommonGC_BC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_DiscConfigCommon_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_L2U2N_Relay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_NonRelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_L3U2N_RelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_TimersAndConstantsRemoteUE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_DRX_ConfigCommonGC_BC_r17 optional
			if sl_DRX_ConfigCommonGC_BC_r17Present {
				ie.sl_DRX_ConfigCommonGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
				if err = ie.sl_DRX_ConfigCommonGC_BC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_DRX_ConfigCommonGC_BC_r17", err)
				}
			}
			// decode sl_DiscConfigCommon_r17 optional
			if sl_DiscConfigCommon_r17Present {
				ie.sl_DiscConfigCommon_r17 = new(SL_DiscConfigCommon_r17)
				if err = ie.sl_DiscConfigCommon_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_DiscConfigCommon_r17", err)
				}
			}
			// decode sl_L2U2N_Relay_r17 optional
			if sl_L2U2N_Relay_r17Present {
				ie.sl_L2U2N_Relay_r17 = new(SIB12_IEs_r16_sl_L2U2N_Relay_r17)
				if err = ie.sl_L2U2N_Relay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_L2U2N_Relay_r17", err)
				}
			}
			// decode sl_NonRelayDiscovery_r17 optional
			if sl_NonRelayDiscovery_r17Present {
				ie.sl_NonRelayDiscovery_r17 = new(SIB12_IEs_r16_sl_NonRelayDiscovery_r17)
				if err = ie.sl_NonRelayDiscovery_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_NonRelayDiscovery_r17", err)
				}
			}
			// decode sl_L3U2N_RelayDiscovery_r17 optional
			if sl_L3U2N_RelayDiscovery_r17Present {
				ie.sl_L3U2N_RelayDiscovery_r17 = new(SIB12_IEs_r16_sl_L3U2N_RelayDiscovery_r17)
				if err = ie.sl_L3U2N_RelayDiscovery_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_L3U2N_RelayDiscovery_r17", err)
				}
			}
			// decode sl_TimersAndConstantsRemoteUE_r17 optional
			if sl_TimersAndConstantsRemoteUE_r17Present {
				ie.sl_TimersAndConstantsRemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17)
				if err = ie.sl_TimersAndConstantsRemoteUE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_TimersAndConstantsRemoteUE_r17", err)
				}
			}
		}
	}
	return nil
}
