package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellReselectionPriorities struct {
	freqPriorityListEUTRA                *FreqPriorityListEUTRA                `optional`
	freqPriorityListNR                   *FreqPriorityListNR                   `optional`
	t320                                 *CellReselectionPriorities_t320       `optional`
	freqPriorityListDedicatedSlicing_r17 *FreqPriorityListDedicatedSlicing_r17 `optional,ext-1`
}

func (ie *CellReselectionPriorities) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.freqPriorityListDedicatedSlicing_r17 != nil
	preambleBits := []bool{hasExtensions, ie.freqPriorityListEUTRA != nil, ie.freqPriorityListNR != nil, ie.t320 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.freqPriorityListEUTRA != nil {
		if err = ie.freqPriorityListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode freqPriorityListEUTRA", err)
		}
	}
	if ie.freqPriorityListNR != nil {
		if err = ie.freqPriorityListNR.Encode(w); err != nil {
			return utils.WrapError("Encode freqPriorityListNR", err)
		}
	}
	if ie.t320 != nil {
		if err = ie.t320.Encode(w); err != nil {
			return utils.WrapError("Encode t320", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.freqPriorityListDedicatedSlicing_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellReselectionPriorities", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.freqPriorityListDedicatedSlicing_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode freqPriorityListDedicatedSlicing_r17 optional
			if ie.freqPriorityListDedicatedSlicing_r17 != nil {
				if err = ie.freqPriorityListDedicatedSlicing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode freqPriorityListDedicatedSlicing_r17", err)
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

func (ie *CellReselectionPriorities) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var freqPriorityListEUTRAPresent bool
	if freqPriorityListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var freqPriorityListNRPresent bool
	if freqPriorityListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var t320Present bool
	if t320Present, err = r.ReadBool(); err != nil {
		return err
	}
	if freqPriorityListEUTRAPresent {
		ie.freqPriorityListEUTRA = new(FreqPriorityListEUTRA)
		if err = ie.freqPriorityListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode freqPriorityListEUTRA", err)
		}
	}
	if freqPriorityListNRPresent {
		ie.freqPriorityListNR = new(FreqPriorityListNR)
		if err = ie.freqPriorityListNR.Decode(r); err != nil {
			return utils.WrapError("Decode freqPriorityListNR", err)
		}
	}
	if t320Present {
		ie.t320 = new(CellReselectionPriorities_t320)
		if err = ie.t320.Decode(r); err != nil {
			return utils.WrapError("Decode t320", err)
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

			freqPriorityListDedicatedSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode freqPriorityListDedicatedSlicing_r17 optional
			if freqPriorityListDedicatedSlicing_r17Present {
				ie.freqPriorityListDedicatedSlicing_r17 = new(FreqPriorityListDedicatedSlicing_r17)
				if err = ie.freqPriorityListDedicatedSlicing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode freqPriorityListDedicatedSlicing_r17", err)
				}
			}
		}
	}
	return nil
}
