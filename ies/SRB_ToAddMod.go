package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRB_ToAddMod struct {
	srb_Identity       SRB_Identity                  `madatory`
	reestablishPDCP    *SRB_ToAddMod_reestablishPDCP `optional`
	discardOnPDCP      *SRB_ToAddMod_discardOnPDCP   `optional`
	pdcp_Config        *PDCP_Config                  `optional`
	srb_Identity_v1700 *SRB_Identity_v1700           `optional,ext-1`
}

func (ie *SRB_ToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.srb_Identity_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.reestablishPDCP != nil, ie.discardOnPDCP != nil, ie.pdcp_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode srb_Identity", err)
	}
	if ie.reestablishPDCP != nil {
		if err = ie.reestablishPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishPDCP", err)
		}
	}
	if ie.discardOnPDCP != nil {
		if err = ie.discardOnPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode discardOnPDCP", err)
		}
	}
	if ie.pdcp_Config != nil {
		if err = ie.pdcp_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_Config", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.srb_Identity_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRB_ToAddMod", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.srb_Identity_v1700 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode srb_Identity_v1700 optional
			if ie.srb_Identity_v1700 != nil {
				if err = ie.srb_Identity_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srb_Identity_v1700", err)
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

func (ie *SRB_ToAddMod) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishPDCPPresent bool
	if reestablishPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var discardOnPDCPPresent bool
	if discardOnPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_ConfigPresent bool
	if pdcp_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode srb_Identity", err)
	}
	if reestablishPDCPPresent {
		ie.reestablishPDCP = new(SRB_ToAddMod_reestablishPDCP)
		if err = ie.reestablishPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishPDCP", err)
		}
	}
	if discardOnPDCPPresent {
		ie.discardOnPDCP = new(SRB_ToAddMod_discardOnPDCP)
		if err = ie.discardOnPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode discardOnPDCP", err)
		}
	}
	if pdcp_ConfigPresent {
		ie.pdcp_Config = new(PDCP_Config)
		if err = ie.pdcp_Config.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_Config", err)
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

			srb_Identity_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srb_Identity_v1700 optional
			if srb_Identity_v1700Present {
				ie.srb_Identity_v1700 = new(SRB_Identity_v1700)
				if err = ie.srb_Identity_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode srb_Identity_v1700", err)
				}
			}
		}
	}
	return nil
}
