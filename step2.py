import re, os
import json
from dataclasses import dataclass, field, asdict
from typing import Optional, List, Dict, Any, Tuple
from pathlib import Path
import sys
from collections import defaultdict


# Default input file
checking = False
if checking:
    json_path = Path(__file__).parent / "cache" / "check.json"
else:
    json_path = Path(__file__).parent / "cache" / "asn1_types.json"


@dataclass
class Asn1Field:
    """Represents a field in a SEQUENCE or CHOICE"""
    name: str
    type_name: str
    lb: Optional[Any] = None  # lower bound
    ub: Optional[Any] = None  # upper bound
    element_type: Optional[str] = None  # for SEQUENCE OF
    element_lb: Optional[Any] = None  # lower bound for SEQUENCE OF
    element_ub: Optional[Any] = None  # upper bound for SEQUENCE OF
    optional: bool = False
    ext: bool = False  # extension marker
    comment: Optional[str] = None
    
    # def __hash__(self):
    #     # Create a hash based on the fields that make the object unique
    #     return hash((
    #         self.name, 
    #         self.type_name, 
    #         self.lb, 
    #         self.ub, 
    #         self.element_type,
    #         self.element_lb,
    #         self.element_ub,
    #         self.optional,
    #         self.ext,
    #         self.comment
    #     ))
    
    def to_dict(self):
        return {k: v for k, v in asdict(self).items() if v is not None}


@dataclass
class Asn1Type:
    """Represents an ASN.1 type definition"""
    name: str
    type_name: str  # 'SEQUENCE', 'CHOICE', 'ENUMERATED', 'INTEGER', 'OCTET_STRING', etc.
    fields: List[Asn1Field] = field(default_factory=list)
    values: List[str] = field(default_factory=list)  # for ENUMERATED
    lb: Optional[Any] = None  # lower bound
    ub: Optional[Any] = None  # upper bound
    parent: Optional[str] = None
    element_type: Optional[str] = None  # for SEQUENCE OF
    element_lb: Optional[Any] = None  # lower bound for SEQUENCE OF
    element_ub: Optional[Any] = None  # upper bound for SEQUENCE OF
    ext: bool = False  # has extension marker "..."
    comment: Optional[str] = None
    
    def to_dict(self):
        result = {"name": self.name, "type_name": self.type_name}
        if self.fields:
            result["fields"] = [f.to_dict() for f in self.fields]
        if self.values:
            result["values"] = self.values
        if self.lb:
            result["lb"] = self.lb
        if self.ub:
            result["ub"] = self.ub
        if self.parent:
            result["parent"] = self.parent
        if self.element_type:
            result["element_type"] = self.element_type
        if self.element_lb:
            result["element_lb"] = self.element_lb
        if self.element_ub:
            result["element_ub"] = self.element_ub
        if self.ext:
            result["ext"] = self.ext
        if self.comment:
            result["comment"] = self.comment
        return result


def load_asn1_types(file_path: str) -> List[Asn1Type]:
    """Load ASN.1 types from JSON file and convert to Asn1Type objects"""
    asn1_types = []
    with open(file_path, 'r', encoding='utf-8') as f:
        data = json.load(f)
    for type_name, type_data in data.items():
        # Create Asn1Field objects for each field
        fields = []
        if 'fields' in type_data and type_data['fields']:
            for field_data in type_data['fields']:
                field = Asn1Field(
                    name=field_data.get('name', ''),
                    type_name=field_data.get('type_name', ''),
                    lb=field_data.get('lb'),
                    ub=field_data.get('ub'),
                    element_type=field_data.get('element_type'),
                    element_lb=field_data.get('element_lb'),
                    element_ub=field_data.get('element_ub'),
                    optional=field_data.get('optional', False),
                    ext=field_data.get('ext', False),
                    comment=field_data.get('comment')
                )
                fields.append(field)
        # Create Asn1Type object
        asn1_type = Asn1Type(
            name=type_data.get('name', type_name),
            type_name=type_data.get('type_name', ''),
            fields=fields,
            values=type_data.get('values', []),
            lb=type_data.get('lb'),
            ub=type_data.get('ub'),
            parent=type_data.get('parent'),
            element_type=type_data.get('element_type'),
            element_lb=type_data.get('element_lb'),
            element_ub=type_data.get('element_ub'),
            ext=type_data.get('ext', False)
        )
        asn1_types.append(asn1_type)
    return asn1_types

def get_fields_signature(fields: List[Asn1Field]) -> Tuple:
    """Create a hashable signature of fields for comparison"""
    signature = []
    for field in fields:
        field_sig = (
            field.name,
            field.type_name,
            field.lb,
            field.ub,
            field.element_type,
            field.element_lb,
            field.element_ub,
            field.optional,
            field.ext,
        )
        signature.append(field_sig)
    return tuple(signature)


def to_go_name(name: str) -> str:
    """Convert ASN.1 name to Go name (replace dashes with underscores)"""
    return name.replace('-', '_')


def is_setuprelease_type(type_name: str) -> bool:
    """Check if a type is a SetupRelease pattern"""
    return "SetupRelease" in type_name or "setuprelease" in type_name.lower()


def extract_setuprelease_inner_type(type_name: str) -> str:
    """Extract the inner type from SetupRelease pattern"""
    # Handle both "SetupRelease of Type" and "SetupRelease{Type}" formats
    if "SetupRelease of" in type_name:
        return type_name.replace("SetupRelease of", "").strip()
    elif "SetupRelease{" in type_name:
        return type_name.replace("SetupRelease{", "").replace("}", "").strip()
    return type_name


def get_type_dependencies(asn1_type: Asn1Type) -> List[str]:
    """Get all types that this type depends on"""
    dependencies = []
    
    # Check fields for dependencies
    for field in asn1_type.fields:
        type_name = field.type_name
        
        # Skip primitive types
        if type_name in ["INTEGER", "ENUMERATED", "BIT_STRING", "BITSTRING", 
                         "OCTET_STRING", "OCTETSTRING", "BOOLEAN", "NULL", 
                         "Struct-Null", "SEQUENCE_OF"]:
            # For SEQUENCE_OF, check element_type
            if type_name == "SEQUENCE_OF" and field.element_type:
                element_type = field.element_type
                if element_type not in ["INTEGER", "ENUMERATED", "BIT_STRING", "BITSTRING", 
                                       "OCTET_STRING", "OCTETSTRING", "BOOLEAN"]:
                    dependencies.append(element_type)
            continue
        
        # Handle SetupRelease pattern
        if is_setuprelease_type(type_name):
            inner_type = extract_setuprelease_inner_type(type_name)
            dependencies.append(inner_type)
        else:
            # Regular custom type
            dependencies.append(type_name)
    
    # For SEQUENCE_OF at type level
    if asn1_type.type_name == "SEQUENCE_OF" and asn1_type.element_type:
        element_type = asn1_type.element_type
        if element_type not in ["INTEGER", "ENUMERATED", "BIT_STRING", "BITSTRING", 
                               "OCTET_STRING", "OCTETSTRING", "BOOLEAN"]:
            dependencies.append(element_type)
    
    return dependencies


def topological_sort_types(asn1_types: List[Asn1Type]) -> List[Asn1Type]:
    """Sort types in dependency order using topological sort"""
    # Create a map of type name to type
    type_map = {t.name: t for t in asn1_types}
    
    # Build dependency graph
    graph = {t.name: get_type_dependencies(t) for t in asn1_types}
    
    # Kahn's algorithm for topological sort
    in_degree = {name: 0 for name in graph}
    for name in graph:
        for dep in graph[name]:
            if dep in in_degree:
                in_degree[dep] += 1
    
    # Find all nodes with no incoming edges
    queue = [name for name in in_degree if in_degree[name] == 0]
    result = []
    
    while queue:
        current = queue.pop(0)
        result.append(type_map[current])
        
        # Remove edges from current
        for neighbor in graph[current]:
            if neighbor in in_degree:
                in_degree[neighbor] -= 1
                if in_degree[neighbor] == 0:
                    queue.append(neighbor)
    
    # If we couldn't sort all types, there's a cycle or missing types
    # In that case, just append the remaining types
    sorted_names = {t.name for t in result}
    for t in asn1_types:
        if t.name not in sorted_names:
            result.append(t)
    
    return result


def format_constraint_value(value: Any) -> str:
    """Format constraint value for Go code (for tags and string output)"""
    if value is None:
        return ""
    if isinstance(value, str):
        # Check if string represents a number (including negative numbers)
        # Try to parse as number first
        try:
            # If it's a numeric string, return as-is (preserve negative sign)
            float(value)  # Test if it's numeric
            return value
        except ValueError:
            # Not a number, replace dashes with underscores for Go identifiers
            return value.replace('-', '_')
    return str(value)


def format_constraint_for_code(value: Any) -> Any:
    """Format constraint value for use in Go code (in constraint structs)
    Returns integer if numeric, formatted string if string variable name"""
    if value is None:
        return 0
    if isinstance(value, (int, float)):
        return int(value)
    if isinstance(value, str):
        # Check if string represents a number (including negative numbers)
        try:
            # If it's a numeric string, convert to int/float
            num_value = float(value)
            return int(num_value) if num_value.is_integer() else num_value
        except ValueError:
            # Not a number, replace dashes with underscores for Go identifiers
            return value.replace('-', '_')
    return value


def normalize_octetstring_constraints(lb: Any, ub: Any) -> tuple:
    """Normalize OCTETSTRING constraints - if CONTAINING is in lb or ub, set both to 0"""
    lb_str = str(lb) if lb is not None else ""
    ub_str = str(ub) if ub is not None else ""
    
    # if "CONTAINING" in lb_str.upper() or "CONTAINING" in ub_str.upper():
    #     return (0, 0)
    
    return (lb, ub)


def get_go_type_name(type_name: str, field: Optional[Asn1Field] = None) -> str:
    """Convert ASN.1 type name to Go type name"""
    # Check for SetupRelease pattern using helper
    if is_setuprelease_type(type_name):
        inner_type = extract_setuprelease_inner_type(type_name)
        return to_go_name(inner_type)
    
    # Handle primitive types
    type_map = {
        "INTEGER": "int64",
        "ENUMERATED": "uper.Enumerated",
        "BIT_STRING": "uper.BitString",
        "BITSTRING": "uper.BitString",
        "OCTET_STRING": "[]byte",
        "OCTETSTRING": "[]byte",
        "OctetString": "[]byte",  # Alternative naming
        "BOOLEAN": "bool",
        "NULL": "uper.NULL",
        "Struct-Null": "interface{}",
    }
    
    if type_name in type_map:
        return type_map[type_name]
    
    # For SEQUENCE_OF, handle element type
    if type_name == "SEQUENCE_OF" and field and field.element_type:
        return f"[]{to_go_name(field.element_type)}"
    
    # Default: convert name format
    return to_go_name(type_name)


