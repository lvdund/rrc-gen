package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters struct {
	supportedROHC_Profiles             PDCP_Parameters_supportedROHC_Profiles              `madatory`
	maxNumberROHC_ContextSessions      PDCP_Parameters_maxNumberROHC_ContextSessions       `madatory`
	uplinkOnlyROHC_Profiles            *PDCP_Parameters_uplinkOnlyROHC_Profiles            `optional`
	continueROHC_Context               *PDCP_Parameters_continueROHC_Context               `optional`
	outOfOrderDelivery                 *PDCP_Parameters_outOfOrderDelivery                 `optional`
	shortSN                            *PDCP_Parameters_shortSN                            `optional`
	pdcp_DuplicationSRB                *PDCP_Parameters_pdcp_DuplicationSRB                `optional`
	pdcp_DuplicationMCG_OrSCG_DRB      *PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB      `optional`
	drb_IAB_r16                        *PDCP_Parameters_drb_IAB_r16                        `optional,ext-1`
	non_DRB_IAB_r16                    *PDCP_Parameters_non_DRB_IAB_r16                    `optional,ext-1`
	extendedDiscardTimer_r16           *PDCP_Parameters_extendedDiscardTimer_r16           `optional,ext-1`
	continueEHC_Context_r16            *PDCP_Parameters_continueEHC_Context_r16            `optional,ext-1`
	ehc_r16                            *PDCP_Parameters_ehc_r16                            `optional,ext-1`
	maxNumberEHC_Contexts_r16          *PDCP_Parameters_maxNumberEHC_Contexts_r16          `optional,ext-1`
	jointEHC_ROHC_Config_r16           *PDCP_Parameters_jointEHC_ROHC_Config_r16           `optional,ext-1`
	pdcp_DuplicationMoreThanTwoRLC_r16 *PDCP_Parameters_pdcp_DuplicationMoreThanTwoRLC_r16 `optional,ext-1`
	longSN_RedCap_r17                  *PDCP_Parameters_longSN_RedCap_r17                  `optional,ext-2`
	udc_r17                            *PDCP_Parameters_udc_r17                            `lb:0,ub:15,optional,ext-2`
}

