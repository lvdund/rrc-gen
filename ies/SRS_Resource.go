package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource struct {
	srs_ResourceId               SRS_ResourceId                             `madatory`
	nrofSRS_Ports                SRS_Resource_nrofSRS_Ports                 `madatory`
	ptrs_PortIndex               *SRS_Resource_ptrs_PortIndex               `optional`
	transmissionComb             SRS_Resource_transmissionComb              `lb:0,ub:1,madatory`
	resourceMapping              SRS_Resource_resourceMapping               `lb:0,ub:5,madatory`
	freqDomainPosition           int64                                      `lb:0,ub:67,madatory`
	freqDomainShift              int64                                      `lb:0,ub:268,madatory`
	freqHopping                  SRS_Resource_freqHopping                   `lb:0,ub:63,madatory`
	groupOrSequenceHopping       SRS_Resource_groupOrSequenceHopping        `madatory`
	resourceType                 SRS_Resource_resourceType                  `madatory`
	sequenceId                   int64                                      `lb:0,ub:1023,madatory,ext`
	spatialRelationInfo          *SRS_SpatialRelationInfo                   `optional,ext`
	resourceMapping_r16          *SRS_Resource_resourceMapping_r16          `lb:0,ub:13,optional,ext-1`
	spatialRelationInfo_PDC_r17  *SpatialRelationInfo_PDC_r17               `optional,ext-2,setuprelease`
	resourceMapping_r17          *SRS_Resource_resourceMapping_r17          `lb:0,ub:13,optional,ext-2`
	partialFreqSounding_r17      *SRS_Resource_partialFreqSounding_r17      `lb:0,ub:1,optional,ext-2`
	transmissionComb_n8_r17      *SRS_Resource_transmissionComb_n8_r17      `lb:0,ub:7,optional,ext-2`
	srs_TCI_State_r17            *SRS_Resource_srs_TCI_State_r17            `optional,ext-2`
	repetitionFactor_v1730       *SRS_Resource_repetitionFactor_v1730       `optional,ext-3`
	srs_DLorJointTCI_State_v1730 *SRS_Resource_srs_DLorJointTCI_State_v1730 `optional,ext-3`
}

