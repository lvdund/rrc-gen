package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters struct {
	supportedBandListNR                                    []BandNR                                   `lb:1,ub:maxBands,madatory`
	supportedBandCombinationList                           *BandCombinationList                       `optional`
	appliedFreqBandListFilter                              *FreqBandList                              `optional`
	supportedBandCombinationList_v1540                     *BandCombinationList_v1540                 `optional,ext-1`
	srs_SwitchingTimeRequested                             *RF_Parameters_srs_SwitchingTimeRequested  `optional,ext-1`
	supportedBandCombinationList_v1550                     *BandCombinationList_v1550                 `optional,ext-2`
	supportedBandCombinationList_v1560                     *BandCombinationList_v1560                 `optional,ext-3`
	supportedBandCombinationList_v1610                     *BandCombinationList_v1610                 `optional,ext-4`
	supportedBandCombinationListSidelinkEUTRA_NR_r16       *BandCombinationListSidelinkEUTRA_NR_r16   `optional,ext-4`
	supportedBandCombinationList_UplinkTxSwitch_r16        *BandCombinationList_UplinkTxSwitch_r16    `optional,ext-4`
	supportedBandCombinationList_v1630                     *BandCombinationList_v1630                 `optional,ext-5`
	supportedBandCombinationListSidelinkEUTRA_NR_v1630     *BandCombinationListSidelinkEUTRA_NR_v1630 `optional,ext-5`
	supportedBandCombinationList_UplinkTxSwitch_v1630      *BandCombinationList_UplinkTxSwitch_v1630  `optional,ext-5`
	supportedBandCombinationList_v1640                     *BandCombinationList_v1640                 `optional,ext-6`
	supportedBandCombinationList_UplinkTxSwitch_v1640      *BandCombinationList_UplinkTxSwitch_v1640  `optional,ext-6`
	supportedBandCombinationList_v1650                     *BandCombinationList_v1650                 `optional,ext-7`
	supportedBandCombinationList_UplinkTxSwitch_v1650      *BandCombinationList_UplinkTxSwitch_v1650  `optional,ext-7`
	extendedBand_n77_r16                                   *RF_Parameters_extendedBand_n77_r16        `optional,ext-8`
	supportedBandCombinationList_UplinkTxSwitch_v1670      *BandCombinationList_UplinkTxSwitch_v1670  `optional,ext-9`
	supportedBandCombinationList_v1680                     *BandCombinationList_v1680                 `optional,ext-10`
	supportedBandCombinationList_v1690                     *BandCombinationList_v1690                 `optional,ext-11`
	supportedBandCombinationList_UplinkTxSwitch_v1690      *BandCombinationList_UplinkTxSwitch_v1690  `optional,ext-11`
	supportedBandCombinationList_v1700                     *BandCombinationList_v1700                 `optional,ext-12`
	supportedBandCombinationList_UplinkTxSwitch_v1700      *BandCombinationList_UplinkTxSwitch_v1700  `optional,ext-12`
	supportedBandCombinationListSL_RelayDiscovery_r17      *[]byte                                    `optional,ext-12`
	supportedBandCombinationListSL_NonRelayDiscovery_r17   *[]byte                                    `optional,ext-12`
	supportedBandCombinationListSidelinkEUTRA_NR_v1710     *BandCombinationListSidelinkEUTRA_NR_v1710 `optional,ext-12`
	sidelinkRequested_r17                                  *RF_Parameters_sidelinkRequested_r17       `optional,ext-12`
	extendedBand_n77_2_r17                                 *RF_Parameters_extendedBand_n77_2_r17      `optional,ext-12`
	supportedBandCombinationList_v1720                     *BandCombinationList_v1720                 `optional,ext-13`
	supportedBandCombinationList_UplinkTxSwitch_v1720      *BandCombinationList_UplinkTxSwitch_v1720  `optional,ext-13`
	supportedBandCombinationList_v1730                     *BandCombinationList_v1730                 `optional,ext-14`
	supportedBandCombinationList_UplinkTxSwitch_v1730      *BandCombinationList_UplinkTxSwitch_v1730  `optional,ext-14`
	supportedBandCombinationListSL_RelayDiscovery_v1730    *BandCombinationListSL_Discovery_r17       `optional,ext-14`
	supportedBandCombinationListSL_NonRelayDiscovery_v1730 *BandCombinationListSL_Discovery_r17       `optional,ext-14`
}

