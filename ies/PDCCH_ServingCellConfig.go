package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ServingCellConfig struct {
	slotFormatIndicator          *SlotFormatIndicator       `optional,setuprelease`
	availabilityIndicator_r16    *AvailabilityIndicator_r16 `optional,ext-1,setuprelease`
	searchSpaceSwitchTimer_r16   *int64                     `lb:1,ub:80,optional,ext-1`
	searchSpaceSwitchTimer_v1710 *int64                     `lb:81,ub:1280,optional,ext-2`
}

func (ie *PDCCH_ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.availabilityIndicator_r16 != nil || ie.searchSpaceSwitchTimer_r16 != nil || ie.searchSpaceSwitchTimer_v1710 != nil
	preambleBits := []bool{hasExtensions, ie.slotFormatIndicator != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.slotFormatIndicator != nil {
		tmp_slotFormatIndicator := utils.SetupRelease[*SlotFormatIndicator]{
			Setup: ie.slotFormatIndicator,
		}
		if err = tmp_slotFormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode slotFormatIndicator", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.availabilityIndicator_r16 != nil || ie.searchSpaceSwitchTimer_r16 != nil, ie.searchSpaceSwitchTimer_v1710 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.availabilityIndicator_r16 != nil, ie.searchSpaceSwitchTimer_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode availabilityIndicator_r16 optional
			if ie.availabilityIndicator_r16 != nil {
				tmp_availabilityIndicator_r16 := utils.SetupRelease[*AvailabilityIndicator_r16]{
					Setup: ie.availabilityIndicator_r16,
				}
				if err = tmp_availabilityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availabilityIndicator_r16", err)
				}
			}
			// encode searchSpaceSwitchTimer_r16 optional
			if ie.searchSpaceSwitchTimer_r16 != nil {
				if err = extWriter.WriteInteger(*ie.searchSpaceSwitchTimer_r16, &uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
					return utils.WrapError("Encode searchSpaceSwitchTimer_r16", err)
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
			optionals_ext_2 := []bool{ie.searchSpaceSwitchTimer_v1710 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode searchSpaceSwitchTimer_v1710 optional
			if ie.searchSpaceSwitchTimer_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.searchSpaceSwitchTimer_v1710, &uper.Constraint{Lb: 81, Ub: 1280}, false); err != nil {
					return utils.WrapError("Encode searchSpaceSwitchTimer_v1710", err)
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

func (ie *PDCCH_ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var slotFormatIndicatorPresent bool
	if slotFormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if slotFormatIndicatorPresent {
		tmp_slotFormatIndicator := utils.SetupRelease[*SlotFormatIndicator]{}
		if err = tmp_slotFormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode slotFormatIndicator", err)
		}
		ie.slotFormatIndicator = tmp_slotFormatIndicator.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			availabilityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceSwitchTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode availabilityIndicator_r16 optional
			if availabilityIndicator_r16Present {
				tmp_availabilityIndicator_r16 := utils.SetupRelease[*AvailabilityIndicator_r16]{}
				if err = tmp_availabilityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode availabilityIndicator_r16", err)
				}
				ie.availabilityIndicator_r16 = tmp_availabilityIndicator_r16.Setup
			}
			// decode searchSpaceSwitchTimer_r16 optional
			if searchSpaceSwitchTimer_r16Present {
				var tmp_int_searchSpaceSwitchTimer_r16 int64
				if tmp_int_searchSpaceSwitchTimer_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
					return utils.WrapError("Decode searchSpaceSwitchTimer_r16", err)
				}
				ie.searchSpaceSwitchTimer_r16 = &tmp_int_searchSpaceSwitchTimer_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			searchSpaceSwitchTimer_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode searchSpaceSwitchTimer_v1710 optional
			if searchSpaceSwitchTimer_v1710Present {
				var tmp_int_searchSpaceSwitchTimer_v1710 int64
				if tmp_int_searchSpaceSwitchTimer_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 81, Ub: 1280}, false); err != nil {
					return utils.WrapError("Decode searchSpaceSwitchTimer_v1710", err)
				}
				ie.searchSpaceSwitchTimer_v1710 = &tmp_int_searchSpaceSwitchTimer_v1710
			}
		}
	}
	return nil
}
