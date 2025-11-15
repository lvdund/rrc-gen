package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2 struct {
	cellReselectionInfoCommon      *SIB2_cellReselectionInfoCommon      `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	cellReselectionServingFreqInfo *SIB2_cellReselectionServingFreqInfo `optional,ext`
	intraFreqCellReselectionInfo   *SIB2_intraFreqCellReselectionInfo   `optional,ext`
	relaxedMeasurement_r16         *SIB2_relaxedMeasurement_r16         `optional,ext-5`
	cellEquivalentSize_r17         *int64                               `lb:2,ub:16,optional,ext-6`
	relaxedMeasurement_r17         *SIB2_relaxedMeasurement_r17         `optional,ext-6`
}

func (ie *SIB2) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.relaxedMeasurement_r16 != nil || ie.cellEquivalentSize_r17 != nil || ie.relaxedMeasurement_r17 != nil
	preambleBits := []bool{hasExtensions, ie.cellReselectionInfoCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellReselectionInfoCommon != nil {
		if err = ie.cellReselectionInfoCommon.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionInfoCommon", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.relaxedMeasurement_r16 != nil, ie.cellEquivalentSize_r17 != nil || ie.relaxedMeasurement_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB2", err)
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.relaxedMeasurement_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode relaxedMeasurement_r16 optional
			if ie.relaxedMeasurement_r16 != nil {
				if err = ie.relaxedMeasurement_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode relaxedMeasurement_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.cellEquivalentSize_r17 != nil, ie.relaxedMeasurement_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cellEquivalentSize_r17 optional
			if ie.cellEquivalentSize_r17 != nil {
				if err = extWriter.WriteInteger(*ie.cellEquivalentSize_r17, &uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
					return utils.WrapError("Encode cellEquivalentSize_r17", err)
				}
			}
			// encode relaxedMeasurement_r17 optional
			if ie.relaxedMeasurement_r17 != nil {
				if err = ie.relaxedMeasurement_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode relaxedMeasurement_r17", err)
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

func (ie *SIB2) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReselectionInfoCommonPresent bool
	if cellReselectionInfoCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cellReselectionInfoCommonPresent {
		ie.cellReselectionInfoCommon = new(SIB2_cellReselectionInfoCommon)
		if err = ie.cellReselectionInfoCommon.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionInfoCommon", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			relaxedMeasurement_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode relaxedMeasurement_r16 optional
			if relaxedMeasurement_r16Present {
				ie.relaxedMeasurement_r16 = new(SIB2_relaxedMeasurement_r16)
				if err = ie.relaxedMeasurement_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode relaxedMeasurement_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			cellEquivalentSize_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			relaxedMeasurement_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cellEquivalentSize_r17 optional
			if cellEquivalentSize_r17Present {
				var tmp_int_cellEquivalentSize_r17 int64
				if tmp_int_cellEquivalentSize_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
					return utils.WrapError("Decode cellEquivalentSize_r17", err)
				}
				ie.cellEquivalentSize_r17 = &tmp_int_cellEquivalentSize_r17
			}
			// decode relaxedMeasurement_r17 optional
			if relaxedMeasurement_r17Present {
				ie.relaxedMeasurement_r17 = new(SIB2_relaxedMeasurement_r17)
				if err = ie.relaxedMeasurement_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode relaxedMeasurement_r17", err)
				}
			}
		}
	}
	return nil
}
