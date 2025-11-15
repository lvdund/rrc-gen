import re
import json
from dataclasses import dataclass, field, asdict
from typing import Optional, List, Dict, Any
from pathlib import Path
import sys

# Default input file
checking = False
if checking:
    asn1_file = Path(__file__).parent / "cache" / "check.asn"
    output_file = Path(__file__).parent / "cache" / "check.json"
else:
    asn1_file = Path(__file__).parent / "cache" / "nr-rrc-17.3.0.asn1"
    output_file = Path(__file__).parent / "cache" / "asn1_types.json"





def uppercase_first_char(s):
    if not s:
        return s
    return s[0].upper() + s[1:]

def get_after_dash(var):
    """Get everything before, after '--' in a string"""
    after_dash = ""
    if isinstance(var, str) and "--" in var:
        after_dash = var.split("--")[1].strip()
        if after_dash == "]]" or after_dash == "]]," or after_dash == "[[":
            return False
        return var.split("--")[0].strip(), after_dash
    return False

def has_ellipsis_at_root_level(data: str) -> bool:
    """
    Check if the data has "..." at the root level
    """
    lines = data.splitlines()
    if not lines:
        return False

    # Find the first non-empty line to treat as "line 1"
    first_line = None
    first_indent = None
    for line in lines:
        stripped = line.lstrip()
        if stripped:  # non-empty
            first_line = line
            first_indent = len(line) - len(stripped)
            break

    if first_line is None:
        return False  # all lines empty

    # Now check all lines for same indent level and content "..."
    for line in lines:
        stripped = line.lstrip()
        if not stripped:
            continue  # skip empty lines

        current_indent = len(line) - len(stripped)
        if current_indent == first_indent and stripped == "...":
            return True

    return False

def extract_sequence_bounds(text) -> dict:
    types = ['INTEGER', 'BITSTRING', 'OCTETSTRING', 'BOOLEAN', 'ENUMERATED']
    types_pattern = '|'.join(types)
    
    # Match bounds as generic tokens: digits, letters, underscores, hyphens
    # e.g., "1", "16", "max", "max-abc", "MAX_INT"
    bound_pattern = r'([a-zA-Z0-9_\-]+)'
    
    pattern = re.compile(
        rf'SEQUENCE\s*\(\s*SIZE\s*\(\s*{bound_pattern}\s*\.\.\s*{bound_pattern}\s*\)\s*\)\s*OF\s+({types_pattern})\s*\(\s*{bound_pattern}\s*\.\.\s*{bound_pattern}\s*\)',
        re.IGNORECASE
    )
    
    match = pattern.search(text)
    if match:
        lb, ub, typ, lb_e, ub_e = match.groups()
        return {
            'check': True,
            'lb': lb,
            'ub': ub,
            'type': typ,
            'lb_e': lb_e,
            'ub_e': ub_e
        }
    else:
        return {
            'check': False,
            'lb': '',
            'ub': '',
            'type': '',
            'lb_e': '',
            'ub_e': ''
        }

def parse_optional_string(input_string: str) -> tuple[str, bool]:
    """
    Simple version: Remove OPTIONAL and return True if found.
    """
    if not input_string:
        return input_string, False
    
    # Remove trailing whitespace
    s = input_string.rstrip()
    
    # Check for OPTIONAL pattern (case insensitive)
    if 'OPTIONAL' in s.upper():
        # Split by '--' to handle comments
        parts = s.split('--', 1)
        main_part = parts[0]
        comment_part = parts[1] if len(parts) > 1 else ""
        
        # Remove OPTIONAL from main part
        main_part = re.sub(r'\s+OPTIONAL\s*(?:,)?\s*$', '', main_part, flags=re.IGNORECASE)
        
        # Check comment length (max 3 words)
        if comment_part:
            comment_words = comment_part.strip().split()
            if len(comment_words) > 3:
                # Too many words in comment, don't treat as OPTIONAL
                return input_string, False
        
        # Reconstruct string without OPTIONAL
        if comment_part:
            result = f"{main_part.strip()} -- {comment_part}".strip()
        else:
            result = main_part.strip()
            
        return result, True
    
    return input_string, False





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
    optional: bool = False
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
        if self.lb is not None:
            result["lb"] = self.lb
        if self.ub is not None:
            result["ub"] = self.ub
        if self.optional:
            result["optional"] = self.optional
        if self.parent:
            result["parent"] = self.parent
        if self.element_type:
            result["element_type"] = self.element_type
        if self.element_lb is not None:
            result["element_lb"] = self.element_lb
        if self.element_ub is not None:
            result["element_ub"] = self.element_ub
        if self.ext:
            result["ext"] = self.ext
        if self.comment is not None:
            result["comment"] = self.comment
        return result


