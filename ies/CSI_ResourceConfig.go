package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ResourceConfig struct {
	csi_ResourceConfigId           CSI_ResourceConfigId                       `madatory`
	csi_RS_ResourceSetList         *CSI_ResourceConfig_csi_RS_ResourceSetList `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	bwp_Id                         BWP_Id                                     `madatory`
	resourceType                   CSI_ResourceConfig_resourceType            `madatory`
	csi_SSB_ResourceSetListExt_r17 *CSI_SSB_ResourceSetId                     `optional,ext-1`
}

func (ie *CSI_ResourceConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.csi_SSB_ResourceSetListExt_r17 != nil
	preambleBits := []bool{hasExtensions, ie.csi_RS_ResourceSetList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_ResourceConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode csi_ResourceConfigId", err)
	}
	if ie.csi_RS_ResourceSetList != nil {
		if err = ie.csi_RS_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_ResourceSetList", err)
		}
	}
	if err = ie.bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode bwp_Id", err)
	}
	if err = ie.resourceType.Encode(w); err != nil {
		return utils.WrapError("Encode resourceType", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.csi_SSB_ResourceSetListExt_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_ResourceConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.csi_SSB_ResourceSetListExt_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode csi_SSB_ResourceSetListExt_r17 optional
			if ie.csi_SSB_ResourceSetListExt_r17 != nil {
				if err = ie.csi_SSB_ResourceSetListExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_SSB_ResourceSetListExt_r17", err)
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

func (ie *CSI_ResourceConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_ResourceSetListPresent bool
	if csi_RS_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_ResourceConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode csi_ResourceConfigId", err)
	}
	if csi_RS_ResourceSetListPresent {
		ie.csi_RS_ResourceSetList = new(CSI_ResourceConfig_csi_RS_ResourceSetList)
		if err = ie.csi_RS_ResourceSetList.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_ResourceSetList", err)
		}
	}
	if err = ie.bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode bwp_Id", err)
	}
	if err = ie.resourceType.Decode(r); err != nil {
		return utils.WrapError("Decode resourceType", err)
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

			csi_SSB_ResourceSetListExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode csi_SSB_ResourceSetListExt_r17 optional
			if csi_SSB_ResourceSetListExt_r17Present {
				ie.csi_SSB_ResourceSetListExt_r17 = new(CSI_SSB_ResourceSetId)
				if err = ie.csi_SSB_ResourceSetListExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_SSB_ResourceSetListExt_r17", err)
				}
			}
		}
	}
	return nil
}
