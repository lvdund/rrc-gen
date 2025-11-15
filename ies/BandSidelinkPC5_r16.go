package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkPC5_r16 struct {
	freqBandSidelink_r16                         FreqBandIndicatorNR                                               `madatory`
	sl_Reception_r16                             *BandSidelinkPC5_r16_sl_Reception_r16                             `optional`
	sl_Tx_256QAM_r16                             *BandSidelinkPC5_r16_sl_Tx_256QAM_r16                             `optional`
	lowSE_64QAM_MCS_TableSidelink_r16            *BandSidelinkPC5_r16_lowSE_64QAM_MCS_TableSidelink_r16            `optional`
	csi_ReportSidelink_r16                       *BandSidelinkPC5_r16_csi_ReportSidelink_r16                       `optional,ext-1`
	rankTwoReception_r16                         *BandSidelinkPC5_r16_rankTwoReception_r16                         `optional,ext-1`
	sl_openLoopPC_RSRP_ReportSidelink_r16        *BandSidelinkPC5_r16_sl_openLoopPC_RSRP_ReportSidelink_r16        `optional,ext-1`
	sl_Rx_256QAM_r16                             *BandSidelinkPC5_r16_sl_Rx_256QAM_r16                             `optional,ext-1`
	rx_IUC_Scheme1_PreferredMode2Sidelink_r17    *BandSidelinkPC5_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17    `optional,ext-2`
	rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 *BandSidelinkPC5_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 `optional,ext-2`
	rx_IUC_Scheme2_Mode2Sidelink_r17             *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17             `optional,ext-2`
	rx_IUC_Scheme1_SCI_r17                       *BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_r17                       `optional,ext-2`
	rx_IUC_Scheme1_SCI_ExplicitReq_r17           *BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17           `optional,ext-2`
	scheme2_ConflictDeterminationRSRP_r17        *BandSidelinkPC5_r16_scheme2_ConflictDeterminationRSRP_r17        `optional,ext-2`
}

