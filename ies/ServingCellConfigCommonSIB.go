package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommonSIB struct {
	downlinkConfigCommon             DownlinkConfigCommonSIB                                      `madatory`
	uplinkConfigCommon               *UplinkConfigCommonSIB                                       `optional`
	supplementaryUplink              *UplinkConfigCommonSIB                                       `optional`
	n_TimingAdvanceOffset            *ServingCellConfigCommonSIB_n_TimingAdvanceOffset            `optional`
	ssb_PositionsInBurst             *ServingCellConfigCommonSIB_ssb_PositionsInBurst             `lb:8,ub:8,optional`
	ssb_PeriodicityServingCell       ServingCellConfigCommonSIB_ssb_PeriodicityServingCell        `madatory`
	tdd_UL_DL_ConfigurationCommon    *TDD_UL_DL_ConfigCommon                                      `optional`
	ss_PBCH_BlockPower               int64                                                        `lb:-60,ub:50,madatory`
	channelAccessMode_r16            *ServingCellConfigCommonSIB_channelAccessMode_r16            `optional,ext-1`
	discoveryBurstWindowLength_r16   *ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16   `optional,ext-1`
	highSpeedConfig_r16              *HighSpeedConfig_r16                                         `optional,ext-1`
	channelAccessMode2_r17           *ServingCellConfigCommonSIB_channelAccessMode2_r17           `optional,ext-2`
	discoveryBurstWindowLength_v1700 *ServingCellConfigCommonSIB_discoveryBurstWindowLength_v1700 `optional,ext-2`
	highSpeedConfigFR2_r17           *HighSpeedConfigFR2_r17                                      `optional,ext-2`
	uplinkConfigCommon_v1700         *UplinkConfigCommonSIB_v1700                                 `optional,ext-2`
	enhancedMeasurementLEO_r17       *ServingCellConfigCommonSIB_enhancedMeasurementLEO_r17       `optional,ext-3`
}

