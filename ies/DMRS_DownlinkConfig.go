package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_DownlinkConfig struct {
	dmrs_Type               *DMRS_DownlinkConfig_dmrs_Type               `optional`
	dmrs_AdditionalPosition *DMRS_DownlinkConfig_dmrs_AdditionalPosition `optional`
	maxLength               *DMRS_DownlinkConfig_maxLength               `optional`
	scramblingID0           *int64                                       `lb:0,ub:65535,optional`
	scramblingID1           *int64                                       `lb:0,ub:65535,optional`
	phaseTrackingRS         *PTRS_DownlinkConfig                         `optional,setuprelease`
	dmrs_Downlink_r16       *DMRS_DownlinkConfig_dmrs_Downlink_r16       `optional,ext-1`
}

func (ie *DMRS_DownlinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dmrs_Downlink_r16 != nil
	preambleBits := []bool{hasExtensions, ie.dmrs_Type != nil, ie.dmrs_AdditionalPosition != nil, ie.maxLength != nil, ie.scramblingID0 != nil, ie.scramblingID1 != nil, ie.phaseTrackingRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dmrs_Type != nil {
		if err = ie.dmrs_Type.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_Type", err)
		}
	}
	if ie.dmrs_AdditionalPosition != nil {
		if err = ie.dmrs_AdditionalPosition.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_AdditionalPosition", err)
		}
	}
	if ie.maxLength != nil {
		if err = ie.maxLength.Encode(w); err != nil {
			return utils.WrapError("Encode maxLength", err)
		}
	}
	if ie.scramblingID0 != nil {
		if err = w.WriteInteger(*ie.scramblingID0, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode scramblingID0", err)
		}
	}
	if ie.scramblingID1 != nil {
		if err = w.WriteInteger(*ie.scramblingID1, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode scramblingID1", err)
		}
	}
	if ie.phaseTrackingRS != nil {
		tmp_phaseTrackingRS := utils.SetupRelease[*PTRS_DownlinkConfig]{
			Setup: ie.phaseTrackingRS,
		}
		if err = tmp_phaseTrackingRS.Encode(w); err != nil {
			return utils.WrapError("Encode phaseTrackingRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dmrs_Downlink_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DMRS_DownlinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dmrs_Downlink_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dmrs_Downlink_r16 optional
			if ie.dmrs_Downlink_r16 != nil {
				if err = ie.dmrs_Downlink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_Downlink_r16", err)
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

func (ie *DMRS_DownlinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_TypePresent bool
	if dmrs_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_AdditionalPositionPresent bool
	if dmrs_AdditionalPositionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxLengthPresent bool
	if maxLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scramblingID0Present bool
	if scramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scramblingID1Present bool
	if scramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phaseTrackingRSPresent bool
	if phaseTrackingRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dmrs_TypePresent {
		ie.dmrs_Type = new(DMRS_DownlinkConfig_dmrs_Type)
		if err = ie.dmrs_Type.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_Type", err)
		}
	}
	if dmrs_AdditionalPositionPresent {
		ie.dmrs_AdditionalPosition = new(DMRS_DownlinkConfig_dmrs_AdditionalPosition)
		if err = ie.dmrs_AdditionalPosition.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_AdditionalPosition", err)
		}
	}
	if maxLengthPresent {
		ie.maxLength = new(DMRS_DownlinkConfig_maxLength)
		if err = ie.maxLength.Decode(r); err != nil {
			return utils.WrapError("Decode maxLength", err)
		}
	}
	if scramblingID0Present {
		var tmp_int_scramblingID0 int64
		if tmp_int_scramblingID0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode scramblingID0", err)
		}
		ie.scramblingID0 = &tmp_int_scramblingID0
	}
	if scramblingID1Present {
		var tmp_int_scramblingID1 int64
		if tmp_int_scramblingID1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode scramblingID1", err)
		}
		ie.scramblingID1 = &tmp_int_scramblingID1
	}
	if phaseTrackingRSPresent {
		tmp_phaseTrackingRS := utils.SetupRelease[*PTRS_DownlinkConfig]{}
		if err = tmp_phaseTrackingRS.Decode(r); err != nil {
			return utils.WrapError("Decode phaseTrackingRS", err)
		}
		ie.phaseTrackingRS = tmp_phaseTrackingRS.Setup
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

			dmrs_Downlink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dmrs_Downlink_r16 optional
			if dmrs_Downlink_r16Present {
				ie.dmrs_Downlink_r16 = new(DMRS_DownlinkConfig_dmrs_Downlink_r16)
				if err = ie.dmrs_Downlink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_Downlink_r16", err)
				}
			}
		}
	}
	return nil
}