func (ie *BandSidelinkPC5_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.csi_ReportSidelink_r16 != nil || ie.rankTwoReception_r16 != nil || ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.sl_Rx_256QAM_r16 != nil || ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_SCI_r17 != nil || ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil || ie.scheme2_ConflictDeterminationRSRP_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sl_Reception_r16 != nil, ie.sl_Tx_256QAM_r16 != nil, ie.lowSE_64QAM_MCS_TableSidelink_r16 != nil}
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
	if ie.sl_Tx_256QAM_r16 != nil {
		if err = ie.sl_Tx_256QAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Tx_256QAM_r16", err)
		}
	}
	if ie.lowSE_64QAM_MCS_TableSidelink_r16 != nil {
		if err = ie.lowSE_64QAM_MCS_TableSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode lowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.csi_ReportSidelink_r16 != nil || ie.rankTwoReception_r16 != nil || ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.sl_Rx_256QAM_r16 != nil, ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.rx_IUC_Scheme1_SCI_r17 != nil || ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil || ie.scheme2_ConflictDeterminationRSRP_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandSidelinkPC5_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.csi_ReportSidelink_r16 != nil, ie.rankTwoReception_r16 != nil, ie.sl_openLoopPC_RSRP_ReportSidelink_r16 != nil, ie.sl_Rx_256QAM_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode csi_ReportSidelink_r16 optional
			if ie.csi_ReportSidelink_r16 != nil {
				if err = ie.csi_ReportSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_ReportSidelink_r16", err)
				}
			}
			// encode rankTwoReception_r16 optional
			if ie.rankTwoReception_r16 != nil {
				if err = ie.rankTwoReception_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rankTwoReception_r16", err)
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
			optionals_ext_2 := []bool{ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil, ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil, ie.rx_IUC_Scheme2_Mode2Sidelink_r17 != nil, ie.rx_IUC_Scheme1_SCI_r17 != nil, ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil, ie.scheme2_ConflictDeterminationRSRP_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			// encode scheme2_ConflictDeterminationRSRP_r17 optional
			if ie.scheme2_ConflictDeterminationRSRP_r17 != nil {
				if err = ie.scheme2_ConflictDeterminationRSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode scheme2_ConflictDeterminationRSRP_r17", err)
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

func (ie *BandSidelinkPC5_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Reception_r16Present bool
	if sl_Reception_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Tx_256QAM_r16Present bool
	if sl_Tx_256QAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lowSE_64QAM_MCS_TableSidelink_r16Present bool
	if lowSE_64QAM_MCS_TableSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.freqBandSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandSidelink_r16", err)
	}
	if sl_Reception_r16Present {
		ie.sl_Reception_r16 = new(BandSidelinkPC5_r16_sl_Reception_r16)
		if err = ie.sl_Reception_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Reception_r16", err)
		}
	}
	if sl_Tx_256QAM_r16Present {
		ie.sl_Tx_256QAM_r16 = new(BandSidelinkPC5_r16_sl_Tx_256QAM_r16)
		if err = ie.sl_Tx_256QAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Tx_256QAM_r16", err)
		}
	}
	if lowSE_64QAM_MCS_TableSidelink_r16Present {
		ie.lowSE_64QAM_MCS_TableSidelink_r16 = new(BandSidelinkPC5_r16_lowSE_64QAM_MCS_TableSidelink_r16)
		if err = ie.lowSE_64QAM_MCS_TableSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode lowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			csi_ReportSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rankTwoReception_r16Present, err := extReader.ReadBool()
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
			// decode csi_ReportSidelink_r16 optional
			if csi_ReportSidelink_r16Present {
				ie.csi_ReportSidelink_r16 = new(BandSidelinkPC5_r16_csi_ReportSidelink_r16)
				if err = ie.csi_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_ReportSidelink_r16", err)
				}
			}
			// decode rankTwoReception_r16 optional
			if rankTwoReception_r16Present {
				ie.rankTwoReception_r16 = new(BandSidelinkPC5_r16_rankTwoReception_r16)
				if err = ie.rankTwoReception_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rankTwoReception_r16", err)
				}
			}
			// decode sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if sl_openLoopPC_RSRP_ReportSidelink_r16Present {
				ie.sl_openLoopPC_RSRP_ReportSidelink_r16 = new(BandSidelinkPC5_r16_sl_openLoopPC_RSRP_ReportSidelink_r16)
				if err = ie.sl_openLoopPC_RSRP_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// decode sl_Rx_256QAM_r16 optional
			if sl_Rx_256QAM_r16Present {
				ie.sl_Rx_256QAM_r16 = new(BandSidelinkPC5_r16_sl_Rx_256QAM_r16)
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
			scheme2_ConflictDeterminationRSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present {
				ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present {
				ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if rx_IUC_Scheme2_Mode2Sidelink_r17Present {
				ie.rx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17)
				if err = ie.rx_IUC_Scheme2_Mode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_SCI_r17 optional
			if rx_IUC_Scheme1_SCI_r17Present {
				ie.rx_IUC_Scheme1_SCI_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_r17)
				if err = ie.rx_IUC_Scheme1_SCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// decode rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if rx_IUC_Scheme1_SCI_ExplicitReq_r17Present {
				ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17)
				if err = ie.rx_IUC_Scheme1_SCI_ExplicitReq_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
				}
			}
			// decode scheme2_ConflictDeterminationRSRP_r17 optional
			if scheme2_ConflictDeterminationRSRP_r17Present {
				ie.scheme2_ConflictDeterminationRSRP_r17 = new(BandSidelinkPC5_r16_scheme2_ConflictDeterminationRSRP_r17)
				if err = ie.scheme2_ConflictDeterminationRSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode scheme2_ConflictDeterminationRSRP_r17", err)
				}
			}
		}
	}
	return nil
}
