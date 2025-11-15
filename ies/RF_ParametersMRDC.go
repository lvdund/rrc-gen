package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_ParametersMRDC struct {
	supportedBandCombinationList                      *BandCombinationList                                           `optional`
	appliedFreqBandListFilter                         *FreqBandList                                                  `optional`
	srs_SwitchingTimeRequested                        *RF_ParametersMRDC_srs_SwitchingTimeRequested                  `optional,ext-1`
	supportedBandCombinationList_v1540                *BandCombinationList_v1540                                     `optional,ext-7`
	supportedBandCombinationList_v1550                *BandCombinationList_v1550                                     `optional,ext-2`
	supportedBandCombinationList_v1560                *BandCombinationList_v1560                                     `optional,ext-7`
	supportedBandCombinationListNEDC_Only             *BandCombinationList                                           `optional,ext-3`
	supportedBandCombinationList_v1570                *BandCombinationList_v1570                                     `optional,ext-7`
	supportedBandCombinationList_v1580                *BandCombinationList_v1580                                     `optional,ext-7`
	supportedBandCombinationList_v1590                *BandCombinationList_v1590                                     `optional,ext-7`
	supportedBandCombinationListNEDC_Only_v15a0       *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v15a0 `optional,ext-7`
	supportedBandCombinationList_v1610                *BandCombinationList_v1610                                     `optional,ext-8`
	supportedBandCombinationListNEDC_Only_v1610       *BandCombinationList_v1610                                     `optional,ext-8`
	supportedBandCombinationList_UplinkTxSwitch_r16   *BandCombinationList_UplinkTxSwitch_r16                        `optional,ext-8`
	supportedBandCombinationList_v1630                *BandCombinationList_v1630                                     `optional,ext-9`
	supportedBandCombinationListNEDC_Only_v1630       *BandCombinationList_v1630                                     `optional,ext-9`
	supportedBandCombinationList_UplinkTxSwitch_v1630 *BandCombinationList_UplinkTxSwitch_v1630                      `optional,ext-9`
	supportedBandCombinationList_v1640                *BandCombinationList_v1640                                     `optional,ext-10`
	supportedBandCombinationListNEDC_Only_v1640       *BandCombinationList_v1640                                     `optional,ext-10`
	supportedBandCombinationList_UplinkTxSwitch_v1640 *BandCombinationList_UplinkTxSwitch_v1640                      `optional,ext-10`
	supportedBandCombinationList_UplinkTxSwitch_v1670 *BandCombinationList_UplinkTxSwitch_v1670                      `optional,ext-11`
	supportedBandCombinationList_v1700                *BandCombinationList_v1700                                     `optional,ext-13`
	supportedBandCombinationList_UplinkTxSwitch_v1700 *BandCombinationList_UplinkTxSwitch_v1700                      `optional,ext-12`
	supportedBandCombinationList_v1720                *BandCombinationList_v1720                                     `optional,ext-13`
	supportedBandCombinationListNEDC_Only_v1720       *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720 `optional,ext-13`
	supportedBandCombinationList_UplinkTxSwitch_v1720 *BandCombinationList_UplinkTxSwitch_v1720                      `optional,ext-13`
	supportedBandCombinationList_v1730                *BandCombinationList_v1730                                     `optional,ext-14`
	supportedBandCombinationListNEDC_Only_v1730       *BandCombinationList_v1730                                     `optional,ext-14`
	supportedBandCombinationList_UplinkTxSwitch_v1730 *BandCombinationList_UplinkTxSwitch_v1730                      `optional,ext-14`
}