def is_base_type_element(element_type: str) -> bool:
    """Check if element_type is a base type (INTEGER, OCTETSTRING, BITSTRING)"""
    return element_type in ["INTEGER", "OCTET_STRING", "OCTETSTRING", "OctetString", 
                           "BIT_STRING", "BITSTRING", "BitString"]


def get_sequence_of_go_type(element_type: str) -> tuple[str, str]:
    """
    Get Go type and utils type for SEQUENCE_OF with base types.
    Returns (go_type, utils_type) tuple.
    If element_type is base type, returns special types, otherwise returns normal types.
    """
    if element_type == "INTEGER":
        return ("[]int64", "utils.INTEGER")
    elif element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
        return ("[][]byte", "utils.OCTETSTRING")
    elif element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
        return ("[]uper.BitString", "utils.BITSTRING")
    else:
        # Normal custom type
        go_type = to_go_name(element_type)
        return (f"[]{go_type}", go_type)


def is_pointer_type(field: Asn1Field, asn1_type: Optional[Asn1Type] = None) -> bool:
    """Determine if a field should be a pointer type"""
    # Optional fields are pointers
    if field.optional:
        return True
    
    # CHOICE alternatives are pointers (except primitives with constraints)
    if asn1_type and asn1_type.type_name == "CHOICE":
        # If it's a primitive with constraints, it might not be a pointer
        if field.type_name in ["INTEGER", "ENUMERATED", "BIT_STRING", "BITSTRING", "BitString", "OCTET_STRING", "OCTETSTRING", "OctetString", "BOOLEAN"]:
            return False
        return True
    
    return False


def get_field_tag(field: Asn1Field, ext_group_num: Optional[int] = None) -> str:
    """Generate field tag string
    
    Args:
        field: The ASN.1 field
        ext_group_num: Extension group number if field is in an extension group
    """
    tags = []
    
    # Add constraints
    if field.lb is not None and field.ub is not None:
        lb_str = format_constraint_value(field.lb)
        ub_str = format_constraint_value(field.ub)
        tags.append(f"lb:{lb_str},ub:{ub_str}")
    elif field.lb is not None:
        lb_str = format_constraint_value(field.lb)
        tags.append(f"lb:{lb_str}")
    elif field.ub is not None:
        ub_str = format_constraint_value(field.ub)
        tags.append(f"ub:{ub_str}")
    
    # Add element constraints for SEQUENCE_OF
    # For base types (INTEGER, OCTETSTRING, BITSTRING), always include e_lb and e_ub
    if field.type_name == "SEQUENCE_OF":
        is_base = is_base_type_element(field.element_type) if field.element_type else False
        if is_base:
            # For base types, always include e_lb and e_ub (default to 0 if not specified)
            e_lb = field.element_lb if field.element_lb is not None else 0
            e_ub = field.element_ub if field.element_ub is not None else 0
            e_lb_str = format_constraint_value(e_lb)
            e_ub_str = format_constraint_value(e_ub)
            tags.append(f"e_lb:{e_lb_str},e_ub:{e_ub_str}")
        elif field.element_lb is not None and field.element_ub is not None:
            e_lb_str = format_constraint_value(field.element_lb)
            e_ub_str = format_constraint_value(field.element_ub)
            tags.append(f"e_lb:{e_lb_str},e_ub:{e_ub_str}")
        elif field.element_lb is not None:
            e_lb_str = format_constraint_value(field.element_lb)
            tags.append(f"e_lb:{e_lb_str}")
        elif field.element_ub is not None:
            e_ub_str = format_constraint_value(field.element_ub)
            tags.append(f"e_ub:{e_ub_str}")
    
    # Add optional/mandatory tag
    if not field.optional:
        tags.append("madatory")
    else:
        tags.append("optional")
    
    # Add extension group tag (ext-{group_num}) instead of just "ext"
    if ext_group_num is not None:
        tags.append(f"ext-{ext_group_num}")
    elif field.ext:
        # If ext is true but no group number provided, try to parse from comment
        if field.comment:
            match = re.match(r'ext-(\d+)', field.comment.strip())
            if match:
                tags.append(f"ext-{match.group(1)}")
            else:
                tags.append("ext")
        else:
            tags.append("ext")
    
    # Add SetupRelease tag
    if is_setuprelease_type(field.type_name):
        tags.append("setuprelease")
    
    return ",".join(tags) if tags else ""


def generate_integer_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for INTEGER type"""
    go_name = to_go_name(asn1_type.name)
    lb = format_constraint_value(asn1_type.lb) if asn1_type.lb is not None else "0"
    ub = format_constraint_value(asn1_type.ub) if asn1_type.ub is not None else ""
    
    tag = f"`lb:{lb},ub:{ub},madatory`" if ub else f"`lb:{lb},madatory`"
    
    constraint_lb = format_constraint_for_code(asn1_type.lb)
    constraint_ub = format_constraint_for_code(asn1_type.ub)
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value uint64 {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = uint64(v)
	return nil
}}
"""
    return code


def generate_enumerated_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for ENUMERATED type"""
    go_name = to_go_name(asn1_type.name)
    lb = asn1_type.lb if asn1_type.lb is not None else 0
    ub = asn1_type.ub if asn1_type.ub is not None else (len(asn1_type.values) - 1 if asn1_type.values else 0)
    
    # Generate constants
    constants = []
    for i, value in enumerate(asn1_type.values):
        const_name = f"{go_name}_Enum_{to_go_name(value)}"
        constants.append(f"\t{const_name} uper.Enumerated = {i}")
    
    const_block = "\n".join(constants) if constants else ""
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
{const_block}
)

type {go_name} struct {{
	Value uper.Enumerated `lb:{lb},ub:{ub},madatory`
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = uper.Enumerated(v)
	return nil
}}
"""
    return code


def generate_bitstring_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for BIT_STRING type"""
    go_name = to_go_name(asn1_type.name)
    lb = format_constraint_value(asn1_type.lb) if asn1_type.lb is not None else ""
    ub = format_constraint_value(asn1_type.ub) if asn1_type.ub is not None else ""
    
    tag = f"`lb:{lb},ub:{ub},madatory`" if lb and ub else (f"`lb:{lb},madatory`" if lb else "`madatory`")
    
    constraint_lb = format_constraint_for_code(asn1_type.lb)
    constraint_ub = format_constraint_for_code(asn1_type.ub)
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value uper.BitString {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = uper.BitString{{
		Bytes:   v,
		NumBits: uint64(n),
	}}
	return nil
}}
"""
    return code


def generate_octetstring_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for OCTET_STRING type"""
    go_name = to_go_name(asn1_type.name)
    # Normalize constraints for CONTAINING
    normalized_lb, normalized_ub = normalize_octetstring_constraints(asn1_type.lb, asn1_type.ub)
    lb = format_constraint_value(normalized_lb) if normalized_lb is not None else ""
    ub = format_constraint_value(normalized_ub) if normalized_ub is not None else ""
    
    tag = f"`lb:{lb},ub:{ub},madatory`" if lb and ub else (f"`lb:{lb},madatory`" if lb else "`madatory`")
    
    constraint_lb = format_constraint_for_code(normalized_lb)
    constraint_ub = format_constraint_for_code(normalized_ub)
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value []byte {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteOctetString(ie.Value, &uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = v
	return nil
}}
"""
    return code


def generate_boolean_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for BOOLEAN type"""
    go_name = to_go_name(asn1_type.name)
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value bool `lb:1,ub:1,madatory`
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteBoolean(ie.Value); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	var v bool
	if v, err = r.ReadBoolean(); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = v
	return nil
}}
"""
    return code


def generate_null_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for NULL type"""
    go_name = to_go_name(asn1_type.name)
    
    code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value uper.NULL
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	if err := w.WriteNull(); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	if err := r.ReadNull(); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	return nil
}}
"""
    return code


