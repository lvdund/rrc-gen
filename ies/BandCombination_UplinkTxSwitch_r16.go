package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_r16 struct {
	bandCombination_r16                        BandCombination                                                                `madatory`
	bandCombination_v1540                      *BandCombination_v1540                                                         `optional`
	bandCombination_v1560                      *BandCombination_v1560                                                         `optional`
	bandCombination_v1570                      *BandCombination_v1570                                                         `optional`
	bandCombination_v1580                      *BandCombination_v1580                                                         `optional`
	bandCombination_v1590                      *BandCombination_v1590                                                         `optional`
	bandCombination_v1610                      *BandCombination_v1610                                                         `optional`
	supportedBandPairListNR_r16                []ULTxSwitchingBandPair_r16                                                    `lb:1,ub:maxULTxSwitchingBandPairs,madatory`
	uplinkTxSwitching_OptionSupport_r16        *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_OptionSupport_r16        `optional`
	uplinkTxSwitching_PowerBoosting_r16        *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PowerBoosting_r16        `optional`
	uplinkTxSwitching_PUSCH_TransCoherence_r16 *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16 `optional,ext-1`
}

func (ie *BandCombination_UplinkTxSwitch_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.uplinkTxSwitching_PUSCH_TransCoherence_r16 != nil
	preambleBits := []bool{hasExtensions, ie.bandCombination_v1540 != nil, ie.bandCombination_v1560 != nil, ie.bandCombination_v1570 != nil, ie.bandCombination_v1580 != nil, ie.bandCombination_v1590 != nil, ie.bandCombination_v1610 != nil, ie.uplinkTxSwitching_OptionSupport_r16 != nil, ie.uplinkTxSwitching_PowerBoosting_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bandCombination_r16.Encode(w); err != nil {
		return utils.WrapError("Encode bandCombination_r16", err)
	}
	if ie.bandCombination_v1540 != nil {
		if err = ie.bandCombination_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1540", err)
		}
	}
	if ie.bandCombination_v1560 != nil {
		if err = ie.bandCombination_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1560", err)
		}
	}
	if ie.bandCombination_v1570 != nil {
		if err = ie.bandCombination_v1570.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1570", err)
		}
	}
	if ie.bandCombination_v1580 != nil {
		if err = ie.bandCombination_v1580.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1580", err)
		}
	}
	if ie.bandCombination_v1590 != nil {
		if err = ie.bandCombination_v1590.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1590", err)
		}
	}
	if ie.bandCombination_v1610 != nil {
		if err = ie.bandCombination_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1610", err)
		}
	}
	tmp_supportedBandPairListNR_r16 := utils.NewSequence[*ULTxSwitchingBandPair_r16]([]*ULTxSwitchingBandPair_r16{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
	for _, i := range ie.supportedBandPairListNR_r16 {
		tmp_supportedBandPairListNR_r16.Value = append(tmp_supportedBandPairListNR_r16.Value, &i)
	}
	if err = tmp_supportedBandPairListNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandPairListNR_r16", err)
	}
	if ie.uplinkTxSwitching_OptionSupport_r16 != nil {
		if err = ie.uplinkTxSwitching_OptionSupport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitching_OptionSupport_r16", err)
		}
	}
	if ie.uplinkTxSwitching_PowerBoosting_r16 != nil {
		if err = ie.uplinkTxSwitching_PowerBoosting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitching_PowerBoosting_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.uplinkTxSwitching_PUSCH_TransCoherence_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandCombination_UplinkTxSwitch_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.uplinkTxSwitching_PUSCH_TransCoherence_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode uplinkTxSwitching_PUSCH_TransCoherence_r16 optional
			if ie.uplinkTxSwitching_PUSCH_TransCoherence_r16 != nil {
				if err = ie.uplinkTxSwitching_PUSCH_TransCoherence_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
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

func (ie *BandCombination_UplinkTxSwitch_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1540Present bool
	if bandCombination_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1560Present bool
	if bandCombination_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1570Present bool
	if bandCombination_v1570Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1580Present bool
	if bandCombination_v1580Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1590Present bool
	if bandCombination_v1590Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandCombination_v1610Present bool
	if bandCombination_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkTxSwitching_OptionSupport_r16Present bool
	if uplinkTxSwitching_OptionSupport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkTxSwitching_PowerBoosting_r16Present bool
	if uplinkTxSwitching_PowerBoosting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bandCombination_r16.Decode(r); err != nil {
		return utils.WrapError("Decode bandCombination_r16", err)
	}
	if bandCombination_v1540Present {
		ie.bandCombination_v1540 = new(BandCombination_v1540)
		if err = ie.bandCombination_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1540", err)
		}
	}
	if bandCombination_v1560Present {
		ie.bandCombination_v1560 = new(BandCombination_v1560)
		if err = ie.bandCombination_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1560", err)
		}
	}
	if bandCombination_v1570Present {
		ie.bandCombination_v1570 = new(BandCombination_v1570)
		if err = ie.bandCombination_v1570.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1570", err)
		}
	}
	if bandCombination_v1580Present {
		ie.bandCombination_v1580 = new(BandCombination_v1580)
		if err = ie.bandCombination_v1580.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1580", err)
		}
	}
	if bandCombination_v1590Present {
		ie.bandCombination_v1590 = new(BandCombination_v1590)
		if err = ie.bandCombination_v1590.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1590", err)
		}
	}
	if bandCombination_v1610Present {
		ie.bandCombination_v1610 = new(BandCombination_v1610)
		if err = ie.bandCombination_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1610", err)
		}
	}
	tmp_supportedBandPairListNR_r16 := utils.NewSequence[*ULTxSwitchingBandPair_r16]([]*ULTxSwitchingBandPair_r16{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
	fn_supportedBandPairListNR_r16 := func() *ULTxSwitchingBandPair_r16 {
		return new(ULTxSwitchingBandPair_r16)
	}
	if err = tmp_supportedBandPairListNR_r16.Decode(r, fn_supportedBandPairListNR_r16); err != nil {
		return utils.WrapError("Decode supportedBandPairListNR_r16", err)
	}
	ie.supportedBandPairListNR_r16 = []ULTxSwitchingBandPair_r16{}
	for _, i := range tmp_supportedBandPairListNR_r16.Value {
		ie.supportedBandPairListNR_r16 = append(ie.supportedBandPairListNR_r16, *i)
	}
	if uplinkTxSwitching_OptionSupport_r16Present {
		ie.uplinkTxSwitching_OptionSupport_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_OptionSupport_r16)
		if err = ie.uplinkTxSwitching_OptionSupport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxSwitching_OptionSupport_r16", err)
		}
	}
	if uplinkTxSwitching_PowerBoosting_r16Present {
		ie.uplinkTxSwitching_PowerBoosting_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PowerBoosting_r16)
		if err = ie.uplinkTxSwitching_PowerBoosting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxSwitching_PowerBoosting_r16", err)
		}
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

			uplinkTxSwitching_PUSCH_TransCoherence_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode uplinkTxSwitching_PUSCH_TransCoherence_r16 optional
			if uplinkTxSwitching_PUSCH_TransCoherence_r16Present {
				ie.uplinkTxSwitching_PUSCH_TransCoherence_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16)
				if err = ie.uplinkTxSwitching_PUSCH_TransCoherence_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
				}
			}
		}
	}
	return nil
}