func (ie *RF_ParametersMRDC) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.srs_SwitchingTimeRequested != nil || ie.supportedBandCombinationList_v1550 != nil || ie.supportedBandCombinationListNEDC_Only != nil || ie.supportedBandCombinationList_v1540 != nil || ie.supportedBandCombinationList_v1560 != nil || ie.supportedBandCombinationList_v1570 != nil || ie.supportedBandCombinationList_v1580 != nil || ie.supportedBandCombinationList_v1590 != nil || ie.supportedBandCombinationListNEDC_Only_v15a0 != nil || ie.supportedBandCombinationList_v1610 != nil || ie.supportedBandCombinationListNEDC_Only_v1610 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil || ie.supportedBandCombinationList_v1630 != nil || ie.supportedBandCombinationListNEDC_Only_v1630 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil || ie.supportedBandCombinationList_v1640 != nil || ie.supportedBandCombinationListNEDC_Only_v1640 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.supportedBandCombinationList_v1700 != nil || ie.supportedBandCombinationList_v1720 != nil || ie.supportedBandCombinationListNEDC_Only_v1720 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil || ie.supportedBandCombinationList_v1730 != nil || ie.supportedBandCombinationListNEDC_Only_v1730 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.supportedBandCombinationList != nil, ie.appliedFreqBandListFilter != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
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
		// Extension bitmap: 11 bits for 11 extension groups
		extBitmap := []bool{ie.srs_SwitchingTimeRequested != nil, ie.supportedBandCombinationList_v1550 != nil, ie.supportedBandCombinationListNEDC_Only != nil, ie.supportedBandCombinationList_v1540 != nil || ie.supportedBandCombinationList_v1560 != nil || ie.supportedBandCombinationList_v1570 != nil || ie.supportedBandCombinationList_v1580 != nil || ie.supportedBandCombinationList_v1590 != nil || ie.supportedBandCombinationListNEDC_Only_v15a0 != nil, ie.supportedBandCombinationList_v1610 != nil || ie.supportedBandCombinationListNEDC_Only_v1610 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil, ie.supportedBandCombinationList_v1630 != nil || ie.supportedBandCombinationListNEDC_Only_v1630 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil, ie.supportedBandCombinationList_v1640 != nil || ie.supportedBandCombinationListNEDC_Only_v1640 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil, ie.supportedBandCombinationList_v1700 != nil || ie.supportedBandCombinationList_v1720 != nil || ie.supportedBandCombinationListNEDC_Only_v1720 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil, ie.supportedBandCombinationList_v1730 != nil || ie.supportedBandCombinationListNEDC_Only_v1730 != nil || ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RF_ParametersMRDC", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.srs_SwitchingTimeRequested != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			optionals_ext_3 := []bool{ie.supportedBandCombinationListNEDC_Only != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationListNEDC_Only optional
			if ie.supportedBandCombinationListNEDC_Only != nil {
				if err = ie.supportedBandCombinationListNEDC_Only.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only", err)
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
			optionals_ext_7 := []bool{ie.supportedBandCombinationList_v1540 != nil, ie.supportedBandCombinationList_v1560 != nil, ie.supportedBandCombinationList_v1570 != nil, ie.supportedBandCombinationList_v1580 != nil, ie.supportedBandCombinationList_v1590 != nil, ie.supportedBandCombinationListNEDC_Only_v15a0 != nil}
			for _, bit := range optionals_ext_7 {
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
			// encode supportedBandCombinationList_v1560 optional
			if ie.supportedBandCombinationList_v1560 != nil {
				if err = ie.supportedBandCombinationList_v1560.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1560", err)
				}
			}
			// encode supportedBandCombinationList_v1570 optional
			if ie.supportedBandCombinationList_v1570 != nil {
				if err = ie.supportedBandCombinationList_v1570.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1570", err)
				}
			}
			// encode supportedBandCombinationList_v1580 optional
			if ie.supportedBandCombinationList_v1580 != nil {
				if err = ie.supportedBandCombinationList_v1580.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1580", err)
				}
			}
			// encode supportedBandCombinationList_v1590 optional
			if ie.supportedBandCombinationList_v1590 != nil {
				if err = ie.supportedBandCombinationList_v1590.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1590", err)
				}
			}
			// encode supportedBandCombinationListNEDC_Only_v15a0 optional
			if ie.supportedBandCombinationListNEDC_Only_v15a0 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v15a0.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v15a0", err)
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
			optionals_ext_8 := []bool{ie.supportedBandCombinationList_v1610 != nil, ie.supportedBandCombinationListNEDC_Only_v1610 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_r16 != nil}
			for _, bit := range optionals_ext_8 {
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
			// encode supportedBandCombinationListNEDC_Only_v1610 optional
			if ie.supportedBandCombinationListNEDC_Only_v1610 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v1610", err)
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

		// encode extension group 9
		if extBitmap[8] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.supportedBandCombinationList_v1630 != nil, ie.supportedBandCombinationListNEDC_Only_v1630 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1630 != nil}
			for _, bit := range optionals_ext_9 {
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
			// encode supportedBandCombinationListNEDC_Only_v1630 optional
			if ie.supportedBandCombinationListNEDC_Only_v1630 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v1630", err)
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

		// encode extension group 10
		if extBitmap[9] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.supportedBandCombinationList_v1640 != nil, ie.supportedBandCombinationListNEDC_Only_v1640 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1640 != nil}
			for _, bit := range optionals_ext_10 {
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
			// encode supportedBandCombinationListNEDC_Only_v1640 optional
			if ie.supportedBandCombinationListNEDC_Only_v1640 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v1640", err)
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

		// encode extension group 11
		if extBitmap[10] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.supportedBandCombinationList_UplinkTxSwitch_v1670 != nil}
			for _, bit := range optionals_ext_11 {
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

		// encode extension group 12
		if extBitmap[11] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 12
			optionals_ext_12 := []bool{ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil}
			for _, bit := range optionals_ext_12 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1700 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1700", err)
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
			optionals_ext_13 := []bool{ie.supportedBandCombinationList_v1700 != nil, ie.supportedBandCombinationList_v1720 != nil, ie.supportedBandCombinationListNEDC_Only_v1720 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1720 != nil}
			for _, bit := range optionals_ext_13 {
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
			// encode supportedBandCombinationList_v1720 optional
			if ie.supportedBandCombinationList_v1720 != nil {
				if err = ie.supportedBandCombinationList_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_v1720", err)
				}
			}
			// encode supportedBandCombinationListNEDC_Only_v1720 optional
			if ie.supportedBandCombinationListNEDC_Only_v1720 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v1720", err)
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
			optionals_ext_14 := []bool{ie.supportedBandCombinationList_v1730 != nil, ie.supportedBandCombinationListNEDC_Only_v1730 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil}
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
			// encode supportedBandCombinationListNEDC_Only_v1730 optional
			if ie.supportedBandCombinationListNEDC_Only_v1730 != nil {
				if err = ie.supportedBandCombinationListNEDC_Only_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationListNEDC_Only_v1730", err)
				}
			}
			// encode supportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if ie.supportedBandCombinationList_UplinkTxSwitch_v1730 != nil {
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v1730", err)
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

