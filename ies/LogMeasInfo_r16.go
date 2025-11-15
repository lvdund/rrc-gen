package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasInfo_r16 struct {
	locationInfo_r16             *LocationInfo_r16                             `optional`
	relativeTimeStamp_r16        int64                                         `lb:0,ub:7200,madatory`
	servCellIdentity_r16         *CGI_Info_Logging_r16                         `optional`
	measResultServingCell_r16    *MeasResultServingCell_r16                    `optional`
	measResultNeighCells_r16     *LogMeasInfo_r16_measResultNeighCells_r16     `optional`
	anyCellSelectionDetected_r16 *LogMeasInfo_r16_anyCellSelectionDetected_r16 `optional`
	inDeviceCoexDetected_r17     *LogMeasInfo_r16_inDeviceCoexDetected_r17     `optional,ext-1`
}

func (ie *LogMeasInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.inDeviceCoexDetected_r17 != nil
	preambleBits := []bool{hasExtensions, ie.locationInfo_r16 != nil, ie.servCellIdentity_r16 != nil, ie.measResultServingCell_r16 != nil, ie.measResultNeighCells_r16 != nil, ie.anyCellSelectionDetected_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.locationInfo_r16 != nil {
		if err = ie.locationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode locationInfo_r16", err)
		}
	}
	if err = w.WriteInteger(ie.relativeTimeStamp_r16, &uper.Constraint{Lb: 0, Ub: 7200}, false); err != nil {
		return utils.WrapError("WriteInteger relativeTimeStamp_r16", err)
	}
	if ie.servCellIdentity_r16 != nil {
		if err = ie.servCellIdentity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode servCellIdentity_r16", err)
		}
	}
	if ie.measResultServingCell_r16 != nil {
		if err = ie.measResultServingCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultServingCell_r16", err)
		}
	}
	if ie.measResultNeighCells_r16 != nil {
		if err = ie.measResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCells_r16", err)
		}
	}
	if ie.anyCellSelectionDetected_r16 != nil {
		if err = ie.anyCellSelectionDetected_r16.Encode(w); err != nil {
			return utils.WrapError("Encode anyCellSelectionDetected_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.inDeviceCoexDetected_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap LogMeasInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.inDeviceCoexDetected_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode inDeviceCoexDetected_r17 optional
			if ie.inDeviceCoexDetected_r17 != nil {
				if err = ie.inDeviceCoexDetected_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode inDeviceCoexDetected_r17", err)
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

func (ie *LogMeasInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var locationInfo_r16Present bool
	if locationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var servCellIdentity_r16Present bool
	if servCellIdentity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultServingCell_r16Present bool
	if measResultServingCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCells_r16Present bool
	if measResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var anyCellSelectionDetected_r16Present bool
	if anyCellSelectionDetected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if locationInfo_r16Present {
		ie.locationInfo_r16 = new(LocationInfo_r16)
		if err = ie.locationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode locationInfo_r16", err)
		}
	}
	var tmp_int_relativeTimeStamp_r16 int64
	if tmp_int_relativeTimeStamp_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7200}, false); err != nil {
		return utils.WrapError("ReadInteger relativeTimeStamp_r16", err)
	}
	ie.relativeTimeStamp_r16 = tmp_int_relativeTimeStamp_r16
	if servCellIdentity_r16Present {
		ie.servCellIdentity_r16 = new(CGI_Info_Logging_r16)
		if err = ie.servCellIdentity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servCellIdentity_r16", err)
		}
	}
	if measResultServingCell_r16Present {
		ie.measResultServingCell_r16 = new(MeasResultServingCell_r16)
		if err = ie.measResultServingCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultServingCell_r16", err)
		}
	}
	if measResultNeighCells_r16Present {
		ie.measResultNeighCells_r16 = new(LogMeasInfo_r16_measResultNeighCells_r16)
		if err = ie.measResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCells_r16", err)
		}
	}
	if anyCellSelectionDetected_r16Present {
		ie.anyCellSelectionDetected_r16 = new(LogMeasInfo_r16_anyCellSelectionDetected_r16)
		if err = ie.anyCellSelectionDetected_r16.Decode(r); err != nil {
			return utils.WrapError("Decode anyCellSelectionDetected_r16", err)
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

			inDeviceCoexDetected_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode inDeviceCoexDetected_r17 optional
			if inDeviceCoexDetected_r17Present {
				ie.inDeviceCoexDetected_r17 = new(LogMeasInfo_r16_inDeviceCoexDetected_r17)
				if err = ie.inDeviceCoexDetected_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode inDeviceCoexDetected_r17", err)
				}
			}
		}
	}
	return nil
}
