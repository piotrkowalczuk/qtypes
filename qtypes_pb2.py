# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: qtypes.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='qtypes.proto',
  package='qtypes',
  syntax='proto3',
  serialized_pb=_b('\n\x0cqtypes.proto\x12\x06qtypes\x1a\x1fgoogle/protobuf/timestamp.proto\"o\n\x06String\x12\x0e\n\x06values\x18\x01 \x03(\t\x12\r\n\x05valid\x18\x02 \x01(\x08\x12\x10\n\x08negation\x18\x03 \x01(\x08\x12\x1f\n\x04type\x18\x04 \x01(\x0e\x32\x11.qtypes.QueryType\x12\x13\n\x0binsensitive\x18\x05 \x01(\x08\"Y\n\x05Int64\x12\x0e\n\x06values\x18\x01 \x03(\x03\x12\r\n\x05valid\x18\x02 \x01(\x08\x12\x10\n\x08negation\x18\x03 \x01(\x08\x12\x1f\n\x04type\x18\x04 \x01(\x0e\x32\x11.qtypes.QueryType\"Z\n\x06Uint64\x12\x0e\n\x06values\x18\x01 \x03(\x04\x12\r\n\x05valid\x18\x02 \x01(\x08\x12\x10\n\x08negation\x18\x03 \x01(\x08\x12\x1f\n\x04type\x18\x04 \x01(\x0e\x32\x11.qtypes.QueryType\"[\n\x07\x46loat64\x12\x0e\n\x06values\x18\x01 \x03(\x01\x12\r\n\x05valid\x18\x02 \x01(\x08\x12\x10\n\x08negation\x18\x03 \x01(\x08\x12\x1f\n\x04type\x18\x04 \x01(\x0e\x32\x11.qtypes.QueryType\"y\n\tTimestamp\x12*\n\x06values\x18\x01 \x03(\x0b\x32\x1a.google.protobuf.Timestamp\x12\r\n\x05valid\x18\x02 \x01(\x08\x12\x10\n\x08negation\x18\x03 \x01(\x08\x12\x1f\n\x04type\x18\x04 \x01(\x0e\x32\x11.qtypes.QueryType*\xb7\x02\n\tQueryType\x12\x08\n\x04NULL\x10\x00\x12\t\n\x05\x45QUAL\x10\x01\x12\x0b\n\x07GREATER\x10\x02\x12\x11\n\rGREATER_EQUAL\x10\x03\x12\x08\n\x04LESS\x10\x04\x12\x0e\n\nLESS_EQUAL\x10\x05\x12\x06\n\x02IN\x10\x06\x12\x0b\n\x07\x42\x45TWEEN\x10\x07\x12\x0e\n\nHAS_PREFIX\x10\x08\x12\x0e\n\nHAS_SUFFIX\x10\t\x12\r\n\tSUBSTRING\x10\n\x12\x0b\n\x07PATTERN\x10\x0b\x12\x0e\n\nMIN_LENGTH\x10\x0c\x12\x0e\n\nMAX_LENGTH\x10\r\x12\x0b\n\x07OVERLAP\x10\x0e\x12\x0c\n\x08\x43ONTAINS\x10\x0f\x12\x13\n\x0fIS_CONTAINED_BY\x10\x10\x12\x0f\n\x0bHAS_ELEMENT\x10\x11\x12\x13\n\x0fHAS_ANY_ELEMENT\x10\x12\x12\x14\n\x10HAS_ALL_ELEMENTS\x10\x13\x42\"Z github.com/piotrkowalczuk/qtypesb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_QUERYTYPE = _descriptor.EnumDescriptor(
  name='QueryType',
  full_name='qtypes.QueryType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NULL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EQUAL', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GREATER', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GREATER_EQUAL', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LESS', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LESS_EQUAL', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='IN', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BETWEEN', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HAS_PREFIX', index=8, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HAS_SUFFIX', index=9, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBSTRING', index=10, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PATTERN', index=11, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MIN_LENGTH', index=12, number=12,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MAX_LENGTH', index=13, number=13,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OVERLAP', index=14, number=14,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONTAINS', index=15, number=15,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='IS_CONTAINED_BY', index=16, number=16,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HAS_ELEMENT', index=17, number=17,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HAS_ANY_ELEMENT', index=18, number=18,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HAS_ALL_ELEMENTS', index=19, number=19,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=570,
  serialized_end=881,
)
_sym_db.RegisterEnumDescriptor(_QUERYTYPE)

