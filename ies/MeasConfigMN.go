package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigMN struct {
	measuredFrequenciesMN []NR_FreqInfo                    `lb:1,ub:maxMeasFreqsMN,optional`
	measGapConfig         *GapConfig                       `optional,setuprelease`
	gapPurpose            *MeasConfigMN_gapPurpose         `optional`
	measGapConfigFR2      *GapConfig                       `optional,ext-1,setuprelease`
	interFreqNoGap_r16    *MeasConfigMN_interFreqNoGap_r16 `optional,ext-2`
}

func (ie *MeasConfigMN) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.measGapConfigFR2 != nil || ie.interFreqNoGap_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.measuredFrequenciesMN) > 0, ie.measGapConfig != nil, ie.gapPurpose != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.measuredFrequenciesMN) > 0 {
		tmp_measuredFrequenciesMN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
		for _, i := range ie.measuredFrequenciesMN {
			tmp_measuredFrequenciesMN.Value = append(tmp_measuredFrequenciesMN.Value, &i)
		}
		if err = tmp_measuredFrequenciesMN.Encode(w); err != nil {
			return utils.WrapError("Encode measuredFrequenciesMN", err)
		}
	}
	if ie.measGapConfig != nil {
		tmp_measGapConfig := utils.SetupRelease[*GapConfig]{
			Setup: ie.measGapConfig,
		}
		if err = tmp_measGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode measGapConfig", err)
		}
	}
	if ie.gapPurpose != nil {
		if err = ie.gapPurpose.Encode(w); err != nil {
			return utils.WrapError("Encode gapPurpose", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.measGapConfigFR2 != nil, ie.interFreqNoGap_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasConfigMN", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.measGapConfigFR2 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measGapConfigFR2 optional
			if ie.measGapConfigFR2 != nil {
				tmp_measGapConfigFR2 := utils.SetupRelease[*GapConfig]{
					Setup: ie.measGapConfigFR2,
				}
				if err = tmp_measGapConfigFR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measGapConfigFR2", err)
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
			optionals_ext_2 := []bool{ie.interFreqNoGap_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFreqNoGap_r16 optional
			if ie.interFreqNoGap_r16 != nil {
				if err = ie.interFreqNoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFreqNoGap_r16", err)
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

func (ie *MeasConfigMN) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var measuredFrequenciesMNPresent bool
	if measuredFrequenciesMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measGapConfigPresent bool
	if measGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var gapPurposePresent bool
	if gapPurposePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measuredFrequenciesMNPresent {
		tmp_measuredFrequenciesMN := utils.NewSequence[*NR_FreqInfo]([]*NR_FreqInfo{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
		fn_measuredFrequenciesMN := func() *NR_FreqInfo {
			return new(NR_FreqInfo)
		}
		if err = tmp_measuredFrequenciesMN.Decode(r, fn_measuredFrequenciesMN); err != nil {
			return utils.WrapError("Decode measuredFrequenciesMN", err)
		}
		ie.measuredFrequenciesMN = []NR_FreqInfo{}
		for _, i := range tmp_measuredFrequenciesMN.Value {
			ie.measuredFrequenciesMN = append(ie.measuredFrequenciesMN, *i)
		}
	}
	if measGapConfigPresent {
		tmp_measGapConfig := utils.SetupRelease[*GapConfig]{}
		if err = tmp_measGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode measGapConfig", err)
		}
		ie.measGapConfig = tmp_measGapConfig.Setup
	}
	if gapPurposePresent {
		ie.gapPurpose = new(MeasConfigMN_gapPurpose)
		if err = ie.gapPurpose.Decode(r); err != nil {
			return utils.WrapError("Decode gapPurpose", err)
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

			measGapConfigFR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measGapConfigFR2 optional
			if measGapConfigFR2Present {
				tmp_measGapConfigFR2 := utils.SetupRelease[*GapConfig]{}
				if err = tmp_measGapConfigFR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode measGapConfigFR2", err)
				}
				ie.measGapConfigFR2 = tmp_measGapConfigFR2.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			interFreqNoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFreqNoGap_r16 optional
			if interFreqNoGap_r16Present {
				ie.interFreqNoGap_r16 = new(MeasConfigMN_interFreqNoGap_r16)
				if err = ie.interFreqNoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFreqNoGap_r16", err)
				}
			}
		}
	}
	return nil
}
