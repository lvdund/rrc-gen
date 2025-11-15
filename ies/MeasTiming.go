package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasTiming struct {
	frequencyAndTiming *MeasTiming_frequencyAndTiming `optional`
	ssb_ToMeasure      *SSB_ToMeasure                 `optional,ext-1`
	physCellId         *PhysCellId                    `optional,ext-1`
}

func (ie *MeasTiming) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ssb_ToMeasure != nil || ie.physCellId != nil
	preambleBits := []bool{hasExtensions, ie.frequencyAndTiming != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyAndTiming != nil {
		if err = ie.frequencyAndTiming.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyAndTiming", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ssb_ToMeasure != nil || ie.physCellId != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasTiming", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ssb_ToMeasure != nil, ie.physCellId != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ssb_ToMeasure optional
			if ie.ssb_ToMeasure != nil {
				if err = ie.ssb_ToMeasure.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_ToMeasure", err)
				}
			}
			// encode physCellId optional
			if ie.physCellId != nil {
				if err = ie.physCellId.Encode(extWriter); err != nil {
					return utils.WrapError("Encode physCellId", err)
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

func (ie *MeasTiming) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyAndTimingPresent bool
	if frequencyAndTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyAndTimingPresent {
		ie.frequencyAndTiming = new(MeasTiming_frequencyAndTiming)
		if err = ie.frequencyAndTiming.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyAndTiming", err)
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

			ssb_ToMeasurePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			physCellIdPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ssb_ToMeasure optional
			if ssb_ToMeasurePresent {
				ie.ssb_ToMeasure = new(SSB_ToMeasure)
				if err = ie.ssb_ToMeasure.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_ToMeasure", err)
				}
			}
			// decode physCellId optional
			if physCellIdPresent {
				ie.physCellId = new(PhysCellId)
				if err = ie.physCellId.Decode(extReader); err != nil {
					return utils.WrapError("Decode physCellId", err)
				}
			}
		}
	}
	return nil
}
