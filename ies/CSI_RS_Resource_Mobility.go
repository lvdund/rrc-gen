package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_Resource_Mobility struct {
	csi_RS_Index                CSI_RS_Index                                       `madatory`
	slotConfig                  CSI_RS_Resource_Mobility_slotConfig                `lb:0,ub:31,madatory`
	associatedSSB               *CSI_RS_Resource_Mobility_associatedSSB            `optional`
	frequencyDomainAllocation   CSI_RS_Resource_Mobility_frequencyDomainAllocation `lb:4,ub:4,madatory`
	firstOFDMSymbolInTimeDomain int64                                              `lb:0,ub:13,madatory`
	sequenceGenerationConfig    int64                                              `lb:0,ub:1023,madatory`
	slotConfig_r17              *CSI_RS_Resource_Mobility_slotConfig_r17           `lb:0,ub:255,optional,ext-1`
}

func (ie *CSI_RS_Resource_Mobility) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.slotConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.associatedSSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_RS_Index.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_Index", err)
	}
	if err = ie.slotConfig.Encode(w); err != nil {
		return utils.WrapError("Encode slotConfig", err)
	}
	if ie.associatedSSB != nil {
		if err = ie.associatedSSB.Encode(w); err != nil {
			return utils.WrapError("Encode associatedSSB", err)
		}
	}
	if err = ie.frequencyDomainAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyDomainAllocation", err)
	}
	if err = w.WriteInteger(ie.firstOFDMSymbolInTimeDomain, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger firstOFDMSymbolInTimeDomain", err)
	}
	if err = w.WriteInteger(ie.sequenceGenerationConfig, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger sequenceGenerationConfig", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.slotConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_RS_Resource_Mobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.slotConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode slotConfig_r17 optional
			if ie.slotConfig_r17 != nil {
				if err = ie.slotConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode slotConfig_r17", err)
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

func (ie *CSI_RS_Resource_Mobility) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var associatedSSBPresent bool
	if associatedSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_RS_Index.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_Index", err)
	}
	if err = ie.slotConfig.Decode(r); err != nil {
		return utils.WrapError("Decode slotConfig", err)
	}
	if associatedSSBPresent {
		ie.associatedSSB = new(CSI_RS_Resource_Mobility_associatedSSB)
		if err = ie.associatedSSB.Decode(r); err != nil {
			return utils.WrapError("Decode associatedSSB", err)
		}
	}
	if err = ie.frequencyDomainAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyDomainAllocation", err)
	}
	var tmp_int_firstOFDMSymbolInTimeDomain int64
	if tmp_int_firstOFDMSymbolInTimeDomain, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger firstOFDMSymbolInTimeDomain", err)
	}
	ie.firstOFDMSymbolInTimeDomain = tmp_int_firstOFDMSymbolInTimeDomain
	var tmp_int_sequenceGenerationConfig int64
	if tmp_int_sequenceGenerationConfig, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger sequenceGenerationConfig", err)
	}
	ie.sequenceGenerationConfig = tmp_int_sequenceGenerationConfig

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

			slotConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode slotConfig_r17 optional
			if slotConfig_r17Present {
				ie.slotConfig_r17 = new(CSI_RS_Resource_Mobility_slotConfig_r17)
				if err = ie.slotConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode slotConfig_r17", err)
				}
			}
		}
	}
	return nil
}
