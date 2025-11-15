package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16 struct {
	freqBandSidelink_r16                             FreqBandIndicatorNR                                                `madatory`
	sl_Reception_r16                                 *BandSidelink_r16_sl_Reception_r16                                 `optional`
	sl_TransmissionMode1_r16                         *BandSidelink_r16_sl_TransmissionMode1_r16                         `optional`
	sync_Sidelink_r16                                *BandSidelink_r16_sync_Sidelink_r16                                `optional`
	sl_Tx_256QAM_r16                                 *BandSidelink_r16_sl_Tx_256QAM_r16                                 `optional`
	psfch_FormatZeroSidelink_r16                     *BandSidelink_r16_psfch_FormatZeroSidelink_r16                     `optional`
	lowSE_64QAM_MCS_TableSidelink_r16                *BandSidelink_r16_lowSE_64QAM_MCS_TableSidelink_r16                `optional`
	enb_sync_Sidelink_r16                            *BandSidelink_r16_enb_sync_Sidelink_r16                            `optional`
	sl_TransmissionMode2_r16                         *BandSidelink_r16_sl_TransmissionMode2_r16                         `optional,ext-1`
	congestionControlSidelink_r16                    *BandSidelink_r16_congestionControlSidelink_r16                    `optional,ext-1`
	fewerSymbolSlotSidelink_r16                      *BandSidelink_r16_fewerSymbolSlotSidelink_r16                      `optional,ext-1`
	sl_openLoopPC_RSRP_ReportSidelink_r16            *BandSidelink_r16_sl_openLoopPC_RSRP_ReportSidelink_r16            `optional,ext-1`
	sl_Rx_256QAM_r16                                 *BandSidelink_r16_sl_Rx_256QAM_r16                                 `optional,ext-1`
	ue_PowerClassSidelink_r16                        *BandSidelink_r16_ue_PowerClassSidelink_r16                        `optional,ext-2`
	sl_TransmissionMode2_RandomResourceSelection_r17 *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17 `optional,ext-3`
	sync_Sidelink_v1710                              *BandSidelink_r16_sync_Sidelink_v1710                              `optional,ext-3`
	enb_sync_Sidelink_v1710                          *BandSidelink_r16_enb_sync_Sidelink_v1710                          `optional,ext-3`
	rx_IUC_Scheme1_PreferredMode2Sidelink_r17        *BandSidelink_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17        `optional,ext-3`
	rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17     *BandSidelink_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17     `optional,ext-3`
	rx_IUC_Scheme2_Mode2Sidelink_r17                 *BandSidelink_r16_rx_IUC_Scheme2_Mode2Sidelink_r17                 `optional,ext-3`
	rx_IUC_Scheme1_SCI_r17                           *BandSidelink_r16_rx_IUC_Scheme1_SCI_r17                           `optional,ext-3`
	rx_IUC_Scheme1_SCI_ExplicitReq_r17               *BandSidelink_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17               `optional,ext-3`
}