func (ie *ServingCellConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.channelAccessMode_r16 != nil || ie.discoveryBurstWindowLength_r16 != nil || ie.highSpeedConfig_r16 != nil || ie.channelAccessMode2_r17 != nil || ie.discoveryBurstWindowLength_v1700 != nil || ie.highSpeedConfigFR2_r17 != nil || ie.uplinkConfigCommon_v1700 != nil || ie.enhancedMeasurementLEO_r17 != nil
	preambleBits := []bool{hasExtensions, ie.uplinkConfigCommon != nil, ie.supplementaryUplink != nil, ie.n_TimingAdvanceOffset != nil, ie.ssb_PositionsInBurst != nil, ie.tdd_UL_DL_ConfigurationCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.downlinkConfigCommon.Encode(w); err != nil {
		return utils.WrapError("Encode downlinkConfigCommon", err)
	}
	if ie.uplinkConfigCommon != nil {
		if err = ie.uplinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkConfigCommon", err)
		}
	}
	if ie.supplementaryUplink != nil {
		if err = ie.supplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode supplementaryUplink", err)
		}
	}
	if ie.n_TimingAdvanceOffset != nil {
		if err = ie.n_TimingAdvanceOffset.Encode(w); err != nil {
			return utils.WrapError("Encode n_TimingAdvanceOffset", err)
		}
	}
	if ie.ssb_PositionsInBurst != nil {
		if err = ie.ssb_PositionsInBurst.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionsInBurst", err)
		}
	}
	if err = ie.ssb_PeriodicityServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_PeriodicityServingCell", err)
	}
	if ie.tdd_UL_DL_ConfigurationCommon != nil {
		if err = ie.tdd_UL_DL_ConfigurationCommon.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	if err = w.WriteInteger(ie.ss_PBCH_BlockPower, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger ss_PBCH_BlockPower", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.channelAccessMode_r16 != nil || ie.discoveryBurstWindowLength_r16 != nil || ie.highSpeedConfig_r16 != nil, ie.channelAccessMode2_r17 != nil || ie.discoveryBurstWindowLength_v1700 != nil || ie.highSpeedConfigFR2_r17 != nil || ie.uplinkConfigCommon_v1700 != nil, ie.enhancedMeasurementLEO_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfigCommonSIB", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.channelAccessMode_r16 != nil, ie.discoveryBurstWindowLength_r16 != nil, ie.highSpeedConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelAccessMode_r16 optional
			if ie.channelAccessMode_r16 != nil {
				if err = ie.channelAccessMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode_r16", err)
				}
			}
			// encode discoveryBurstWindowLength_r16 optional
			if ie.discoveryBurstWindowLength_r16 != nil {
				if err = ie.discoveryBurstWindowLength_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discoveryBurstWindowLength_r16", err)
				}
			}
			// encode highSpeedConfig_r16 optional
			if ie.highSpeedConfig_r16 != nil {
				if err = ie.highSpeedConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode highSpeedConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.channelAccessMode2_r17 != nil, ie.discoveryBurstWindowLength_v1700 != nil, ie.highSpeedConfigFR2_r17 != nil, ie.uplinkConfigCommon_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelAccessMode2_r17 optional
			if ie.channelAccessMode2_r17 != nil {
				if err = ie.channelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode2_r17", err)
				}
			}
			// encode discoveryBurstWindowLength_v1700 optional
			if ie.discoveryBurstWindowLength_v1700 != nil {
				if err = ie.discoveryBurstWindowLength_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discoveryBurstWindowLength_v1700", err)
				}
			}
			// encode highSpeedConfigFR2_r17 optional
			if ie.highSpeedConfigFR2_r17 != nil {
				if err = ie.highSpeedConfigFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode highSpeedConfigFR2_r17", err)
				}
			}
			// encode uplinkConfigCommon_v1700 optional
			if ie.uplinkConfigCommon_v1700 != nil {
				if err = ie.uplinkConfigCommon_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkConfigCommon_v1700", err)
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
			optionals_ext_3 := []bool{ie.enhancedMeasurementLEO_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enhancedMeasurementLEO_r17 optional
			if ie.enhancedMeasurementLEO_r17 != nil {
				if err = ie.enhancedMeasurementLEO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedMeasurementLEO_r17", err)
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

func (ie *ServingCellConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkConfigCommonPresent bool
	if uplinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supplementaryUplinkPresent bool
	if supplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var n_TimingAdvanceOffsetPresent bool
	if n_TimingAdvanceOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionsInBurstPresent bool
	if ssb_PositionsInBurstPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_UL_DL_ConfigurationCommonPresent bool
	if tdd_UL_DL_ConfigurationCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.downlinkConfigCommon.Decode(r); err != nil {
		return utils.WrapError("Decode downlinkConfigCommon", err)
	}
	if uplinkConfigCommonPresent {
		ie.uplinkConfigCommon = new(UplinkConfigCommonSIB)
		if err = ie.uplinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkConfigCommon", err)
		}
	}
	if supplementaryUplinkPresent {
		ie.supplementaryUplink = new(UplinkConfigCommonSIB)
		if err = ie.supplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode supplementaryUplink", err)
		}
	}
	if n_TimingAdvanceOffsetPresent {
		ie.n_TimingAdvanceOffset = new(ServingCellConfigCommonSIB_n_TimingAdvanceOffset)
		if err = ie.n_TimingAdvanceOffset.Decode(r); err != nil {
			return utils.WrapError("Decode n_TimingAdvanceOffset", err)
		}
	}
	if ssb_PositionsInBurstPresent {
		ie.ssb_PositionsInBurst = new(ServingCellConfigCommonSIB_ssb_PositionsInBurst)
		if err = ie.ssb_PositionsInBurst.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionsInBurst", err)
		}
	}
	if err = ie.ssb_PeriodicityServingCell.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_PeriodicityServingCell", err)
	}
	if tdd_UL_DL_ConfigurationCommonPresent {
		ie.tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
		if err = ie.tdd_UL_DL_ConfigurationCommon.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	var tmp_int_ss_PBCH_BlockPower int64
	if tmp_int_ss_PBCH_BlockPower, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger ss_PBCH_BlockPower", err)
	}
	ie.ss_PBCH_BlockPower = tmp_int_ss_PBCH_BlockPower

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

			channelAccessMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			discoveryBurstWindowLength_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			highSpeedConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelAccessMode_r16 optional
			if channelAccessMode_r16Present {
				ie.channelAccessMode_r16 = new(ServingCellConfigCommonSIB_channelAccessMode_r16)
				if err = ie.channelAccessMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode_r16", err)
				}
			}
			// decode discoveryBurstWindowLength_r16 optional
			if discoveryBurstWindowLength_r16Present {
				ie.discoveryBurstWindowLength_r16 = new(ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16)
				if err = ie.discoveryBurstWindowLength_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode discoveryBurstWindowLength_r16", err)
				}
			}
			// decode highSpeedConfig_r16 optional
			if highSpeedConfig_r16Present {
				ie.highSpeedConfig_r16 = new(HighSpeedConfig_r16)
				if err = ie.highSpeedConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode highSpeedConfig_r16", err)
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

			channelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			discoveryBurstWindowLength_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			highSpeedConfigFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkConfigCommon_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelAccessMode2_r17 optional
			if channelAccessMode2_r17Present {
				ie.channelAccessMode2_r17 = new(ServingCellConfigCommonSIB_channelAccessMode2_r17)
				if err = ie.channelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode2_r17", err)
				}
			}
			// decode discoveryBurstWindowLength_v1700 optional
			if discoveryBurstWindowLength_v1700Present {
				ie.discoveryBurstWindowLength_v1700 = new(ServingCellConfigCommonSIB_discoveryBurstWindowLength_v1700)
				if err = ie.discoveryBurstWindowLength_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode discoveryBurstWindowLength_v1700", err)
				}
			}
			// decode highSpeedConfigFR2_r17 optional
			if highSpeedConfigFR2_r17Present {
				ie.highSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
				if err = ie.highSpeedConfigFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode highSpeedConfigFR2_r17", err)
				}
			}
			// decode uplinkConfigCommon_v1700 optional
			if uplinkConfigCommon_v1700Present {
				ie.uplinkConfigCommon_v1700 = new(UplinkConfigCommonSIB_v1700)
				if err = ie.uplinkConfigCommon_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkConfigCommon_v1700", err)
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

			enhancedMeasurementLEO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enhancedMeasurementLEO_r17 optional
			if enhancedMeasurementLEO_r17Present {
				ie.enhancedMeasurementLEO_r17 = new(ServingCellConfigCommonSIB_enhancedMeasurementLEO_r17)
				if err = ie.enhancedMeasurementLEO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedMeasurementLEO_r17", err)
				}
			}
		}
	}
	return nil
}