func (ie *RF_ParametersMRDC) Decode(r *uper.UperReader) error {
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
		// Read extension bitmap: 11 bits for 11 extension groups
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

			srs_SwitchingTimeRequestedPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srs_SwitchingTimeRequested optional
			if srs_SwitchingTimeRequestedPresent {
				ie.srs_SwitchingTimeRequested = new(RF_ParametersMRDC_srs_SwitchingTimeRequested)
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

			supportedBandCombinationListNEDC_OnlyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationListNEDC_Only optional
			if supportedBandCombinationListNEDC_OnlyPresent {
				ie.supportedBandCombinationListNEDC_Only = new(BandCombinationList)
				if err = ie.supportedBandCombinationListNEDC_Only.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only", err)
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

			supportedBandCombinationList_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_v1560Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_v1570Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_v1580Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_v1590Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListNEDC_Only_v15a0Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationList_v1560 optional
			if supportedBandCombinationList_v1560Present {
				ie.supportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
				if err = ie.supportedBandCombinationList_v1560.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1560", err)
				}
			}
			// decode supportedBandCombinationList_v1570 optional
			if supportedBandCombinationList_v1570Present {
				ie.supportedBandCombinationList_v1570 = new(BandCombinationList_v1570)
				if err = ie.supportedBandCombinationList_v1570.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1570", err)
				}
			}
			// decode supportedBandCombinationList_v1580 optional
			if supportedBandCombinationList_v1580Present {
				ie.supportedBandCombinationList_v1580 = new(BandCombinationList_v1580)
				if err = ie.supportedBandCombinationList_v1580.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1580", err)
				}
			}
			// decode supportedBandCombinationList_v1590 optional
			if supportedBandCombinationList_v1590Present {
				ie.supportedBandCombinationList_v1590 = new(BandCombinationList_v1590)
				if err = ie.supportedBandCombinationList_v1590.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1590", err)
				}
			}
			// decode supportedBandCombinationListNEDC_Only_v15a0 optional
			if supportedBandCombinationListNEDC_Only_v15a0Present {
				ie.supportedBandCombinationListNEDC_Only_v15a0 = new(RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v15a0)
				if err = ie.supportedBandCombinationListNEDC_Only_v15a0.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v15a0", err)
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

			supportedBandCombinationList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListNEDC_Only_v1610Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationListNEDC_Only_v1610 optional
			if supportedBandCombinationListNEDC_Only_v1610Present {
				ie.supportedBandCombinationListNEDC_Only_v1610 = new(BandCombinationList_v1610)
				if err = ie.supportedBandCombinationListNEDC_Only_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v1610", err)
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
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListNEDC_Only_v1630Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationListNEDC_Only_v1630 optional
			if supportedBandCombinationListNEDC_Only_v1630Present {
				ie.supportedBandCombinationListNEDC_Only_v1630 = new(BandCombinationList_v1630)
				if err = ie.supportedBandCombinationListNEDC_Only_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v1630", err)
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
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListNEDC_Only_v1640Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationListNEDC_Only_v1640 optional
			if supportedBandCombinationListNEDC_Only_v1640Present {
				ie.supportedBandCombinationListNEDC_Only_v1640 = new(BandCombinationList_v1640)
				if err = ie.supportedBandCombinationListNEDC_Only_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v1640", err)
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
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
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
		// decode extension group 12
		if len(extBitmap) > 11 && extBitmap[11] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedBandCombinationList_UplinkTxSwitch_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1700Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1700 = new(BandCombinationList_UplinkTxSwitch_v1700)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1700", err)
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

			supportedBandCombinationList_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationListNEDC_Only_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1720Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationList_v1720 optional
			if supportedBandCombinationList_v1720Present {
				ie.supportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
				if err = ie.supportedBandCombinationList_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_v1720", err)
				}
			}
			// decode supportedBandCombinationListNEDC_Only_v1720 optional
			if supportedBandCombinationListNEDC_Only_v1720Present {
				ie.supportedBandCombinationListNEDC_Only_v1720 = new(RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720)
				if err = ie.supportedBandCombinationListNEDC_Only_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v1720", err)
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
			supportedBandCombinationListNEDC_Only_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedBandCombinationList_UplinkTxSwitch_v1730Present, err := extReader.ReadBool()
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
			// decode supportedBandCombinationListNEDC_Only_v1730 optional
			if supportedBandCombinationListNEDC_Only_v1730Present {
				ie.supportedBandCombinationListNEDC_Only_v1730 = new(BandCombinationList_v1730)
				if err = ie.supportedBandCombinationListNEDC_Only_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationListNEDC_Only_v1730", err)
				}
			}
			// decode supportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if supportedBandCombinationList_UplinkTxSwitch_v1730Present {
				ie.supportedBandCombinationList_UplinkTxSwitch_v1730 = new(BandCombinationList_UplinkTxSwitch_v1730)
				if err = ie.supportedBandCombinationList_UplinkTxSwitch_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v1730", err)
				}
			}
		}
	}
	return nil
}