QueryType = enum_type_wrapper.EnumTypeWrapper(_QUERYTYPE)
NULL = 0
EQUAL = 1
GREATER = 2
GREATER_EQUAL = 3
LESS = 4
LESS_EQUAL = 5
IN = 6
BETWEEN = 7
HAS_PREFIX = 8
HAS_SUFFIX = 9
SUBSTRING = 10
PATTERN = 11
MIN_LENGTH = 12
MAX_LENGTH = 13
OVERLAP = 14
CONTAINS = 15
IS_CONTAINED_BY = 16
HAS_ELEMENT = 17
HAS_ANY_ELEMENT = 18
HAS_ALL_ELEMENTS = 19



_STRING = _descriptor.Descriptor(
  name='String',
  full_name='qtypes.String',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='qtypes.String.values', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='qtypes.String.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='negation', full_name='qtypes.String.negation', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='qtypes.String.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='insensitive', full_name='qtypes.String.insensitive', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=168,
)


_INT64 = _descriptor.Descriptor(
  name='Int64',
  full_name='qtypes.Int64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='qtypes.Int64.values', index=0,
      number=1, type=3, cpp_type=2, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='qtypes.Int64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='negation', full_name='qtypes.Int64.negation', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='qtypes.Int64.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=170,
  serialized_end=259,
)


_UINT64 = _descriptor.Descriptor(
  name='Uint64',
  full_name='qtypes.Uint64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='qtypes.Uint64.values', index=0,
      number=1, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='qtypes.Uint64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='negation', full_name='qtypes.Uint64.negation', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='qtypes.Uint64.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=261,
  serialized_end=351,
)


_FLOAT64 = _descriptor.Descriptor(
  name='Float64',
  full_name='qtypes.Float64',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='qtypes.Float64.values', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='qtypes.Float64.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='negation', full_name='qtypes.Float64.negation', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='qtypes.Float64.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=353,
  serialized_end=444,
)


_TIMESTAMP = _descriptor.Descriptor(
  name='Timestamp',
  full_name='qtypes.Timestamp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='qtypes.Timestamp.values', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='valid', full_name='qtypes.Timestamp.valid', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='negation', full_name='qtypes.Timestamp.negation', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='qtypes.Timestamp.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=446,
  serialized_end=567,
)

_STRING.fields_by_name['type'].enum_type = _QUERYTYPE
_INT64.fields_by_name['type'].enum_type = _QUERYTYPE
_UINT64.fields_by_name['type'].enum_type = _QUERYTYPE
_FLOAT64.fields_by_name['type'].enum_type = _QUERYTYPE
_TIMESTAMP.fields_by_name['values'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMESTAMP.fields_by_name['type'].enum_type = _QUERYTYPE
DESCRIPTOR.message_types_by_name['String'] = _STRING
DESCRIPTOR.message_types_by_name['Int64'] = _INT64
DESCRIPTOR.message_types_by_name['Uint64'] = _UINT64
DESCRIPTOR.message_types_by_name['Float64'] = _FLOAT64
DESCRIPTOR.message_types_by_name['Timestamp'] = _TIMESTAMP
DESCRIPTOR.enum_types_by_name['QueryType'] = _QUERYTYPE

String = _reflection.GeneratedProtocolMessageType('String', (_message.Message,), dict(
  DESCRIPTOR = _STRING,
  __module__ = 'qtypes_pb2'
  # @@protoc_insertion_point(class_scope:qtypes.String)
  ))
_sym_db.RegisterMessage(String)

Int64 = _reflection.GeneratedProtocolMessageType('Int64', (_message.Message,), dict(
  DESCRIPTOR = _INT64,
  __module__ = 'qtypes_pb2'
  # @@protoc_insertion_point(class_scope:qtypes.Int64)
  ))
_sym_db.RegisterMessage(Int64)

Uint64 = _reflection.GeneratedProtocolMessageType('Uint64', (_message.Message,), dict(
  DESCRIPTOR = _UINT64,
  __module__ = 'qtypes_pb2'
  # @@protoc_insertion_point(class_scope:qtypes.Uint64)
  ))
_sym_db.RegisterMessage(Uint64)

Float64 = _reflection.GeneratedProtocolMessageType('Float64', (_message.Message,), dict(
  DESCRIPTOR = _FLOAT64,
  __module__ = 'qtypes_pb2'
  # @@protoc_insertion_point(class_scope:qtypes.Float64)
  ))
_sym_db.RegisterMessage(Float64)

Timestamp = _reflection.GeneratedProtocolMessageType('Timestamp', (_message.Message,), dict(
  DESCRIPTOR = _TIMESTAMP,
  __module__ = 'qtypes_pb2'
  # @@protoc_insertion_point(class_scope:qtypes.Timestamp)
  ))
_sym_db.RegisterMessage(Timestamp)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z github.com/piotrkowalczuk/qtypes'))
# @@protoc_insertion_point(module_scope)
