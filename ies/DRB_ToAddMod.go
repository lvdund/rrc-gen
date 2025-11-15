package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_ToAddMod struct {
	cnAssociation   *DRB_ToAddMod_cnAssociation   `lb:0,ub:15,optional`
	drb_Identity    DRB_Identity                  `madatory`
	reestablishPDCP *DRB_ToAddMod_reestablishPDCP `optional`
	recoverPDCP     *DRB_ToAddMod_recoverPDCP     `optional`
	pdcp_Config     *PDCP_Config                  `optional`
	daps_Config_r16 *DRB_ToAddMod_daps_Config_r16 `optional,ext-1`
}

func (ie *DRB_ToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.daps_Config_r16 != nil
	preambleBits := []bool{hasExtensions, ie.cnAssociation != nil, ie.reestablishPDCP != nil, ie.recoverPDCP != nil, ie.pdcp_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cnAssociation != nil {
		if err = ie.cnAssociation.Encode(w); err != nil {
			return utils.WrapError("Encode cnAssociation", err)
		}
	}
	if err = ie.drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode drb_Identity", err)
	}
	if ie.reestablishPDCP != nil {
		if err = ie.reestablishPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishPDCP", err)
		}
	}
	if ie.recoverPDCP != nil {
		if err = ie.recoverPDCP.Encode(w); err != nil {
			return utils.WrapError("Encode recoverPDCP", err)
		}
	}
	if ie.pdcp_Config != nil {
		if err = ie.pdcp_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_Config", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.daps_Config_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DRB_ToAddMod", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.daps_Config_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode daps_Config_r16 optional
			if ie.daps_Config_r16 != nil {
				if err = ie.daps_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode daps_Config_r16", err)
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

func (ie *DRB_ToAddMod) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cnAssociationPresent bool
	if cnAssociationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishPDCPPresent bool
	if reestablishPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var recoverPDCPPresent bool
	if recoverPDCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_ConfigPresent bool
	if pdcp_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cnAssociationPresent {
		ie.cnAssociation = new(DRB_ToAddMod_cnAssociation)
		if err = ie.cnAssociation.Decode(r); err != nil {
			return utils.WrapError("Decode cnAssociation", err)
		}
	}
	if err = ie.drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode drb_Identity", err)
	}
	if reestablishPDCPPresent {
		ie.reestablishPDCP = new(DRB_ToAddMod_reestablishPDCP)
		if err = ie.reestablishPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishPDCP", err)
		}
	}
	if recoverPDCPPresent {
		ie.recoverPDCP = new(DRB_ToAddMod_recoverPDCP)
		if err = ie.recoverPDCP.Decode(r); err != nil {
			return utils.WrapError("Decode recoverPDCP", err)
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

			daps_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode daps_Config_r16 optional
			if daps_Config_r16Present {
				ie.daps_Config_r16 = new(DRB_ToAddMod_daps_Config_r16)
				if err = ie.daps_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode daps_Config_r16", err)
				}
			}
		}
	}
	return nil
}
