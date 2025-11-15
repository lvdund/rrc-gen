package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AS_Context struct {
	reestablishmentInfo              *ReestablishmentInfo          `optional`
	configRestrictInfo               *ConfigRestrictInfoSCG        `optional`
	ran_NotificationAreaInfo         *RAN_NotificationAreaInfo     `optional,ext-1`
	ueAssistanceInformation          *[]byte                       `optional,ext-2`
	selectedBandCombinationSN        *BandCombinationInfoSN        `optional,ext-3`
	configRestrictInfoDAPS_r16       *ConfigRestrictInfoDAPS_r16   `optional,ext-4`
	sidelinkUEInformationNR_r16      *[]byte                       `optional,ext-4`
	sidelinkUEInformationEUTRA_r16   *[]byte                       `optional,ext-4`
	ueAssistanceInformationEUTRA_r16 *[]byte                       `optional,ext-4`
	ueAssistanceInformationSCG_r16   *[]byte                       `optional,ext-4`
	needForGapsInfoNR_r16            *NeedForGapsInfoNR_r16        `optional,ext-4`
	configRestrictInfoDAPS_v1640     *ConfigRestrictInfoDAPS_v1640 `optional,ext-5`
	needForGapNCSG_InfoNR_r17        *NeedForGapNCSG_InfoNR_r17    `optional,ext-6`
	needForGapNCSG_InfoEUTRA_r17     *NeedForGapNCSG_InfoEUTRA_r17 `optional,ext-6`
	mbsInterestIndication_r17        *[]byte                       `optional,ext-6`
}