func (ie *BandSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_TransmissionMode2_r16 != nil || ie.congestionControlSidelink_r16 != nil || ie.fewerSymbolSlotSidelink_r16 != nil || ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.sl_Rx_256QAM_r16 != nil || ie.ue_PowerClassSidelink_r16 != nil || ie.sl_TransmissionMode2_RandomResourceSelection_r17 != nil || ie.sync_Sidelink_v1710 != nil || ie.enb_sync_Sidelink_v1710 != nil || ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_SCI_r17 != nil || ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sl_Reception_r16 != nil, ie.sl_TransmissionMode1_r16 != nil, ie.sync_Sidelink_r16 != nil, ie.sl_Tx_256QAM_r16 != nil, ie.psfch_FormatZeroSidelink_r16 != nil, ie.lowSE_64QAM_MCS_TableSidelink_r16 != nil, ie.enb_sync_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.freqBandSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode freqBandSidelink_r16", err)
	}
	if ie.sl_Reception_r16 != nil {
		if err = ie.sl_Reception_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Reception_r16", err)
		}
	}
	if ie.sl_TransmissionMode1_r16 != nil {
		if err = ie.sl_TransmissionMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TransmissionMode1_r16", err)
		}
	}
	if ie.sync_Sidelink_r16 != nil {
		if err = ie.sync_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sync_Sidelink_r16", err)
		}
	}
	if ie.sl_Tx_256QAM_r16 != nil {
		if err = ie.sl_Tx_256QAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Tx_256QAM_r16", err)
		}
	}
	if ie.psfch_FormatZeroSidelink_r16 != nil {
		if err = ie.psfch_FormatZeroSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode psfch_FormatZeroSidelink_r16", err)
		}
	}
	if ie.lowSE_64QAM_MCS_TableSidelink_r16 != nil {
		if err = ie.lowSE_64QAM_MCS_TableSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode lowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if ie.enb_sync_Sidelink_r16 != nil {
		if err = ie.enb_sync_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode enb_sync_Sidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.sl_TransmissionMode2_r16 != nil || ie.congestionControlSidelink_r16 != nil || ie.fewerSymbolSlotSidelink_r16 != nil || ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.sl_Rx_256QAM_r16 != nil, ie.ue_PowerClassSidelink_r16 != nil, ie.sl_TransmissionMode2_RandomResourceSelection_r17 != nil || ie.sync_Sidelink_v1710 != nil || ie.enb_sync_Sidelink_v1710 != nil || ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_SCI_r17 != nil || ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandSidelink_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_TransmissionMode2_r16 != nil, ie.congestionControlSidelink_r16 != nil, ie.fewerSymbolSlotSidelink_r16 != nil, ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil, ie.sl_Rx_256QAM_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_TransmissionMode2_r16 optional
			if ie.sl_TransmissionMode2_r16 != nil {
				if err = ie.sl_TransmissionMode2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_TransmissionMode2_r16", err)
				}
			}
			// encode congestionControlSidelink_r16 optional
			if ie.congestionControlSidelink_r16 != nil {
				if err = ie.congestionControlSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode congestionControlSidelink_r16", err)
				}
			}
			// encode fewerSymbolSlotSidelink_r16 optional
			if ie.fewerSymbolSlotSidelink_r16 != nil {
				if err = ie.fewerSymbolSlotSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fewerSymbolSlotSidelink_r16", err)
				}
			}
			// encode sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil {
				if err = ie.sl_openLoopPC_RSRP_ReportSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// encode sl_Rx_256QAM_r16 optional
			if ie.sl_Rx_256QAM_r16 != nil {
				if err = ie.sl_Rx_256QAM_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_Rx_256QAM_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.ue_PowerClassSidelink_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ue_PowerClassSidelink_r16 optional
			if ie.ue_PowerClassSidelink_r16 != nil {
				if err = ie.ue_PowerClassSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ue_PowerClassSidelink_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.sl_TransmissionMode2_RandomResourceSelection_r17 != nil, ie.sync_Sidelink_v1710 != nil, ie.enb_sync_Sidelink_v1710 != nil, ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil, ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil, ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil, ie.rx_IUC_Scheme1_SCI_r17 != nil, ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_TransmissionMode2_RandomResourceSelection_r17 optional
			if ie.sl_TransmissionMode2_RandomResourceSelection_r17 != nil {
				if err = ie.sl_TransmissionMode2_RandomResourceSelection_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_TransmissionMode2_RandomResourceSelection_r17", err)
				}
			}
			// encode sync_Sidelink_v1710 optional
			if ie.sync_Sidelink_v1710 != nil {
				if err = ie.sync_Sidelink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sync_Sidelink_v1710", err)
				}
			}
			// encode enb_sync_Sidelink_v1710 optional
			if ie.enb_sync_Sidelink_v1710 != nil {
				if err = ie.enb_sync_Sidelink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enb_sync_Sidelink_v1710", err)
				}
			}
			// encode rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil {
				if err = ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// encode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil {
				if err = ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// encode rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
				if err = ie.rx_IUC_Scheme2_Mode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// encode rx_IUC_Scheme1_SCI_r17 optional
			if ie.rx_IUC_Scheme1_SCI_r17 != nil {
				if err = ie.rx_IUC_Scheme1_SCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// encode rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil {
				if err = ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
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

func (ie *BandSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Reception_r16Present bool
	if sl_Reception_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TransmissionMode1_r16Present bool
	if sl_TransmissionMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sync_Sidelink_r16Present bool
	if sync_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Tx_256QAM_r16Present bool
	if sl_Tx_256QAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var psfch_FormatZeroSidelink_r16Present bool
	if psfch_FormatZeroSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lowSE_64QAM_MCS_TableSidelink_r16Present bool
	if lowSE_64QAM_MCS_TableSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var enb_sync_Sidelink_r16Present bool
	if enb_sync_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.freqBandSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandSidelink_r16", err)
	}
	if sl_Reception_r16Present {
		ie.sl_Reception_r16 = new(BandSidelink_r16_sl_Reception_r16)
		if err = ie.sl_Reception_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Reception_r16", err)
		}
	}
	if sl_TransmissionMode1_r16Present {
		ie.sl_TransmissionMode1_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16)
		if err = ie.sl_TransmissionMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TransmissionMode1_r16", err)
		}
	}
	if sync_Sidelink_r16Present {
		ie.sync_Sidelink_r16 = new(BandSidelink_r16_sync_Sidelink_r16)
		if err = ie.sync_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sync_Sidelink_r16", err)
		}
	}
	if sl_Tx_256QAM_r16Present {
		ie.sl_Tx_256QAM_r16 = new(BandSidelink_r16_sl_Tx_256QAM_r16)
		if err = ie.sl_Tx_256QAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Tx_256QAM_r16", err)
		}
	}
	if psfch_FormatZeroSidelink_r16Present {
		ie.psfch_FormatZeroSidelink_r16 = new(BandSidelink_r16_psfch_FormatZeroSidelink_r16)
		if err = ie.psfch_FormatZeroSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode psfch_FormatZeroSidelink_r16", err)
		}
	}
	if lowSE_64QAM_MCS_TableSidelink_r16Present {
		ie.lowSE_64QAM_MCS_TableSidelink_r16 = new(BandSidelink_r16_lowSE_64QAM_MCS_TableSidelink_r16)
		if err = ie.lowSE_64QAM_MCS_TableSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode lowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if enb_sync_Sidelink_r16Present {
		ie.enb_sync_Sidelink_r16 = new(BandSidelink_r16_enb_sync_Sidelink_r16)
		if err = ie.enb_sync_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode enb_sync_Sidelink_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			sl_TransmissionMode2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			congestionControlSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			fewerSymbolSlotSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_openLoopPC_RSRP_ReportSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_Rx_256QAM_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_TransmissionMode2_r16 optional
			if sl_TransmissionMode2_r16Present {
				ie.sl_TransmissionMode2_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16)
				if err = ie.sl_TransmissionMode2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_TransmissionMode2_r16", err)
				}
			}
			// decode congestionControlSidelink_r16 optional
			if congestionControlSidelink_r16Present {
				ie.congestionControlSidelink_r16 = new(BandSidelink_r16_congestionControlSidelink_r16)
				if err = ie.congestionControlSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode congestionControlSidelink_r16", err)
				}
			}
			// decode fewerSymbolSlotSidelink_r16 optional
			if fewerSymbolSlotSidelink_r16Present {
				ie.fewerSymbolSlotSidelink_r16 = new(BandSidelink_r16_fewerSymbolSlotSidelink_r16)
				if err = ie.fewerSymbolSlotSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode fewerSymbolSlotSidelink_r16", err)
				}
			}
			// decode sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if sl_openLoopPC_RSRP_ReportSidelink_r16Present {
				ie.sl_openLoopPC_RSRP_ReportSidelink_r16 = new(BandSidelink_r16_sl_openLoopPC_RSRP_ReportSidelink_r16)
				if err = ie.sl_openLoopPC_RSRP_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// decode sl_Rx_256QAM_r16 optional
			if sl_Rx_256QAM_r16Present {
				ie.sl_Rx_256QAM_r16 = new(BandSidelink_r16_sl_Rx_256QAM_r16)
				if err = ie.sl_Rx_256QAM_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_Rx_256QAM_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ue_PowerClassSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ue_PowerClassSidelink_r16 optional
			if ue_PowerClassSidelink_r16Present {
				ie.ue_PowerClassSidelink_r16 = new(BandSidelink_r16_ue_PowerClassSidelink_r16)
				if err = ie.ue_PowerClassSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ue_PowerClassSidelink_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_TransmissionMode2_RandomResourceSelection_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sync_Sidelink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enb_sync_Sidelink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rx_IUC_Scheme2_Mode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rx_IUC_Scheme1_SCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rx_IUC_Scheme1_SCI_ExplicitReq_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_TransmissionMode2_RandomResourceSelection_r17 optional
			if sl_TransmissionMode2_RandomResourceSelection_r17Present {
				ie.sl_TransmissionMode2_RandomResourceSelection_r17 = new(BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17)
				if err = ie.sl_TransmissionMode2_RandomResourceSelection_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_TransmissionMode2_RandomResourceSelection_r17", err)
				}
			}
			// decode sync_Sidelink_v1710 optional
			if sync_Sidelink_v1710Present {
				ie.sync_Sidelink_v1710 = new(BandSidelink_r16_sync_Sidelink_v1710)
				if err = ie.sync_Sidelink_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode sync_Sidelink_v1710", err)
				}
			}
			// decode enb_sync_Sidelink_v1710 optional
			if enb_sync_Sidelink_v1710Present {
				ie.enb_sync_Sidelink_v1710 = new(BandSidelink_r16_enb_sync_Sidelink_v1710)
				if err = ie.enb_sync_Sidelink_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode enb_sync_Sidelink_v1710", err)
				}
			}
			// decode rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present {
				ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present {
				ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if rx_IUC_Scheme2_Mode2Sidelink_r17Present {
				ie.rx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme2_Mode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme2_Mode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_SCI_r17 optional
			if rx_IUC_Scheme1_SCI_r17Present {
				ie.rx_IUC_Scheme1_SCI_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_SCI_r17)
				if err = ie.rx_IUC_Scheme1_SCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if rx_IUC_Scheme1_SCI_ExplicitReq_r17Present {
				ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17)
				if err = ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
				}
			}
		}
	}
	return nil
}
