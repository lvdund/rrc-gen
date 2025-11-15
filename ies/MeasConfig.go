package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfig struct {
	measObjectToRemoveList         *MeasObjectToRemoveList                    `optional`
	measObjectToAddModList         *MeasObjectToAddModList                    `optional`
	reportConfigToRemoveList       *ReportConfigToRemoveList                  `optional`
	reportConfigToAddModList       *ReportConfigToAddModList                  `optional`
	measIdToRemoveList             *MeasIdToRemoveList                        `optional`
	measIdToAddModList             *MeasIdToAddModList                        `optional`
	s_MeasureConfig                *MeasConfig_s_MeasureConfig                `optional`
	quantityConfig                 *QuantityConfig                            `optional`
	measGapConfig                  *MeasGapConfig                             `optional`
	measGapSharingConfig           *MeasGapSharingConfig                      `optional`
	interFrequencyConfig_NoGap_r16 *MeasConfig_interFrequencyConfig_NoGap_r16 `optional,ext-1`
}

func (ie *MeasConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.interFrequencyConfig_NoGap_r16 != nil
	preambleBits := []bool{hasExtensions, ie.measObjectToRemoveList != nil, ie.measObjectToAddModList != nil, ie.reportConfigToRemoveList != nil, ie.reportConfigToAddModList != nil, ie.measIdToRemoveList != nil, ie.measIdToAddModList != nil, ie.s_MeasureConfig != nil, ie.quantityConfig != nil, ie.measGapConfig != nil, ie.measGapSharingConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measObjectToRemoveList != nil {
		if err = ie.measObjectToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode measObjectToRemoveList", err)
		}
	}
	if ie.measObjectToAddModList != nil {
		if err = ie.measObjectToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode measObjectToAddModList", err)
		}
	}
	if ie.reportConfigToRemoveList != nil {
		if err = ie.reportConfigToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode reportConfigToRemoveList", err)
		}
	}
	if ie.reportConfigToAddModList != nil {
		if err = ie.reportConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode reportConfigToAddModList", err)
		}
	}
	if ie.measIdToRemoveList != nil {
		if err = ie.measIdToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode measIdToRemoveList", err)
		}
	}
	if ie.measIdToAddModList != nil {
		if err = ie.measIdToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode measIdToAddModList", err)
		}
	}
	if ie.s_MeasureConfig != nil {
		if err = ie.s_MeasureConfig.Encode(w); err != nil {
			return utils.WrapError("Encode s_MeasureConfig", err)
		}
	}
	if ie.quantityConfig != nil {
		if err = ie.quantityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode quantityConfig", err)
		}
	}
	if ie.measGapConfig != nil {
		if err = ie.measGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode measGapConfig", err)
		}
	}
	if ie.measGapSharingConfig != nil {
		if err = ie.measGapSharingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode measGapSharingConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.interFrequencyConfig_NoGap_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.interFrequencyConfig_NoGap_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFrequencyConfig_NoGap_r16 optional
			if ie.interFrequencyConfig_NoGap_r16 != nil {
				if err = ie.interFrequencyConfig_NoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFrequencyConfig_NoGap_r16", err)
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

func (ie *MeasConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var measObjectToRemoveListPresent bool
	if measObjectToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measObjectToAddModListPresent bool
	if measObjectToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportConfigToRemoveListPresent bool
	if reportConfigToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportConfigToAddModListPresent bool
	if reportConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measIdToRemoveListPresent bool
	if measIdToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measIdToAddModListPresent bool
	if measIdToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var s_MeasureConfigPresent bool
	if s_MeasureConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var quantityConfigPresent bool
	if quantityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measGapConfigPresent bool
	if measGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measGapSharingConfigPresent bool
	if measGapSharingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measObjectToRemoveListPresent {
		ie.measObjectToRemoveList = new(MeasObjectToRemoveList)
		if err = ie.measObjectToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectToRemoveList", err)
		}
	}
	if measObjectToAddModListPresent {
		ie.measObjectToAddModList = new(MeasObjectToAddModList)
		if err = ie.measObjectToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectToAddModList", err)
		}
	}
	if reportConfigToRemoveListPresent {
		ie.reportConfigToRemoveList = new(ReportConfigToRemoveList)
		if err = ie.reportConfigToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigToRemoveList", err)
		}
	}
	if reportConfigToAddModListPresent {
		ie.reportConfigToAddModList = new(ReportConfigToAddModList)
		if err = ie.reportConfigToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigToAddModList", err)
		}
	}
	if measIdToRemoveListPresent {
		ie.measIdToRemoveList = new(MeasIdToRemoveList)
		if err = ie.measIdToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode measIdToRemoveList", err)
		}
	}
	if measIdToAddModListPresent {
		ie.measIdToAddModList = new(MeasIdToAddModList)
		if err = ie.measIdToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode measIdToAddModList", err)
		}
	}
	if s_MeasureConfigPresent {
		ie.s_MeasureConfig = new(MeasConfig_s_MeasureConfig)
		if err = ie.s_MeasureConfig.Decode(r); err != nil {
			return utils.WrapError("Decode s_MeasureConfig", err)
		}
	}
	if quantityConfigPresent {
		ie.quantityConfig = new(QuantityConfig)
		if err = ie.quantityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode quantityConfig", err)
		}
	}
	if measGapConfigPresent {
		ie.measGapConfig = new(MeasGapConfig)
		if err = ie.measGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode measGapConfig", err)
		}
	}
	if measGapSharingConfigPresent {
		ie.measGapSharingConfig = new(MeasGapSharingConfig)
		if err = ie.measGapSharingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode measGapSharingConfig", err)
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

			interFrequencyConfig_NoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFrequencyConfig_NoGap_r16 optional
			if interFrequencyConfig_NoGap_r16Present {
				ie.interFrequencyConfig_NoGap_r16 = new(MeasConfig_interFrequencyConfig_NoGap_r16)
				if err = ie.interFrequencyConfig_NoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFrequencyConfig_NoGap_r16", err)
				}
			}
		}
	}
	return nil
}