func (ie *AS_Context) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ran_NotificationAreaInfo != nil || ie.ueAssistanceInformation != nil || ie.selectedBandCombinationSN != nil || ie.configRestrictInfoDAPS_r16 != nil || ie.sidelinkUEInformationNR_r16 != nil || ie.sidelinkUEInformationEUTRA_r16 != nil || ie.ueAssistanceInformationEUTRA_r16 != nil || ie.ueAssistanceInformationSCG_r16 != nil || ie.needForGapsInfoNR_r16 != nil || ie.configRestrictInfoDAPS_v1640 != nil || ie.needForGapNCSG_InfoNR_r17 != nil || ie.needForGapNCSG_InfoEUTRA_r17 != nil || ie.mbsInterestIndication_r17 != nil
	preambleBits := []bool{hasExtensions, ie.reestablishmentInfo != nil, ie.configRestrictInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reestablishmentInfo != nil {
		if err = ie.reestablishmentInfo.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishmentInfo", err)
		}
	}
	if ie.configRestrictInfo != nil {
		if err = ie.configRestrictInfo.Encode(w); err != nil {
			return utils.WrapError("Encode configRestrictInfo", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 6 bits for 6 extension groups
		extBitmap := []bool{ie.ran_NotificationAreaInfo != nil, ie.ueAssistanceInformation != nil, ie.selectedBandCombinationSN != nil, ie.configRestrictInfoDAPS_r16 != nil || ie.sidelinkUEInformationNR_r16 != nil || ie.sidelinkUEInformationEUTRA_r16 != nil || ie.ueAssistanceInformationEUTRA_r16 != nil || ie.ueAssistanceInformationSCG_r16 != nil || ie.needForGapsInfoNR_r16 != nil, ie.configRestrictInfoDAPS_v1640 != nil, ie.needForGapNCSG_InfoNR_r17 != nil || ie.needForGapNCSG_InfoEUTRA_r17 != nil || ie.mbsInterestIndication_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap AS_Context", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ran_NotificationAreaInfo != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ran_NotificationAreaInfo optional
			if ie.ran_NotificationAreaInfo != nil {
				if err = ie.ran_NotificationAreaInfo.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ran_NotificationAreaInfo", err)
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
			optionals_ext_2 := []bool{ie.ueAssistanceInformation != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ueAssistanceInformation optional
			if ie.ueAssistanceInformation != nil {
				if err = extWriter.WriteOctetString(*ie.ueAssistanceInformation, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode ueAssistanceInformation", err)
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
			optionals_ext_3 := []bool{ie.selectedBandCombinationSN != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode selectedBandCombinationSN optional
			if ie.selectedBandCombinationSN != nil {
				if err = ie.selectedBandCombinationSN.Encode(extWriter); err != nil {
					return utils.WrapError("Encode selectedBandCombinationSN", err)
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
			optionals_ext_4 := []bool{ie.configRestrictInfoDAPS_r16 != nil, ie.sidelinkUEInformationNR_r16 != nil, ie.sidelinkUEInformationEUTRA_r16 != nil, ie.ueAssistanceInformationEUTRA_r16 != nil, ie.ueAssistanceInformationSCG_r16 != nil, ie.needForGapsInfoNR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode configRestrictInfoDAPS_r16 optional
			if ie.configRestrictInfoDAPS_r16 != nil {
				if err = ie.configRestrictInfoDAPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configRestrictInfoDAPS_r16", err)
				}
			}
			// encode sidelinkUEInformationNR_r16 optional
			if ie.sidelinkUEInformationNR_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.sidelinkUEInformationNR_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sidelinkUEInformationNR_r16", err)
				}
			}
			// encode sidelinkUEInformationEUTRA_r16 optional
			if ie.sidelinkUEInformationEUTRA_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.sidelinkUEInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sidelinkUEInformationEUTRA_r16", err)
				}
			}
			// encode ueAssistanceInformationEUTRA_r16 optional
			if ie.ueAssistanceInformationEUTRA_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.ueAssistanceInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode ueAssistanceInformationEUTRA_r16", err)
				}
			}
			// encode ueAssistanceInformationSCG_r16 optional
			if ie.ueAssistanceInformationSCG_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.ueAssistanceInformationSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode ueAssistanceInformationSCG_r16", err)
				}
			}
			// encode needForGapsInfoNR_r16 optional
			if ie.needForGapsInfoNR_r16 != nil {
				if err = ie.needForGapsInfoNR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode needForGapsInfoNR_r16", err)
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
			optionals_ext_5 := []bool{ie.configRestrictInfoDAPS_v1640 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode configRestrictInfoDAPS_v1640 optional
			if ie.configRestrictInfoDAPS_v1640 != nil {
				if err = ie.configRestrictInfoDAPS_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configRestrictInfoDAPS_v1640", err)
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
			optionals_ext_6 := []bool{ie.needForGapNCSG_InfoNR_r17 != nil, ie.needForGapNCSG_InfoEUTRA_r17 != nil, ie.mbsInterestIndication_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode needForGapNCSG_InfoNR_r17 optional
			if ie.needForGapNCSG_InfoNR_r17 != nil {
				if err = ie.needForGapNCSG_InfoNR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode needForGapNCSG_InfoNR_r17", err)
				}
			}
			// encode needForGapNCSG_InfoEUTRA_r17 optional
			if ie.needForGapNCSG_InfoEUTRA_r17 != nil {
				if err = ie.needForGapNCSG_InfoEUTRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode needForGapNCSG_InfoEUTRA_r17", err)
				}
			}
			// encode mbsInterestIndication_r17 optional
			if ie.mbsInterestIndication_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.mbsInterestIndication_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode mbsInterestIndication_r17", err)
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

func (ie *AS_Context) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishmentInfoPresent bool
	if reestablishmentInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configRestrictInfoPresent bool
	if configRestrictInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if reestablishmentInfoPresent {
		ie.reestablishmentInfo = new(ReestablishmentInfo)
		if err = ie.reestablishmentInfo.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishmentInfo", err)
		}
	}
	if configRestrictInfoPresent {
		ie.configRestrictInfo = new(ConfigRestrictInfoSCG)
		if err = ie.configRestrictInfo.Decode(r); err != nil {
			return utils.WrapError("Decode configRestrictInfo", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 6 bits for 6 extension groups
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

			ran_NotificationAreaInfoPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ran_NotificationAreaInfo optional
			if ran_NotificationAreaInfoPresent {
				ie.ran_NotificationAreaInfo = new(RAN_NotificationAreaInfo)
				if err = ie.ran_NotificationAreaInfo.Decode(extReader); err != nil {
					return utils.WrapError("Decode ran_NotificationAreaInfo", err)
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

			ueAssistanceInformationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ueAssistanceInformation optional
			if ueAssistanceInformationPresent {
				var tmp_os_ueAssistanceInformation []byte
				if tmp_os_ueAssistanceInformation, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode ueAssistanceInformation", err)
				}
				ie.ueAssistanceInformation = &tmp_os_ueAssistanceInformation
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			selectedBandCombinationSNPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode selectedBandCombinationSN optional
			if selectedBandCombinationSNPresent {
				ie.selectedBandCombinationSN = new(BandCombinationInfoSN)
				if err = ie.selectedBandCombinationSN.Decode(extReader); err != nil {
					return utils.WrapError("Decode selectedBandCombinationSN", err)
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

			configRestrictInfoDAPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sidelinkUEInformationNR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sidelinkUEInformationEUTRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ueAssistanceInformationEUTRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ueAssistanceInformationSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			needForGapsInfoNR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode configRestrictInfoDAPS_r16 optional
			if configRestrictInfoDAPS_r16Present {
				ie.configRestrictInfoDAPS_r16 = new(ConfigRestrictInfoDAPS_r16)
				if err = ie.configRestrictInfoDAPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configRestrictInfoDAPS_r16", err)
				}
			}
			// decode sidelinkUEInformationNR_r16 optional
			if sidelinkUEInformationNR_r16Present {
				var tmp_os_sidelinkUEInformationNR_r16 []byte
				if tmp_os_sidelinkUEInformationNR_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sidelinkUEInformationNR_r16", err)
				}
				ie.sidelinkUEInformationNR_r16 = &tmp_os_sidelinkUEInformationNR_r16
			}
			// decode sidelinkUEInformationEUTRA_r16 optional
			if sidelinkUEInformationEUTRA_r16Present {
				var tmp_os_sidelinkUEInformationEUTRA_r16 []byte
				if tmp_os_sidelinkUEInformationEUTRA_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sidelinkUEInformationEUTRA_r16", err)
				}
				ie.sidelinkUEInformationEUTRA_r16 = &tmp_os_sidelinkUEInformationEUTRA_r16
			}
			// decode ueAssistanceInformationEUTRA_r16 optional
			if ueAssistanceInformationEUTRA_r16Present {
				var tmp_os_ueAssistanceInformationEUTRA_r16 []byte
				if tmp_os_ueAssistanceInformationEUTRA_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode ueAssistanceInformationEUTRA_r16", err)
				}
				ie.ueAssistanceInformationEUTRA_r16 = &tmp_os_ueAssistanceInformationEUTRA_r16
			}
			// decode ueAssistanceInformationSCG_r16 optional
			if ueAssistanceInformationSCG_r16Present {
				var tmp_os_ueAssistanceInformationSCG_r16 []byte
				if tmp_os_ueAssistanceInformationSCG_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode ueAssistanceInformationSCG_r16", err)
				}
				ie.ueAssistanceInformationSCG_r16 = &tmp_os_ueAssistanceInformationSCG_r16
			}
			// decode needForGapsInfoNR_r16 optional
			if needForGapsInfoNR_r16Present {
				ie.needForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
				if err = ie.needForGapsInfoNR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode needForGapsInfoNR_r16", err)
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

			configRestrictInfoDAPS_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode configRestrictInfoDAPS_v1640 optional
			if configRestrictInfoDAPS_v1640Present {
				ie.configRestrictInfoDAPS_v1640 = new(ConfigRestrictInfoDAPS_v1640)
				if err = ie.configRestrictInfoDAPS_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode configRestrictInfoDAPS_v1640", err)
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

			needForGapNCSG_InfoNR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			needForGapNCSG_InfoEUTRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mbsInterestIndication_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode needForGapNCSG_InfoNR_r17 optional
			if needForGapNCSG_InfoNR_r17Present {
				ie.needForGapNCSG_InfoNR_r17 = new(NeedForGapNCSG_InfoNR_r17)
				if err = ie.needForGapNCSG_InfoNR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode needForGapNCSG_InfoNR_r17", err)
				}
			}
			// decode needForGapNCSG_InfoEUTRA_r17 optional
			if needForGapNCSG_InfoEUTRA_r17Present {
				ie.needForGapNCSG_InfoEUTRA_r17 = new(NeedForGapNCSG_InfoEUTRA_r17)
				if err = ie.needForGapNCSG_InfoEUTRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode needForGapNCSG_InfoEUTRA_r17", err)
				}
			}
			// decode mbsInterestIndication_r17 optional
			if mbsInterestIndication_r17Present {
				var tmp_os_mbsInterestIndication_r17 []byte
				if tmp_os_mbsInterestIndication_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode mbsInterestIndication_r17", err)
				}
				ie.mbsInterestIndication_r17 = &tmp_os_mbsInterestIndication_r17
			}
		}
	}
	return nil
}
