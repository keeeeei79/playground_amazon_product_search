# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: search.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0csearch.proto\x12\x06search\"\x1a\n\x07Request\x12\x0f\n\x07keyword\x18\x01 \x01(\t\"%\n\x08Response\x12\x19\n\x04\x64ocs\x18\x01 \x03(\x0b\x32\x0b.search.Doc\"i\n\x03\x44oc\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x14\n\x0c\x62ullet_point\x18\x04 \x01(\t\x12\r\n\x05\x62rand\x18\x05 \x01(\t\x12\r\n\x05\x63olor\x18\x06 \x01(\t2<\n\rSearchService\x12+\n\x06Search\x12\x0f.search.Request\x1a\x10.search.ResponseBDZBgithub.com/keeeeei79/playground_amazon_product_search/proto/;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'search_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/keeeeei79/playground_amazon_product_search/proto/;proto'
  _globals['_REQUEST']._serialized_start=24
  _globals['_REQUEST']._serialized_end=50
  _globals['_RESPONSE']._serialized_start=52
  _globals['_RESPONSE']._serialized_end=89
  _globals['_DOC']._serialized_start=91
  _globals['_DOC']._serialized_end=196
  _globals['_SEARCHSERVICE']._serialized_start=198
  _globals['_SEARCHSERVICE']._serialized_end=258
# @@protoc_insertion_point(module_scope)