func (ie *SRS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.resourceMapping_r16 != nil || ie.spatialRelationInfo_PDC_r17 != nil || ie.resourceMapping_r17 != nil || ie.partialFreqSounding_r17 != nil || ie.transmissionComb_n8_r17 != nil || ie.srs_TCI_State_r17 != nil || ie.repetitionFactor_v1730 != nil || ie.srs_DLorJointTCI_State_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.ptrs_PortIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srs_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode srs_ResourceId", err)
	}
	if err = ie.nrofSRS_Ports.Encode(w); err != nil {
		return utils.WrapError("Encode nrofSRS_Ports", err)
	}
	if ie.ptrs_PortIndex != nil {
		if err = ie.ptrs_PortIndex.Encode(w); err != nil {
			return utils.WrapError("Encode ptrs_PortIndex", err)
		}
	}
	if err = ie.transmissionComb.Encode(w); err != nil {
		return utils.WrapError("Encode transmissionComb", err)
	}
	if err = ie.resourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode resourceMapping", err)
	}
	if err = w.WriteInteger(ie.freqDomainPosition, &uper.Constraint{Lb: 0, Ub: 67}, false); err != nil {
		return utils.WrapError("WriteInteger freqDomainPosition", err)
	}
	if err = w.WriteInteger(ie.freqDomainShift, &uper.Constraint{Lb: 0, Ub: 268}, false); err != nil {
		return utils.WrapError("WriteInteger freqDomainShift", err)
	}
	if err = ie.freqHopping.Encode(w); err != nil {
		return utils.WrapError("Encode freqHopping", err)
	}
	if err = ie.groupOrSequenceHopping.Encode(w); err != nil {
		return utils.WrapError("Encode groupOrSequenceHopping", err)
	}
	if err = ie.resourceType.Encode(w); err != nil {
		return utils.WrapError("Encode resourceType", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.resourceMapping_r16 != nil, ie.spatialRelationInfo_PDC_r17 != nil || ie.resourceMapping_r17 != nil || ie.partialFreqSounding_r17 != nil || ie.transmissionComb_n8_r17 != nil || ie.srs_TCI_State_r17 != nil, ie.repetitionFactor_v1730 != nil || ie.srs_DLorJointTCI_State_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_Resource", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.resourceMapping_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode resourceMapping_r16 optional
			if ie.resourceMapping_r16 != nil {
				if err = ie.resourceMapping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceMapping_r16", err)
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
			optionals_ext_2 := []bool{ie.spatialRelationInfo_PDC_r17 != nil, ie.resourceMapping_r17 != nil, ie.partialFreqSounding_r17 != nil, ie.transmissionComb_n8_r17 != nil, ie.srs_TCI_State_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spatialRelationInfo_PDC_r17 optional
			if ie.spatialRelationInfo_PDC_r17 != nil {
				tmp_spatialRelationInfo_PDC_r17 := utils.SetupRelease[*SpatialRelationInfo_PDC_r17]{
					Setup: ie.spatialRelationInfo_PDC_r17,
				}
				if err = tmp_spatialRelationInfo_PDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationInfo_PDC_r17", err)
				}
			}
			// encode resourceMapping_r17 optional
			if ie.resourceMapping_r17 != nil {
				if err = ie.resourceMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceMapping_r17", err)
				}
			}
			// encode partialFreqSounding_r17 optional
			if ie.partialFreqSounding_r17 != nil {
				if err = ie.partialFreqSounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode partialFreqSounding_r17", err)
				}
			}
			// encode transmissionComb_n8_r17 optional
			if ie.transmissionComb_n8_r17 != nil {
				if err = ie.transmissionComb_n8_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode transmissionComb_n8_r17", err)
				}
			}
			// encode srs_TCI_State_r17 optional
			if ie.srs_TCI_State_r17 != nil {
				if err = ie.srs_TCI_State_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_TCI_State_r17", err)
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
			optionals_ext_3 := []bool{ie.repetitionFactor_v1730 != nil, ie.srs_DLorJointTCI_State_v1730 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode repetitionFactor_v1730 optional
			if ie.repetitionFactor_v1730 != nil {
				if err = ie.repetitionFactor_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode repetitionFactor_v1730", err)
				}
			}
			// encode srs_DLorJointTCI_State_v1730 optional
			if ie.srs_DLorJointTCI_State_v1730 != nil {
				if err = ie.srs_DLorJointTCI_State_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_DLorJointTCI_State_v1730", err)
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

func (ie *SRS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ptrs_PortIndexPresent bool
	if ptrs_PortIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srs_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode srs_ResourceId", err)
	}
	if err = ie.nrofSRS_Ports.Decode(r); err != nil {
		return utils.WrapError("Decode nrofSRS_Ports", err)
	}
	if ptrs_PortIndexPresent {
		ie.ptrs_PortIndex = new(SRS_Resource_ptrs_PortIndex)
		if err = ie.ptrs_PortIndex.Decode(r); err != nil {
			return utils.WrapError("Decode ptrs_PortIndex", err)
		}
	}
	if err = ie.transmissionComb.Decode(r); err != nil {
		return utils.WrapError("Decode transmissionComb", err)
	}
	if err = ie.resourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode resourceMapping", err)
	}
	var tmp_int_freqDomainPosition int64
	if tmp_int_freqDomainPosition, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 67}, false); err != nil {
		return utils.WrapError("ReadInteger freqDomainPosition", err)
	}
	ie.freqDomainPosition = tmp_int_freqDomainPosition
	var tmp_int_freqDomainShift int64
	if tmp_int_freqDomainShift, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 268}, false); err != nil {
		return utils.WrapError("ReadInteger freqDomainShift", err)
	}
	ie.freqDomainShift = tmp_int_freqDomainShift
	if err = ie.freqHopping.Decode(r); err != nil {
		return utils.WrapError("Decode freqHopping", err)
	}
	if err = ie.groupOrSequenceHopping.Decode(r); err != nil {
		return utils.WrapError("Decode groupOrSequenceHopping", err)
	}
	if err = ie.resourceType.Decode(r); err != nil {
		return utils.WrapError("Decode resourceType", err)
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

			resourceMapping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode resourceMapping_r16 optional
			if resourceMapping_r16Present {
				ie.resourceMapping_r16 = new(SRS_Resource_resourceMapping_r16)
				if err = ie.resourceMapping_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceMapping_r16", err)
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

			spatialRelationInfo_PDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			partialFreqSounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			transmissionComb_n8_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_TCI_State_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spatialRelationInfo_PDC_r17 optional
			if spatialRelationInfo_PDC_r17Present {
				tmp_spatialRelationInfo_PDC_r17 := utils.SetupRelease[*SpatialRelationInfo_PDC_r17]{}
				if err = tmp_spatialRelationInfo_PDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelationInfo_PDC_r17", err)
				}
				ie.spatialRelationInfo_PDC_r17 = tmp_spatialRelationInfo_PDC_r17.Setup
			}
			// decode resourceMapping_r17 optional
			if resourceMapping_r17Present {
				ie.resourceMapping_r17 = new(SRS_Resource_resourceMapping_r17)
				if err = ie.resourceMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceMapping_r17", err)
				}
			}
			// decode partialFreqSounding_r17 optional
			if partialFreqSounding_r17Present {
				ie.partialFreqSounding_r17 = new(SRS_Resource_partialFreqSounding_r17)
				if err = ie.partialFreqSounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode partialFreqSounding_r17", err)
				}
			}
			// decode transmissionComb_n8_r17 optional
			if transmissionComb_n8_r17Present {
				ie.transmissionComb_n8_r17 = new(SRS_Resource_transmissionComb_n8_r17)
				if err = ie.transmissionComb_n8_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode transmissionComb_n8_r17", err)
				}
			}
			// decode srs_TCI_State_r17 optional
			if srs_TCI_State_r17Present {
				ie.srs_TCI_State_r17 = new(SRS_Resource_srs_TCI_State_r17)
				if err = ie.srs_TCI_State_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_TCI_State_r17", err)
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

			repetitionFactor_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_DLorJointTCI_State_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode repetitionFactor_v1730 optional
			if repetitionFactor_v1730Present {
				ie.repetitionFactor_v1730 = new(SRS_Resource_repetitionFactor_v1730)
				if err = ie.repetitionFactor_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode repetitionFactor_v1730", err)
				}
			}
			// decode srs_DLorJointTCI_State_v1730 optional
			if srs_DLorJointTCI_State_v1730Present {
				ie.srs_DLorJointTCI_State_v1730 = new(SRS_Resource_srs_DLorJointTCI_State_v1730)
				if err = ie.srs_DLorJointTCI_State_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_DLorJointTCI_State_v1730", err)
				}
			}
		}
	}
	return nil
}