def generate_sequence_of_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for SEQUENCE_OF type"""
    go_name = to_go_name(asn1_type.name)
    element_type_raw = asn1_type.element_type if asn1_type.element_type else "interface{}"
    lb = format_constraint_value(asn1_type.lb) if asn1_type.lb is not None else ""
    ub = format_constraint_value(asn1_type.ub) if asn1_type.ub is not None else ""
    
    # Build tag with sequence constraints
    tag_parts = []
    if lb and ub:
        tag_parts.append(f"lb:{lb},ub:{ub}")
    elif lb:
        tag_parts.append(f"lb:{lb}")
    elif ub:
        tag_parts.append(f"ub:{ub}")
    
    # Add element constraints for base types
    if is_base_type_element(element_type_raw):
        element_lb = format_constraint_value(asn1_type.element_lb) if asn1_type.element_lb is not None else ""
        element_ub = format_constraint_value(asn1_type.element_ub) if asn1_type.element_ub is not None else ""
        if element_lb and element_ub:
            tag_parts.append(f"e_lb:{element_lb},e_ub:{element_ub}")
        elif element_lb:
            tag_parts.append(f"e_lb:{element_lb}")
        elif element_ub:
            tag_parts.append(f"e_ub:{element_ub}")
    
    tag_parts.append("madatory")
    tag = f"`{','.join(tag_parts)}`" if tag_parts else "``"
    
    constraint_lb = format_constraint_for_code(asn1_type.lb)
    constraint_ub = format_constraint_for_code(asn1_type.ub)
    
    # Check if element_type is a base type
    is_base = is_base_type_element(element_type_raw)
    go_type, utils_type = get_sequence_of_go_type(element_type_raw)
    
    if is_base:
        # Special handling for base types
        if element_type_raw == "INTEGER":
            # Need element constraints - use default for now, should come from field
            element_lb = format_constraint_for_code(asn1_type.element_lb) if asn1_type.element_lb is not None else 0
            element_ub = format_constraint_for_code(asn1_type.element_ub) if asn1_type.element_ub is not None else 0
            code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value {go_type} {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	for _, i := range ie.Value {{
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}}
	if err = tmp.Encode(w); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	fn := func() *{utils_type} {{
		ie := utils.NewINTEGER(0, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		return &ie
	}}
	if err = tmp.Decode(r, fn); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = {go_type}{{}}
	for _, i := range tmp.Value {{
		ie.Value = append(ie.Value, int64(i.Value))
	}}
	return err
}}
"""
        elif element_type_raw in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
            element_lb = format_constraint_for_code(asn1_type.element_lb) if asn1_type.element_lb is not None else 0
            element_ub = format_constraint_for_code(asn1_type.element_ub) if asn1_type.element_ub is not None else 0
            code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value {go_type} {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	for _, i := range ie.Value {{
		tmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}}
	if err = tmp.Encode(w); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	fn := func() *{utils_type} {{
		ie := utils.NewOCTETSTRING([]byte{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		return &ie
	}}
	if err = tmp.Decode(r, fn); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = {go_type}{{}}
	for _, i := range tmp.Value {{
		ie.Value = append(ie.Value, i.Value)
	}}
	return err
}}
"""
        elif element_type_raw in ["BIT_STRING", "BITSTRING", "BitString"]:
            element_lb = format_constraint_for_code(asn1_type.element_lb) if asn1_type.element_lb is not None else 0
            element_ub = format_constraint_for_code(asn1_type.element_ub) if asn1_type.element_ub is not None else 0
            code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value {go_type} {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	for _, i := range ie.Value {{
		tmp_ie := utils.NewBITSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}}
	if err = tmp.Encode(w); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	tmp := utils.NewSequence[*{utils_type}]([]*{utils_type}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	fn := func() *{utils_type} {{
		ie := utils.NewBITSTRING(uper.BitString{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
		return &ie
	}}
	if err = tmp.Decode(r, fn); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = {go_type}{{}}
	for _, i := range tmp.Value {{
		ie.Value = append(ie.Value, i.Value)
	}}
	return err
}}
"""
        else:
            # Fallback (shouldn't happen)
            go_type_normal = to_go_name(element_type_raw)
            code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value []{go_type_normal} {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	tmp := utils.NewSequence[*{go_type_normal}]([]*{go_type_normal}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	for _, i := range ie.Value {{
		tmp.Value = append(tmp.Value, &i)
	}}
	if err = tmp.Encode(w); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	tmp := utils.NewSequence[*{go_type_normal}]([]*{go_type_normal}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	fn := func() *{go_type_normal} {{
		return new({go_type_normal})
	}}
	if err = tmp.Decode(r, fn); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = []{go_type_normal}{{}}
	for _, i := range tmp.Value {{
		ie.Value = append(ie.Value, *i)
	}}
	return err
}}
"""
    else:
        # Normal custom type
        go_type_normal = to_go_name(element_type_raw)
        code = f"""package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type {go_name} struct {{
	Value []{go_type_normal} {tag}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	tmp := utils.NewSequence[*{go_type_normal}]([]*{go_type_normal}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	for _, i := range ie.Value {{
		tmp.Value = append(tmp.Value, &i)
	}}
	if err = tmp.Encode(w); err != nil {{
		return utils.WrapError("Encode {go_name}", err)
	}}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	tmp := utils.NewSequence[*{go_type_normal}]([]*{go_type_normal}{{}}, uper.Constraint{{Lb: {constraint_lb}, Ub: {constraint_ub}}}, false)
	fn := func() *{go_type_normal} {{
		return new({go_type_normal})
	}}
	if err = tmp.Decode(r, fn); err != nil {{
		return utils.WrapError("Decode {go_name}", err)
	}}
	ie.Value = []{go_type_normal}{{}}
	for _, i := range tmp.Value {{
		ie.Value = append(ie.Value, *i)
	}}
	return err
}}
"""
    return code


def generate_choice_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for CHOICE type"""
    go_name = to_go_name(asn1_type.name)
    
    # Generate constants
    constants = [f"\t{go_name}_Choice_nothing uint64 = iota"]
    for field in asn1_type.fields:
        const_name = f"{go_name}_Choice_{to_go_name(field.name)}"
        constants.append(f"\t{const_name}")
    
    const_block = "\n".join(constants)
    
    # Generate struct fields
    struct_fields = ["\tChoice uint64"]
    for field in asn1_type.fields:
        field_name = to_go_name(field.name)
        go_type = get_go_type_name(field.type_name, field)
        
        # Handle SEQUENCE_OF in CHOICE
        if field.type_name == "SEQUENCE_OF" and field.element_type:
            go_type_seq, _ = get_sequence_of_go_type(field.element_type)
            tag = get_field_tag(field)
            struct_fields.append(f"\t{field_name} {go_type_seq} `{tag}`" if tag else f"\t{field_name} {go_type_seq}")
        elif field.type_name == "NULL":
            # NULL in CHOICE is value type, not pointer
            tag = get_field_tag(field)
            struct_fields.append(f"\t{field_name} uper.NULL `{tag}`" if tag else f"\t{field_name} uper.NULL")
        elif field.type_name in ["INTEGER", "ENUMERATED", "BIT_STRING", "BITSTRING", "BitString", "OCTET_STRING", "OCTETSTRING", "OctetString", "BOOLEAN"]:
            # Primitives in CHOICE are value types, not pointers
            tag = get_field_tag(field)
            struct_fields.append(f"\t{field_name} {go_type} `{tag}`" if tag else f"\t{field_name} {go_type}")
        elif go_type == "interface{}":
            # interface{} should not be a pointer
            tag = get_field_tag(field)
            struct_fields.append(f"\t{field_name} {go_type} `{tag}`" if tag else f"\t{field_name} {go_type}")
        else:
            # Custom types in CHOICE are pointers
            struct_fields.append(f"\t{field_name} *{go_type}")
    
    struct_block = "\n".join(struct_fields)
    
    # Generate encode switch
    encode_cases = []
    for i, field in enumerate(asn1_type.fields, 1):
        field_name = to_go_name(field.name)
        const_name = f"{go_name}_Choice_{field_name}"
        
        if field.type_name == "SEQUENCE_OF" and field.element_type:
            is_base = is_base_type_element(field.element_type)
            go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type)
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
            element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
            
            if is_base:
                if field.element_type == "INTEGER":
                    encode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		for _, i := range ie.{field_name} {{
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			tmp.Value = append(tmp.Value, &tmp_ie)
		}}
		if err = tmp.Encode(w); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
                elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    encode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		for _, i := range ie.{field_name} {{
			tmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			tmp.Value = append(tmp.Value, &tmp_ie)
		}}
		if err = tmp.Encode(w); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
                elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                    encode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		for _, i := range ie.{field_name} {{
			tmp_ie := utils.NewBITSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			tmp.Value = append(tmp.Value, &tmp_ie)
		}}
		if err = tmp.Encode(w); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
            else:
                element_type = to_go_name(field.element_type)
                encode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		for _, i := range ie.{field_name} {{
			tmp.Value = append(tmp.Value, &i)
		}}
		if err = tmp.Encode(w); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name == "INTEGER":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            encode_cases.append(f"""\tcase {const_name}:
		if err = w.WriteInteger(int64(ie.{field_name}), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name == "ENUMERATED":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            encode_cases.append(f"""\tcase {const_name}:
		if err = w.WriteEnumerate(uint64(ie.{field_name}), uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            encode_cases.append(f"""\tcase {const_name}:
		if err = w.WriteBitString(ie.{field_name}.Bytes, uint(ie.{field_name}.NumBits), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
            # Normalize constraints for CONTAINING
            normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
            lb = format_constraint_for_code(normalized_lb)
            ub = format_constraint_for_code(normalized_ub)
            encode_cases.append(f"""\tcase {const_name}:
		if err = w.WriteOctetString(ie.{field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name == "BOOLEAN":
            encode_cases.append(f"""\tcase {const_name}:
		if err = w.WriteBoolean(ie.{field_name}); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif field.type_name == "NULL":
            encode_cases.append(f"""\tcase {const_name}:
		if err := w.WriteNull(); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
        elif "interface{}" in get_go_type_name(field.type_name, field):
            # interface{} has no methods, skip encoding
            encode_cases.append(f"\tcase {const_name}:\n\t\t// interface{{}} field of choice - nothing to encode")
        else:
            encode_cases.append(f"""\tcase {const_name}:
		if err = ie.{field_name}.Encode(w); err != nil {{
			err = utils.WrapError("Encode {field_name}", err)
		}}""")
    
    encode_switch = "\n".join(encode_cases)
    
        # Generate decode switch
    decode_cases = []
    for i, field in enumerate(asn1_type.fields, 1):
        field_name = to_go_name(field.name)
        const_name = f"{go_name}_Choice_{field_name}"
        go_type = get_go_type_name(field.type_name, field)
        
        if field.type_name == "SEQUENCE_OF" and field.element_type:
            is_base = is_base_type_element(field.element_type)
            go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type)
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            
            if is_base:
                element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
                element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
                if field.element_type == "INTEGER":
                    decode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		fn := func() *{utils_type_seq} {{
			ie := utils.NewINTEGER(0, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			return &ie
		}}
		if err = tmp.Decode(r, fn); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = {go_type_seq}{{}}
		for _, i := range tmp.Value {{
			ie.{field_name} = append(ie.{field_name}, int64(i.Value))
		}}""")
                elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    decode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		fn := func() *{utils_type_seq} {{
			ie := utils.NewOCTETSTRING([]byte{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			return &ie
		}}
		if err = tmp.Decode(r, fn); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = {go_type_seq}{{}}
		for _, i := range tmp.Value {{
			ie.{field_name} = append(ie.{field_name}, i.Value)
		}}""")
                elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                    decode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		fn := func() *{utils_type_seq} {{
			ie := utils.NewBITSTRING(uper.BitString{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)
			return &ie
		}}
		if err = tmp.Decode(r, fn); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = {go_type_seq}{{}}
		for _, i := range tmp.Value {{
			ie.{field_name} = append(ie.{field_name}, i.Value)
		}}""")
            else:
                element_type = to_go_name(field.element_type)
                decode_cases.append(f"""\tcase {const_name}:
		tmp := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)
		fn := func() *{element_type} {{
			return new({element_type})
		}}
		if err = tmp.Decode(r, fn); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = []{element_type}{{}}
		for _, i := range tmp.Value {{
			ie.{field_name} = append(ie.{field_name}, *i)
		}}""")
        elif field.type_name == "INTEGER":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            decode_cases.append(f"""\tcase {const_name}:
		var tmp_int_{field_name} int64
		if tmp_int_{field_name}, err = r.ReadInteger(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = tmp_int_{field_name}""")
        elif field.type_name == "ENUMERATED":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            decode_cases.append(f"""\tcase {const_name}:
		var tmp_enum_{field_name} uint64
		if tmp_enum_{field_name}, err = r.ReadEnumerate(uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = uper.Enumerated(tmp_enum_{field_name})""")
        elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            decode_cases.append(f"""\tcase {const_name}:
		var tmp_bs_{field_name} []byte
		var n_{field_name} uint
		if tmp_bs_{field_name}, n_{field_name}, err = r.ReadBitString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = uper.BitString{{
			Bytes:   tmp_bs_{field_name},
			NumBits: uint64(n_{field_name}),
		}}""")
        elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
            # Normalize constraints for CONTAINING
            normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
            lb = format_constraint_for_code(normalized_lb)
            ub = format_constraint_for_code(normalized_ub)
            decode_cases.append(f"""\tcase {const_name}:
		var tmp_os_{field_name} []byte
		if tmp_os_{field_name}, err = r.ReadOctetString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = tmp_os_{field_name}""")
        elif field.type_name == "BOOLEAN":
            decode_cases.append(f"""\tcase {const_name}:
		var tmp_bool_{field_name} bool
		if tmp_bool_{field_name}, err = r.ReadBoolean(); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}
		ie.{field_name} = tmp_bool_{field_name}""")
        elif field.type_name == "NULL":
            decode_cases.append(f"""\tcase {const_name}:
		if err := r.ReadNull(); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}""")
        else:
            go_type = get_go_type_name(field.type_name, field)
            if "interface{}" in go_type:
                # interface{} has no methods, skip decoding
                decode_cases.append(f"\tcase {const_name}:\n\t\t// interface{{}} field of choice - nothing to decode")
            else:
                decode_cases.append(f"""\tcase {const_name}:
		ie.{field_name} = new({go_type})
		if err = ie.{field_name}.Decode(r); err != nil {{
			return  utils.WrapError("Decode {field_name}", err)
		}}""")
    
    decode_switch = "\n".join(decode_cases)
    
    num_choices = len(asn1_type.fields)
    
    code = f"""package ies

import (
	"fmt"
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
{const_block}
)

type {go_name} struct {{
{struct_block}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
	if err = w.WriteChoice(ie.Choice, {num_choices}, false); err != nil {{
		return err
	}}
	switch ie.Choice {{
{encode_switch}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}}
	return err
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
	if ie.Choice, err = r.ReadChoice({num_choices}, false); err != nil {{
		return err
	}}
	switch ie.Choice {{
{decode_switch}
	default:
		return  fmt.Errorf("invalid choice: %d", ie.Choice)
	}}
	return nil
}}
"""
    return code


def parse_extension_groups(fields: List[Asn1Field]) -> Dict[int, List[Asn1Field]]:
    """Parse extension groups from fields. Returns dict mapping group number to list of fields."""
    ext_groups = {}
    for field in fields:
        if field.ext and field.comment:
            # Parse comment like "ext-1", "ext-2", etc.
            match = re.match(r'ext-(\d+)', field.comment.strip())
            if match:
                group_num = int(match.group(1))
                if group_num not in ext_groups:
                    ext_groups[group_num] = []
                ext_groups[group_num].append(field)
    return ext_groups


def separate_root_and_extension_fields(fields: List[Asn1Field]) -> Tuple[List[Asn1Field], Dict[int, List[Asn1Field]]]:
    """Separate root fields (ext: false) from extension fields (ext: true)."""
    root_fields = []
    for field in fields:
        if not field.ext:
            root_fields.append(field)
    
    ext_groups = parse_extension_groups(fields)
    return root_fields, ext_groups


def generate_sequence_code(asn1_type: Asn1Type) -> str:
    """Generate Go code for SEQUENCE type"""
    go_name = to_go_name(asn1_type.name)
    
    # Separate root and extension fields
    root_fields, ext_groups = separate_root_and_extension_fields(asn1_type.fields)
    has_extensions = len(ext_groups) > 0
    
    # Generate struct fields (include all fields: root + extension)
    struct_fields = []
    optional_count = 0
    
    # Build a map from field name to extension group number for quick lookup
    field_to_ext_group = {}
    for group_num, ext_fields in ext_groups.items():
        for ext_field in ext_fields:
            field_to_ext_group[ext_field.name] = group_num
    
    for field in asn1_type.fields:
        field_name = to_go_name(field.name)
        go_type = get_go_type_name(field.type_name, field)
        
        # Get extension group number if field is in an extension group
        ext_group_num = field_to_ext_group.get(field.name, None)
        
        # Handle SetupRelease
        if is_setuprelease_type(field.type_name):
            inner_type = extract_setuprelease_inner_type(field.type_name)
            go_type = to_go_name(inner_type)
            # Get all tags (optional, ext, setuprelease, etc.)
            tag = get_field_tag(field, ext_group_num)
            struct_fields.append(f"\t{field_name} *{go_type} `{tag}`" if tag else f"\t{field_name} *{go_type}")
            if field.optional:
                optional_count += 1
            continue
        
        # Handle SEQUENCE_OF - should not be pointer even if optional
        if field.type_name == "SEQUENCE_OF":
            if field.element_type:
                go_type_seq, _ = get_sequence_of_go_type(field.element_type)
            else:
                go_type_seq = "[]interface{}"
            tag = get_field_tag(field, ext_group_num)
            if field.optional:
                optional_count += 1
            struct_fields.append(f"\t{field_name} {go_type_seq} `{tag}`" if tag else f"\t{field_name} {go_type_seq}")
        # Handle NULL - should not be pointer
        elif field.type_name == "NULL":
            tag = get_field_tag(field, ext_group_num)
            if field.optional:
                optional_count += 1
            struct_fields.append(f"\t{field_name} uper.NULL `{tag}`" if tag else f"\t{field_name} uper.NULL")
        # Handle optional fields
        elif field.optional:
            # Optional fields are pointers, except interface{} which should not be a pointer
            tag = get_field_tag(field, ext_group_num)
            if go_type == "interface{}":
                struct_fields.append(f"\t{field_name} {go_type} `{tag}`" if tag else f"\t{field_name} {go_type}")
            else:
                struct_fields.append(f"\t{field_name} *{go_type} `{tag}`" if tag else f"\t{field_name} *{go_type}")
            optional_count += 1
        else:
            tag = get_field_tag(field, ext_group_num)
            struct_fields.append(f"\t{field_name} {go_type} `{tag}`" if tag else f"\t{field_name} {go_type}")
    
    struct_block = "\n".join(struct_fields) if struct_fields else ""
    
    # Generate encode code
    encode_lines = []
    
    # Count optional root fields (excluding Struct-Null and interface{})
    optional_root_fields = [f for f in root_fields if f.optional and f.type_name != "Struct-Null" and f.type_name != "interface{}"]
    optional_root_count = len(optional_root_fields)
    
    # Generate hasExtensions check (check if any extension field is present)
    if has_extensions:
        ext_field_checks = []
        for group_num in sorted(ext_groups.keys()):
            for ext_field in ext_groups[group_num]:
                # Skip Struct-Null and interface{} fields (they don't contribute to hasExtensions)
                if ext_field.type_name == "Struct-Null" or ext_field.type_name == "interface{}":
                    continue
                field_name = to_go_name(ext_field.name)
                if ext_field.type_name == "SEQUENCE_OF":
                    ext_field_checks.append(f"len(ie.{field_name}) > 0")
                else:
                    ext_field_checks.append(f"ie.{field_name} != nil")
        has_extensions_expr = " || ".join(ext_field_checks) if ext_field_checks else "false"
        encode_lines.append(f"\thasExtensions := {has_extensions_expr}")
    
    # Generate preamble bits
    preamble_bits = []
    if has_extensions:
        preamble_bits.append("hasExtensions")
    for field in root_fields:
        if field.optional and field.type_name != "Struct-Null" and field.type_name != "interface{}":
            field_name = to_go_name(field.name)
            if field.type_name == "SEQUENCE_OF":
                preamble_bits.append(f"len(ie.{field_name}) > 0")
            else:
                preamble_bits.append(f"ie.{field_name} != nil")
    
    if preamble_bits:
        encode_lines.append(f"\tpreambleBits := []bool{{{', '.join(preamble_bits)}}}")
        encode_lines.append("\tfor _, bit := range preambleBits {")
        encode_lines.append("\t\tif err = w.WriteBool(bit); err != nil {")
        encode_lines.append("\t\t\treturn err")
        encode_lines.append("\t\t}")
        encode_lines.append("\t}")
    
    # Encode root fields only (mandatory first, then optional)
    for field in root_fields:
        field_name = to_go_name(field.name)
        
        # Handle optional fields
        if field.optional:
            # Handle optional SetupRelease
            if is_setuprelease_type(field.type_name):
                encode_lines.append(f"\tif ie.{field_name} != nil {{")
                inner_type = extract_setuprelease_inner_type(field.type_name)
                go_type = to_go_name(inner_type)
                encode_lines.append(f"\t\ttmp_{field_name} := utils.SetupRelease[*{go_type}]{{")
                encode_lines.append(f"\t\t\tSetup: ie.{field_name},")
                encode_lines.append(f"\t\t}}")
                encode_lines.append(f"\t\tif err = tmp_{field_name}.Encode(w); err != nil {{")
                encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                encode_lines.append(f"\t\t}}")
                encode_lines.append(f"\t}}")
                continue
            
            # For optional SEQUENCE_OF, check if slice is not nil/empty
            if field.type_name == "SEQUENCE_OF":
                encode_lines.append(f"\tif len(ie.{field_name}) > 0 {{")
                is_base = is_base_type_element(field.element_type) if field.element_type else False
                go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type) if field.element_type else ("[]interface{}", "interface{}")
                lb = format_constraint_for_code(field.lb)
                ub = format_constraint_for_code(field.ub)
                element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
                element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
                
                if is_base:
                    if field.element_type == "INTEGER":
                        encode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        encode_lines.append(f"\t\tfor _, i := range ie.{field_name} {{")
                        encode_lines.append(f"\t\t\ttmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        encode_lines.append(f"\t\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                        encode_lines.append(f"\t\t}}")
                    elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                        encode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        encode_lines.append(f"\t\tfor _, i := range ie.{field_name} {{")
                        encode_lines.append(f"\t\t\ttmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        encode_lines.append(f"\t\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                        encode_lines.append(f"\t\t}}")
                    elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                        encode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        encode_lines.append(f"\t\tfor _, i := range ie.{field_name} {{")
                        encode_lines.append(f"\t\t\ttmp_ie := utils.NewBITSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        encode_lines.append(f"\t\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                        encode_lines.append(f"\t\t}}")
                else:
                    element_type = to_go_name(field.element_type) if field.element_type else "interface{}"
                    encode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    encode_lines.append(f"\t\tfor _, i := range ie.{field_name} {{")
                    encode_lines.append(f"\t\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &i)")
                    encode_lines.append(f"\t\t}}")
                encode_lines.append(f"\t\tif err = tmp_{field_name}.Encode(w); err != nil {{")
                encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                encode_lines.append(f"\t\t}}}}")
            else:
                # Only generate if block if field is not Struct-Null and not interface{}
                if field.type_name != "Struct-Null" and field.type_name != "interface{}":
                    encode_lines.append(f"\tif ie.{field_name} != nil {{")
                
                # Encode the optional field based on its type
                if field.type_name == "INTEGER":
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    encode_lines.append(f"\t\tif err = w.WriteInteger(*ie.{field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name == "ENUMERATED":
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    encode_lines.append(f"\t\tif err = w.WriteEnumerate(uint64(*ie.{field_name}), uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    encode_lines.append(f"\t\tif err = w.WriteBitString(ie.{field_name}.Bytes, uint(ie.{field_name}.NumBits), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    # Normalize constraints for CONTAINING
                    normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
                    lb = format_constraint_for_code(normalized_lb)
                    ub = format_constraint_for_code(normalized_ub)
                    encode_lines.append(f"\t\tif err = w.WriteOctetString(*ie.{field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name == "BOOLEAN":
                    encode_lines.append(f"\t\tif err = w.WriteBoolean(*ie.{field_name}); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name == "NULL":
                    encode_lines.append(f"\t\tif err := w.WriteNull(); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                elif field.type_name == "Struct-Null":
                    # Struct-Null is empty, nothing to encode
                    pass
                elif field.type_name == "interface{}":
                    # interface{} has no methods, skip encoding
                    pass
                else:
                    # Custom type - call Encode method
                    encode_lines.append(f"\t\tif err = ie.{field_name}.Encode(w); err != nil {{")
                    encode_lines.append(f"\t\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
                    encode_lines.append(f"\t\t}}")
                
                # Only close if block if it was opened (not Struct-Null and not interface{})
                if field.type_name != "Struct-Null" and field.type_name != "interface{}":
                    encode_lines.append(f"\t}}")
            continue
        
        # Handle mandatory SetupRelease
        if is_setuprelease_type(field.type_name):
            inner_type = extract_setuprelease_inner_type(field.type_name)
            go_type = to_go_name(inner_type)
            encode_lines.append(f"\ttmp_{field_name} := utils.SetupRelease[*{go_type}]{{")
            encode_lines.append(f"\t\tSetup: ie.{field_name},")
            encode_lines.append(f"\t}}")
            encode_lines.append(f"\tif err = tmp_{field_name}.Encode(w); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
            encode_lines.append(f"\t}}")
            continue
        
        # Handle different field types
        if field.type_name == "INTEGER":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            encode_lines.append(f"\tif err = w.WriteInteger(ie.{field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"WriteInteger {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            encode_lines.append(f"\tif err = w.WriteBitString(ie.{field_name}.Bytes, uint(ie.{field_name}.NumBits), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"WriteBitString {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
            # Normalize constraints for CONTAINING
            normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
            lb = format_constraint_for_code(normalized_lb)
            ub = format_constraint_for_code(normalized_ub)
            encode_lines.append(f"\tif err = w.WriteOctetString(ie.{field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"WriteOctetString {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif field.type_name == "BOOLEAN":
            encode_lines.append(f"\tif err = w.WriteBoolean(ie.{field_name}); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"WriteBoolean {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif field.type_name == "NULL":
            encode_lines.append(f"\tif err := w.WriteNull(); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif field.type_name == "SEQUENCE_OF":
            is_base = is_base_type_element(field.element_type) if field.element_type else False
            go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type) if field.element_type else ("[]interface{}", "interface{}")
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
            element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
            
            if is_base:
                if field.element_type == "INTEGER":
                    encode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    encode_lines.append(f"\tfor _, i := range ie.{field_name} {{")
                    encode_lines.append(f"\t\ttmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    encode_lines.append(f"\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                    encode_lines.append(f"\t}}")
                elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    encode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    encode_lines.append(f"\tfor _, i := range ie.{field_name} {{")
                    encode_lines.append(f"\t\ttmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    encode_lines.append(f"\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                    encode_lines.append(f"\t}}")
                elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                    encode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    encode_lines.append(f"\tfor _, i := range ie.{field_name} {{")
                    encode_lines.append(f"\t\ttmp_ie := utils.NewBITSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    encode_lines.append(f"\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &tmp_ie)")
                    encode_lines.append(f"\t}}")
            else:
                element_type = to_go_name(field.element_type) if field.element_type else "interface{}"
                encode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                encode_lines.append(f"\tfor _, i := range ie.{field_name} {{")
                encode_lines.append(f"\t\ttmp_{field_name}.Value = append(tmp_{field_name}.Value, &i)")
                encode_lines.append(f"\t}}")
            encode_lines.append(f"\tif err = tmp_{field_name}.Encode(w); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
            encode_lines.append(f"\t}}")
        elif go_type == "interface{}":
            # interface{} has no methods, skip encoding
            pass
        else:
            # Custom type - call Encode method
            encode_lines.append(f"\tif err = ie.{field_name}.Encode(w); err != nil {{")
            encode_lines.append(f"\t\treturn utils.WrapError(\"Encode {field_name}\", err)")
            encode_lines.append(f"\t}}")
    
    # Encode extension groups if hasExtensions
    if has_extensions:
        # Generate extension bitmap
        ext_bitmap_exprs = []
        for group_num in sorted(ext_groups.keys()):
            ext_fields = ext_groups[group_num]
            group_checks = []
            for ext_field in ext_fields:
                # Skip Struct-Null and interface{} fields (they don't contribute to extension bitmap)
                if ext_field.type_name == "Struct-Null" or ext_field.type_name == "interface{}":
                    continue
                field_name = to_go_name(ext_field.name)
                if ext_field.type_name == "SEQUENCE_OF":
                    group_checks.append(f"len(ie.{field_name}) > 0")
                else:
                    group_checks.append(f"ie.{field_name} != nil")
            # If all fields in group are Struct-Null/interface{}, the group is never present
            ext_bitmap_exprs.append(" || ".join(group_checks) if group_checks else "false")
        
        encode_lines.append("\tif hasExtensions {")
        encode_lines.append(f"\t\t// Extension bitmap: {len(ext_groups)} bits for {len(ext_groups)} extension groups")
        encode_lines.append(f"\t\textBitmap := []bool{{{', '.join(ext_bitmap_exprs)}}}")
        encode_lines.append("\t\tif err := w.WriteExtBitMap(extBitmap); err != nil {")
        encode_lines.append("\t\t\treturn utils.WrapError(\"WriteExtBitMap " + go_name + "\", err)")
        encode_lines.append("\t\t}")
        
        # Encode each extension group
        for group_num in sorted(ext_groups.keys()):
            ext_fields = ext_groups[group_num]
            encode_lines.append(f"\n\t\t// encode extension group {group_num}")
            encode_lines.append(f"\t\tif extBitmap[{group_num-1}] {{")
            encode_lines.append("\t\t\textBuf := new(bytes.Buffer)")
            encode_lines.append("\t\t\textWriter := uper.NewWriter(extBuf)")
            encode_lines.append("")
            
            # Write preamble bits for optional fields in this extension group
            # Skip Struct-Null and interface{} fields
            optional_ext_fields = [f for f in ext_fields if f.optional and f.type_name != "Struct-Null" and f.type_name != "interface{}"]
            if optional_ext_fields:
                ext_preamble_bits = []
                for ext_field in ext_fields:
                    if ext_field.optional and ext_field.type_name != "Struct-Null" and ext_field.type_name != "interface{}":
                        field_name = to_go_name(ext_field.name)
                        if ext_field.type_name == "SEQUENCE_OF":
                            ext_preamble_bits.append(f"len(ie.{field_name}) > 0")
                        else:
                            ext_preamble_bits.append(f"ie.{field_name} != nil")
                encode_lines.append(f"\t\t\t// Write preamble bits for optional fields in extension group {group_num}")
                encode_lines.append(f"\t\t\toptionals_ext_{group_num} := []bool{{{', '.join(ext_preamble_bits)}}}")
                encode_lines.append(f"\t\t\tfor _, bit := range optionals_ext_{group_num} {{")
                encode_lines.append(f"\t\t\t\tif err := extWriter.WriteBool(bit); err != nil {{")
                encode_lines.append(f"\t\t\t\t\treturn err")
                encode_lines.append(f"\t\t\t\t}}")
                encode_lines.append(f"\t\t\t}}")
                encode_lines.append("")
            
            # Encode extension fields (using extWriter)
            for ext_field in ext_fields:
                ext_field_name = to_go_name(ext_field.name)
                ext_go_type = get_go_type_name(ext_field.type_name, ext_field)
                
                # Skip Struct-Null and interface{} fields (nothing to encode)
                if ext_field.type_name == "Struct-Null" or ext_field.type_name == "interface{}":
                    continue
                
                if ext_field.optional:
                    if ext_field.type_name == "SEQUENCE_OF":
                        encode_lines.append(f"\t\t\t// encode {ext_field_name} optional")
                        encode_lines.append(f"\t\t\tif len(ie.{ext_field_name}) > 0 {{")
                    else:
                        encode_lines.append(f"\t\t\t// encode {ext_field_name} optional")
                        encode_lines.append(f"\t\t\tif ie.{ext_field_name} != nil {{")
                else:
                    encode_lines.append(f"\t\t\t// encode {ext_field_name}")
                
                # Generate encoding code for this extension field (similar to root fields but using extWriter)
                if ext_field.type_name == "SEQUENCE_OF":
                    is_base = is_base_type_element(ext_field.element_type) if ext_field.element_type else False
                    go_type_seq, utils_type_seq = get_sequence_of_go_type(ext_field.element_type) if ext_field.element_type else ("[]interface{}", "interface{}")
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    element_lb = format_constraint_for_code(ext_field.element_lb) if ext_field.element_lb is not None else 0
                    element_ub = format_constraint_for_code(ext_field.element_ub) if ext_field.element_ub is not None else 0
                    
                    if is_base:
                        if ext_field.element_type == "INTEGER":
                            encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                            encode_lines.append(f"\t\t\t\tfor _, i := range ie.{ext_field_name} {{")
                            encode_lines.append(f"\t\t\t\t\ttmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                            encode_lines.append(f"\t\t\t\t\ttmp_{ext_field_name}.Value = append(tmp_{ext_field_name}.Value, &tmp_ie)")
                            encode_lines.append(f"\t\t\t\t}}")
                        elif ext_field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                            encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                            encode_lines.append(f"\t\t\t\tfor _, i := range ie.{ext_field_name} {{")
                            encode_lines.append(f"\t\t\t\t\ttmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                            encode_lines.append(f"\t\t\t\t\ttmp_{ext_field_name}.Value = append(tmp_{ext_field_name}.Value, &tmp_ie)")
                            encode_lines.append(f"\t\t\t\t}}")
                        elif ext_field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                            encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                            encode_lines.append(f"\t\t\t\tfor _, i := range ie.{ext_field_name} {{")
                            encode_lines.append(f"\t\t\t\t\ttmp_ie := utils.NewBITSTRING(i, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                            encode_lines.append(f"\t\t\t\t\ttmp_{ext_field_name}.Value = append(tmp_{ext_field_name}.Value, &tmp_ie)")
                            encode_lines.append(f"\t\t\t\t}}")
                    else:
                        element_type = to_go_name(ext_field.element_type) if ext_field.element_type else "interface{}"
                        encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        encode_lines.append(f"\t\t\t\tfor _, i := range ie.{ext_field_name} {{")
                        encode_lines.append(f"\t\t\t\t\ttmp_{ext_field_name}.Value = append(tmp_{ext_field_name}.Value, &i)")
                        encode_lines.append(f"\t\t\t\t}}")
                    encode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Encode(extWriter); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name == "INTEGER":
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteInteger(*ie.{ext_field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteInteger(ie.{ext_field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name == "ENUMERATED":
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteEnumerate(uint64(*ie.{ext_field_name}), uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteEnumerate(uint64(ie.{ext_field_name}), uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteBitString(ie.{ext_field_name}.Bytes, uint(ie.{ext_field_name}.NumBits), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteBitString(ie.{ext_field_name}.Bytes, uint(ie.{ext_field_name}.NumBits), &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    normalized_lb, normalized_ub = normalize_octetstring_constraints(ext_field.lb, ext_field.ub)
                    lb = format_constraint_for_code(normalized_lb)
                    ub = format_constraint_for_code(normalized_ub)
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteOctetString(*ie.{ext_field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteOctetString(ie.{ext_field_name}, &uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name == "BOOLEAN":
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteBoolean(*ie.{ext_field_name}); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = extWriter.WriteBoolean(ie.{ext_field_name}); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name == "NULL":
                    encode_lines.append(f"\t\t\t\tif err := extWriter.WriteNull(); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                elif is_setuprelease_type(ext_field.type_name):
                    inner_type = extract_setuprelease_inner_type(ext_field.type_name)
                    go_type = to_go_name(inner_type)
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.SetupRelease[*{go_type}]{{")
                        encode_lines.append(f"\t\t\t\t\tSetup: ie.{ext_field_name},")
                        encode_lines.append(f"\t\t\t\t}}")
                    else:
                        encode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.SetupRelease[*{go_type}]{{")
                        encode_lines.append(f"\t\t\t\t\tSetup: ie.{ext_field_name},")
                        encode_lines.append(f"\t\t\t\t}}")
                    encode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Encode(extWriter); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                else:
                    # Custom type - call Encode method
                    if ext_field.optional:
                        encode_lines.append(f"\t\t\t\tif err = ie.{ext_field_name}.Encode(extWriter); err != nil {{")
                    else:
                        encode_lines.append(f"\t\t\t\tif err = ie.{ext_field_name}.Encode(extWriter); err != nil {{")
                    encode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Encode {ext_field_name}\", err)")
                    encode_lines.append(f"\t\t\t\t}}")
                
                if ext_field.optional:
                    encode_lines.append(f"\t\t\t}}")
            
            encode_lines.append("")
            encode_lines.append("\t\t\tif err := extWriter.Close(); err != nil {")
            encode_lines.append("\t\t\t\treturn err")
            encode_lines.append("\t\t\t}")
            encode_lines.append("")
            encode_lines.append("\t\t\tif err := w.WriteOpenType(extBuf.Bytes()); err != nil {")
            encode_lines.append("\t\t\t\treturn err")
            encode_lines.append("\t\t\t}")
            encode_lines.append("\t\t}")
        
        encode_lines.append("\t}")
    
    encode_block = "\n".join(encode_lines)
    
    # Generate decode code
    decode_lines = []
    
    # Read preamble bits
    preamble_count = (1 if has_extensions else 0) + optional_root_count
    if preamble_count > 0:
        if has_extensions:
            decode_lines.append("\tvar extensionBit bool")
            decode_lines.append("\tif extensionBit, err = r.ReadBool(); err != nil {")
            decode_lines.append("\t\treturn err")
            decode_lines.append("\t}")
        
        # Read optional root field bits (excluding Struct-Null and interface{})
        for field in root_fields:
            if field.optional and field.type_name != "Struct-Null" and field.type_name != "interface{}":
                field_name = to_go_name(field.name)
                decode_lines.append(f"\tvar {field_name}Present bool")
                decode_lines.append(f"\tif {field_name}Present, err = r.ReadBool(); err != nil {{")
                decode_lines.append("\t\treturn err")
                decode_lines.append("\t}")
    
    # Decode root fields only
    for field in root_fields:
        field_name = to_go_name(field.name)
        
        # Handle optional fields
        if field.optional:
            # Handle optional SetupRelease
            if is_setuprelease_type(field.type_name):
                decode_lines.append(f"\tif {field_name}Present {{")
                inner_type = extract_setuprelease_inner_type(field.type_name)
                go_type = to_go_name(inner_type)
                decode_lines.append(f"\t\ttmp_{field_name} := utils.SetupRelease[*{go_type}]{{}}")
                decode_lines.append(f"\t\tif err = tmp_{field_name}.Decode(r); err != nil {{")
                decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                decode_lines.append(f"\t\t}}")
                decode_lines.append(f"\t\tie.{field_name} = tmp_{field_name}.Setup")
                decode_lines.append(f"\t}}")
                continue
            
            # For optional SEQUENCE_OF, check if present and decode
            if field.type_name == "SEQUENCE_OF":
                decode_lines.append(f"\tif {field_name}Present {{")
                is_base = is_base_type_element(field.element_type) if field.element_type else False
                go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type) if field.element_type else ("[]interface{}", "interface{}")
                lb = format_constraint_for_code(field.lb)
                ub = format_constraint_for_code(field.ub)
                
                if is_base:
                    element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
                    element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
                    decode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    decode_lines.append(f"\t\tfn_{field_name} := func() *{utils_type_seq} {{")
                    if field.element_type == "INTEGER":
                        decode_lines.append(f"\t\t\tie := utils.NewINTEGER(0, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                        decode_lines.append(f"\t\t\tie := utils.NewOCTETSTRING([]byte{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                        decode_lines.append(f"\t\t\tie := utils.NewBITSTRING(uper.BitString{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                    decode_lines.append(f"\t\t\treturn &ie")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tif err = tmp_{field_name}.Decode(r, fn_{field_name}); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = {go_type_seq}{{}}")
                    if field.element_type == "INTEGER":
                        decode_lines.append(f"\t\tfor _, i := range tmp_{field_name}.Value {{")
                        decode_lines.append(f"\t\t\tie.{field_name} = append(ie.{field_name}, int64(i.Value))")
                        decode_lines.append(f"\t\t}}")
                    else:
                        decode_lines.append(f"\t\tfor _, i := range tmp_{field_name}.Value {{")
                        decode_lines.append(f"\t\t\tie.{field_name} = append(ie.{field_name}, i.Value)")
                        decode_lines.append(f"\t\t}}")
                else:
                    element_type = to_go_name(field.element_type) if field.element_type else "interface{}"
                    decode_lines.append(f"\t\ttmp_{field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                    decode_lines.append(f"\t\tfn_{field_name} := func() *{element_type} {{")
                    decode_lines.append(f"\t\t\treturn new({element_type})")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tif err = tmp_{field_name}.Decode(r, fn_{field_name}); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = []{element_type}{{}}")
                    decode_lines.append(f"\t\tfor _, i := range tmp_{field_name}.Value {{")
                    decode_lines.append(f"\t\t\tie.{field_name} = append(ie.{field_name}, *i)")
                    decode_lines.append(f"\t\t}}")
                decode_lines.append(f"\t\t}}")
            else:
                # Only generate if block if field is not Struct-Null and not interface{}
                if field.type_name != "Struct-Null" and field.type_name != "interface{}":
                    decode_lines.append(f"\tif {field_name}Present {{")
                
                # Decode the optional field based on its type
                if field.type_name == "INTEGER":
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    decode_lines.append(f"\t\tvar tmp_int_{field_name} int64")
                    decode_lines.append(f"\t\tif tmp_int_{field_name}, err = r.ReadInteger(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = &tmp_int_{field_name}")
                elif field.type_name == "ENUMERATED":
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    decode_lines.append(f"\t\tvar tmp_enum_{field_name} uint64")
                    decode_lines.append(f"\t\tif tmp_enum_{field_name}, err = r.ReadEnumerate(uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\ttmp_enum_val := uper.Enumerated(tmp_enum_{field_name})")
                    decode_lines.append(f"\t\tie.{field_name} = &tmp_enum_val")
                elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
                    lb = format_constraint_for_code(field.lb)
                    ub = format_constraint_for_code(field.ub)
                    decode_lines.append(f"\t\tvar tmp_bs_{field_name} []byte")
                    decode_lines.append(f"\t\tvar n_{field_name} uint")
                    decode_lines.append(f"\t\tif tmp_bs_{field_name}, n_{field_name}, err = r.ReadBitString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\ttmp_bitstring := uper.BitString{{")
                    decode_lines.append(f"\t\t\tBytes:   tmp_bs_{field_name},")
                    decode_lines.append(f"\t\t\tNumBits: uint64(n_{field_name}),")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = &tmp_bitstring")
                elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    # Normalize constraints for CONTAINING
                    normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
                    lb = format_constraint_for_code(normalized_lb)
                    ub = format_constraint_for_code(normalized_ub)
                    decode_lines.append(f"\t\tvar tmp_os_{field_name} []byte")
                    decode_lines.append(f"\t\tif tmp_os_{field_name}, err = r.ReadOctetString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = &tmp_os_{field_name}")
                elif field.type_name == "BOOLEAN":
                    decode_lines.append(f"\t\tvar tmp_bool_{field_name} bool")
                    decode_lines.append(f"\t\tif tmp_bool_{field_name}, err = r.ReadBoolean(); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                    decode_lines.append(f"\t\tie.{field_name} = &tmp_bool_{field_name}")
                elif field.type_name == "NULL":
                    decode_lines.append(f"\t\tif err := r.ReadNull(); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                elif field.type_name == "Struct-Null":
                    # Struct-Null is empty
                    pass
                elif field.type_name == "interface{}":
                    # interface{} has no methods, skip decoding
                    pass
                else:
                    # Custom type - call Decode method (optional fields are pointers, so new() is correct)
                    decode_lines.append(f"\t\tie.{field_name} = new({get_go_type_name(field.type_name, field)})")
                    decode_lines.append(f"\t\tif err = ie.{field_name}.Decode(r); err != nil {{")
                    decode_lines.append(f"\t\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                    decode_lines.append(f"\t\t}}")
                
                # Only close if block if it was opened (not Struct-Null and not interface{})
                if field.type_name != "Struct-Null" and field.type_name != "interface{}":
                    decode_lines.append(f"\t}}")
            continue
        
        # Handle mandatory SetupRelease
        if is_setuprelease_type(field.type_name):
            inner_type = extract_setuprelease_inner_type(field.type_name)
            go_type = to_go_name(inner_type)
            decode_lines.append(f"\ttmp_{field_name} := utils.SetupRelease[*{go_type}]{{}}")
            decode_lines.append(f"\tif err = tmp_{field_name}.Decode(r); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
            decode_lines.append(f"\t}}")
            decode_lines.append(f"\tie.{field_name} = tmp_{field_name}.Setup")
            continue
        
        # Handle different field types
        if field.type_name == "INTEGER":
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            decode_lines.append(f"\tvar tmp_int_{field_name} int64")
            decode_lines.append(f"\tif tmp_int_{field_name}, err = r.ReadInteger(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"ReadInteger {field_name}\", err)")
            decode_lines.append(f"\t}}")
            decode_lines.append(f"\tie.{field_name} = tmp_int_{field_name}")
        elif field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            decode_lines.append(f"\tvar tmp_bs_{field_name} []byte")
            decode_lines.append(f"\tvar n_{field_name} uint")
            decode_lines.append(f"\tif tmp_bs_{field_name}, n_{field_name}, err = r.ReadBitString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"ReadBitString {field_name}\", err)")
            decode_lines.append(f"\t}}")
            decode_lines.append(f"\tie.{field_name} = uper.BitString{{")
            decode_lines.append(f"\t\tBytes:   tmp_bs_{field_name},")
            decode_lines.append(f"\t\tNumBits: uint64(n_{field_name}),")
            decode_lines.append(f"\t}}")
        elif field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
            # Normalize constraints for CONTAINING
            normalized_lb, normalized_ub = normalize_octetstring_constraints(field.lb, field.ub)
            lb = format_constraint_for_code(normalized_lb)
            ub = format_constraint_for_code(normalized_ub)
            decode_lines.append(f"\tvar tmp_os_{field_name} []byte")
            decode_lines.append(f"\tif tmp_os_{field_name}, err = r.ReadOctetString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"ReadOctetString {field_name}\", err)")
            decode_lines.append(f"\t}}")
            decode_lines.append(f"\tie.{field_name} = tmp_os_{field_name}")
        elif field.type_name == "BOOLEAN":
            decode_lines.append(f"\tvar tmp_bool_{field_name} bool")
            decode_lines.append(f"\tif tmp_bool_{field_name}, err = r.ReadBoolean(); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"ReadBoolean {field_name}\", err)")
            decode_lines.append(f"\t}}")
            decode_lines.append(f"\tie.{field_name} = tmp_bool_{field_name}")
        elif field.type_name == "NULL":
            decode_lines.append(f"\tif err := r.ReadNull(); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
            decode_lines.append(f"\t}}")
        elif field.type_name == "SEQUENCE_OF":
            is_base = is_base_type_element(field.element_type) if field.element_type else False
            go_type_seq, utils_type_seq = get_sequence_of_go_type(field.element_type) if field.element_type else ("[]interface{}", "interface{}")
            lb = format_constraint_for_code(field.lb)
            ub = format_constraint_for_code(field.ub)
            
            if is_base:
                element_lb = format_constraint_for_code(field.element_lb) if field.element_lb is not None else 0
                element_ub = format_constraint_for_code(field.element_ub) if field.element_ub is not None else 0
                decode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                decode_lines.append(f"\tfn_{field_name} := func() *{utils_type_seq} {{")
                if field.element_type == "INTEGER":
                    decode_lines.append(f"\t\tie := utils.NewINTEGER(0, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                elif field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    decode_lines.append(f"\t\tie := utils.NewOCTETSTRING([]byte{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                elif field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                    decode_lines.append(f"\t\tie := utils.NewBITSTRING(uper.BitString{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                decode_lines.append(f"\t\treturn &ie")
                decode_lines.append(f"\t}}")
                decode_lines.append(f"\tif err = tmp_{field_name}.Decode(r, fn_{field_name}); err != nil {{")
                decode_lines.append(f"\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                decode_lines.append(f"\t}}")
                decode_lines.append(f"\tie.{field_name} = {go_type_seq}{{}}")
                if field.element_type == "INTEGER":
                    decode_lines.append(f"\tfor _, i := range tmp_{field_name}.Value {{")
                    decode_lines.append(f"\t\tie.{field_name} = append(ie.{field_name}, int64(i.Value))")
                    decode_lines.append(f"\t}}")
                else:
                    decode_lines.append(f"\tfor _, i := range tmp_{field_name}.Value {{")
                    decode_lines.append(f"\t\tie.{field_name} = append(ie.{field_name}, i.Value)")
                    decode_lines.append(f"\t}}")
            else:
                element_type = to_go_name(field.element_type) if field.element_type else "interface{}"
                decode_lines.append(f"\ttmp_{field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                decode_lines.append(f"\tfn_{field_name} := func() *{element_type} {{")
                decode_lines.append(f"\t\treturn new({element_type})")
                decode_lines.append(f"\t}}")
                decode_lines.append(f"\tif err = tmp_{field_name}.Decode(r, fn_{field_name}); err != nil {{")
                decode_lines.append(f"\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
                decode_lines.append(f"\t}}")
                decode_lines.append(f"\tie.{field_name} = []{element_type}{{}}")
                decode_lines.append(f"\tfor _, i := range tmp_{field_name}.Value {{")
                decode_lines.append(f"\t\tie.{field_name} = append(ie.{field_name}, *i)")
                decode_lines.append(f"\t}}")
        elif go_type == "interface{}":
            # interface{} has no methods, skip decoding
            pass
        else:
            # Custom type - call Decode method directly (mandatory fields are value types, not pointers)
            decode_lines.append(f"\tif err = ie.{field_name}.Decode(r); err != nil {{")
            decode_lines.append(f"\t\treturn utils.WrapError(\"Decode {field_name}\", err)")
            decode_lines.append(f"\t}}")
    
    # Decode extension groups if extensionBit is true
    if has_extensions:
        decode_lines.append("")
        decode_lines.append("\tif extensionBit {")
        decode_lines.append(f"\t\t// Read extension bitmap: {len(ext_groups)} bits for {len(ext_groups)} extension groups")
        decode_lines.append(f"\t\textBitmap, err := r.ReadExtBitMap()")
        decode_lines.append("\t\tif err != nil {")
        decode_lines.append("\t\t\treturn err")
        decode_lines.append("\t\t}")
        decode_lines.append("")
        
        # Decode each extension group
        for group_num in sorted(ext_groups.keys()):
            ext_fields = ext_groups[group_num]
            decode_lines.append(f"\t\t// decode extension group {group_num}")
            decode_lines.append(f"\t\tif len(extBitmap) > {group_num-1} && extBitmap[{group_num-1}] {{")
            decode_lines.append("\t\t\textBytes, err := r.ReadOpenType()")
            decode_lines.append("\t\t\tif err != nil {")
            decode_lines.append("\t\t\t\treturn err")
            decode_lines.append("\t\t\t}")
            decode_lines.append("")
            decode_lines.append("\t\t\textReader := uper.NewReader(bytes.NewReader(extBytes))")
            decode_lines.append("")
            
            # Read preamble bits for optional fields in this extension group
            # Skip Struct-Null and interface{} fields
            optional_ext_fields = [f for f in ext_fields if f.optional and f.type_name != "Struct-Null" and f.type_name != "interface{}"]
            if optional_ext_fields:
                for ext_field in ext_fields:
                    if ext_field.optional and ext_field.type_name != "Struct-Null" and ext_field.type_name != "interface{}":
                        ext_field_name = to_go_name(ext_field.name)
                        decode_lines.append(f"\t\t\t{ext_field_name}Present, err := extReader.ReadBool()")
                        decode_lines.append("\t\t\tif err != nil {")
                        decode_lines.append("\t\t\t\treturn err")
                        decode_lines.append("\t\t\t}")
            
            # Decode extension fields
            for ext_field in ext_fields:
                ext_field_name = to_go_name(ext_field.name)
                ext_go_type = get_go_type_name(ext_field.type_name, ext_field)
                
                # Skip Struct-Null and interface{} fields (nothing to decode)
                if ext_field.type_name == "Struct-Null" or ext_field.type_name == "interface{}":
                    continue
                
                if ext_field.optional:
                    decode_lines.append(f"\t\t\t// decode {ext_field_name} optional")
                    decode_lines.append(f"\t\t\tif {ext_field_name}Present {{")
                else:
                    decode_lines.append(f"\t\t\t// decode {ext_field_name}")
                
                # Generate decoding code for this extension field (similar to root fields but using extReader)
                if ext_field.type_name == "SEQUENCE_OF":
                    is_base = is_base_type_element(ext_field.element_type) if ext_field.element_type else False
                    go_type_seq, utils_type_seq = get_sequence_of_go_type(ext_field.element_type) if ext_field.element_type else ("[]interface{}", "interface{}")
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    
                    if is_base:
                        element_lb = format_constraint_for_code(ext_field.element_lb) if ext_field.element_lb is not None else 0
                        element_ub = format_constraint_for_code(ext_field.element_ub) if ext_field.element_ub is not None else 0
                        decode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{utils_type_seq}]([]*{utils_type_seq}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        decode_lines.append(f"\t\t\t\tfn_{ext_field_name} := func() *{utils_type_seq} {{")
                        if ext_field.element_type == "INTEGER":
                            decode_lines.append(f"\t\t\t\t\tie := utils.NewINTEGER(0, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        elif ext_field.element_type in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                            decode_lines.append(f"\t\t\t\t\tie := utils.NewOCTETSTRING([]byte{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        elif ext_field.element_type in ["BIT_STRING", "BITSTRING", "BitString"]:
                            decode_lines.append(f"\t\t\t\t\tie := utils.NewBITSTRING(uper.BitString{{}}, uper.Constraint{{Lb: {element_lb}, Ub: {element_ub}}}, false)")
                        decode_lines.append(f"\t\t\t\t\treturn &ie")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Decode(extReader, fn_{ext_field_name}); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = {go_type_seq}{{}}")
                        if ext_field.element_type == "INTEGER":
                            decode_lines.append(f"\t\t\t\tfor _, i := range tmp_{ext_field_name}.Value {{")
                            decode_lines.append(f"\t\t\t\t\tie.{ext_field_name} = append(ie.{ext_field_name}, int64(i.Value))")
                            decode_lines.append(f"\t\t\t\t}}")
                        else:
                            decode_lines.append(f"\t\t\t\tfor _, i := range tmp_{ext_field_name}.Value {{")
                            decode_lines.append(f"\t\t\t\t\tie.{ext_field_name} = append(ie.{ext_field_name}, i.Value)")
                            decode_lines.append(f"\t\t\t\t}}")
                    else:
                        element_type = to_go_name(ext_field.element_type) if ext_field.element_type else "interface{}"
                        decode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.NewSequence[*{element_type}]([]*{element_type}{{}}, uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false)")
                        decode_lines.append(f"\t\t\t\tfn_{ext_field_name} := func() *{element_type} {{")
                        decode_lines.append(f"\t\t\t\t\treturn new({element_type})")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Decode(extReader, fn_{ext_field_name}); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = []{element_type}{{}}")
                        decode_lines.append(f"\t\t\t\tfor _, i := range tmp_{ext_field_name}.Value {{")
                        decode_lines.append(f"\t\t\t\t\tie.{ext_field_name} = append(ie.{ext_field_name}, *i)")
                        decode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name == "INTEGER":
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tvar tmp_int_{ext_field_name} int64")
                        decode_lines.append(f"\t\t\t\tif tmp_int_{ext_field_name}, err = extReader.ReadInteger(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = &tmp_int_{ext_field_name}")
                    else:
                        decode_lines.append(f"\t\t\t\tvar tmp_int_{ext_field_name} int64")
                        decode_lines.append(f"\t\t\t\tif tmp_int_{ext_field_name}, err = extReader.ReadInteger(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = tmp_int_{ext_field_name}")
                elif ext_field.type_name == "ENUMERATED":
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tvar tmp_enum_{ext_field_name} uint64")
                        decode_lines.append(f"\t\t\t\tif tmp_enum_{ext_field_name}, err = extReader.ReadEnumerate(uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\ttmp_enum_val := uper.Enumerated(tmp_enum_{ext_field_name})")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = &tmp_enum_val")
                    else:
                        decode_lines.append(f"\t\t\t\tvar tmp_enum_{ext_field_name} uint64")
                        decode_lines.append(f"\t\t\t\tif tmp_enum_{ext_field_name}, err = extReader.ReadEnumerate(uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = uper.Enumerated(tmp_enum_{ext_field_name})")
                elif ext_field.type_name in ["BIT_STRING", "BITSTRING", "BitString"]:
                    lb = format_constraint_for_code(ext_field.lb)
                    ub = format_constraint_for_code(ext_field.ub)
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tvar tmp_bs_{ext_field_name} []byte")
                        decode_lines.append(f"\t\t\t\tvar n_{ext_field_name} uint")
                        decode_lines.append(f"\t\t\t\tif tmp_bs_{ext_field_name}, n_{ext_field_name}, err = extReader.ReadBitString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\ttmp_bitstring := uper.BitString{{")
                        decode_lines.append(f"\t\t\t\t\tBytes:   tmp_bs_{ext_field_name},")
                        decode_lines.append(f"\t\t\t\t\tNumBits: uint64(n_{ext_field_name}),")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = &tmp_bitstring")
                    else:
                        decode_lines.append(f"\t\t\t\tvar tmp_bs_{ext_field_name} []byte")
                        decode_lines.append(f"\t\t\t\tvar n_{ext_field_name} uint")
                        decode_lines.append(f"\t\t\t\tif tmp_bs_{ext_field_name}, n_{ext_field_name}, err = extReader.ReadBitString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = uper.BitString{{")
                        decode_lines.append(f"\t\t\t\t\tBytes:   tmp_bs_{ext_field_name},")
                        decode_lines.append(f"\t\t\t\t\tNumBits: uint64(n_{ext_field_name}),")
                        decode_lines.append(f"\t\t\t\t}}")
                elif ext_field.type_name in ["OCTET_STRING", "OCTETSTRING", "OctetString"]:
                    normalized_lb, normalized_ub = normalize_octetstring_constraints(ext_field.lb, ext_field.ub)
                    lb = format_constraint_for_code(normalized_lb)
                    ub = format_constraint_for_code(normalized_ub)
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tvar tmp_os_{ext_field_name} []byte")
                        decode_lines.append(f"\t\t\t\tif tmp_os_{ext_field_name}, err = extReader.ReadOctetString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = &tmp_os_{ext_field_name}")
                    else:
                        decode_lines.append(f"\t\t\t\tvar tmp_os_{ext_field_name} []byte")
                        decode_lines.append(f"\t\t\t\tif tmp_os_{ext_field_name}, err = extReader.ReadOctetString(&uper.Constraint{{Lb: {lb}, Ub: {ub}}}, false); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = tmp_os_{ext_field_name}")
                elif ext_field.type_name == "BOOLEAN":
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tvar tmp_bool_{ext_field_name} bool")
                        decode_lines.append(f"\t\t\t\tif tmp_bool_{ext_field_name}, err = extReader.ReadBoolean(); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = &tmp_bool_{ext_field_name}")
                    else:
                        decode_lines.append(f"\t\t\t\tvar tmp_bool_{ext_field_name} bool")
                        decode_lines.append(f"\t\t\t\tif tmp_bool_{ext_field_name}, err = extReader.ReadBoolean(); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = tmp_bool_{ext_field_name}")
                elif ext_field.type_name == "NULL":
                    decode_lines.append(f"\t\t\t\tif err := extReader.ReadNull(); err != nil {{")
                    decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                    decode_lines.append(f"\t\t\t\t}}")
                elif is_setuprelease_type(ext_field.type_name):
                    inner_type = extract_setuprelease_inner_type(ext_field.type_name)
                    go_type = to_go_name(inner_type)
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.SetupRelease[*{go_type}]{{}}")
                        decode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Decode(extReader); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = tmp_{ext_field_name}.Setup")
                    else:
                        decode_lines.append(f"\t\t\t\ttmp_{ext_field_name} := utils.SetupRelease[*{go_type}]{{}}")
                        decode_lines.append(f"\t\t\t\tif err = tmp_{ext_field_name}.Decode(extReader); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = tmp_{ext_field_name}.Setup")
                else:
                    # Custom type - call Decode method
                    if ext_field.optional:
                        decode_lines.append(f"\t\t\t\tie.{ext_field_name} = new({ext_go_type})")
                        decode_lines.append(f"\t\t\t\tif err = ie.{ext_field_name}.Decode(extReader); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                    else:
                        decode_lines.append(f"\t\t\t\tif err = ie.{ext_field_name}.Decode(extReader); err != nil {{")
                        decode_lines.append(f"\t\t\t\t\treturn utils.WrapError(\"Decode {ext_field_name}\", err)")
                        decode_lines.append(f"\t\t\t\t}}")
                
                if ext_field.optional:
                    decode_lines.append(f"\t\t\t}}")
            
            decode_lines.append("\t\t}")
        
        decode_lines.append("\t}")
    
    decode_block = "\n".join(decode_lines)
    
    # Build imports
    imports = ['"github.com/lvdund/asn1go/uper"', '"github.com/lvdund/rrc/utils"']
    if has_extensions:
        imports.append('"bytes"')
    imports_str = "\n\t".join(imports)
    
    code = f"""package ies

import (
	{imports_str}
)

type {go_name} struct {{
{struct_block}
}}

func (ie *{go_name}) Encode(w *uper.UperWriter) error {{
	var err error
{encode_block}
	return nil
}}

func (ie *{go_name}) Decode(r *uper.UperReader) error {{
	var err error
{decode_block}
	return nil
}}
"""
    return code


def convert_to_golang(asn1_types: List[Asn1Type], output_dir: Path = None):
    """Convert ASN.1 types to Go code files"""
    if output_dir is None:
        output_dir = Path(__file__).parent / "ies"
    else:
        output_dir = Path(output_dir)
    
    output_dir.mkdir(exist_ok=True)
    
    # Sort types by dependencies (topological sort)
    print("Sorting types by dependencies...")
    sorted_types = topological_sort_types(asn1_types)
    print(f"Sorted {len(sorted_types)} types")
    
    # Create a map for quick lookup
    type_map = {asn1_type.name: asn1_type for asn1_type in sorted_types}
    
    generated_count = 0
    skipped_count = 0
    
    for asn1_type in sorted_types:
        # Generate all types, including those with parent (they're still valid types)        
        go_name = to_go_name(asn1_type.name)
        file_name = f"{go_name}.go"
        file_path = output_dir / file_name
        if go_name in {'PDCP_Config_moreThanTwoRLC_DRB_r16', 
                        'MeasurementTimingConfiguration_v1610_IEs',
                        'PDCP_Config_moreThanTwoRLC_DRB_r16',
                        'RLF_TimersAndConstants',
                        'PDSCH_HARQ_ACK_CodebookList_r16'}:
            continue
        
        # Generate code based on type_name
        code = ""
        try:
            if asn1_type.type_name == "INTEGER":
                code = generate_integer_code(asn1_type)
            elif asn1_type.type_name == "ENUMERATED":
                code = generate_enumerated_code(asn1_type)
            elif asn1_type.type_name in ["BIT_STRING", "BITSTRING"]:
                code = generate_bitstring_code(asn1_type)
            elif asn1_type.type_name in ["OCTET_STRING", "OCTETSTRING"]:
                code = generate_octetstring_code(asn1_type)
            elif asn1_type.type_name == "BOOLEAN":
                code = generate_boolean_code(asn1_type)
            elif asn1_type.type_name == "NULL":
                code = generate_null_code(asn1_type)
            elif asn1_type.type_name == "SEQUENCE_OF":
                code = generate_sequence_of_code(asn1_type)
            elif asn1_type.type_name == "CHOICE":
                code = generate_choice_code(asn1_type)
            elif asn1_type.type_name == "SEQUENCE":
                code = generate_sequence_code(asn1_type)
            else:
                # Unknown type - skip
                skipped_count += 1
                continue
            
            if code:
                with open(file_path, 'w', encoding='utf-8') as f:
                    f.write(code)
                generated_count += 1
        except Exception as e:
            print(f"Error generating code for {asn1_type.name}: {e}")
            skipped_count += 1
            continue
    
    print(f"Generated {generated_count} Go files")
    print(f"Skipped {skipped_count} types")





def main():
    print("Files that would be deleted:")
    for filename in os.listdir("ies"):
        file_path = os.path.join("ies", filename)
        if os.path.isfile(file_path) and not filename.startswith("00"):
            os.remove(file_path)
            
    asn1_types = load_asn1_types(str(json_path))
    print(f"Loaded {len(asn1_types)} ASN.1 types")
    
    # Convert to Go code
    convert_to_golang(asn1_types)

    os.system("make format")


if __name__ == "__main__":
    main()

