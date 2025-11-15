package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DownlinkConfigCommonSIB struct {
	frequencyInfoDL               FrequencyInfoDL_SIB `madatory`
	initialDownlinkBWP            BWP_DownlinkCommon  `madatory`
	bcch_Config                   BCCH_Config         `madatory`
	pcch_Config                   PCCH_Config         `madatory`
	pei_Config_r17                *PEI_Config_r17     `optional,ext-1`
	initialDownlinkBWP_RedCap_r17 *BWP_DownlinkCommon `optional,ext-1`
}

func (ie *DownlinkConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.pei_Config_r17 != nil || ie.initialDownlinkBWP_RedCap_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.frequencyInfoDL.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyInfoDL", err)
	}
	if err = ie.initialDownlinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode initialDownlinkBWP", err)
	}
	if err = ie.bcch_Config.Encode(w); err != nil {
		return utils.WrapError("Encode bcch_Config", err)
	}
	if err = ie.pcch_Config.Encode(w); err != nil {
		return utils.WrapError("Encode pcch_Config", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.pei_Config_r17 != nil || ie.initialDownlinkBWP_RedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DownlinkConfigCommonSIB", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.pei_Config_r17 != nil, ie.initialDownlinkBWP_RedCap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pei_Config_r17 optional
			if ie.pei_Config_r17 != nil {
				if err = ie.pei_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pei_Config_r17", err)
				}
			}
			// encode initialDownlinkBWP_RedCap_r17 optional
			if ie.initialDownlinkBWP_RedCap_r17 != nil {
				if err = ie.initialDownlinkBWP_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode initialDownlinkBWP_RedCap_r17", err)
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

func (ie *DownlinkConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.frequencyInfoDL.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyInfoDL", err)
	}
	if err = ie.initialDownlinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode initialDownlinkBWP", err)
	}
	if err = ie.bcch_Config.Decode(r); err != nil {
		return utils.WrapError("Decode bcch_Config", err)
	}
	if err = ie.pcch_Config.Decode(r); err != nil {
		return utils.WrapError("Decode pcch_Config", err)
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

			pei_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			initialDownlinkBWP_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pei_Config_r17 optional
			if pei_Config_r17Present {
				ie.pei_Config_r17 = new(PEI_Config_r17)
				if err = ie.pei_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pei_Config_r17", err)
				}
			}
			// decode initialDownlinkBWP_RedCap_r17 optional
			if initialDownlinkBWP_RedCap_r17Present {
				ie.initialDownlinkBWP_RedCap_r17 = new(BWP_DownlinkCommon)
				if err = ie.initialDownlinkBWP_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode initialDownlinkBWP_RedCap_r17", err)
				}
			}
		}
	}
	return nil
}
