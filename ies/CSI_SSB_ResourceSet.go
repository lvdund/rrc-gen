package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SSB_ResourceSet struct {
	csi_SSB_ResourceSetId        CSI_SSB_ResourceSetId           `madatory`
	csi_SSB_ResourceList         []SSB_Index                     `lb:1,ub:maxNrofCSI_SSB_ResourcePerSet,madatory`
	servingAdditionalPCIList_r17 []ServingAdditionalPCIIndex_r17 `lb:1,ub:maxNrofCSI_SSB_ResourcePerSet,optional,ext-1`
}

func (ie *CSI_SSB_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.servingAdditionalPCIList_r17) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_SSB_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode csi_SSB_ResourceSetId", err)
	}
	tmp_csi_SSB_ResourceList := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
	for _, i := range ie.csi_SSB_ResourceList {
		tmp_csi_SSB_ResourceList.Value = append(tmp_csi_SSB_ResourceList.Value, &i)
	}
	if err = tmp_csi_SSB_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode csi_SSB_ResourceList", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.servingAdditionalPCIList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_SSB_ResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.servingAdditionalPCIList_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode servingAdditionalPCIList_r17 optional
			if len(ie.servingAdditionalPCIList_r17) > 0 {
				tmp_servingAdditionalPCIList_r17 := utils.NewSequence[*ServingAdditionalPCIIndex_r17]([]*ServingAdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
				for _, i := range ie.servingAdditionalPCIList_r17 {
					tmp_servingAdditionalPCIList_r17.Value = append(tmp_servingAdditionalPCIList_r17.Value, &i)
				}
				if err = tmp_servingAdditionalPCIList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode servingAdditionalPCIList_r17", err)
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

func (ie *CSI_SSB_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_SSB_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode csi_SSB_ResourceSetId", err)
	}
	tmp_csi_SSB_ResourceList := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
	fn_csi_SSB_ResourceList := func() *SSB_Index {
		return new(SSB_Index)
	}
	if err = tmp_csi_SSB_ResourceList.Decode(r, fn_csi_SSB_ResourceList); err != nil {
		return utils.WrapError("Decode csi_SSB_ResourceList", err)
	}
	ie.csi_SSB_ResourceList = []SSB_Index{}
	for _, i := range tmp_csi_SSB_ResourceList.Value {
		ie.csi_SSB_ResourceList = append(ie.csi_SSB_ResourceList, *i)
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

			servingAdditionalPCIList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode servingAdditionalPCIList_r17 optional
			if servingAdditionalPCIList_r17Present {
				tmp_servingAdditionalPCIList_r17 := utils.NewSequence[*ServingAdditionalPCIIndex_r17]([]*ServingAdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
				fn_servingAdditionalPCIList_r17 := func() *ServingAdditionalPCIIndex_r17 {
					return new(ServingAdditionalPCIIndex_r17)
				}
				if err = tmp_servingAdditionalPCIList_r17.Decode(extReader, fn_servingAdditionalPCIList_r17); err != nil {
					return utils.WrapError("Decode servingAdditionalPCIList_r17", err)
				}
				ie.servingAdditionalPCIList_r17 = []ServingAdditionalPCIIndex_r17{}
				for _, i := range tmp_servingAdditionalPCIList_r17.Value {
					ie.servingAdditionalPCIList_r17 = append(ie.servingAdditionalPCIList_r17, *i)
				}
			}
		}
	}
	return nil
}
