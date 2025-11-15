package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VisitedCellInfo_r16 struct {
	visitedCellId_r16               *VisitedCellInfo_r16_visitedCellId_r16 `optional`
	timeSpent_r16                   int64                                  `lb:0,ub:4095,madatory`
	visitedPSCellInfoListReport_r17 *VisitedPSCellInfoList_r17             `optional,ext-1`
}

func (ie *VisitedCellInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.visitedPSCellInfoListReport_r17 != nil
	preambleBits := []bool{hasExtensions, ie.visitedCellId_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.visitedCellId_r16 != nil {
		if err = ie.visitedCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode visitedCellId_r16", err)
		}
	}
	if err = w.WriteInteger(ie.timeSpent_r16, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("WriteInteger timeSpent_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.visitedPSCellInfoListReport_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap VisitedCellInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.visitedPSCellInfoListReport_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode visitedPSCellInfoListReport_r17 optional
			if ie.visitedPSCellInfoListReport_r17 != nil {
				if err = ie.visitedPSCellInfoListReport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode visitedPSCellInfoListReport_r17", err)
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

func (ie *VisitedCellInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var visitedCellId_r16Present bool
	if visitedCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if visitedCellId_r16Present {
		ie.visitedCellId_r16 = new(VisitedCellInfo_r16_visitedCellId_r16)
		if err = ie.visitedCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode visitedCellId_r16", err)
		}
	}
	var tmp_int_timeSpent_r16 int64
	if tmp_int_timeSpent_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("ReadInteger timeSpent_r16", err)
	}
	ie.timeSpent_r16 = tmp_int_timeSpent_r16

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

			visitedPSCellInfoListReport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode visitedPSCellInfoListReport_r17 optional
			if visitedPSCellInfoListReport_r17Present {
				ie.visitedPSCellInfoListReport_r17 = new(VisitedPSCellInfoList_r17)
				if err = ie.visitedPSCellInfoListReport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode visitedPSCellInfoListReport_r17", err)
				}
			}
		}
	}
	return nil
}
