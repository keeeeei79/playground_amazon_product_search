from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Request(_message.Message):
    __slots__ = ["keyword"]
    KEYWORD_FIELD_NUMBER: _ClassVar[int]
    keyword: str
    def __init__(self, keyword: _Optional[str] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ["docs"]
    DOCS_FIELD_NUMBER: _ClassVar[int]
    docs: _containers.RepeatedCompositeFieldContainer[Doc]
    def __init__(self, docs: _Optional[_Iterable[_Union[Doc, _Mapping]]] = ...) -> None: ...

class Doc(_message.Message):
    __slots__ = ["id", "title", "description", "bullet_point", "brand", "color"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    BULLET_POINT_FIELD_NUMBER: _ClassVar[int]
    BRAND_FIELD_NUMBER: _ClassVar[int]
    COLOR_FIELD_NUMBER: _ClassVar[int]
    id: str
    title: str
    description: str
    bullet_point: str
    brand: str
    color: str
    def __init__(self, id: _Optional[str] = ..., title: _Optional[str] = ..., description: _Optional[str] = ..., bullet_point: _Optional[str] = ..., brand: _Optional[str] = ..., color: _Optional[str] = ...) -> None: ...