func (ie *RF_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.supportedBandCombinationList_v1540 != nil || ie.srs_SwitchingTimeRequested != nil || ie.supportedBandCombinationList_v1550 != nil || ie.supportedBandCombinationList_v1560 != nil || ie.supportedBandCombinationList_v1610 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_r16 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil || ie.supportedBandCombinationList_v1630 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil || ie.supportedBandCombinationList_v1640 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil || ie.supportedBandCombinationList_v1650 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1650 != nil || ie.extendedBand_n77_r16 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil || ie.supportedBandCombinationList_v1680 != nil || ie.supportedBandCombinationList_v1690 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1690 != nil || ie.supportedBandCombinationList_v1700 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.supportedBandCombinationListSL_RelayDiscovery_r17 != nil || ie.supportedBandCombinationListSL_NonRelayDiscovery_r17 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil || ie.sidelinkRequested_r17 != nil || ie.extendedBand_n77_2_r17 != nil || ie.supportedBandCombinationList_v1720 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil || ie.supportedBandCombinationList_v1730 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil || ie.supportedBandCombinationListSL_RelayDiscovery_v1730 != nil || ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.supportedBandCombinationList != nil, ie.appliedFreqBandListFilter != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_supportedBandListNR := utils.NewSequence[*BandNR]([]*BandNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.supportedBandListNR {
		tmp_supportedBandListNR.Value = append(tmp_supportedBandListNR.Value, &i)
	}
	if err = tmp_supportedBandListNR.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandListNR", err)
	}
	if ie.supportedBandCombinationList != nil {
		if err = ie.supportedBandCombinationList.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList", err)
		}
	}
	if ie.appliedFreqBandListFilter != nil {
		if err = ie.appliedFreqBandListFilter.Encode(w); err != nil {
			return utils.WrapError("Encode appliedFreqBandListFilter", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 14 bits for 14 extension groups
		extBitmap := []bool{ie.supportedBandCombinationList_v1540 != nil || ie.srs_SwitchingTimeRequested != nil, ie.supportedBandCombinationList_v1550 != nil, ie.supportedBandCombinationList_v1560 != nil, ie.supportedBandCombinationList_v1610 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_r16 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil, ie.supportedBandCombinationList_v1630 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil, ie.supportedBandCombinationList_v1640 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil, ie.supportedBandCombinationList_v1650 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1650 != nil, ie.extendedBand_n77_r16 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil, ie.supportedBandCombinationList_v1680 != nil, ie.supportedBandCombinationList_v1690 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1690 != nil, ie.supportedBandCombinationList_v1700 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.supportedBandCombinationListSL_RelayDiscovery_r17 != nil || ie.supportedBandCombinationListSL_NonRelayDiscovery_r17 != nil || ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil || ie.sidelinkRequested_r17 != nil || ie.extendedBand_n77_2_r17 != nil, ie.supportedBandCombinationList_v1720 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil, ie.supportedBandCombinationList_v1730 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil || ie.supportedBandCombinationListSL_RelayDiscovery_v1730 != nil || ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RF_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.supportedBandCombinationList_v1540 != nil, ie.srs_SwitchingTimeRequested != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1540 optional
			if ie.supportedBandCombinationList_v1540 != nil {
				if err = ie.supportedBandCombinationList_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1540", err)
				}
			}
			// encode srs_SwitchingTimeRequested optional
			if ie.srs_SwitchingTimeRequested != nil {
				if err = ie.srs_SwitchingTimeRequested.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_SwitchingTimeRequested", err)
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
			optionals_ext_2 := []bool{ie.supportedBandCombinationList_v1550 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1550 optional
			if ie.supportedBandCombinationList_v1550 != nil {
				if err = ie.supportedBandCombinationList_v1550.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1550", err)
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
			optionals_ext_3 := []bool{ie.supportedBandCombinationList_v1560 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1560 optional
			if ie.supportedBandCombinationList_v1560 != nil {
				if err = ie.supportedBandCombinationList_v1560.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1560", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.supportedBandCombinationList_v1610 != nil, ie.supportedBandCombinationListSidelinkEUTRA_NR_r16 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1610 optional
			if ie.supportedBandCombinationList_v1610 != nil {
				if err = ie.supportedBandCombinationList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1610", err)
				}
			}
			// encode supportedBandCombinationListSidelinkEUTRA_NR_r16 optional
			if ie.supportedBandCombinationListSidelinkEUTRA_NR_r16 != nil {
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSidelinkEUTRA_NR_r16", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_r16 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.supportedBandCombinationList_v1630 != nil, ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1630 optional
			if ie.supportedBandCombinationList_v1630 != nil {
				if err = ie.supportedBandCombinationList_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1630", err)
				}
			}
			// encode supportedBandCombinationListSidelinkEUTRA_NR_v1630 optional
			if ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil {
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSidelinkEUTRA_NR_v1630", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1630 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1630", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.supportedBandCombinationList_v1640 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1640 optional
			if ie.supportedBandCombinationList_v1640 != nil {
				if err = ie.supportedBandCombinationList_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1640", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1640 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1640", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.supportedBandCombinationList_v1650 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1650 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1650 optional
			if ie.supportedBandCombinationList_v1650 != nil {
				if err = ie.supportedBandCombinationList_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1650", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1650 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1650 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1650", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.extendedBand_n77_r16 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode extendedBand_n77_r16 optional
			if ie.extendedBand_n77_r16 != nil {
				if err = ie.extendedBand_n77_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedBand_n77_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 9
		if extBitmap[8] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_UplinkTxSwitch_v1670 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1670.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1670", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 10
		if extBitmap[9] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.supportedBandCombinationList_v1680 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1680 optional
			if ie.supportedBandCombinationList_v1680 != nil {
				if err = ie.supportedBandCombinationList_v1680.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1680", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 11
		if extBitmap[10] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.supportedBandCombinationList_v1690 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1690 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1690 optional
			if ie.supportedBandCombinationList_v1690 != nil {
				if err = ie.supportedBandCombinationList_v1690.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1690", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1690 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1690 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1690.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1690", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 12
		if extBitmap[11] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 12
			optionals_ext_12 := []bool{ie.supportedBandCombinationList_v1700 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil, ie.supportedBandCombinationListSL_RelayDiscovery_r17 != nil, ie.supportedBandCombinationListSL_NonRelayDiscovery_r17 != nil, ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil, ie.sidelinkRequested_r17 != nil, ie.extendedBand_n77_2_r17 != nil}
			for _, bit := range optionals_ext_12 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1700 optional
			if ie.supportedBandCombinationList_v1700 != nil {
				if err = ie.supportedBandCombinationList_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1700", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1700", err)
				}
			}
			// encode supportedBandCombinationListSL_RelayDiscovery_r17 optional
			if ie.supportedBandCombinationListSL_RelayDiscovery_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.supportedBandCombinationListSL_RelayDiscovery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSL_RelayDiscovery_r17", err)
				}
			}
			// encode supportedBandCombinationListSL_NonRelayDiscovery_r17 optional
			if ie.supportedBandCombinationListSL_NonRelayDiscovery_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.supportedBandCombinationListSL_NonRelayDiscovery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSL_NonRelayDiscovery_r17", err)
				}
			}
			// encode supportedBandCombinationListSidelinkEUTRA_NR_v1710 optional
			if ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil {
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSidelinkEUTRA_NR_v1710", err)
				}
			}
			// encode sidelinkRequested_r17 optional
			if ie.sidelinkRequested_r17 != nil {
				if err = ie.sidelinkRequested_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sidelinkRequested_r17", err)
				}
			}
			// encode extendedBand_n77_2_r17 optional
			if ie.extendedBand_n77_2_r17 != nil {
				if err = ie.extendedBand_n77_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedBand_n77_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 13
		if extBitmap[12] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 13
			optionals_ext_13 := []bool{ie.supportedBandCombinationList_v1720 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil}
			for _, bit := range optionals_ext_13 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1720 optional
			if ie.supportedBandCombinationList_v1720 != nil {
				if err = ie.supportedBandCombinationList_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1720", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1720 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 14
		if extBitmap[13] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 14
			optionals_ext_14 := []bool{ie.supportedBandCombinationList_v1730 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil, ie.supportedBandCombinationListSL_RelayDiscovery_v1730 != nil, ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil}
			for _, bit := range optionals_ext_14 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_v1730 optional
			if ie.supportedBandCombinationList_v1730 != nil {
				if err = ie.supportedBandCombinationList_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1730", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1730", err)
				}
			}
			// encode supportedBandCombinationListSL_RelayDiscovery_v1730 optional
			if ie.supportedBandCombinationListSL_RelayDiscovery_v1730 != nil {
				if err = ie.supportedBandCombinationListSL_RelayDiscovery_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSL_RelayDiscovery_v1730", err)
				}
			}
			// encode supportedBandCombinationListSL_NonRelayDiscovery_v1730 optional
			if ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil {
				if err = ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListSL_NonRelayDiscovery_v1730", err)
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

func (ie *RF_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationListPresent bool
	if supportedBandCombinationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var appliedFreqBandListFilterPresent bool
	if appliedFreqBandListFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_supportedBandListNR := utils.NewSequence[*BandNR]([]*BandNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_supportedBandListNR := func() *BandNR {
		return new(BandNR)
	}
	if err = tmp_supportedBandListNR.Decode(r, fn_supportedBandListNR); err != nil {
		return utils.WrapError("Decode supportedBandListNR", err)
	}
	ie.supportedBandListNR = []BandNR{}
	for _, i := range tmp_supportedBandListNR.Value {
		ie.supportedBandListNR = append(ie.supportedBandListNR, *i)
	}
	if supportedBandCombinationListPresent {
		ie.supportedBandCombinationList = new(BandCombinationList)
		if err = ie.supportedBandCombinationList.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList", err)
		}
	}
	if appliedFreqBandListFilterPresent {
		ie.appliedFreqBandListFilter = new(FreqBandList)
		if err = ie.appliedFreqBandListFilter.Decode(r); err != nil {
			return utils.WrapError("Decode appliedFreqBandListFilter", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 14 bits for 14 extension groups
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

			supportedBandCombinationList_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_SwitchingTimeRequestedPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1540 optional
			if supportedBandCombinationList_v1540Present {
				ie.supportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
				if err = ie.supportedBandCombinationList_v1540.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1540", err)
				}
			}
			// decode srs_SwitchingTimeRequested optional
			if srs_SwitchingTimeRequestedPresent {
				ie.srs_SwitchingTimeRequested = new(RF_Parameters_srs_SwitchingTimeRequested)
				if err = ie.srs_SwitchingTimeRequested.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_SwitchingTimeRequested", err)
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

			supportedBandCombinationList_v1550Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1550 optional
			if supportedBandCombinationList_v1550Present {
				ie.supportedBandCombinationList_v1550 = new(BandCombinationList_v1550)
				if err = ie.supportedBandCombinationList_v1550.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1550", err)
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

			supportedBandCombinationList_v1560Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1560 optional
			if supportedBandCombinationList_v1560Present {
				ie.supportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
				if err = ie.supportedBandCombinationList_v1560.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1560", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSidelinkEUTRA_NR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1610 optional
			if supportedBandCombinationList_v1610Present {
				ie.supportedBandCombinationList_v1610 = new(BandCombinationList_v1610)
				if err = ie.supportedBandCombinationList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1610", err)
				}
			}
			// decode supportedBandCombinationListSidelinkEUTRA_NR_r16 optional
			if supportedBandCombinationListSidelinkEUTRA_NR_r16Present {
				ie.supportedBandCombinationListSidelinkEUTRA_NR_r16 = new(BandCombinationListSidelinkEUTRA_NR_r16)
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSidelinkEUTRA_NR_r16", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_r16 optional
			if supportedBandCombinationList_UplinkTxSwitch_r16Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_r16 = new(BandCombinationList_UplinkTxSwitch_r16)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSidelinkEUTRA_NR_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1630 optional
			if supportedBandCombinationList_v1630Present {
				ie.supportedBandCombinationList_v1630 = new(BandCombinationList_v1630)
				if err = ie.supportedBandCombinationList_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1630", err)
				}
			}
			// decode supportedBandCombinationListSidelinkEUTRA_NR_v1630 optional
			if supportedBandCombinationListSidelinkEUTRA_NR_v1630Present {
				ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630 = new(BandCombinationListSidelinkEUTRA_NR_v1630)
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSidelinkEUTRA_NR_v1630", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1630 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1630Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1630 = new(BandCombinationList_UplinkTxSwitch_v1630)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1630", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1640 optional
			if supportedBandCombinationList_v1640Present {
				ie.supportedBandCombinationList_v1640 = new(BandCombinationList_v1640)
				if err = ie.supportedBandCombinationList_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1640", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1640 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1640Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1640 = new(BandCombinationList_UplinkTxSwitch_v1640)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1640", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1650 optional
			if supportedBandCombinationList_v1650Present {
				ie.supportedBandCombinationList_v1650 = new(BandCombinationList_v1650)
				if err = ie.supportedBandCombinationList_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1650", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1650 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1650Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1650 = new(BandCombinationList_UplinkTxSwitch_v1650)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1650", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			extendedBand_n77_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode extendedBand_n77_r16 optional
			if extendedBand_n77_r16Present {
				ie.extendedBand_n77_r16 = new(RF_Parameters_extendedBand_n77_r16)
				if err = ie.extendedBand_n77_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedBand_n77_r16", err)
				}
			}
		}
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_UplinkTxSwitch_v1670Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1670 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1670Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1670 = new(BandCombinationList_UplinkTxSwitch_v1670)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1670.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1670", err)
				}
			}
		}
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1680Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1680 optional
			if supportedBandCombinationList_v1680Present {
				ie.supportedBandCombinationList_v1680 = new(BandCombinationList_v1680)
				if err = ie.supportedBandCombinationList_v1680.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1680", err)
				}
			}
		}
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1690Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1690Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1690 optional
			if supportedBandCombinationList_v1690Present {
				ie.supportedBandCombinationList_v1690 = new(BandCombinationList_v1690)
				if err = ie.supportedBandCombinationList_v1690.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1690", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1690 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1690Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1690 = new(BandCombinationList_UplinkTxSwitch_v1690)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1690.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1690", err)
				}
			}
		}
		// decode extension group 12
		if len(extBitmap) > 11 && extBitmap[11] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSL_RelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSL_NonRelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSidelinkEUTRA_NR_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sidelinkRequested_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedBand_n77_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1700 optional
			if supportedBandCombinationList_v1700Present {
				ie.supportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
				if err = ie.supportedBandCombinationList_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1700", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1700Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1700 = new(BandCombinationList_UplinkTxSwitch_v1700)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1700", err)
				}
			}
			// decode supportedBandCombinationListSL_RelayDiscovery_r17 optional
			if supportedBandCombinationListSL_RelayDiscovery_r17Present {
				var tmp_os_supportedBandCombinationListSL_RelayDiscovery_r17 []byte
				if tmp_os_supportedBandCombinationListSL_RelayDiscovery_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSL_RelayDiscovery_r17", err)
				}
				ie.supportedBandCombinationListSL_RelayDiscovery_r17 = &tmp_os_supportedBandCombinationListSL_RelayDiscovery_r17
			}
			// decode supportedBandCombinationListSL_NonRelayDiscovery_r17 optional
			if supportedBandCombinationListSL_NonRelayDiscovery_r17Present {
				var tmp_os_supportedBandCombinationListSL_NonRelayDiscovery_r17 []byte
				if tmp_os_supportedBandCombinationListSL_NonRelayDiscovery_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSL_NonRelayDiscovery_r17", err)
				}
				ie.supportedBandCombinationListSL_NonRelayDiscovery_r17 = &tmp_os_supportedBandCombinationListSL_NonRelayDiscovery_r17
			}
			// decode supportedBandCombinationListSidelinkEUTRA_NR_v1710 optional
			if supportedBandCombinationListSidelinkEUTRA_NR_v1710Present {
				ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710 = new(BandCombinationListSidelinkEUTRA_NR_v1710)
				if err = ie.supportedBandCombinationListSidelinkEUTRA_NR_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSidelinkEUTRA_NR_v1710", err)
				}
			}
			// decode sidelinkRequested_r17 optional
			if sidelinkRequested_r17Present {
				ie.sidelinkRequested_r17 = new(RF_Parameters_sidelinkRequested_r17)
				if err = ie.sidelinkRequested_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sidelinkRequested_r17", err)
				}
			}
			// decode extendedBand_n77_2_r17 optional
			if extendedBand_n77_2_r17Present {
				ie.extendedBand_n77_2_r17 = new(RF_Parameters_extendedBand_n77_2_r17)
				if err = ie.extendedBand_n77_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedBand_n77_2_r17", err)
				}
			}
		}
		// decode extension group 13
		if len(extBitmap) > 12 && extBitmap[12] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1720 optional
			if supportedBandCombinationList_v1720Present {
				ie.supportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
				if err = ie.supportedBandCombinationList_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1720", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1720 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1720Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1720 = new(BandCombinationList_UplinkTxSwitch_v1720)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1720", err)
				}
			}
		}
		// decode extension group 14
		if len(extBitmap) > 13 && extBitmap[13] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSL_RelayDiscovery_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListSL_NonRelayDiscovery_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_v1730 optional
			if supportedBandCombinationList_v1730Present {
				ie.supportedBandCombinationList_v1730 = new(BandCombinationList_v1730)
				if err = ie.supportedBandCombinationList_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1730", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1730Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1730 = new(BandCombinationList_UplinkTxSwitch_v1730)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1730", err)
				}
			}
			// decode supportedBandCombinationListSL_RelayDiscovery_v1730 optional
			if supportedBandCombinationListSL_RelayDiscovery_v1730Present {
				ie.supportedBandCombinationListSL_RelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
				if err = ie.supportedBandCombinationListSL_RelayDiscovery_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSL_RelayDiscovery_v1730", err)
				}
			}
			// decode supportedBandCombinationListSL_NonRelayDiscovery_v1730 optional
			if supportedBandCombinationListSL_NonRelayDiscovery_v1730Present {
				ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
				if err = ie.supportedBandCombinationListSL_NonRelayDiscovery_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListSL_NonRelayDiscovery_v1730", err)
				}
			}
		}
	}
	return nil
}