class Asn1Parser:
    """Parser for ASN.1 RRC specification files"""
    
    def __init__(self):
        self.types: Dict[str, Asn1Type] = {}
        self.constants: Dict[str, int] = {}
        
    def parse_file(self, filepath: str):
        """Parse an ASN.1 file and load all type definitions"""
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Remove comments but keep inline ones for field metadata
        # First, let's process the content line by line
        self.parse_content(content)
        
    def parse_content(self, content: str):
        """Parse ASN.1 content"""
        # Remove multi-line comments /* ... */
        content = re.sub(r'/\*.*?\*/', '', content, flags=re.DOTALL)
        
        # Find all type definitions (Type-Name ::= ...)
        # Pattern: identifier ::= type_definition
        pattern = r'([a-zA-Z][a-zA-Z0-9\-]*)\s*::=\s*'
        
        # Split content by definitions
        parts = re.split(pattern, content)
        
        # Process pairs: parts[i] = name, parts[i+1] = definition
        for i in range(1, len(parts), 2):
            if i + 1 >= len(parts):
                break
                
            type_name = parts[i].strip()
            definition = parts[i + 1].strip()
            
            # Extract definition until the next type or end
            # Find the end of this definition (before next ::= or END)
            definition = self._extract_definition(definition)
            try:
                asn1_type = self.parse_type_definition(type_name, definition)

                if asn1_type:
                    normalized_name = type_name
                    asn1_type.name = normalized_name
                    self.types[normalized_name] = asn1_type
            except Exception as e:
                print(f"Error parsing {type_name}: {e}")
                
    def _extract_definition(self, text: str) -> str:
        """Extract a single type definition"""
        # Find the complete definition by counting braces
        brace_count = 0
        result = []
        
        for i, char in enumerate(text):
            result.append(char)
            
            if char == '{':
                brace_count += 1
            elif char == '}':
                brace_count -= 1
                if brace_count == 0:
                    # Definition complete
                    return ''.join(result)
            elif char == '\n' and brace_count == 0:
                # Simple type without braces
                return ''.join(result).split('\n')[0]
                
        return ''.join(result)
    
    def _extract_braced_content(self, text: str, keyword: str) -> Optional[str]:
        """
        Extract content between balanced braces after a keyword.
        E.g., for 'SEQUENCE { ... }' returns the content between braces.
        """
        # Find the keyword
        pattern = re.search(rf'{keyword}\s*\{{', text)
        if not pattern:
            return None
        
        # Start after the opening brace
        start = pattern.end() - 1  # Position of '{'
        brace_count = 0
        content_start = start + 1
        
        for i in range(start, len(text)):
            if text[i] == '{':
                brace_count += 1
            elif text[i] == '}':
                brace_count -= 1
                if brace_count == 0:
                    # Found matching closing brace
                    return text[content_start:i]
        
        # No matching closing brace found
        return None
        
    def parse_type_definition(self, name: str, definition: str) -> Optional[Asn1Type]:
        """Parse a type definition and return Asn1Type"""
        definition = definition.strip()

        # Check for SEQUENCE
        if definition.startswith('SEQUENCE'):
            return self._parse_sequence(name, definition)
        
        # Check for CHOICE
        elif definition.startswith('CHOICE'):
            return self._parse_choice(name, definition)
        
        # Check for ENUMERATED
        elif definition.startswith('ENUMERATED'):
            return self._parse_enumerated(name, definition)
        
        # Check for INTEGER
        elif definition.startswith('INTEGER'):
            return self._parse_integer(name, definition)
        
        # Check for BIT STRING
        elif 'BIT STRING' in definition or 'BITSTRING' in definition.upper():
            return self._parse_bit_string(name, definition)
        
        # Check for OCTET STRING
        elif 'OCTET STRING' in definition or 'OCTETSTRING' in definition:
            return self._parse_octet_string(name, definition)
        
        # Check for BOOLEAN
        elif definition.startswith('BOOLEAN'):
            return Asn1Type(name=name, type_name='BOOLEAN')
        
        # Check for NULL
        elif definition.startswith('NULL'):
            return None  # Skip NULL types
        
        # Otherwise, it's a type reference
        else:
            return self._parse_type_reference(name, definition)
    
    def _parse_sequence(self, name: str, definition: str) -> Optional[Asn1Type]:
        """Parse SEQUENCE or SEQUENCE OF"""
        # Check for SEQUENCE OF (must match at the start of definition, not in the middle)
        # Pattern 1: SEQUENCE (SIZE(...)) OF Type
        match = re.match(r'^SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+([a-zA-Z][a-zA-Z0-9\-]*)', definition)
        if match:
            size_constraint = match.group(1)
            element_type = match.group(2)

            asn1_type = Asn1Type(
                name=name,
                type_name='SEQUENCE_OF',
                element_type=element_type,
            )
            
            # Check if it's a range (contains ..) or a single value
            if '..' in size_constraint:
                constraints = self._parse_constraint(size_constraint)
                asn1_type.lb = constraints['lb']
                asn1_type.ub = constraints['ub']
            else:
                # Single value: both lb and ub are the same
                value = size_constraint.strip()
                asn1_type.lb = value
                asn1_type.ub = value
            
            return asn1_type

        # Pattern 2: SEQUENCE OF Type (without SIZE)
        match = re.match(r'^SEQUENCE\s+OF\s+([a-zA-Z][a-zA-Z0-9\-]*)', definition)
        if match:
            element_type = match.group(1)

            asn1_type = Asn1Type(
                name=name,
                type_name='SEQUENCE_OF',
                element_type=element_type,
            )
            
            return asn1_type
        
        # Regular SEQUENCE with fields
        # Extract content between { } with balanced braces
        content = self._extract_braced_content(definition, 'SEQUENCE')
        if content is None:
            return asn1_type
        
        # Check if it's empty
        if not content.strip() or content.strip() == '':
            return None
        
        asn1_type = Asn1Type(name=name, type_name='SEQUENCE')
        
        # Parse fields
        fields = self._parse_fields(content, name)
        asn1_type.fields = fields
        
        # Check for extension marker
        if has_ellipsis_at_root_level(content):
            asn1_type.ext = True

        return asn1_type
    
    def _parse_choice(self, name: str, definition: str) -> Optional[Asn1Type]:
        """Parse CHOICE type"""
        # Extract content between { } with balanced braces
        content = self._extract_braced_content(definition, 'CHOICE')
        if content is None:
            return None
        
        asn1_type = Asn1Type(name=name, type_name='CHOICE')
        
        # Parse alternatives (same as fields)
        fields = self._parse_fields(content, name)
        asn1_type.fields = fields
        
        # Check for extension marker
        if '...' in content:
            asn1_type.ext = True
        
        return asn1_type
    
    def _parse_fields(self, content: str, parent_name: str) -> List[Asn1Field]:
        """Parse fields from SEQUENCE or CHOICE content"""
        fields = []
        
        # Track extension state
        after_extension_marker = False
        in_extension_group = False
        
        # Build a map of field names to their extension status
        lines = content.split('\n')
        field_extension_map = {}
        field_comment_map = {}
        current_field_name = None
        current_ext_group_index = 0
        active_ext_group_index = 0
        
        for line in lines:
            stripped = line.strip()
            
            # Check for extension marker
            if stripped == '...' or stripped == '...,':
                after_extension_marker = True
                continue
            
            # Check for extension group markers
            if stripped.startswith('-- [[') or '-- [[' in stripped:
                if after_extension_marker:
                    current_ext_group_index += 1
                    active_ext_group_index = current_ext_group_index
                in_extension_group = True
                continue
            elif stripped.startswith('-- ]]') or '-- ]]' in stripped:
                in_extension_group = False
                active_ext_group_index = 0
                continue
            
            # Skip other comment-only lines
            if stripped.startswith('--'):
                continue
            
            # Try to extract field name from this line
            # Field names are the first identifier on a line (not part of a comment)
            if get_after_dash(line):
                line_without_comment = get_after_dash(line)[0]
            else:
                line_without_comment = line
            
            # Check if this line starts a field definition
            tokens = line_without_comment.strip().split()

            if tokens and tokens[0] and tokens[0][0].islower():  # Field names start with lowercase
                current_field_name = tokens[0]
                is_extension = after_extension_marker or active_ext_group_index > 0
                field_extension_map[current_field_name] = is_extension
                comment_value = ""
                if active_ext_group_index > 0:
                    comment_value = f"ext-{active_ext_group_index}"
                field_comment_map[current_field_name] = comment_value

        # Now process fields normally but use the extension map
        processed_lines = []

        for line in lines:
            # Check for inline comment
            if line.startswith('--'):
                continue
            elif "-- [[" in line or "-- ]]" in line:
                continue
            elif get_after_dash(line):
                line = get_after_dash(line)[0]
            
            processed_lines.append(line)

        content = '\n'.join(processed_lines)

        # Split by comma, but respect braces
        field_strings = self._split_by_comma(content)

        for field_str in field_strings:
            field_str = field_str.strip()
            # Skip empty, extension markers
            if not field_str or field_str == '...':
                continue
            
            field = self._parse_field(field_str, parent_name)
            if field:
                # Apply extension status from the map
                if field.name in field_extension_map:
                    field.ext = field_extension_map[field.name]
                    comment_value = field_comment_map.get(field.name, "")
                    if comment_value:
                        field.comment = comment_value
                if field.comment is None:
                    field.comment = ""
                
                fields.append(field)
        
        return fields
    
    def _parse_field(self, field_str: str, parent_name: str) -> Optional[Asn1Field]:
        """Parse a single field definition"""
        # Remove leading comments (from previous field)
        # Comments start with -- at the beginning of a line or after whitespace
        lines = field_str.split('\n')

        cleaned_lines = []
        for line in lines:
            stripped = line.strip()
            # Skip lines that are only comments
            if stripped.startswith('--'):
                continue
            cleaned_lines.append(line)
        
        field_str = '\n'.join(cleaned_lines).strip()
        
        if not field_str:
            return None

        # Extract inline comment (after the field definition)
        comment = None
        if '--' in field_str:
            parts = field_str.split('--', 1)
            field_str = parts[0].strip()
            comment = parts[1].strip()
                
        # Check for OPTIONAL
        field_str, optional = parse_optional_string(field_str)

        # Parse field: fieldName Type [constraints]
        # Pattern: identifier whitespace Type
        tokens = field_str.split()
        
        if len(tokens) < 2:
            return None
        
        field_name = tokens[0]
        # The rest is the type (might include constraints)
        type_part = ' '.join(tokens[1:])
        # Normalize field name
        normalized_field_name = field_name

        # Check for SEQUENCE OF (must match at the start of definition, not in the middle)
        # Pattern 0: SEQUENCE (SIZE) OF BASE_Type (...)
        extract_result = extract_sequence_bounds(type_part)
        if extract_result['check']:
            asn1_type = Asn1Type(
                name=normalized_field_name,
                type_name='SEQUENCE_OF',
                optional=optional,
                lb=extract_result['lb'],
                ub=extract_result['ub'],
                element_type=extract_result['type'],
                element_lb=extract_result['lb_e'],
                element_ub=extract_result['ub_e']
            )
            return asn1_type
        # Pattern 1: SEQUENCE (SIZE(...)) OF Type
        match = re.match(r'^SEQUENCE\s*\(SIZE\s*\(([^)]+)\)\)\s*OF\s+([a-zA-Z][a-zA-Z0-9\-]*)', type_part)
        if match:
            size_constraint = match.group(1)
            element_type = match.group(2)
            constraints = self._parse_constraint(size_constraint)
            lb = ""
            ub = ""
            if '..' in size_constraint:
                parts = size_constraint.split('..')
                lb = parts[0].strip()
                ub = parts[1].strip()
            else:
                lb = size_constraint
                ub = size_constraint
            
            field = Asn1Field(
                name=normalized_field_name,
                type_name="SEQUENCE_OF",
                lb=lb,
                ub=ub,
                element_type=element_type,
                optional=optional,
            )
            return field
        # Pattern 2: SEQUENCE OF Type (without SIZE)
        match = re.match(r'^SEQUENCE\s+OF\s+([a-zA-Z][a-zA-Z0-9\-]*)', type_part)
        if match:
            element_type = match.group(1)
            field = Asn1Field(
                name=normalized_field_name,
                type_name="SEQUENCE_OF",
                element_type=element_type,
                optional=optional,
            )
            return field
        
        # Handle inline SEQUENCE or CHOICE FIRST (before simple types)
        # This prevents BITSTRING/OCTETSTRING inside CHOICE from being detected as the field type
        if 'SEQUENCE { ... }' == type_part or 'SEQUENCE {}' == type_part or 'SEQUENCE{}' == type_part:
            type_name = "Struct-Null"
        elif 'SEQUENCE' in type_part and '{' in type_part:
            # Inline SEQUENCE - need to extract and create nested type
            if normalized_field_name == 'c1':
                nested_name = f"{parent_name}-C1"
            elif normalized_field_name == 'criticalExtensions':
                nested_name = f"{parent_name}-{uppercase_first_char(normalized_field_name)}"
            elif normalized_field_name == 'C2':
                nested_name = f"{parent_name}-C2"
            elif normalized_field_name == 'messageClassExtension':
                nested_name = f"{parent_name}-MessageClassExtension"
            elif normalized_field_name == 'codebookType':
                nested_name = f"{parent_name}-codebookType"
            elif normalized_field_name == 'type1':
                nested_name = f"{parent_name}-type1"
            elif normalized_field_name == 'type2':
                nested_name = f"{parent_name}-type2"
            elif normalized_field_name == 'subType':
                nested_name = f"{parent_name}-subType"
            elif normalized_field_name == 'periodicity':
                nested_name = f"{parent_name}-periodicity"
            elif normalized_field_name.startswith('dummy'):
                nested_name = f"{parent_name}-{normalized_field_name}"
            else:
                nested_name = f"{parent_name}-{normalized_field_name}"
            nested_type = self.parse_type_definition(nested_name, type_part)
            if nested_type:
                nested_type.parent = parent_name
                self.types[nested_name] = nested_type
            type_name = nested_name
        elif 'CHOICE' in type_part and '{' in type_part:
            # Inline CHOICE - need to extract and create nested type
            if normalized_field_name == 'c1':
                nested_name = f"{parent_name}-C1"
            elif normalized_field_name == 'c2':
                nested_name = f"{parent_name}-C2"
            else:
                nested_name = f"{parent_name}-{normalized_field_name}"
            nested_type = self.parse_type_definition(nested_name, type_part)
            if nested_type:
                nested_type.parent = parent_name
                self.types[nested_name] = nested_type
            type_name = nested_name
        # Handle inline ENUMERATED
        elif 'ENUMERATED' in type_part:
            # nested_name = f"{parent_name}{normalized_field_name.capitalize()}" #########################
            nested_name = f"{parent_name}-{normalized_field_name}"
            nested_type = self._parse_enumerated(nested_name, type_part)
            if nested_type:
                nested_type.parent = parent_name
                self.types[nested_name] = nested_type
            type_name = nested_name
        # Handle SetupRelease
        elif 'SetupRelease' in type_part:
            # Extract the contained type
            match = re.search(r'SetupRelease\s*\{\s*([a-zA-Z][a-zA-Z0-9\-]*)\s*\}', type_part)
            if match:
                contained_type = match.group(1).strip()
                type_name = f"SetupRelease of {contained_type}"
            else:
                type_name = "SetupRelease"
        # Handle OCTET STRING (CONTAINING ...) or OCTETSTRING (CONTAINING ...)
        elif 'CONTAINING' in type_part:
            match = re.search(r'CONTAINING\s+([a-zA-Z][a-zA-Z0-9\-]*)', type_part)
            if match:
                comment = f"Containing {match.group(1)}"
                type_name = "OctetString"
            else:
                type_name = "OctetString"
        # Handle plain OCTETSTRING or OCTET STRING
        elif re.search(r'\bOCTET\s*STRING\b', type_part, re.IGNORECASE) or re.search(r'\bOCTETSTRING\b', type_part, re.IGNORECASE):
            type_name = "OctetString"
        # Handle plain BITSTRING or BIT STRING  
        elif re.search(r'\bBIT\s*STRING\b', type_part, re.IGNORECASE) or re.search(r'\bBITSTRING\b', type_part, re.IGNORECASE):
            type_name = "BitString"
        else:
            # Extract type name (first token or before constraints)
            type_match = re.match(r'([a-zA-Z][a-zA-Z0-9\-]*)', type_part)
            if type_match:
                type_name = type_match.group(1)
            else:
                # Handle special types
                type_part_upper = type_part.upper()
                if 'OCTETSTRING' in type_part_upper:
                    type_name = 'OctetString'
                elif 'BITSTRING' in type_part_upper:
                    type_name = 'BitString'
                elif 'INTEGER' in type_part_upper:
                    type_name = 'Integer'
                elif 'BOOLEAN' in type_part_upper:
                    type_name = 'Boolean'
                else:
                    return None
        
        # Parse constraints
        lb, ub = self._parse_field_constraints(type_part)
        field = Asn1Field(
            name=normalized_field_name,
            type_name=type_name,
            optional=optional,
            lb=lb,
            ub=ub,
            comment=comment if comment is not None else ""
        )
        return field
    
    def _parse_enumerated(self, name: str, definition: str) -> Optional[Asn1Type]:
        """Parse ENUMERATED type"""
        # Extract values between { } with balanced braces
        content = self._extract_braced_content(definition, 'ENUMERATED')
        if content is None:
            return None
        
        # Check if there's an extension marker
        has_extension = '...' in content
        
        # Split content by extension marker first
        if has_extension:
            # Find the extension marker position
            # Handle cases: "val1, val2, ..." or "val1, val2, ..., val3" or "val1, val2, ... val3"
            parts = content.split('...')
            root_part = parts[0] if len(parts) > 0 else ""
            ext_part = parts[1] if len(parts) > 1 else ""
        else:
            root_part = content
            ext_part = ""
        
        # Parse root values (before ...)
        root_values = []
        for value in root_part.split(','):
            value = value.strip()
            # Remove comments
            if '--' in value:
                value = value.split('--')[0].strip()
            # Skip extension group markers
            if value.startswith('[[') or value.startswith(']]'):
                continue
            if value:
                # Normalize value name
                root_values.append(value)
        
        # Parse extension values (after ...)
        ext_values = []
        if ext_part:
            for value in ext_part.split(','):
                value = value.strip()
                # Remove comments
                if '--' in value:
                    value = value.split('--')[0].strip()
                # Skip extension group markers
                if value.startswith('[[') or value.startswith(']]'):
                    continue
                if value:
                    # Normalize value name
                    ext_values.append(value)
        
        # Combine all values
        all_values = root_values + ext_values
        
        asn1_type = Asn1Type(
            name=name,
            type_name='ENUMERATED',
            values=all_values
        )
        
        # Set lb/ub based on root values (before extension)
        if root_values:
            asn1_type.lb = 0
            asn1_type.ub = len(root_values) - 1
        
        # Mark if there's an extension
        if has_extension:
            asn1_type.ext = True
        
        return asn1_type

    def _parse_integer(self, name: str, definition: str) -> Asn1Type:
        """Parse INTEGER type"""
        asn1_type = Asn1Type(name=name, type_name='INTEGER')
        
        # Extract constraints
        match = re.search(r'\(([^)]+)\)', definition)
        if match:
            constraint_str = match.group(1)
            constraints = self._parse_constraint(constraint_str)
            asn1_type.lb = constraints['lb']
            asn1_type.ub = constraints['ub']
        
        return asn1_type
    
    def _parse_bit_string(self, name: str, definition: str) -> Asn1Type:
        """Parse BIT STRING type"""
        asn1_type = Asn1Type(name=name, type_name='BIT_STRING')
        # Extract SIZE constraint
        match = re.search(r'SIZE\s*\(([^)]+)\)', definition)
        if match:
            size_str = match.group(1)
            constraints = self._parse_constraint(size_str)
            asn1_type.lb = constraints['lb']
            asn1_type.ub = constraints['ub']
        
        return asn1_type
    
    def _parse_octet_string(self, name: str, definition: str) -> Asn1Type:
        """Parse OCTET STRING type"""
        # Check for CONTAINING
        if 'CONTAINING' in definition:
            match = re.search(r'CONTAINING\s+([a-zA-Z][a-zA-Z0-9\-]*)', definition)
            if match:
                contained_type = match.group(1)
                asn1_type = Asn1Type(
                    name=name,
                    type_name='OCTET_STRING',
                    element_type=contained_type
                )
                return asn1_type
        
        asn1_type = Asn1Type(name=name, type_name='OCTET_STRING')
        
        # Extract SIZE constraint
        match = re.search(r'SIZE\s*\(([^)]+)\)', definition)
        if match:
            size_str = match.group(1)
            constraints = self._parse_constraint(size_str)
            asn1_type.lb = constraints['lb']
            asn1_type.ub = constraints['ub']
        
        return asn1_type
    
    def _parse_type_reference(self, name: str, definition: str) -> Asn1Type:
        """Parse a type reference (alias)"""
        # Extract the referenced type name
        type_match = re.match(r'([a-zA-Z][a-zA-Z0-9\-]*)', definition)
        if type_match:
            ref_type = type_match.group(1)
            asn1_type = Asn1Type(
                name=name,
                type_name='TYPE_REFERENCE',
                element_type=ref_type
            )
            return asn1_type
        return None
    
    def _parse_constraint(self, constraint_str: str) -> Dict[str, Any]:
        """Parse constraint string and return dict with lb/ub"""
        constraints = {}
        
        # Pattern: x..y or x or MIN..MAX
        if '..' in constraint_str:
            parts = constraint_str.split('..')
            lb = parts[0].strip()
            ub = parts[1].strip()
            
            # Try to convert to int
            try:
                constraints['lb'] = int(lb)
            except ValueError:
                constraints['lb'] = lb
            
            try:
                constraints['ub'] = int(ub)
            except ValueError:
                constraints['ub'] = ub
        else:
            # Single value
            try:
                val = int(constraint_str.strip())
                constraints['lb'] = val
                constraints['ub'] = val
            except ValueError:
                constraints['value'] = constraint_str.strip()
        
        return constraints
    
    def _parse_field_constraints(self, type_part: str) -> tuple:
        """Parse constraints from field type part"""
        lb = None
        ub = None

        if 'CONTAINING' in type_part:
            return None, None
        
        # Look for (SIZE (...)) or (...)
        match = re.search(r'\(SIZE\s*\(([^)]+)\)\)', type_part)
        if not match:
            match = re.search(r'\(([^)]+)\)', type_part)
        
        if match:
            constraint_str = match.group(1).strip()
            
            # Check if it's a range (contains ..)
            if '..' in constraint_str:
                parts = constraint_str.split('..')
                lb = parts[0].strip()
                ub = parts[1].strip()
            else:
                # Single value: both lb and ub are the same
                lb = constraint_str
                ub = constraint_str

        return lb, ub
    
    def _split_by_comma(self, content: str) -> List[str]:
        """Split content by comma, respecting braces and brackets"""
        parts = []
        current = []
        brace_count = 0
        bracket_count = 0
        
        for char in content:
            if char == '{':
                brace_count += 1
            elif char == '}':
                brace_count -= 1
            elif char == '[':
                bracket_count += 1
            elif char == ']':
                bracket_count -= 1
            elif char == ',' and brace_count == 0 and bracket_count == 0:
                parts.append(''.join(current))
                current = []
                continue
            
            current.append(char)
        
        if current:
            parts.append(''.join(current))
        
        return parts
    
    def export_json(self, output_file: str):
        """Export parsed types to JSON file"""
        output = {name: typ.to_dict() for name, typ in self.types.items()}
        
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(output, f, indent=2)
        
        print(f"Exported {len(self.types)} types to {output_file}")
    
    def print_summary(self):
        """Print summary statistics"""
        kind_counts = {}
        for typ in self.types.values():
            kind_counts[typ.type_name] = kind_counts.get(typ.type_name, 0) + 1
        
        print("\n=== ASN.1 Parser Summary ===")
        print(f"Total types parsed: {len(self.types)}")
        print("\nBreakdown by kind:")
        for kind, count in sorted(kind_counts.items()):
            print(f"  {kind}: {count}")


def main():    
    # Check if file exists
    if not asn1_file.exists():
        print(f"Error: ASN.1 file not found: {asn1_file}")
        sys.exit(1)
    
    print(f"Parsing ASN.1 file: {asn1_file}")
    
    # Parse the file
    parser = Asn1Parser()
    parser.parse_file(str(asn1_file))
    
    # Print summary
    parser.print_summary()
    
    # Export to JSON
    parser.export_json(str(output_file))
    
    print(f"\nDone! Check {output_file} for results.")


if __name__ == '__main__':
    main()

