package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReconfigurationWithSync struct {
	spCellConfigCommon         *ServingCellConfigCommon                      `optional`
	newUE_Identity             RNTI_Value                                    `madatory`
	t304                       ReconfigurationWithSync_t304                  `madatory`
	rach_ConfigDedicated       *ReconfigurationWithSync_rach_ConfigDedicated `optional`
	smtc                       *SSB_MTC                                      `optional,ext-1`
	daps_UplinkPowerConfig_r16 *DAPS_UplinkPowerConfig_r16                   `optional,ext-2`
	sl_PathSwitchConfig_r17    *SL_PathSwitchConfig_r17                      `optional,ext-3`
}

func (ie *ReconfigurationWithSync) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.smtc != nil || ie.daps_UplinkPowerConfig_r16 != nil || ie.sl_PathSwitchConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.spCellConfigCommon != nil, ie.rach_ConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.spCellConfigCommon != nil {
		if err = ie.spCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode spCellConfigCommon", err)
		}
	}
	if err = ie.newUE_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode newUE_Identity", err)
	}
	if err = ie.t304.Encode(w); err != nil {
		return utils.WrapError("Encode t304", err)
	}
	if ie.rach_ConfigDedicated != nil {
		if err = ie.rach_ConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode rach_ConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.smtc != nil, ie.daps_UplinkPowerConfig_r16 != nil, ie.sl_PathSwitchConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ReconfigurationWithSync", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.smtc != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode smtc optional
			if ie.smtc != nil {
				if err = ie.smtc.Encode(extWriter); err != nil {
					return utils.WrapError("Encode smtc", err)
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
			optionals_ext_2 := []bool{ie.daps_UplinkPowerConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode daps_UplinkPowerConfig_r16 optional
			if ie.daps_UplinkPowerConfig_r16 != nil {
				if err = ie.daps_UplinkPowerConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode daps_UplinkPowerConfig_r16", err)
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
			optionals_ext_3 := []bool{ie.sl_PathSwitchConfig_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_PathSwitchConfig_r17 optional
			if ie.sl_PathSwitchConfig_r17 != nil {
				if err = ie.sl_PathSwitchConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PathSwitchConfig_r17", err)
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

func (ie *ReconfigurationWithSync) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var spCellConfigCommonPresent bool
	if spCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rach_ConfigDedicatedPresent bool
	if rach_ConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if spCellConfigCommonPresent {
		ie.spCellConfigCommon = new(ServingCellConfigCommon)
		if err = ie.spCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode spCellConfigCommon", err)
		}
	}
	if err = ie.newUE_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode newUE_Identity", err)
	}
	if err = ie.t304.Decode(r); err != nil {
		return utils.WrapError("Decode t304", err)
	}
	if rach_ConfigDedicatedPresent {
		ie.rach_ConfigDedicated = new(ReconfigurationWithSync_rach_ConfigDedicated)
		if err = ie.rach_ConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode rach_ConfigDedicated", err)
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

			smtcPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode smtc optional
			if smtcPresent {
				ie.smtc = new(SSB_MTC)
				if err = ie.smtc.Decode(extReader); err != nil {
					return utils.WrapError("Decode smtc", err)
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

			daps_UplinkPowerConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode daps_UplinkPowerConfig_r16 optional
			if daps_UplinkPowerConfig_r16Present {
				ie.daps_UplinkPowerConfig_r16 = new(DAPS_UplinkPowerConfig_r16)
				if err = ie.daps_UplinkPowerConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode daps_UplinkPowerConfig_r16", err)
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

			sl_PathSwitchConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_PathSwitchConfig_r17 optional
			if sl_PathSwitchConfig_r17Present {
				ie.sl_PathSwitchConfig_r17 = new(SL_PathSwitchConfig_r17)
				if err = ie.sl_PathSwitchConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PathSwitchConfig_r17", err)
				}
			}
		}
	}
	return nil
}