func (ie *PDCP_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.drb_IAB_r16 != nil || ie.non_DRB_IAB_r16 != nil || ie.extendedDiscardTimer_r16 != nil || ie.continueEHC_Context_r16 != nil || ie.ehc_r16 != nil || ie.maxNumberEHC_Contexts_r16 != nil || ie.jointEHC_ROHC_Config_r16 != nil || ie.pdcp_DuplicationMoreThanTwoRLC_r16 != nil || ie.longSN_RedCap_r17 != nil || ie.udc_r17 != nil
	preambleBits := []bool{hasExtensions, ie.uplinkOnlyROHC_Profiles != nil, ie.continueROHC_Context != nil, ie.outOfOrderDelivery != nil, ie.shortSN != nil, ie.pdcp_DuplicationSRB != nil, ie.pdcp_DuplicationMCG_OrSCG_DRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedROHC_Profiles.Encode(w); err != nil {
		return utils.WrapError("Encode supportedROHC_Profiles", err)
	}
	if err = ie.maxNumberROHC_ContextSessions.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberROHC_ContextSessions", err)
	}
	if ie.uplinkOnlyROHC_Profiles != nil {
		if err = ie.uplinkOnlyROHC_Profiles.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkOnlyROHC_Profiles", err)
		}
	}
	if ie.continueROHC_Context != nil {
		if err = ie.continueROHC_Context.Encode(w); err != nil {
			return utils.WrapError("Encode continueROHC_Context", err)
		}
	}
	if ie.outOfOrderDelivery != nil {
		if err = ie.outOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode outOfOrderDelivery", err)
		}
	}
	if ie.shortSN != nil {
		if err = ie.shortSN.Encode(w); err != nil {
			return utils.WrapError("Encode shortSN", err)
		}
	}
	if ie.pdcp_DuplicationSRB != nil {
		if err = ie.pdcp_DuplicationSRB.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_DuplicationSRB", err)
		}
	}
	if ie.pdcp_DuplicationMCG_OrSCG_DRB != nil {
		if err = ie.pdcp_DuplicationMCG_OrSCG_DRB.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_DuplicationMCG_OrSCG_DRB", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.drb_IAB_r16 != nil || ie.non_DRB_IAB_r16 != nil || ie.extendedDiscardTimer_r16 != nil || ie.continueEHC_Context_r16 != nil || ie.ehc_r16 != nil || ie.maxNumberEHC_Contexts_r16 != nil || ie.jointEHC_ROHC_Config_r16 != nil || ie.pdcp_DuplicationMoreThanTwoRLC_r16 != nil, ie.longSN_RedCap_r17 != nil || ie.udc_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCP_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.drb_IAB_r16 != nil, ie.non_DRB_IAB_r16 != nil, ie.extendedDiscardTimer_r16 != nil, ie.continueEHC_Context_r16 != nil, ie.ehc_r16 != nil, ie.maxNumberEHC_Contexts_r16 != nil, ie.jointEHC_ROHC_Config_r16 != nil, ie.pdcp_DuplicationMoreThanTwoRLC_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode drb_IAB_r16 optional
			if ie.drb_IAB_r16 != nil {
				if err = ie.drb_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drb_IAB_r16", err)
				}
			}
			// encode non_DRB_IAB_r16 optional
			if ie.non_DRB_IAB_r16 != nil {
				if err = ie.non_DRB_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode non_DRB_IAB_r16", err)
				}
			}
			// encode extendedDiscardTimer_r16 optional
			if ie.extendedDiscardTimer_r16 != nil {
				if err = ie.extendedDiscardTimer_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedDiscardTimer_r16", err)
				}
			}
			// encode continueEHC_Context_r16 optional
			if ie.continueEHC_Context_r16 != nil {
				if err = ie.continueEHC_Context_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode continueEHC_Context_r16", err)
				}
			}
			// encode ehc_r16 optional
			if ie.ehc_r16 != nil {
				if err = ie.ehc_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ehc_r16", err)
				}
			}
			// encode maxNumberEHC_Contexts_r16 optional
			if ie.maxNumberEHC_Contexts_r16 != nil {
				if err = ie.maxNumberEHC_Contexts_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberEHC_Contexts_r16", err)
				}
			}
			// encode jointEHC_ROHC_Config_r16 optional
			if ie.jointEHC_ROHC_Config_r16 != nil {
				if err = ie.jointEHC_ROHC_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode jointEHC_ROHC_Config_r16", err)
				}
			}
			// encode pdcp_DuplicationMoreThanTwoRLC_r16 optional
			if ie.pdcp_DuplicationMoreThanTwoRLC_r16 != nil {
				if err = ie.pdcp_DuplicationMoreThanTwoRLC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcp_DuplicationMoreThanTwoRLC_r16", err)
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
			optionals_ext_2 := []bool{ie.longSN_RedCap_r17 != nil, ie.udc_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode longSN_RedCap_r17 optional
			if ie.longSN_RedCap_r17 != nil {
				if err = ie.longSN_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode longSN_RedCap_r17", err)
				}
			}
			// encode udc_r17 optional
			if ie.udc_r17 != nil {
				if err = ie.udc_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode udc_r17", err)
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

func (ie *PDCP_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkOnlyROHC_ProfilesPresent bool
	if uplinkOnlyROHC_ProfilesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var continueROHC_ContextPresent bool
	if continueROHC_ContextPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var outOfOrderDeliveryPresent bool
	if outOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var shortSNPresent bool
	if shortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_DuplicationSRBPresent bool
	if pdcp_DuplicationSRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_DuplicationMCG_OrSCG_DRBPresent bool
	if pdcp_DuplicationMCG_OrSCG_DRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedROHC_Profiles.Decode(r); err != nil {
		return utils.WrapError("Decode supportedROHC_Profiles", err)
	}
	if err = ie.maxNumberROHC_ContextSessions.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberROHC_ContextSessions", err)
	}
	if uplinkOnlyROHC_ProfilesPresent {
		ie.uplinkOnlyROHC_Profiles = new(PDCP_Parameters_uplinkOnlyROHC_Profiles)
		if err = ie.uplinkOnlyROHC_Profiles.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkOnlyROHC_Profiles", err)
		}
	}
	if continueROHC_ContextPresent {
		ie.continueROHC_Context = new(PDCP_Parameters_continueROHC_Context)
		if err = ie.continueROHC_Context.Decode(r); err != nil {
			return utils.WrapError("Decode continueROHC_Context", err)
		}
	}
	if outOfOrderDeliveryPresent {
		ie.outOfOrderDelivery = new(PDCP_Parameters_outOfOrderDelivery)
		if err = ie.outOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode outOfOrderDelivery", err)
		}
	}
	if shortSNPresent {
		ie.shortSN = new(PDCP_Parameters_shortSN)
		if err = ie.shortSN.Decode(r); err != nil {
			return utils.WrapError("Decode shortSN", err)
		}
	}
	if pdcp_DuplicationSRBPresent {
		ie.pdcp_DuplicationSRB = new(PDCP_Parameters_pdcp_DuplicationSRB)
		if err = ie.pdcp_DuplicationSRB.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_DuplicationSRB", err)
		}
	}
	if pdcp_DuplicationMCG_OrSCG_DRBPresent {
		ie.pdcp_DuplicationMCG_OrSCG_DRB = new(PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB)
		if err = ie.pdcp_DuplicationMCG_OrSCG_DRB.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_DuplicationMCG_OrSCG_DRB", err)
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

			drb_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			non_DRB_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedDiscardTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			continueEHC_Context_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ehc_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberEHC_Contexts_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			jointEHC_ROHC_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcp_DuplicationMoreThanTwoRLC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode drb_IAB_r16 optional
			if drb_IAB_r16Present {
				ie.drb_IAB_r16 = new(PDCP_Parameters_drb_IAB_r16)
				if err = ie.drb_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode drb_IAB_r16", err)
				}
			}
			// decode non_DRB_IAB_r16 optional
			if non_DRB_IAB_r16Present {
				ie.non_DRB_IAB_r16 = new(PDCP_Parameters_non_DRB_IAB_r16)
				if err = ie.non_DRB_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode non_DRB_IAB_r16", err)
				}
			}
			// decode extendedDiscardTimer_r16 optional
			if extendedDiscardTimer_r16Present {
				ie.extendedDiscardTimer_r16 = new(PDCP_Parameters_extendedDiscardTimer_r16)
				if err = ie.extendedDiscardTimer_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedDiscardTimer_r16", err)
				}
			}
			// decode continueEHC_Context_r16 optional
			if continueEHC_Context_r16Present {
				ie.continueEHC_Context_r16 = new(PDCP_Parameters_continueEHC_Context_r16)
				if err = ie.continueEHC_Context_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode continueEHC_Context_r16", err)
				}
			}
			// decode ehc_r16 optional
			if ehc_r16Present {
				ie.ehc_r16 = new(PDCP_Parameters_ehc_r16)
				if err = ie.ehc_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ehc_r16", err)
				}
			}
			// decode maxNumberEHC_Contexts_r16 optional
			if maxNumberEHC_Contexts_r16Present {
				ie.maxNumberEHC_Contexts_r16 = new(PDCP_Parameters_maxNumberEHC_Contexts_r16)
				if err = ie.maxNumberEHC_Contexts_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberEHC_Contexts_r16", err)
				}
			}
			// decode jointEHC_ROHC_Config_r16 optional
			if jointEHC_ROHC_Config_r16Present {
				ie.jointEHC_ROHC_Config_r16 = new(PDCP_Parameters_jointEHC_ROHC_Config_r16)
				if err = ie.jointEHC_ROHC_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode jointEHC_ROHC_Config_r16", err)
				}
			}
			// decode pdcp_DuplicationMoreThanTwoRLC_r16 optional
			if pdcp_DuplicationMoreThanTwoRLC_r16Present {
				ie.pdcp_DuplicationMoreThanTwoRLC_r16 = new(PDCP_Parameters_pdcp_DuplicationMoreThanTwoRLC_r16)
				if err = ie.pdcp_DuplicationMoreThanTwoRLC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcp_DuplicationMoreThanTwoRLC_r16", err)
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

			longSN_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			udc_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode longSN_RedCap_r17 optional
			if longSN_RedCap_r17Present {
				ie.longSN_RedCap_r17 = new(PDCP_Parameters_longSN_RedCap_r17)
				if err = ie.longSN_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode longSN_RedCap_r17", err)
				}
			}
			// decode udc_r17 optional
			if udc_r17Present {
				ie.udc_r17 = new(PDCP_Parameters_udc_r17)
				if err = ie.udc_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode udc_r17", err)
				}
			}
		}
	}
	return nil
}
