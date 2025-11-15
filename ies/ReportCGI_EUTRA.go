package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportCGI_EUTRA struct {
	cellForWhichToReportCGI EUTRA_PhysCellId                       `madatory`
	useAutonomousGaps_r16   *ReportCGI_EUTRA_useAutonomousGaps_r16 `optional,ext-1`
}

func (ie *ReportCGI_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.useAutonomousGaps_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.cellForWhichToReportCGI.Encode(w); err != nil {
		return utils.WrapError("Encode cellForWhichToReportCGI", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.useAutonomousGaps_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ReportCGI_EUTRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.useAutonomousGaps_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode useAutonomousGaps_r16 optional
			if ie.useAutonomousGaps_r16 != nil {
				if err = ie.useAutonomousGaps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode useAutonomousGaps_r16", err)
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

func (ie *ReportCGI_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.cellForWhichToReportCGI.Decode(r); err != nil {
		return utils.WrapError("Decode cellForWhichToReportCGI", err)
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

			useAutonomousGaps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode useAutonomousGaps_r16 optional
			if useAutonomousGaps_r16Present {
				ie.useAutonomousGaps_r16 = new(ReportCGI_EUTRA_useAutonomousGaps_r16)
				if err = ie.useAutonomousGaps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode useAutonomousGaps_r16", err)
				}
			}
		}
	}
	return nil
}
